package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"xiaolehuigo/accountserver/middlewares"
	"xiaolehuigo/accountserver/model"
	"xiaolehuigo/accountserver/model/datamodel"
	"xiaolehuigo/accountserver/service"
	"xiaolehuigo/accountserver/util"
)

var Login = func(context *gin.Context) {
	baseResponse := model.NewBaseResponse()
	mobile := context.PostForm("mobile")
	password := context.PostForm("password")
	if mobile == "" || password == "" {
		baseResponse.GetFailureResponse(model.QUERY_PARAM_ERROR)
	} else {
		user := service.QueryUserByMobile(mobile)
		logrus.Info("login mobile:s%,password:%s", mobile, password)
		if password != "" && password != "null" && password == user.Password {
			baseResponse.GetSuccessResponse()
			baseResponse.Message = "登录成功"
			user := service.QueryUserByMobile(mobile)
			token := datamodel.Token{
				UserId: user.UserId,
			}
			err, ok := service.AddToken(&token)
			if err != nil {
				baseResponse.RespError = err.(model.RespError)
			} else {
				if ok {
					user.Token = token.Token
					baseResponse.Data = util.TurnUserInfoLogin(user)
				} else {
					baseResponse.GetFailureResponse(model.SYSTEM_ERROE)
				}
			}
		} else {
			baseResponse.GetFailureResponse(model.PASSWORD_ERROR)
		}
	}
	context.JSON(http.StatusOK, baseResponse)
}

var Regist = func(context *gin.Context) {
	baseResponse := model.NewBaseResponse()
	var userInfo datamodel.UserInfo
	userInfo.Mobile = context.PostForm("mobile")
	userInfo.Password = context.PostForm("password")
	userInfo.NickName = context.PostForm("nickName")
	captchaId := context.PostForm("captchaId")
	captchaValue := context.PostForm("captchaValue")
	logrus.Info("createUser:", userInfo)
	logrus.Printf("createUser captchaId: %s captchaValue:%s", captchaId, captchaValue)
	if userInfo.Password != "" && userInfo.Mobile != "" &&
		userInfo.NickName != "" && captchaId != "" && captchaValue != "" {
		userCheck := service.QueryUserByMobile(userInfo.Mobile)
		if userCheck.UserId > 0 {
			baseResponse.GetFailureResponse(model.USER_EXIST)
		} else if util.VerifyCaptcha(captchaId, captchaValue) {
			userInfo.Money = 20000
			err, ok := service.CreateUser(userInfo)
			if err != nil {
				baseResponse.RespError = err.(model.RespError)
			} else {
				if ok {
					baseResponse.GetSuccessResponse()
					user := service.QueryUserByMobile(userInfo.Mobile)
					baseResponse.Data = util.TurnUserInfoResp(user)
				} else {
					baseResponse.GetFailureResponse(model.SYSTEM_ERROE)
				}
			}
		} else {
			baseResponse.GetFailureResponse(model.CAPTCHA_ERROR)
		}
	} else {
		baseResponse.GetFailureResponse(model.QUERY_PARAM_ERROR)
	}
	context.JSON(http.StatusOK, baseResponse)
}

var UpdatePassword = func(context *gin.Context) {
	baseResponse := model.NewBaseResponse()
	var userInfo datamodel.UserInfo
	userInfo.Password = context.PostForm("newPassword")
	userInfo.Mobile = context.PostForm("mobile")
	captchaId := context.PostForm("captchaId")
	captchaValue := context.PostForm("captchaValue")
	logrus.Info("UpdatePassword:=", userInfo)
	logrus.Printf("UpdatePassword captchaId: %s captchaValue:%s", captchaId, captchaValue)
	if userInfo.Password != "" && userInfo.Mobile != "" &&
		captchaId != "" && captchaValue != "" {
		userCheck := service.QueryUserByMobile(userInfo.Mobile)
		if userCheck.UserId > 0 {
			if util.VerifyCaptcha(captchaId, captchaValue) {
				err, ok := service.UpdatePassword(userInfo)
				if err != nil {
					baseResponse.RespError = err.(model.RespError)
				} else {
					if ok {
						baseResponse.GetSuccessResponse()
						baseResponse.Message = "修改成功"
					} else {
						baseResponse.GetFailureResponse(model.SYSTEM_ERROE)
					}
				}
			} else {
				baseResponse.GetFailureResponse(model.CAPTCHA_ERROR)
			}
		} else {
			baseResponse.GetFailureResponse(model.USER_NOT_EXIST)
		}
	} else {
		baseResponse.GetFailureResponse(model.QUERY_PARAM_ERROR)
	}
	context.JSON(http.StatusOK, baseResponse)
}

var LoginOut = func(context *gin.Context) {
	baseResponse := model.NewBaseResponse()
	head := context.Request.Header
	token := head.Get(middlewares.ACCESS_TOKEN_KEY)
	logrus.Info("LoginOut:=", token)
	err, ok := service.DeleteTokenByToken(token)
	if err != nil {
		baseResponse.RespError = err.(model.RespError)
	} else {
		if ok {
			baseResponse.GetSuccessResponse()

		} else {
			baseResponse.GetFailureResponse(model.SYSTEM_ERROE)
		}
	}
	context.JSON(http.StatusOK, baseResponse)
}

var GetUserInfo = func(context *gin.Context) {
	head := context.Request.Header
	token := head.Get(middlewares.ACCESS_TOKEN_KEY)
	logrus.Info("GetUserInfo:=", token)
	tokenQ := service.GetTokenByTokenStr(token)
	userId := tokenQ.UserId
	userInfo := service.QueryUserByUserId(userId)
	baseResponse := model.NewBaseResponse()
	if userInfo.UserId <= 0 {
		baseResponse.GetFailureResponse(model.QUERY_NO_DATA)
	} else {
		baseResponse.GetSuccessResponse()
		baseResponse.Data = util.TurnUserInfoResp(userInfo)
	}
	context.JSON(http.StatusOK, baseResponse)
}
