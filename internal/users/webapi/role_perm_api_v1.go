package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xiaohangshuhub/admin/internal/users/app/roleperm"
	rpd "github.com/xiaohangshuhub/admin/internal/users/domain/roleperm"
	"go.uber.org/zap"
)

func RolePermApiV1EndPoint(router *gin.Engine, log *zap.Logger, app *roleperm.App) {

	group := router.Group("/role")
	{
		group.POST("", RoleCreate(app, log))
		group.PUT("", RoleUpdate(app, log))
		group.DELETE("", RoleDelete(app, log))

	}
	group = router.Group("/perm")
	{
		group.POST("", PermCreate(app, log))
		group.PUT("", PermUpdate(app, log))
		group.DELETE("", PermDelete(app, log))
	}
}

// RoleCreate
// @Summary 创建角色
// @Description 创建角色
// @Tags RolePerm
// @Accept json
// @Produce json
// @Param request body roleperm.RoleCreateCmd true "创建角色参数"
// @Success 200 {object} Response[bool] "创建成功"
// @Failure 400 {object} Response[bool] "请求参数错误"
// @Failure 500 {object} Response[bool] "服务器内部错误"
// @Router /roleperm [post]
func RoleCreate(app *roleperm.App, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmd roleperm.RoleCreateCmd

		if err := c.ShouldBindJSON(&cmd); err != nil {
			log.Warn("参数绑定失败", zap.Error(err))
			c.JSON(http.StatusBadRequest, BadRequest())
			return
		}

		result, err := app.RoleCreateCmdHandler.Handle(c, cmd)

		if err == nil {
			c.JSON(http.StatusOK, Success(result))
			return
		}

		if userErr, ok := err.(*rpd.Error); ok {
			log.Error("创建角色失败", zap.String("role", cmd.Role), zap.Error(err))
			c.JSON(http.StatusInternalServerError, Fail(userErr.Code, userErr.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, InternalServerError())
	}
}

// RoleUpdate
// @Summary 更新角色
// @Description 更新角色
// @Tags RolePerm
// @Accept json
// @Produce json
// @Param request body roleperm.RoleUpdateCmd true "更新角色参数"
// @Success 200 {object} Response[bool] "更新成功"
// @Failure 400 {object} Response[bool] "请求参数错误"
// @Failure 500 {object} Response[bool] "服务器内部错误"
// @Router /roleperm [put]
func RoleUpdate(app *roleperm.App, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var cmd roleperm.RoleUpdateCmd

		if err := c.ShouldBindJSON(&cmd); err != nil {
			log.Warn("参数绑定失败", zap.Error(err))
			c.JSON(http.StatusBadRequest, BadRequest())
			return
		}

		result, err := app.RoleUpdateCmdHandler.Handle(c, cmd)

		if err == nil {
			c.JSON(http.StatusOK, Success(result))
			return
		}

		if userErr, ok := err.(*rpd.Error); ok {
			log.Error("更新角色失败", zap.String("ID", cmd.ID.String()), zap.Error(err))
			c.JSON(http.StatusInternalServerError, Fail(userErr.Code, userErr.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, InternalServerError())
	}
}

// RoleDelete
// @Summary 删除角色
// @Description 删除角色
// @Tags RolePerm
// @Accept json
// @Produce json
// @Param request body roleperm.RoleDeleteCmd true "删除角色参数"
// @Success 200 {object} Response[bool] "删除成功"
// @Failure 400 {object} Response[bool] "请求参数错误"
// @Failure 500 {object} Response[bool] "服务器内部错误"
// @Router /roleperm [delete]
func RoleDelete(app *roleperm.App, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var cmd roleperm.RoleDeleteCmd

		if err := c.ShouldBindJSON(&cmd); err != nil {
			log.Warn("参数绑定失败", zap.Error(err))
			c.JSON(http.StatusBadRequest, BadRequest())
			return
		}

		result, err := app.RoleDeleteCmdHandler.Handle(c, cmd)

		if err == nil {
			c.JSON(http.StatusOK, Success(result))
			return
		}

		if userErr, ok := err.(*rpd.Error); ok {
			log.Error("删除角色失败", zap.String("ID", cmd.ID.String()), zap.Error(err))
			c.JSON(http.StatusInternalServerError, Fail(userErr.Code, userErr.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, InternalServerError())
	}
}

// PermCreate
// @Summary 创建权限
// @Description 创建权限
// @Tags RolePerm
// @Accept json
// @Produce json
// @Param request body roleperm.PermCreateCmd true "创建权限参数"
// @Success 200 {object} Response[bool] "创建成功"
// @Failure 400 {object} Response[bool] "请求参数错误"
// @Failure 500 {object} Response[bool] "服务器内部错误"
// @Router /roleperm [post]
func PermCreate(app *roleperm.App, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var cmd roleperm.PermCreateCmd

		if err := c.ShouldBindJSON(&cmd); err != nil {
			log.Warn("参数绑定失败", zap.Error(err))
			c.JSON(http.StatusBadRequest, BadRequest())
			return
		}

		result, err := app.PermCreateCmdHandler.Handle(c, cmd)

		if err == nil {
			c.JSON(http.StatusOK, Success(result))
			return
		}

		if userErr, ok := err.(*rpd.Error); ok {
			log.Error("创建权限失败", zap.String("Perm", cmd.Perm), zap.Error(err))
			c.JSON(http.StatusInternalServerError, Fail(userErr.Code, userErr.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, InternalServerError())
	}
}

// PermUpdate
// @Summary 更新权限
// @Description 更新权限
// @Tags RolePerm
// @Accept json
// @Produce json
// @Param request body roleperm.PermUpdateCmd true "更新权限参数"
// @Success 200 {object} Response[bool] "更新成功"
// @Failure 400 {object} Response[bool] "请求参数错误"
// @Failure 500 {object} Response[bool] "服务器内部错误"
// @Router /roleperm [put]
func PermUpdate(app *roleperm.App, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var cmd roleperm.PermUpdateCmd

		if err := c.ShouldBindJSON(&cmd); err != nil {
			log.Warn("参数绑定失败", zap.Error(err))
			c.JSON(http.StatusBadRequest, BadRequest())
			return
		}

		result, err := app.PermUpdateCmdHandler.Handle(c, cmd)

		if err == nil {
			c.JSON(http.StatusOK, Success(result))
			return
		}

		if userErr, ok := err.(*rpd.Error); ok {
			log.Error("更新权限失败", zap.String("ID", cmd.ID.String()), zap.Error(err))
			c.JSON(http.StatusInternalServerError, Fail(userErr.Code, userErr.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, InternalServerError())
	}
}

// PermDelete
// @Summary 删除权限
// @Description 删除权限
// @Tags RolePerm
// @Accept json
// @Produce json
// @Param request body roleperm.PermDeleteCmd true "删除权限参数"
// @Success 200 {object} Response[bool] "删除成功"
// @Failure 400 {object} Response[bool] "请求参数错误"
// @Failure 500 {object} Response[bool] "服务器内部错误"
// @Router /roleperm [delete]
func PermDelete(app *roleperm.App, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var cmd roleperm.PermDeleteCmd

		if err := c.ShouldBindJSON(&cmd); err != nil {
			log.Warn("参数绑定失败", zap.Error(err))
			c.JSON(http.StatusBadRequest, BadRequest())
			return
		}

		result, err := app.PermDeleteCmdHandler.Handle(c, cmd)

		if err == nil {
			c.JSON(http.StatusOK, Success(result))
			return
		}

		if userErr, ok := err.(*rpd.Error); ok {
			log.Error("删除权限失败", zap.String("ID", cmd.ID.String()), zap.Error(err))
			c.JSON(http.StatusInternalServerError, Fail(userErr.Code, userErr.Error()))
			return
		}

		c.JSON(http.StatusInternalServerError, InternalServerError())
	}
}
