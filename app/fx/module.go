package fx

import (
	"github.com/maiaaraujo5/consul-golang/app/server/handler"
	"github.com/maiaaraujo5/consul-golang/consul"
	"github.com/maiaaraujo5/consul-golang/echo"
	"go.uber.org/fx"
)

func Module() fx.Option {
	return fx.Options(
		fx.Provide(
			handler.NewCreateServer,
			consul.Connect,
		),
		fx.Invoke(echo.New),
	)
}
