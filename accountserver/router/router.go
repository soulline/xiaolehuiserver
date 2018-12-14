package router

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaolehuigo/accountserver/controller"
	"xiaolehuigo/accountserver/middlewares"
)

func ConfigRouter(router *gin.Engine) {
	router.GET("/healthcheck", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"code": 200})
	})
	userGroup := router.Group("/account", middlewares.ParseUserIdMiddleware())
	{
		userGroup.POST("/login", controller.Login)
		userGroup.POST("/create", controller.Regist)
		userGroup.PUT("/modifyPassword", controller.UpdatePassword)
		userGroup.DELETE("/loginOut", controller.LoginOut)
		userGroup.GET("/userInfo", controller.GetUserInfo)
	}
	captchaGroup := router.Group("/captcha")
	{
		captchaGroup.GET("/getCaptcha", controller.GetCaptcha)
		captchaGroup.GET("/verifyCaptcha", controller.VerifyCaptcha)
		captchaGroup.GET("/show/:source", controller.GetCaptchaPng)
	}
}
