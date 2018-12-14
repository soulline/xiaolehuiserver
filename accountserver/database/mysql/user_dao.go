package mysql

import (
	"github.com/sirupsen/logrus"
	"xiaolehuigo/accountserver/model"
	"xiaolehuigo/accountserver/model/datamodel"
	"xiaolehuigo/accountserver/util"
)

type UserDao interface {
	CreateUser(user datamodel.UserInfo) (error, bool)
	QueryUserByMobile(mobile string) (user datamodel.UserInfo)
	QueryUserByUserId(userId int) (user datamodel.UserInfo)
	UpdateUser(User datamodel.UserInfo) (error, bool)
	UpdatePassword(User datamodel.UserInfo) (error, bool)
	UpdateMoney(User datamodel.UserInfo) (error, bool)
}

var UserDaoImplDB *UserDaoImpl

func init() {
	UserDaoImplDB = new(UserDaoImpl)
}

type UserDaoImpl struct {
}

func (UserDaoImpl *UserDaoImpl) CreateUser(User datamodel.UserInfo) (error, bool) {
	db := DB().Add(&User)
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

func (UserDaoImpl *UserDaoImpl) QueryUserByMobile(mobile string) (user datamodel.UserInfo) {
	var userQuery datamodel.UserInfo
	db := DB().Where("mobile = ?", mobile).First(&userQuery)
	affectRow := db.RowsAffected
	if db.Error != nil {
		util.CheckErr(db.Error)
		return datamodel.UserInfo{}
	}
	if affectRow >= 1 {
		return userQuery
	}
	return datamodel.UserInfo{}
}

func (UserDaoImpl *UserDaoImpl) QueryUserByUserId(userId int) (user datamodel.UserInfo) {
	var userQuery datamodel.UserInfo
	db := DB().Where("user_id = ?", userId).First(&userQuery)
	affectRow := db.RowsAffected
	if db.Error != nil {
		util.CheckErr(db.Error)
		return datamodel.UserInfo{}
	}
	if affectRow >= 1 {
		return userQuery
	}
	return datamodel.UserInfo{}
}

func (UserDaoImpl *UserDaoImpl) UpdateUser(user datamodel.UserInfo) (error, bool) {
	db := DB().Model(&user).Where("user_id = ?", user.UserId).Updates(map[string]interface{}{
		"mobile": user.Mobile, "password": user.Password, "nick_name": user.Password,
		"money": user.Money, "token": user.Token})
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

func (UserDaoImpl *UserDaoImpl) UpdatePassword(User datamodel.UserInfo) (error, bool) {
	db := DB().Model(&User).Where("mobile = ?", User.Mobile).Update("password", User.Password)
	db.Close()
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

func (UserDaoImpl *UserDaoImpl) UpdateMoney(user datamodel.UserInfo) (error, bool) {
	db := DB().Model(&user).Where("user_id = ?", user.UserId).Update("money", user.Money)
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

/**
* 批量更新(事务)
 */
func (UserDaoImpl *UserDaoImpl) UpdateUsersMoney(users []datamodel.UserInfo) bool {
	isSuccess := true
	tx := DB().Begin()
	logrus.Println("users : ", users)
	for _, user := range users {
		err := tx.Model(&user).Where("user_id = ?", user.UserId).Update("money", user.Money).Error
		if err != nil {
			tx.Rollback()
			logrus.Error(err)
			isSuccess = false
			break
		}
	}
	tx.Commit()
	return isSuccess
}
