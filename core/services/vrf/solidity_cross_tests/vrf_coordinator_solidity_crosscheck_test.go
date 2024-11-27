package solidity_cross_tests_test

import (
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	proof2 "github.com/goplugin/pluginv3.0/v2/core/services/vrf/proof"
	"github.com/goplugin/pluginv3.0/v2/core/services/vrf/solidity_cross_tests"
	"github.com/goplugin/pluginv3.0/v2/core/services/vrf/vrftesthelpers"

	"github.com/goplugin/pluginv3.0/v2/core/chains/evm/utils"
	"github.com/goplugin/pluginv3.0/v2/core/services/keystore/keys/vrfkey"
	"github.com/goplugin/pluginv3.0/v2/core/services/signatures/secp256k1"

	"github.com/goplugin/pluginv3.0/v2/core/internal/cltest"
)

const defaultGasLimit uint32 = 500000

func TestRequestIDMatches(t *testing.T) {
	keyHash := common.HexToHash("0x01")
	key := cltest.MustGenerateRandomKey(t)
	baseContract := vrftesthelpers.NewVRFCoordinatorUniverse(t, key).RequestIDBase
	var seed = big.NewInt(1)
	solidityRequestID, err := baseContract.MakeRequestId(nil, keyHash, seed)
	require.NoError(t, err, "failed to calculate VRF requestID on simulated ethereum blockchain")
	goRequestLog := &solidity_cross_tests.RandomnessRequestLog{KeyHash: keyHash, Seed: seed}
	assert.Equal(t, common.Hash(solidityRequestID), goRequestLog.ComputedRequestID(),
		"solidity VRF requestID differs from golang requestID!")
}

var (
	rawSecretKey = big.NewInt(1) // never do this in production!
	secretKey    = vrfkey.MustNewV2XXXTestingOnly(rawSecretKey)
	publicKey    = (&secp256k1.Secp256k1{}).Point().Mul(secp256k1.IntToScalar(
		rawSecretKey), nil)
	hardcodedSeed = big.NewInt(0)
	vrfFee        = big.NewInt(7)
)

// registerProvingKey registers keyHash to neil in the VRFCoordinator universe
// represented by coordinator, with the given jobID and fee.
func registerProvingKey(t *testing.T, coordinator vrftesthelpers.CoordinatorUniverse) (
	keyHash [32]byte, jobID [32]byte, fee *big.Int) {
	copy(jobID[:], []byte("exactly 32 characters in length."))
	_, err := coordinator.RootContract.RegisterProvingKey(
		coordinator.Neil, vrfFee, coordinator.Neil.From, pair(secp256k1.Coordinates(publicKey)), jobID)
	require.NoError(t, err, "failed to register VRF proving key on VRFCoordinator contract")
	coordinator.Backend.Commit()
	keyHash = utils.MustHash(string(secp256k1.LongMarshal(publicKey)))
	return keyHash, jobID, vrfFee
}

func TestRegisterProvingKey(t *testing.T) {
	key := cltest.MustGenerateRandomKey(t)
	coord := vrftesthelpers.NewVRFCoordinatorUniverse(t, key)
	keyHash, jobID, fee := registerProvingKey(t, coord)
	log, err := coord.RootContract.FilterNewServiceAgreement(nil)
	require.NoError(t, err, "failed to subscribe to NewServiceAgreement logs on simulated ethereum blockchain")
	logCount := 0
	for log.Next() {
		logCount++
		assert.Equal(t, log.Event.KeyHash, keyHash, "VRFCoordinator logged a different keyHash than was registered")
		assert.True(t, fee.Cmp(log.Event.Fee) == 0, "VRFCoordinator logged a different fee than was registered")
	}
	require.Equal(t, 1, logCount, "unexpected NewServiceAgreement log generated by key VRF key registration")
	serviceAgreement, err := coord.RootContract.ServiceAgreements(nil, keyHash)
	require.NoError(t, err, "failed to retrieve previously registered VRF service agreement from VRFCoordinator")
	assert.Equal(t, coord.Neil.From, serviceAgreement.VRFOracle,
		"VRFCoordinator registered wrong provider, on service agreement!")
	assert.Equal(t, jobID, serviceAgreement.JobID,
		"VRFCoordinator registered wrong jobID, on service agreement!")
	assert.True(t, fee.Cmp(serviceAgreement.Fee) == 0,
		"VRFCoordinator registered wrong fee, on service agreement!")
}

func TestFailToRegisterProvingKeyFromANonOwnerAddress(t *testing.T) {
	key := cltest.MustGenerateRandomKey(t)
	coordinator := vrftesthelpers.NewVRFCoordinatorUniverse(t, key)

	var jobID [32]byte
	copy(jobID[:], []byte("exactly 32 characters in length."))
	_, err := coordinator.RootContract.RegisterProvingKey(
		coordinator.Ned, vrfFee, coordinator.Neil.From, pair(secp256k1.Coordinates(publicKey)), jobID)

	require.Error(t, err, "expected an error")
	require.Contains(t, err.Error(), "Ownable: caller is not the owner")
}

// requestRandomness sends a randomness request via Carol's consuming contract,
// in the VRFCoordinator universe represented by coordinator, specifying the
// given keyHash and seed, and paying the given fee. It returns the log emitted
// from the VRFCoordinator in response to the request
func requestRandomness(t *testing.T, coordinator vrftesthelpers.CoordinatorUniverse,
	keyHash common.Hash, fee *big.Int) *solidity_cross_tests.RandomnessRequestLog {
	_, err := coordinator.ConsumerContract.TestRequestRandomness(coordinator.Carol,
		keyHash, fee)
	require.NoError(t, err, "problem during initial VRF randomness request")
	coordinator.Backend.Commit()
	log, err := coordinator.RootContract.FilterRandomnessRequest(nil, nil)
	require.NoError(t, err, "failed to subscribe to RandomnessRequest logs")
	logCount := 0
	for log.Next() {
		logCount++
	}
	require.Equal(t, 1, logCount, "unexpected log generated by randomness request to VRFCoordinator")
	return solidity_cross_tests.RawRandomnessRequestLogToRandomnessRequestLog(
		(*solidity_cross_tests.RawRandomnessRequestLog)(log.Event))
}

func requestRandomnessV08(t *testing.T, coordinator vrftesthelpers.CoordinatorUniverse,
	keyHash common.Hash, fee *big.Int) *solidity_cross_tests.RandomnessRequestLog {
	_, err := coordinator.ConsumerContractV08.DoRequestRandomness(coordinator.Carol,
		keyHash, fee)
	require.NoError(t, err, "problem during initial VRF randomness request")
	coordinator.Backend.Commit()
	log, err := coordinator.RootContract.FilterRandomnessRequest(nil, nil)
	require.NoError(t, err, "failed to subscribe to RandomnessRequest logs")
	logCount := 0
	for log.Next() {
		if log.Event.Sender == coordinator.ConsumerContractAddressV08 {
			logCount++
		}
	}
	require.Equal(t, 1, logCount, "unexpected log generated by randomness request to VRFCoordinator")
	return solidity_cross_tests.RawRandomnessRequestLogToRandomnessRequestLog(
		(*solidity_cross_tests.RawRandomnessRequestLog)(log.Event))
}

func TestRandomnessRequestLog(t *testing.T) {
	key := cltest.MustGenerateRandomKey(t)
	coord := vrftesthelpers.NewVRFCoordinatorUniverseWithV08Consumer(t, key)
	keyHash_, jobID_, fee := registerProvingKey(t, coord)
	keyHash := common.BytesToHash(keyHash_[:])
	jobID := common.BytesToHash(jobID_[:])
	var tt = []struct {
		rr func(t *testing.T, coordinator vrftesthelpers.CoordinatorUniverse,
			keyHash common.Hash, fee *big.Int) *solidity_cross_tests.RandomnessRequestLog
		ms              func() (*big.Int, error)
		consumerAddress common.Address
	}{
		{
			rr: requestRandomness,
			ms: func() (*big.Int, error) {
				return coord.RequestIDBase.MakeVRFInputSeed(nil, keyHash, hardcodedSeed, coord.ConsumerContractAddress, big.NewInt(0))
			},
			consumerAddress: coord.ConsumerContractAddress,
		},
		{
			rr: requestRandomnessV08,
			ms: func() (*big.Int, error) {
				return coord.RequestIDBaseV08.MakeVRFInputSeed(nil, keyHash, hardcodedSeed, coord.ConsumerContractAddressV08, big.NewInt(0))
			},
			consumerAddress: coord.ConsumerContractAddressV08,
		},
	}
	for _, tc := range tt {
		log := tc.rr(t, coord, keyHash, fee)
		assert.Equal(t, keyHash, log.KeyHash, "VRFCoordinator logged wrong KeyHash for randomness request")
		nonce := big.NewInt(0)
		actualSeed, err := tc.ms()
		require.NoError(t, err, "failure while using VRFCoordinator to calculate actual VRF input seed")
		assert.True(t, actualSeed.Cmp(log.Seed) == 0,
			"VRFCoordinator logged wrong actual input seed from randomness request")
		golangSeed := utils.MustHash(string(append(append(append(
			keyHash[:],
			common.BigToHash(hardcodedSeed).Bytes()...),
			common.BytesToHash(tc.consumerAddress.Bytes()).Bytes()...),
			common.BigToHash(nonce).Bytes()...)))
		assert.Equal(t, golangSeed, common.BigToHash((log.Seed)), "VRFCoordinator logged different actual input seed than expected by golang code!")
		assert.Equal(t, jobID, log.JobID, "VRFCoordinator logged different JobID from randomness request!")
		assert.Equal(t, tc.consumerAddress, log.Sender, "VRFCoordinator logged different requester address from randomness request!")
		assert.True(t, fee.Cmp((*big.Int)(log.Fee)) == 0, "VRFCoordinator logged different fee from randomness request!")
		parsedLog, err := solidity_cross_tests.ParseRandomnessRequestLog(log.Raw.Raw)
		assert.NoError(t, err, "could not parse randomness request log generated by VRFCoordinator")
		assert.True(t, parsedLog.Equal(*log), "got a different randomness request log by parsing the raw data than reported by simulated backend")
	}
}

// fulfillRandomnessRequest is neil fulfilling randomness requested by log.
func fulfillRandomnessRequest(t *testing.T, coordinator vrftesthelpers.CoordinatorUniverse, log solidity_cross_tests.RandomnessRequestLog) vrfkey.Proof {
	preSeed, err := proof2.BigToSeed(log.Seed)
	require.NoError(t, err, "pre-seed %x out of range", preSeed)
	s := proof2.PreSeedData{
		PreSeed:   preSeed,
		BlockHash: log.Raw.Raw.BlockHash,
		BlockNum:  log.Raw.Raw.BlockNumber,
	}
	seed := proof2.FinalSeed(s)
	proof, err := secretKey.GenerateProofWithNonce(seed, big.NewInt(1) /* nonce */)
	require.NoError(t, err)
	proofBlob, err := vrftesthelpers.GenerateProofResponseFromProof(proof, s)
	require.NoError(t, err, "could not generate VRF proof!")
	// Seems to be a bug in the simulated backend: without this extra Commit, the
	// EVM seems to think it's still on the block in which the request was made,
	// which means that the relevant blockhash is unavailable.
	coordinator.Backend.Commit()
	// This is simulating a node response, so set the gas limit as plugin does
	var neil bind.TransactOpts = *coordinator.Neil
	neil.GasLimit = uint64(defaultGasLimit)
	_, err = coordinator.RootContract.FulfillRandomnessRequest(&neil, proofBlob[:])
	require.NoError(t, err, "failed to fulfill randomness request!")
	coordinator.Backend.Commit()
	return proof
}

func TestFulfillRandomness(t *testing.T) {
	key := cltest.MustGenerateRandomKey(t)
	coordinator := vrftesthelpers.NewVRFCoordinatorUniverse(t, key)
	keyHash, _, fee := registerProvingKey(t, coordinator)
	randomnessRequestLog := requestRandomness(t, coordinator, keyHash, fee)
	proof := fulfillRandomnessRequest(t, coordinator, *randomnessRequestLog)
	output, err := coordinator.ConsumerContract.RandomnessOutput(nil)
	require.NoError(t, err, "failed to get VRF output from consuming contract, "+
		"after randomness request was fulfilled")
	assert.True(t, proof.Output.Cmp(output) == 0, "VRF output from randomness "+
		"request fulfillment was different than provided! Expected %d, got %d. "+
		"This can happen if you update the VRFCoordinator wrapper without a "+
		"corresponding update to the VRFConsumer", proof.Output, output)
	requestID, err := coordinator.ConsumerContract.RequestId(nil)
	require.NoError(t, err, "failed to get requestId from VRFConsumer")
	assert.Equal(t, randomnessRequestLog.RequestID, common.Hash(requestID),
		"VRFConsumer has different request ID than logged from randomness request!")
	neilBalance, err := coordinator.RootContract.WithdrawableTokens(
		nil, coordinator.Neil.From)
	require.NoError(t, err, "failed to get neil's token balance, after he "+
		"successfully fulfilled a randomness request")
	assert.True(t, neilBalance.Cmp(fee) == 0, "neil's balance on VRFCoordinator "+
		"was not paid his fee, despite successful fulfillment of randomness request!")
}

func TestWithdraw(t *testing.T) {
	key := cltest.MustGenerateRandomKey(t)
	coordinator := vrftesthelpers.NewVRFCoordinatorUniverse(t, key)
	keyHash, _, fee := registerProvingKey(t, coordinator)
	log := requestRandomness(t, coordinator, keyHash, fee)
	fulfillRandomnessRequest(t, coordinator, *log)
	payment := big.NewInt(4)
	peteThePunter := common.HexToAddress("0xdeadfa11deadfa11deadfa11deadfa11deadfa11")
	_, err := coordinator.RootContract.Withdraw(coordinator.Neil, peteThePunter, payment)
	require.NoError(t, err, "failed to withdraw PLI from neil's balance")
	coordinator.Backend.Commit()
	peteBalance, err := coordinator.LinkContract.BalanceOf(nil, peteThePunter)
	require.NoError(t, err, "failed to get balance of payee on PLI contract, after payment")
	assert.True(t, payment.Cmp(peteBalance) == 0,
		"PLI balance is wrong, following payment")
	neilBalance, err := coordinator.RootContract.WithdrawableTokens(
		nil, coordinator.Neil.From)
	require.NoError(t, err, "failed to get neil's balance on VRFCoordinator")
	assert.True(t, big.NewInt(0).Sub(fee, payment).Cmp(neilBalance) == 0,
		"neil's VRFCoordinator balance is wrong, after he's made a withdrawal!")
	_, err = coordinator.RootContract.Withdraw(coordinator.Neil, peteThePunter, fee)
	assert.Error(t, err, "VRFcoordinator allowed overdraft")
}