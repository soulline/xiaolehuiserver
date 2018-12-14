package cardmodel

import (
	"time"
	"xiaolehuigo/gameserver/enum"
)

type CardShow struct {
	ShowId         string              //出牌Id
	GameId         string              //牌局Id
	ShowTime       time.Time           //出牌时间
	ShowValue      []string            //牌面数组
	CardMap        map[int]int         //牌面计算结果
	MaxCount       int                 //同值牌出现的最大次数
	MaxValues      []int               //同值牌出现的最大次数列表
	CompareValue   int                 //用于比较大小的值
	CompareCount   int                 //用于比较的连续次数
	CardTypeStatus enum.CardTypeStatus //牌面类型
}
