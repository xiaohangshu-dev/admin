package webapi

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaohangshuhub/admin/internal/users/app/roleperm"
	"go.uber.org/zap"
)

func RolePermApiV1EndPoint(router *gin.Engine, log *zap.Logger, userapp *roleperm.App) {

	group := router.Group("/roleperm")
	{
		group.POST("", nil)
	}

}
