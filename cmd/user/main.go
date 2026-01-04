package main

import (
	_ "github.com/xiaohangshuhub/admin/api/users/docs" // swagger 一定要有这行,指向你的文档地址
	"github.com/xiaohangshuhub/admin/internal/users/webapi"

	"github.com/xiaohangshuhub/go-workit/pkg/db"
	"github.com/xiaohangshuhub/go-workit/pkg/webapp"
	"github.com/xiaohangshuhub/go-workit/pkg/webapp/dbctx"
)

func main() {

	builder := webapp.NewBuilder()

	builder.AddServices(webapi.DependencyInjection()...)

	builder.AddDbContext(func(options *dbctx.Options) {
		options.UsePostgresSQL("", func(pco *db.PostgresConfigOptions) {
			pco.PgSQLCfg.DSN = "host=172.16.1.105 user=postgres password=postgres dbname=xiaohangshu port=5432 sslmode=disable TimeZone=Asia/Shanghai"
		})
	})

	app := builder.Build()

	if app.Env().IsDevelopment {
		app.UseSwagger()
	}

	app.MapRoute(webapi.UserApiV1EndPoint)

	app.Run()
}
