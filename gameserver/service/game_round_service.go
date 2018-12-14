package service

import (
	"xiaolehuigo/gameserver/database"
	"xiaolehuigo/gameserver/model"
	"xiaolehuigo/gameserver/util"
)

/**
* 查询牌局
 */
func GetGameRound(gameId string) cardmodel.GameRound {
	mysqlRound := dizhudatabase.GameRoundDaoImplDB.GetGameRound(gameId)
	return GameRoundTranslate(mysqlRound)
}

/**
* 添加或更新牌局信息
 */
func AddGameRound(round *cardmodel.GameRound) (error, bool) {
	isSaved := CheckRoundIsSaved(*round)
	if isSaved {
		err, ok := dizhudatabase.GameRoundDaoImplDB.UpdateGameRound(GameRoundTranslateMysql(*round))
		return err, ok
	} else {
		err, ok := dizhudatabase.GameRoundDaoImplDB.AddGameRound(GameRoundTranslateMysql(*round))
		return err, ok
	}
}

/**
* 删除牌局Id
 */
func DeleteGameRound(gameId string) (error, bool) {
	return dizhudatabase.GameRoundDaoImplDB.DeleteGameRound(gameId)
}

func CheckRoundIsSaved(round cardmodel.GameRound) bool {
	mysqlRound := dizhudatabase.GameRoundDaoImplDB.GetGameRound(round.GameId)
	if mysqlRound.GameId != "" {
		return true
	} else {
		return false
	}
}

func GameRoundTranslate(mysqlRound cardmodel.MysqlGameRound) cardmodel.GameRound {
	players := GetPlayersByGameId(mysqlRound.GameId)
	cardShow := GetCardShow(mysqlRound.LastShow)
	round := cardmodel.GameRound{
		GameId:           mysqlRound.GameId,
		Gamers:           players,
		Landlord:         mysqlRound.Landlord,
		HandCard:         util.StringToArray(mysqlRound.HandCard),
		CurrentOrder:     mysqlRound.CurrentOrder,
		CurrentOrderUser: mysqlRound.CurrentOrderUser,
		GameStatus:       mysqlRound.GameStatus,
		HandMoney:        mysqlRound.HandMoney,
		LastShow:         cardShow,
		ShowCount:        mysqlRound.ShowCount,
	}
	return round
}

func GameRoundTranslateMysql(round cardmodel.GameRound) cardmodel.MysqlGameRound {
	var playerIds []int
	for _, player := range round.Gamers {
		playerIds = append(playerIds, player.UserId)
	}
	mysqlRound := cardmodel.MysqlGameRound{
		GameId:           round.GameId,
		Gamers:           util.IntArrayToString(playerIds),
		Landlord:         round.Landlord,
		HandCard:         util.ArrayToString(round.HandCard),
		CurrentOrder:     round.CurrentOrder,
		CurrentOrderUser: round.CurrentOrderUser,
		GameStatus:       round.GameStatus,
		HandMoney:        round.HandMoney,
		LastShow:         round.LastShow.ShowId,
		ShowCount:        round.ShowCount,
	}
	return mysqlRound
}
