package chaos

import (
	"fmt"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"

	"github.com/goplugin/plugin-testing-framework/blockchain"
	ctf_config "github.com/goplugin/plugin-testing-framework/config"
	"github.com/goplugin/plugin-testing-framework/k8s/chaos"
	"github.com/goplugin/plugin-testing-framework/k8s/environment"
	"github.com/goplugin/plugin-testing-framework/k8s/pkg/helm/plugin"
	"github.com/goplugin/plugin-testing-framework/k8s/pkg/helm/ethereum"
	"github.com/goplugin/plugin-testing-framework/logging"
	"github.com/goplugin/plugin-testing-framework/networks"
	"github.com/goplugin/plugin-testing-framework/utils/ptr"
	"github.com/goplugin/plugin-testing-framework/utils/testcontext"

	"github.com/goplugin/pluginv3.0/integration-tests/actions"
	"github.com/goplugin/pluginv3.0/integration-tests/actions/ocr2vrf_actions"
	"github.com/goplugin/pluginv3.0/integration-tests/actions/ocr2vrf_actions/ocr2vrf_constants"
	"github.com/goplugin/pluginv3.0/integration-tests/client"
	"github.com/goplugin/pluginv3.0/integration-tests/config"
	"github.com/goplugin/pluginv3.0/integration-tests/contracts"
	tc "github.com/goplugin/pluginv3.0/integration-tests/testconfig"
)

func TestOCR2VRFChaos(t *testing.T) {
	t.Parallel()
	l := logging.GetTestLogger(t)
	testconfig, err := tc.GetConfig("Chaos", tc.OCR2VRF)
	if err != nil {
		t.Fatal(err)
	}

	loadedNetwork := networks.MustGetSelectedNetworkConfig(testconfig.Network)[0]

	defaultOCR2VRFSettings := map[string]interface{}{
		"replicas": 6,
		"toml": networks.AddNetworkDetailedConfig(
			config.BaseOCR2Config,
			testconfig.Pyroscope,
			config.DefaultOCR2VRFNetworkDetailTomlConfig,
			loadedNetwork,
		),
	}

	defaultOCR2VRFEthereumSettings := &ethereum.Props{
		NetworkName: loadedNetwork.Name,
		Simulated:   loadedNetwork.Simulated,
		WsURLs:      loadedNetwork.URLs,
	}

	var overrideFn = func(_ interface{}, target interface{}) {
		ctf_config.MustConfigOverridePluginVersion(testconfig.PluginImage, target)
		ctf_config.MightConfigOverridePyroscopeKey(testconfig.Pyroscope, target)
	}

	pluginCfg := plugin.NewWithOverride(0, defaultOCR2VRFSettings, testconfig.PluginImage, overrideFn)

	testCases := map[string]struct {
		networkChart environment.ConnectedChart
		clChart      environment.ConnectedChart
		chaosFunc    chaos.ManifestFunc
		chaosProps   *chaos.Props
	}{
		// network-* and pods-* are split intentionally into 2 parallel groups
		// we can't use chaos.NewNetworkPartition and chaos.NewFailPods in parallel
		// because of jsii runtime bug, see Makefile

		// PodChaosFailMinorityNodes Test description:
		//1. DKG and VRF beacon processes are set and VRF request gets fulfilled
		//2. Apply chaos experiment - take down 2 nodes out of 5 non-bootstrap
		//3. Bring back all nodes to normal
		//4. verify VRF request gets fulfilled
		PodChaosFailMinorityNodes: {
			ethereum.New(defaultOCR2VRFEthereumSettings),
			pluginCfg,
			chaos.NewFailPods,
			&chaos.Props{
				LabelsSelector: &map[string]*string{ChaosGroupMinority: ptr.Ptr("1")},
				DurationStr:    "1m",
			},
		},
		//todo - currently failing, need to investigate deeper
		//PodChaosFailMajorityNodes: {
		//	ethereum.New(defaultOCR2VRFEthereumSettings),
		//	pluginCfg,
		//	chaos.NewFailPods,
		//	&chaos.Props{
		//		LabelsSelector: &map[string]*string{ChaosGroupMajority: ptr.Ptr("1")},
		//		DurationStr:    "1m",
		//	},
		//},
		//todo - do we need these chaos tests?
		//PodChaosFailMajorityDB: {
		//	ethereum.New(defaultOCR2VRFEthereumSettings),
		//	pluginCfg,
		//	chaos.NewFailPods,
		//	&chaos.Props{
		//		LabelsSelector: &map[string]*string{ChaosGroupMajority: ptr.Ptr("1")},
		//		DurationStr:    "1m",
		//		ContainerNames: &[]*string{ptr.Ptr("plugin-db")},
		//	},
		//},
		//NetworkChaosFailMajorityNetwork: {
		//	ethereum.New(defaultOCR2VRFEthereumSettings),
		//	pluginCfg,
		//	chaos.NewNetworkPartition,
		//	&chaos.Props{
		//		FromLabels:  &map[string]*string{ChaosGroupMajority: ptr.Ptr("1")},
		//		ToLabels:    &map[string]*string{ChaosGroupMinority: ptr.Ptr("1")},
		//		DurationStr: "1m",
		//	},
		//},
		//NetworkChaosFailBlockchainNode: {
		//	ethereum.New(defaultOCR2VRFEthereumSettings),
		//	pluginCfg,
		//	chaos.NewNetworkPartition,
		//	&chaos.Props{
		//		FromLabels:  &map[string]*string{"app": ptr.Ptr("geth")},
		//		ToLabels:    &map[string]*string{ChaosGroupMajority: ptr.Ptr("1")},
		//		DurationStr: "1m",
		//	},
		//},
	}

	for testCaseName, tc := range testCases {
		testCase := tc
		t.Run(fmt.Sprintf("OCR2VRF_%s", testCaseName), func(t *testing.T) {
			t.Parallel()
			testNetwork := networks.MustGetSelectedNetworkConfig(testconfig.Network)[0] // Need a new copy of the network for each test
			testEnvironment := environment.
				New(&environment.Config{
					NamespacePrefix: fmt.Sprintf(
						"chaos-ocr2vrf-%s", strings.ReplaceAll(strings.ToLower(testNetwork.Name), " ", "-"),
					),
					Test: t,
				}).
				AddHelm(testCase.networkChart).
				AddHelm(testCase.clChart)
			err := testEnvironment.Run()
			require.NoError(t, err, "Error running test environment")
			if testEnvironment.WillUseRemoteRunner() {
				return
			}

			err = testEnvironment.Client.LabelChaosGroup(testEnvironment.Cfg.Namespace, "instance=node-", 1, 2, ChaosGroupMinority)
			require.NoError(t, err)
			err = testEnvironment.Client.LabelChaosGroup(testEnvironment.Cfg.Namespace, "instance=node-", 3, 5, ChaosGroupMajority)
			require.NoError(t, err)

			chainClient, err := blockchain.NewEVMClient(testNetwork, testEnvironment, l)
			require.NoError(t, err, "Error connecting to blockchain")
			contractDeployer, err := contracts.NewContractDeployer(chainClient, l)
			require.NoError(t, err, "Error building contract deployer")
			pluginNodes, err := client.ConnectPluginNodes(testEnvironment)
			require.NoError(t, err, "Error connecting to Plugin nodes")
			nodeAddresses, err := actions.PluginNodeAddresses(pluginNodes)
			require.NoError(t, err, "Retrieving on-chain wallet addresses for plugin nodes shouldn't fail")

			t.Cleanup(func() {
				err := actions.TeardownSuite(t, testEnvironment, pluginNodes, nil, zapcore.PanicLevel, &testconfig, chainClient)
				require.NoError(t, err, "Error tearing down environment")
			})

			chainClient.ParallelTransactions(true)

			linkToken, err := contractDeployer.DeployLinkTokenContract()
			require.NoError(t, err, "Error deploying PLI token")

			mockETHLinkFeed, err := contractDeployer.DeployMockETHPLIFeed(ocr2vrf_constants.LinkEthFeedResponse)
			require.NoError(t, err, "Error deploying Mock ETH/PLI Feed")

			_, _, vrfBeaconContract, consumerContract, subID := ocr2vrf_actions.SetupOCR2VRFUniverse(
				t,
				linkToken,
				mockETHLinkFeed,
				contractDeployer,
				chainClient,
				nodeAddresses,
				pluginNodes,
				testNetwork,
			)

			//Request and Redeem Randomness to verify that process works fine
			requestID := ocr2vrf_actions.RequestAndRedeemRandomness(
				t,
				consumerContract,
				chainClient,
				vrfBeaconContract,
				ocr2vrf_constants.NumberOfRandomWordsToRequest,
				subID,
				ocr2vrf_constants.ConfirmationDelay,
				ocr2vrf_constants.RandomnessRedeemTransmissionEventTimeout,
			)

			for i := uint16(0); i < ocr2vrf_constants.NumberOfRandomWordsToRequest; i++ {
				randomness, err := consumerContract.GetRandomnessByRequestId(testcontext.Get(t), requestID, big.NewInt(int64(i)))
				require.NoError(t, err)
				l.Info().Interface("Random Number", randomness).Interface("Randomness Number Index", i).Msg("Randomness retrieved from Consumer contract")
				require.NotEqual(t, 0, randomness.Uint64(), "Randomness retrieved from Consumer contract give an answer other than 0")
			}

			id, err := testEnvironment.Chaos.Run(testCase.chaosFunc(testEnvironment.Cfg.Namespace, testCase.chaosProps))
			require.NoError(t, err, "Error running Chaos Experiment")
			l.Info().Msg("Chaos Applied")

			err = testEnvironment.Chaos.WaitForAllRecovered(id, time.Minute)
			require.NoError(t, err, "Error waiting for Chaos Experiment to end")
			l.Info().Msg("Chaos Recovered")

			//Request and Redeem Randomness again to see that after Chaos Experiment whole process is still working
			requestID = ocr2vrf_actions.RequestAndRedeemRandomness(
				t,
				consumerContract,
				chainClient,
				vrfBeaconContract,
				ocr2vrf_constants.NumberOfRandomWordsToRequest,
				subID,
				ocr2vrf_constants.ConfirmationDelay,
				ocr2vrf_constants.RandomnessRedeemTransmissionEventTimeout,
			)

			for i := uint16(0); i < ocr2vrf_constants.NumberOfRandomWordsToRequest; i++ {
				randomness, err := consumerContract.GetRandomnessByRequestId(testcontext.Get(t), requestID, big.NewInt(int64(i)))
				require.NoError(t, err, "Error getting Randomness result from Consumer Contract")
				l.Info().Interface("Random Number", randomness).Interface("Randomness Number Index", i).Msg("Randomness retrieved from Consumer contract")
				require.NotEqual(t, 0, randomness.Uint64(), "Randomness retrieved from Consumer contract give an answer other than 0")
			}
		})
	}
}
