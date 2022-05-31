package driving

import (
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/domain/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	nftgo.NewService,
	wire.Bind(new(service.Query), new(*nftgo.Service)),
)
