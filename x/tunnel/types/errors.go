package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/tunnel module sentinel errors
var (
	ErrInvalidGenesis       = errorsmod.Register(ModuleName, 2, "invalid genesis")
	ErrMaxTunnelChannels    = errorsmod.Register(ModuleName, 3, "max tunnel channels")
	ErrTunnelNotFound       = errorsmod.Register(ModuleName, 4, "tunnel not found")
	ErrPacketNotFound       = errorsmod.Register(ModuleName, 5, "packet not found")
	ErrNoPacketContent      = errorsmod.Register(ModuleName, 6, "no packet content")
	ErrInvalidTunnelCreator = errorsmod.Register(ModuleName, 7, "invalid creator of tunnel")
	ErrAccountAlreadyExist  = errorsmod.Register(ModuleName, 8, "account already exist")
	ErrInvalidVersion       = errorsmod.Register(ModuleName, 9, "invalid ICS20 version")
)
