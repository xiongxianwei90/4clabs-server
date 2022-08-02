package driving

import (
	"4clabs-server/app/4clabs-server/internal/adapter/driving/nftgo"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/repo"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/repo/nft"
	"4clabs-server/app/4clabs-server/internal/adapter/driving/repo/trade"
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

	wire.Bind(new(ports.Register), new(*repo.Register)),
	repo.NewRegister,

	wire.Bind(new(ports.Comic), new(*repo.Comic)),
	repo.NewComic,

	wire.Bind(new(ports.Trade), new(*trade.Trade)),
	trade.NewTrade,

	wire.Bind(new(ports.Script), new(*repo.Script)),
	repo.NewScript,

	nft.NewNft,
)
