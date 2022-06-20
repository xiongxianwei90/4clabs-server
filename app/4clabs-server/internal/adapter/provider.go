package adapter

import (
	"4clabs-server/app/4clabs-server/internal/adapter/driven"
	"4clabs-server/app/4clabs-server/internal/adapter/driving"
	"4clabs-server/app/4clabs-server/internal/data"
	"github.com/google/wire"
)

var Provider = wire.NewSet(
	driving.ProviderSet,
	driven.ProviderSet,
	data.NewData,
)
