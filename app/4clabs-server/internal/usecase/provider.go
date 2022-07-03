package usecase

import "github.com/google/wire"

var ProvderSet = wire.NewSet(
	NewAddress,
	NewNft,
	NewAuth,
	NewTicket,
)
