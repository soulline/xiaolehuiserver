package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xiaolehuigo/accountserver/model"
)

func RecoverMiddleWares() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				resp := model.NewBaseResponse()
				resp.Message = fmt.Sprint(err)
				context.JSON(http.StatusInternalServerError, resp)
			}
		}()
		context.Next()
	}
}
