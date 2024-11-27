package evm

import (
	"context"

	commontypes "github.com/goplugin/plugin-common/pkg/types"
)

type readBinding interface {
	GetLatestValue(ctx context.Context, params, returnVal any) error
	Bind(binding commontypes.BoundContract) error
	SetCodec(codec commontypes.RemoteCodec)
	Register() error
	Unregister() error
}
