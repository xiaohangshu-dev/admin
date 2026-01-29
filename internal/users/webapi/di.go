package webapi

import (
	"github.com/xiaohangshu-dev/admin/internal/users/app"
	"go.uber.org/fx"
)

func DependencyInjection() []fx.Option {
	di := []fx.Option{}

	di = append(di, app.DependencyInjection()...)

	return di
}
