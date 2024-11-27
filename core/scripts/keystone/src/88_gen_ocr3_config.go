package src

import (
	helpers "github.com/goplugin/pluginv3.0/core/scripts/common"
	ksdeploy "github.com/goplugin/pluginv3.0/integration-tests/deployment/keystone"
)

func mustReadConfig(fileName string) (output ksdeploy.TopLevelConfigSource) {
	return mustParseJSON[ksdeploy.TopLevelConfigSource](fileName)
}

func generateOCR3Config(nodeList string, configFile string, chainID int64, pubKeysPath string) ksdeploy.Orc2drOracleConfig {
	topLevelCfg := mustReadConfig(configFile)
	cfg := topLevelCfg.OracleConfig
	nca := downloadNodePubKeys(nodeList, chainID, pubKeysPath)
	c, err := ksdeploy.GenerateOCR3Config(cfg, nca)
	helpers.PanicErr(err)
	return c
}
