package driven

import (
	"4clabs-server/app/4clabs-server/internal/adapter/driven/server"
	"4clabs-server/app/4clabs-server/internal/adapter/driven/service"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(server.ProviderSet,
	service.NewService,
)
