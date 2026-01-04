package app

import (
	"github.com/xiaohangshuhub/admin/internal/users/app/user"
	userDomain "github.com/xiaohangshuhub/admin/internal/users/domain/user"
	"go.uber.org/fx"
)

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(userDomain.NewManager),
		fx.Provide(user.NewCreateCmdHandler),
		fx.Provide(user.NewLoginHandler),
		fx.Provide(user.NewUserApp),
	}

}
