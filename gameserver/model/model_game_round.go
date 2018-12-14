package cardmodel

import (
	"fmt"
	"sync"
)

type GameRound struct {
	GameId           string   //牌局Id
	Gamers           []Player //玩家列表
	Landlord         int      //地主userId
	HandCard         []string //底牌
	CurrentOrder     int      //当前次序0,1,2
	CurrentOrderUser int      //当前次序玩家userId
	GameStatus       int      //牌局状态 1:未开始 2:叫抢地主阶段 3:已开始
	HandMoney        int      //底分
	LastShow         CardShow //上一次出的牌
	ShowCount        int      //所有玩家出牌总次数

	GameLock sync.RWMutex //游戏牌局锁
}

func (round *GameRound) AddPlayer(gamer Player) bool {
	round.GameLock.Lock()
	defer round.GameLock.Unlock()
	fmt.Println("AddPlayer : ", len(round.Gamers))
	if len(round.Gamers) >= 3 {
		return false
	} else if len(round.Gamers) >= 0 {
		round.Gamers = append(round.Gamers, gamer)
		return true
	}
	return false
}

func (round *GameRound) RemovePlayer(gamer Player) bool {
	round.GameLock.Lock()
	defer round.GameLock.Unlock()
	fmt.Println("RemovePlayer : ", len(round.Gamers))
	if len(round.Gamers) <= 0 {
		return false
	} else {
		var removeIndex int
		for i := 0; i < len(round.Gamers); i++ {
			if round.Gamers[i].UserId == gamer.UserId {
				removeIndex = i
			}
		}
		round.Gamers = append(round.Gamers[:removeIndex], round.Gamers[removeIndex+1:]...)
		return true
	}
}

func (round *GameRound) NextOrder() int {
	if round.CurrentOrder >= 2 {
		round.CurrentOrder = 0
	} else {
		round.CurrentOrder += 1
	}
	return round.CurrentOrder
}
