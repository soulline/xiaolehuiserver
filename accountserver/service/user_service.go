package service

import (
	"xiaolehuigo/accountserver/database/mysql"
	"xiaolehuigo/accountserver/model/datamodel"
)

//创建账号
func CreateUser(UserInfo datamodel.UserInfo) (error, bool) {
	err, ok := mysql.UserDaoImplDB.CreateUser(UserInfo)
	return err, ok
}

/**
根据手机号查账号
*/
func QueryUserByMobile(mobile string) (User datamodel.UserInfo) {
	return mysql.UserDaoImplDB.QueryUserByMobile(mobile)
}

/**
根据userId查手机号
*/
func QueryUserByUserId(userId int) (User datamodel.UserInfo) {
	return mysql.UserDaoImplDB.QueryUserByUserId(userId)
}

/**
修改密码
*/
func UpdatePassword(userInfo datamodel.UserInfo) (error, bool) {
	return mysql.UserDaoImplDB.UpdatePassword(userInfo)
}

/**
* 更新用户信息
 */
func UpdateUser(userInfo datamodel.UserInfo) (error, bool) {
	return mysql.UserDaoImplDB.UpdateUser(userInfo)
}

/**
* 更新金币
 */
func UpdateMoney(userInfo datamodel.UserInfo) (error, bool) {
	return mysql.UserDaoImplDB.UpdateMoney(userInfo)
}

/**
* 批量更新金币
 */
func UpdateUsersMoney(users []datamodel.UserInfo) bool {
	return mysql.UserDaoImplDB.UpdateUsersMoney(users)
}
