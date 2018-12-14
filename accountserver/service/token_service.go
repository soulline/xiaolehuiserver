package service

import (
	"fmt"
	"xiaolehuigo/accountserver/database/mysql"
	"xiaolehuigo/accountserver/model/datamodel"
	"xiaolehuigo/accountserver/tokenutil"
	"xiaolehuigo/accountserver/util"
)

func AddToken(tokenInfo *datamodel.Token) (error, bool) {
	tokenQ := mysql.TokenDaoImplDB.QueryToken(tokenInfo.UserId)
	tokenInfo.LoginTime = util.GetNowTime()
	tokenString, err := tokenutil.NewToken()
	if err != nil {
		return err, false
	}
	fmt.Println("tokenString : ", tokenString)
	tokenInfo.Token = tokenString
	fmt.Println("tokenQs%", tokenQ)
	if tokenQ.Token != "" {
		err, ok := mysql.TokenDaoImplDB.UpdateToken(tokenInfo)
		return err, ok
	} else {
		err, ok := mysql.TokenDaoImplDB.AddToken(tokenInfo)
		return err, ok
	}
}

func GetTokenByUserId(userId int) (token datamodel.Token) {
	return mysql.TokenDaoImplDB.QueryToken(userId)
}

func GetTokenByTokenStr(tokenStr string) (token datamodel.Token) {
	return mysql.TokenDaoImplDB.QueryTokenByTokenStr(tokenStr)
}

func DeleteToken(userId int) (error, bool) {
	return mysql.TokenDaoImplDB.DeleteToken(userId)
}

func DeleteTokenByToken(token string) (error, bool) {
	return mysql.TokenDaoImplDB.DeleteTokenByToken(token)
}

func CheckLoginStatus(userId int) bool {
	token := GetTokenByUserId(userId)
	if token.Token != "" {
		return true
	}
	return false
}
