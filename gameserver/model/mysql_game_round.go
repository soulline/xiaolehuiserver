package cardmodel

type MysqlGameRound struct {
	MysqlBaseModel
	GameId           string `gorm:"primary_key" json:"gameId"` //牌局Id
	Gamers           string `json:"gamers"`                    //玩家Id列表
	Landlord         int    `json:"landlord"`                  //地主userId
	HandCard         string `json:"handCard"`                  //底牌
	CurrentOrder     int    `json:"currentOrder"`              //当前次序0,1,2
	CurrentOrderUser int    `json:"currentOrderUser"`          //当前次序玩家userId
	GameStatus       int    `json:"gameStatus"`                //牌局状态 1:未开始 2:叫抢地主阶段 3:已开始
	HandMoney        int    `json:"handMoney"`                 //底分
	LastShow         string `json:"lastShow"`                  //上一次出的牌Id
	ShowCount        int    `json:"showCount"`                 //所有玩家出牌总次数
}

func (MysqlGameRound) TableName() string {
	return "game_round_info"
}
