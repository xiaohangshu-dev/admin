package user

type App struct {
	*CreateCmdHandler
	*LoginHandler
}

func NewUserApp(createHandler *CreateCmdHandler, loginHandler *LoginHandler) *App {
	return &App{
		CreateCmdHandler: createHandler,
		LoginHandler:     loginHandler,
	}
}
