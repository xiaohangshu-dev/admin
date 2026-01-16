package user

type App struct {
	*CreateCmdHandler
	*LoginHandler
	*DeleteCmdHandler
	*UpdateCmdHandler
	*UpdatePwdCmdHandler
	*UserQueryHandler
}

func NewApp(
	createHandler *CreateCmdHandler, loginHandler *LoginHandler,
	deleteHandler *DeleteCmdHandler, updateHandler *UpdateCmdHandler,
	updatePwdHandler *UpdatePwdCmdHandler, queryHandler *UserQueryHandler) *App {

	return &App{
		CreateCmdHandler:    createHandler,
		LoginHandler:        loginHandler,
		DeleteCmdHandler:    deleteHandler,
		UpdateCmdHandler:    updateHandler,
		UpdatePwdCmdHandler: updatePwdHandler,
		UserQueryHandler:    queryHandler,
	}
}
