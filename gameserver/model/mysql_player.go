package cardmodel

type MysqlPlayer struct {
	MysqlBaseModel
	UserId     int    `gorm:"primary_key" json:"userId"` //玩家id
	SessionId  string `json:"sessionId"`                 //sessionId
	GameId     string `json:"gameId"`                    //玩家所在的牌局Id
	Identify   int    `json:"identify"`                  //身份 0: 未分配 1: 贫民 2:地主
	NickName   string `json:"nickName"`                  //昵称
	Money      int    `json:"money"`                     //金币
	AliveCards string `json:"aliveCards"`                //剩余牌面
	IsAway     bool   `json:"isAway"`                    //是否逃跑
	MoneyDiff  int    `json:"moneyDiff"`                 //金币变化
	Status     int    `json:"status"`                    //玩家状态  0：未准备   1：已准备
}

func (MysqlPlayer) TableName() string {
	return "game_player_info"
}
