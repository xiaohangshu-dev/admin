package user

type App struct {
	*CreateCmdHandler
	*LoginHandler
	*DeleteCmdHandler
	*UpdateCmdHandler
	*UpdatePwdCmdHandler
}

func NewApp(
	createHandler *CreateCmdHandler, loginHandler *LoginHandler,
	deleteHandler *DeleteCmdHandler, updateHandler *UpdateCmdHandler,
	updatePwdHandler *UpdatePwdCmdHandler) *App {

	return &App{
		CreateCmdHandler:    createHandler,
		LoginHandler:        loginHandler,
		DeleteCmdHandler:    deleteHandler,
		UpdateCmdHandler:    updateHandler,
		UpdatePwdCmdHandler: updatePwdHandler,
	}
}
