package datamodel

type UserInfo struct {
	MysqlBaseModel
	//用户Id
	UserId int `gorm:"primary_key" json:"userId"`
	//用户手机号
	Mobile string `gorm:"not null" json:"mobile"`
	//用户密码
	Password string `gorm:"not null" json:"password"`
	//用户昵称
	NickName string `gorm:"size:255" json:"nickName"`
	//用户积分
	Money int `json:"money"`

	Token string `gorm:"-" json:"token"`
}

func (UserInfo) TableName() string {
	return "user_info"
}
