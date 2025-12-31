package configs

import "go.uber.org/fx"

func DependencyInjection() []fx.Option {
	return []fx.Option{
		fx.Provide(NewOAuth2),
	}
}
