package datamodel

type UserResp struct {
	//用户Id
	UserId int `json:"userId"`
	//用户手机号
	Mobile string `json:"mobile"`
	//用户昵称
	NickName string `json:"nickName"`
	//用户积分
	Money int `json:"money"`
}
