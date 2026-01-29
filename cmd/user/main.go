package main

import (
	_ "github.com/xiaohangshu-dev/admin/api/users/docs" // swagger 一定要有这行,指向你的文档地址
	"github.com/xiaohangshu-dev/admin/internal/users/webapi"

	"github.com/xiaohangshu-dev/go-workit/pkg/db"
	"github.com/xiaohangshu-dev/go-workit/pkg/webapp"
	"github.com/xiaohangshu-dev/go-workit/pkg/webapp/auth"
	"github.com/xiaohangshu-dev/go-workit/pkg/webapp/auth/scheme/jwt"
	"github.com/xiaohangshu-dev/go-workit/pkg/webapp/authz"
	"github.com/xiaohangshu-dev/go-workit/pkg/webapp/dbctx"
)

func main() {

	builder := webapp.NewBuilder()

	builder.AddServices(webapi.DependencyInjection()...)

	builder.AddDbContext(func(options *dbctx.Options) {
		options.UsePostgresSQL("", func(pco *db.PostgresConfigOptions) {
			pco.PgSQLCfg.DSN = builder.Config().GetString("database.dsn")
		})
	})

	builder.AddAuthentication(func(options *auth.Options) {
		scheme := "oauth2"
		options.DefaultScheme = scheme
		options.AddJwtBearer(scheme, func(jo *jwt.Options) {

		})
	})

	builder.AddAuthorization(func(options *authz.Options) {
		uid_policy := "uid_policy"
		options.DefaultPolicy = uid_policy
		options.RequireHasChaims(uid_policy, "uid")
	})

	app := builder.Build()

	if app.Env().IsDevelopment {
		app.UseSwagger()
	}

	app.UseAuthentication()
	app.UseAuthorization()
	app.MapRoute(webapi.UserApiV1EndPoint, webapi.RolePermApiV1EndPoint)

	app.Run()
}
