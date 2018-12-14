package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaolehuigo/accountserver/model"
	"xiaolehuigo/accountserver/service"
)

const (
	ACCESS_TOKEN_KEY = "access_token"
)

/**
解析头部uid信息，若未解析成功，直接返回401
*/
func ParseUserIdMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		path := context.Request.URL.Path
		fmt.Println(path)
		if path != "/account/login" && path != "/account/create" && path != "/account/modifyPassword" {
			head := context.Request.Header
			token := head.Get(ACCESS_TOKEN_KEY)
			if token == "" {
				response := model.NewFailureResponse(model.FORBIDDEN)
				context.AbortWithStatusJSON(http.StatusUnauthorized, response)
				return
			} else {
				tokenQ := service.GetTokenByTokenStr(token)
				if tokenQ.UserId <= 0 {
					response := model.NewFailureResponse(model.FORBIDDEN)
					context.AbortWithStatusJSON(http.StatusUnauthorized, response)
					return
				}
			}
		}
		context.Next()
	}
}
