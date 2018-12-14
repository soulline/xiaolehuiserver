package service

import (
	"xiaolehuigo/gameserver/database"
	"xiaolehuigo/gameserver/model"
	"xiaolehuigo/gameserver/util"
)

/**
* 查询牌局记录
 */
func GetGameRecordByGameId(gameId string) []cardmodel.MysqlGameRecord {
	mysqlRecord := dizhudatabase.GameRecordDaoImplDB.GetGameRecordByGameId(gameId)
	return mysqlRecord
}

/**
* 查询牌局记录
 */
func GetGameRecord(recordId string) cardmodel.MysqlGameRecord {
	mysqlRecord := dizhudatabase.GameRecordDaoImplDB.GetGameRecord(recordId)
	return mysqlRecord
}

/**
* 检索玩家牌局记录
 */
func GetGameRecordByUserId(userId int) []cardmodel.MysqlGameRecord {
	return dizhudatabase.GameRecordDaoImplDB.GetGameRecordByUserId(userId)
}

/**
* 添加或更新牌局信息
 */
func AddGameRecord(record *cardmodel.MysqlGameRecord) (error, bool) {
	isSaved := CheckRecordIsSaved(*record)
	if isSaved {
		err, ok := dizhudatabase.GameRecordDaoImplDB.UpdateGameRecord(*record)
		return err, ok
	} else {
		err, ok := dizhudatabase.GameRecordDaoImplDB.AddGameRecord(*record)
		return err, ok
	}
}

/**
* 插入游戏记录
 */
func AddGameRecordByRound(round cardmodel.GameRound) (error, bool) {
	record := GameRecordTranslateMysql(round)
	return AddGameRecord(&record)
}

/**
* 删除牌局记录
 */
func DeleteGameRecord(gameId string) (error, bool) {
	return dizhudatabase.GameRecordDaoImplDB.DeleteGameRecord(gameId)
}

/**
* 删除牌局记录
 */
func DeleteGameRecordById(recordId string) (error, bool) {
	return dizhudatabase.GameRecordDaoImplDB.DeleteGameRecordById(recordId)
}

func CheckRecordIsSaved(record cardmodel.MysqlGameRecord) bool {
	mysqlRound := dizhudatabase.GameRecordDaoImplDB.GetGameRecord(record.RecordId)
	if mysqlRound.RecordId != "" {
		return true
	} else {
		return false
	}
}

func GameRecordTranslate(mysqlRecord cardmodel.MysqlGameRecord) cardmodel.GameRound {
	players := GetPlayersByGameId(mysqlRecord.GameId)
	cardShow := GetCardShow(mysqlRecord.LastShow)
	round := cardmodel.GameRound{
		GameId:           mysqlRecord.GameId,
		Gamers:           players,
		Landlord:         mysqlRecord.Landlord,
		HandCard:         util.StringToArray(mysqlRecord.HandCard),
		CurrentOrder:     mysqlRecord.CurrentOrder,
		CurrentOrderUser: mysqlRecord.CurrentOrderUser,
		GameStatus:       mysqlRecord.GameStatus,
		HandMoney:        mysqlRecord.HandMoney,
		LastShow:         cardShow,
		ShowCount:        mysqlRecord.ShowCount,
	}
	return round
}

func GameRecordTranslateMysql(round cardmodel.GameRound) cardmodel.MysqlGameRecord {
	mysqlRound := cardmodel.MysqlGameRecord{
		RecordId:         "record_" + util.GetUUID(),
		GameId:           round.GameId,
		Landlord:         round.Landlord,
		HandCard:         util.ArrayToString(round.HandCard),
		CurrentOrder:     round.CurrentOrder,
		CurrentOrderUser: round.CurrentOrderUser,
		GameStatus:       round.GameStatus,
		HandMoney:        round.HandMoney,
		LastShow:         round.LastShow.ShowId,
		ShowCount:        round.ShowCount,
	}
	for i := 0; i < len(round.Gamers); i++ {
		if i == 0 {
			mysqlRound.PlayerFirst = round.Gamers[i].UserId
		} else if i == 1 {
			mysqlRound.PlayerSecond = round.Gamers[i].UserId
		} else if i == 2 {
			mysqlRound.PlayerThird = round.Gamers[i].UserId
		}
	}
	return mysqlRound
}
