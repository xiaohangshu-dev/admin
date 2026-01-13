package app

import (
	"github.com/xiaohangshuhub/admin/internal/users/app/perm"
	"github.com/xiaohangshuhub/admin/internal/users/app/user"
	permDomin "github.com/xiaohangshuhub/admin/internal/users/domain/perm"
	userDomin "github.com/xiaohangshuhub/admin/internal/users/domain/user"
	"go.uber.org/fx"
)

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(userDomin.NewManager),
		fx.Provide(user.NewCreateCmdHandler),
		fx.Provide(user.NewLoginHandler),
		fx.Provide(user.NewApp),
		fx.Provide(permDomin.NewManager),
		fx.Provide(perm.NewApp),
	}

}
