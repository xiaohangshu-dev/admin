package roleperm

type App struct {
	*PermCreateCmdHandler
	*PermUpdateCmdHandler
	*PermDeleteCmdHandler
	*RoleCreateCmdHandler
	*RoleUpdateCmdHandler
	*RoleDeleteCmdHandler
}

func NewApp(
	permCH *PermCreateCmdHandler, permUH *PermUpdateCmdHandler,
	permDH *PermDeleteCmdHandler, roleCH *RoleCreateCmdHandler,
	roleUH *RoleUpdateCmdHandler, roleDH *RoleDeleteCmdHandler) *App {

	return &App{
		PermCreateCmdHandler: permCH,
		PermUpdateCmdHandler: permUH,
		PermDeleteCmdHandler: permDH,
		RoleCreateCmdHandler: roleCH,
		RoleUpdateCmdHandler: roleUH,
		RoleDeleteCmdHandler: roleDH,
	}
}
