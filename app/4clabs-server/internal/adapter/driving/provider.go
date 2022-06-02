package driving

import (
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/ports"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	nftgo.NewService,
	wire.Bind(new(ports.Query), new(*nftgo.Service)),
)