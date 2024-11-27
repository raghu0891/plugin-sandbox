package soak

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/goplugin/plugin-testing-framework/logging"

	"github.com/goplugin/pluginv3.0/integration-tests/actions"
	tc "github.com/goplugin/pluginv3.0/integration-tests/testconfig"
	"github.com/goplugin/pluginv3.0/integration-tests/testsetups"
)

func TestOCRSoak(t *testing.T) {
	l := logging.GetTestLogger(t)
	// Use this variable to pass in any custom EVM specific TOML values to your Plugin nodes
	customNetworkTOML := ``
	// Uncomment below for debugging TOML issues on the node
	// network := networks.MustGetSelectedNetworksFromEnv()[0]
	// fmt.Println("Using Plugin TOML\n---------------------")
	// fmt.Println(networks.AddNetworkDetailedConfig(config.BaseOCR1Config, customNetworkTOML, network))
	// fmt.Println("---------------------")

	config, err := tc.GetConfig("Soak", tc.OCR)
	require.NoError(t, err, "Error getting config")

	ocrSoakTest, err := testsetups.NewOCRSoakTest(t, &config, false)
	require.NoError(t, err, "Error creating soak test")
	if !ocrSoakTest.Interrupted() {
		ocrSoakTest.DeployEnvironment(customNetworkTOML, &config)
	}
	if ocrSoakTest.Environment().WillUseRemoteRunner() {
		return
	}
	t.Cleanup(func() {
		if err := actions.TeardownRemoteSuite(ocrSoakTest.TearDownVals(t)); err != nil {
			l.Error().Err(err).Msg("Error tearing down environment")
		}
	})
	if ocrSoakTest.Interrupted() {
		err = ocrSoakTest.LoadState()
		require.NoError(t, err, "Error loading state")
		ocrSoakTest.Resume()
	} else {
		ocrSoakTest.Setup(&config)
		ocrSoakTest.Run()
	}
}
