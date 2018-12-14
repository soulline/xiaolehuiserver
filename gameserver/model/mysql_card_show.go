package cardmodel

import "time"

type MysqlCardShow struct {
	MysqlBaseModel
	ShowId         string    `gorm:"primary_key" json:"showId"` //出牌Id
	GameId         string    `gorm:"not null" json:"gameId"`    //对应的牌局Id
	ShowTime       time.Time `json:"showTime"`                  //出牌时间
	ShowValue      string    `json:"showValue"`                 //牌面数组
	CardMap        string    `json:"cardMap"`                   //牌面计算结果
	MaxCount       int       `json:"maxCount"`                  //同值牌出现的最大次数
	MaxValues      string    `json:"maxVlues"`                  //同值牌出现的最大次数列表
	CompareValue   int       `json:"compareValue"`              //用于比较大小的值
	CompareCount   int       `json:"compareCount"`              //用于比较的连续次数
	CardTypeStatus int       `json:"cardTypeStatus"`            //牌面类型
}

func (MysqlCardShow) TableName() string {
	return "card_show_info"
}
