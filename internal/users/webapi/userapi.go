package webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	u "github.com/xiaohangshuhub/admin/internal/users/app/user"
	"github.com/xiaohangshuhub/admin/internal/users/domain/user"
	"go.uber.org/zap"
)

func UserApiV1EndPoint(router *gin.Engine, log *zap.Logger, userapp *u.App) {

	group := router.Group("/user").WithAuthzPolicies("admin")
	{
		group.GET("", Create(userapp, log))
		group.POST("", Create(userapp, log))
		group.PUT("", Create(userapp, log))
		group.DELETE("", Create(userapp, log))
		group.POST("login", Login(userapp, log)).WithAllowAnonymous()
	}

}

// Create godoc
// @Summary 创建用户
// @Description 创建新用户123...
// @Tags User
// @Accept json
// @Produce json
// @Param request body user.CreateCmd true "创建用户参数"
// @Success 200 {object} Response[bool] "创建成功"
// @Failure 400 {object} Response[bool] "请求参数错误"
// @Failure 500 {object} Response[bool] "服务器内部错误"
// @Router /user [post]
func Create(app *u.App, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var cmd u.CreateCmd

		if err := c.ShouldBindJSON(&cmd); err != nil {
			log.Warn("参数绑定失败", zap.Error(err))
			c.JSON(http.StatusBadRequest, BadRequest())
			return
		}

		result, err := app.CreateCmdHandler.Handle(c, cmd)

		if err != nil {

			// 判断是否为可预知的业务错误
			if userErr, ok := err.(user.Error); ok {
				switch err {
				case user.ErrPwdEmpty, user.ErrPhoneAlreadyExist:
					log.Error("创建用户失败", zap.String("loginname", cmd.Username), zap.Error(err))
					c.JSON(http.StatusInternalServerError, Fail(user.ErrPwdEmpty.Code, err.Error()))
				default:
					log.Error("创建用户失败", zap.String("loginname", cmd.Username), zap.Error(err))
					c.JSON(http.StatusInternalServerError, Fail(userErr.Code, userErr.Error()))
				}
				return
			}

			c.JSON(http.StatusInternalServerError, InternalServerError())

			return
		}

		c.JSON(http.StatusOK, Success(result))
	}
}

// Login godoc
// @Summary 用户登录
// @Description 用户登录验证
// @Tags User
// @Accept json
// @Produce json
// @Param request body user.Login true "登录参数"
// @Success 200 {object} Response[user.UserDto] "登录成功"
// @Failure 400 {object} Response[user.UserDto] "请求参数错误"
// @Failure 500 {object} Response[user.UserDto] "服务器内部错误"
// @Router /user/login [post]
func Login(app *u.App, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var query u.Login

		if err := c.ShouldBindJSON(&query); err != nil {
			log.Warn("参数绑定失败", zap.Error(err))
			c.JSON(http.StatusBadRequest, BadRequest())
			return
		}

		result, err := app.LoginHandler.Handle(c, query)

		if err != nil {

			// 判断是否为可预知的业务错误,如果是除密码hash错误外的错误直接返回,hash错误单独处理为密码格式错误
			if userErr, ok := err.(user.Error); ok {
				switch err {
				case user.ErrPasswordInvalid:
					log.Error("登录失败", zap.String("loginname", query.Phone), zap.Error(err))
					c.JSON(http.StatusInternalServerError, Fail(user.ErrPasswordInvalid.Code, err.Error()))
				default:
					log.Error("登录失败", zap.String("loginname", query.Phone), zap.Error(err))
					c.JSON(http.StatusInternalServerError, Fail(userErr.Code, userErr.Error()))
				}
				return
			}

			c.JSON(http.StatusInternalServerError, InternalServerError())

			return
		}

		c.JSON(http.StatusOK, Success(result))
	}
}
