package app

import (
	"github.com/xiaohangshuhub/admin/internal/users/app/roleperm"
	"github.com/xiaohangshuhub/admin/internal/users/app/user"
	rpd "github.com/xiaohangshuhub/admin/internal/users/domain/roleperm"
	ud "github.com/xiaohangshuhub/admin/internal/users/domain/user"
	"go.uber.org/fx"
)

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(ud.NewManager),
		fx.Provide(user.NewCreateCmdHandler),
		fx.Provide(user.NewLoginHandler),
		fx.Provide(user.NewApp),
		fx.Provide(rpd.NewManager),
		fx.Provide(roleperm.NewApp),
	}

}
