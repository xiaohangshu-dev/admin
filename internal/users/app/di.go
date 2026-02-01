package app

import (
	"github.com/xiaohangshu-dev/admin/internal/users/app/roleperm"
	"github.com/xiaohangshu-dev/admin/internal/users/app/user"
	rpd "github.com/xiaohangshu-dev/admin/internal/users/domain/roleperm"
	ud "github.com/xiaohangshu-dev/admin/internal/users/domain/user"
	"go.uber.org/fx"
)

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(ud.NewManager),
		fx.Provide(user.NewCreateCmdHandler),
		fx.Provide(user.NewLoginHandler),
		fx.Provide(user.NewDeleteCmdHandler),
		fx.Provide(user.NewUpdateCmdHandler),
		fx.Provide(user.NewUpdatePwdCmdHandler),
		fx.Provide(user.NewApp),
		fx.Provide(rpd.NewManager),
		fx.Provide(roleperm.NewRoleCreateCmdHandler),
		fx.Provide(roleperm.NewPermCreateCmdHandler),
		fx.Provide(roleperm.NewRoleUpdateCmdHandler),
		fx.Provide(roleperm.NewPermUpdateCmdHandler),
		fx.Provide(roleperm.NewRoleDeleteCmdHandler),
		fx.Provide(roleperm.NewPermDeleteCmdHandler),
		fx.Provide(roleperm.NewApp),
		fx.Provide(user.NewUserQueryHandler),
	}

}
