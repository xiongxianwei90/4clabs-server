package driving

import (
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/repo"
	"4clabs-server/app/4clabs-server/internal/ports"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(
	nftgo.NewService,
	wire.Bind(new(ports.Query), new(*nftgo.Service)),

	wire.Bind(new(ports.Auth), new(*repo.User)),
	repo.NewUser,

	wire.Bind(new(ports.TicketRepo), new(*repo.Ticket)),
	repo.NewTicket,
)

