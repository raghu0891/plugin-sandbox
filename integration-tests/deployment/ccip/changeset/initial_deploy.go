package changeset

import (
	"github.com/goplugin/ccip-owner-contracts/pkg/proposal/timelock"

	"github.com/goplugin/pluginv3.0/integration-tests/deployment"

	ccipdeployment "github.com/goplugin/pluginv3.0/integration-tests/deployment/ccip"
)

// We expect the change set input to be unique per change set.
// TODO: Maybe there's a generics approach here?
// Note if the change set is a deployment and it fails we have 2 options:
// - Just throw away the addresses, fix issue and try again (potentially expensive on mainnet)
func InitialDeployChangeSet(env deployment.Environment, c ccipdeployment.DeployCCIPContractConfig) (deployment.ChangesetOutput, error) {
	ab := deployment.NewMemoryAddressBook()
	err := ccipdeployment.DeployCCIPContracts(env, ab, c)
	if err != nil {
		env.Logger.Errorw("Failed to deploy CCIP contracts", "err", err, "addresses", ab)
		return deployment.ChangesetOutput{AddressBook: ab}, deployment.MaybeDataErr(err)
	}
	js, err := ccipdeployment.NewCCIPJobSpecs(env.NodeIDs, env.Offchain)
	if err != nil {
		return deployment.ChangesetOutput{AddressBook: ab}, err
	}
	return deployment.ChangesetOutput{
		Proposals:   []timelock.MCMSWithTimelockProposal{},
		AddressBook: ab,
		// Mapping of which nodes get which jobs.
		JobSpecs: js,
	}, nil
}
