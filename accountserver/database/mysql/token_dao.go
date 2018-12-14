package mysql

import (
	"xiaolehuigo/accountserver/model"
	"xiaolehuigo/accountserver/model/datamodel"
	"xiaolehuigo/accountserver/util"
)

type TokenDao interface {
	QueryToken(userid int) (token datamodel.Token)
	AddToken(token *datamodel.Token) (error, bool)
	QueryTokenByTokenStr(tokenStr string) (token datamodel.Token)
	UpdateToken(token *datamodel.Token) (error, bool)
	DeleteToken(userId int) (error, bool)
	DeleteTokenByToken(token string) (error, bool)
}

var TokenDaoImplDB *TokenDaoImpl

func init() {
	TokenDaoImplDB = new(TokenDaoImpl)
}

type TokenDaoImpl struct {
}

func (TokenDaoImplDB *TokenDaoImpl) QueryToken(userId int) (token datamodel.Token) {
	var tokenQuery datamodel.Token
	DB().Where("user_id = ?", userId).First(&tokenQuery)
	return tokenQuery
}

func (TokenDaoImplDB *TokenDaoImpl) QueryTokenByTokenStr(tokenStr string) (token datamodel.Token) {
	var tokenQuery datamodel.Token
	DB().Where("token = ?", tokenStr).First(&tokenQuery)
	return tokenQuery
}

func (TokenDaoImplDB *TokenDaoImpl) AddToken(token *datamodel.Token) (error, bool) {
	db := DB().Add(token)
	affectRow := db.RowsAffected
	if db.Error != nil {
		util.CheckErr(db.Error)
		return model.NewRespError(model.SYSTEM_ERROE), false
	}
	if affectRow >= 1 {
		return nil, true
	}
	return nil, false
}

func (TokenDaoImplDB *TokenDaoImpl) UpdateToken(token *datamodel.Token) (error, bool) {
	db := DB().Model(&token).Where("user_id = ?", token.UserId).Updates(map[string]interface{}{
		"token": token.Token, "login_time": token.LoginTime})
	affectRow := db.RowsAffected
	if db.Error != nil {
		util.CheckErr(db.Error)
		return model.NewRespError(model.SYSTEM_ERROE), false
	}
	if affectRow >= 1 {
		return nil, true
	}
	return nil, false
}

func (TokenDaoImplDB *TokenDaoImpl) DeleteToken(userId int) (error, bool) {
	db := DB().Unscoped().Where("user_id LIKE ?", userId).Delete(datamodel.Token{})
	affectRow := db.RowsAffected
	if db.Error != nil {
		util.CheckErr(db.Error)
		return model.NewRespError(model.SYSTEM_ERROE), false
	}
	if affectRow >= 1 {
		return nil, true
	}
	return nil, false
}

func (TokenDaoImplDB *TokenDaoImpl) DeleteTokenByToken(token string) (error, bool) {
	db := DB().Unscoped().Where("token LIKE ?", token).Delete(datamodel.Token{})
	affectRow := db.RowsAffected
	if db.Error != nil {
		util.CheckErr(db.Error)
		return model.NewRespError(model.SYSTEM_ERROE), false
	}
	if affectRow >= 1 {
		return nil, true
	}
	return nil, false
}
