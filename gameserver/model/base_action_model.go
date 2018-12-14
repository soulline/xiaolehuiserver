package cardmodel

type BaseModel struct {
	ActionCode int         `json:"actionCode"`     //指令类型
	UserId     int         `json:"userId"`         //用户Id
	Data       interface{} `json:"data,omitempty"` //指令详情

}

func GetBaseModeByCode(actionCode int) BaseModel {
	return BaseModel{
		ActionCode: actionCode,
	}
}

func GetBaseMode(actionCode int, userId int) BaseModel {
	return BaseModel{
		ActionCode: actionCode,
		UserId:     userId,
	}
}
