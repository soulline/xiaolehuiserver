package service

import (
	"xiaolehuigo/gameserver/database"
	"xiaolehuigo/gameserver/enum"
	"xiaolehuigo/gameserver/model"
	"xiaolehuigo/gameserver/util"
)

/**
* 添加或更新出牌记录
 */
func AddCardShow(cardShow *cardmodel.CardShow) (error, bool) {
	queryCardShow := GetCardShow(cardShow.ShowId)
	if queryCardShow.ShowId != "" && queryCardShow.GameId != "" {
		err, ok := dizhudatabase.CardShowDaoImplDB.UpdateCardShow(CardShowTranslateMysql(*cardShow))
		return err, ok
	} else {
		err, ok := dizhudatabase.CardShowDaoImplDB.AddCardShow(CardShowTranslateMysql(*cardShow))
		return err, ok
	}
}

/**
* 查询出牌记录根据出牌id
 */
func GetCardShow(showId string) cardmodel.CardShow {
	mysqlCardShow := dizhudatabase.CardShowDaoImplDB.GetCardShow(showId)
	return CardShowTranslate(mysqlCardShow)
}

/**
* 查询出牌记录根据牌局id
 */
func GetCardShowByGameId(gameId string) []cardmodel.CardShow {
	mysqlCardShows := dizhudatabase.CardShowDaoImplDB.GetCardShowByGameId(gameId)
	var shows []cardmodel.CardShow
	if len(mysqlCardShows) > 0 {
		for _, show := range mysqlCardShows {
			shows = append(shows, CardShowTranslate(show))
		}
	}
	return shows
}

/**
* 根据出牌Id删除出牌记录
 */
func DeleteCardShow(showId string) (error, bool) {
	return dizhudatabase.CardShowDaoImplDB.DeleteCardShow(showId)
}

/**
* 根据牌局Id删除出牌记录
 */
func DelteCardShowbyGameId(gameId string) (error, bool) {
	return dizhudatabase.CardShowDaoImplDB.DeleteCardShowByGameId(gameId)
}

/**
* 类型互转
 */
func CardShowTranslate(mysqlcardShow cardmodel.MysqlCardShow) cardmodel.CardShow {
	cardShow := cardmodel.CardShow{
		ShowId:         mysqlcardShow.ShowId,
		ShowTime:       mysqlcardShow.ShowTime,
		ShowValue:      util.StringToArray(mysqlcardShow.ShowValue),
		CardMap:        util.StringToMap(mysqlcardShow.CardMap),
		MaxCount:       mysqlcardShow.MaxCount,
		MaxValues:      util.StringToIntArray(mysqlcardShow.MaxValues),
		CompareValue:   mysqlcardShow.CompareValue,
		CompareCount:   mysqlcardShow.CompareCount,
		CardTypeStatus: enum.GetStatus(mysqlcardShow.CardTypeStatus),
	}
	return cardShow
}

/**
* 类型互转
 */
func CardShowTranslateMysql(cardShow cardmodel.CardShow) cardmodel.MysqlCardShow {
	mysqlCardShow := cardmodel.MysqlCardShow{
		ShowId:         cardShow.ShowId,
		ShowTime:       cardShow.ShowTime,
		ShowValue:      util.ArrayToString(cardShow.ShowValue),
		CardMap:        util.MapToString(cardShow.CardMap),
		MaxCount:       cardShow.MaxCount,
		MaxValues:      util.IntArrayToString(cardShow.MaxValues),
		CompareValue:   cardShow.CompareValue,
		CompareCount:   cardShow.CompareCount,
		CardTypeStatus: enum.GetIntValue(cardShow.CardTypeStatus),
	}
	return mysqlCardShow
}
