package main

import (
	_ "github.com/xiaohangshuhub/admin/api/users/docs" // swagger 一定要有这行,指向你的文档地址
	"github.com/xiaohangshuhub/admin/internal/users/webapi"

	"github.com/xiaohangshuhub/go-workit/pkg/webapp"
)

func main() {

	builder := webapp.NewBuilder()

	app := builder.Build()

	if app.Env().IsDevelopment {
		app.UseSwagger()
	}

	app.MapRoute(webapi.UserApiV1EndPoint)

	app.Run()
}
