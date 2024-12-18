package keystest

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/goplugin/pluginv3.0/v2/core/services/keystore/keys/p2pkey"
)

func NewP2PKeyV2(t *testing.T) p2pkey.KeyV2 {
	k, err := p2pkey.NewV2()
	require.NoError(t, err)
	return k
}
