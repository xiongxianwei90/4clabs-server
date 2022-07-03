package server

import (
	v1 "4clabs-server/api/service/v1"
	"4clabs-server/app/4clabs-server/internal/adapter/driven/service"
	"4clabs-server/app/4clabs-server/internal/conf"
	"4clabs-server/pkg/auth"
	middleware "4clabs-server/pkg/midwares"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
)

// NewHTTPServer new a HTTP server.
func NewHTTPServer(
	c *conf.Bootstrap,
	s *service.Service,
	logger log.Logger,
	u *auth.ContextUtils,
	authutils *auth.JwtUtils,
) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			middleware.GetAuthMiddle(
				authutils,
				u,
				map[string]struct{}{
					"/v1/nft/comic/create": {},
					"/v1/nft/register":     {},
				},
			),
			recovery.Recovery(),
		),
	}
	if c.Server.Http.Network != "" {
		opts = append(opts, http.Network(c.Server.Http.Network))
	}
	if c.Server.Http.Addr != "" {
		opts = append(opts, http.Address(c.Server.Http.Addr))
	}
	if c.Server.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Server.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	v1.RegisterNftHTTPServer(srv, s)
	return srv
}
