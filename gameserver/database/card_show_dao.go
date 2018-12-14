package dizhudatabase

import (
	"xiaolehuigo/gameserver/model"
)

type CardShowDao interface {
	GetCardShow(showId string) cardmodel.MysqlCardShow
	GetCardShowByGameId(gameId string) []cardmodel.MysqlCardShow
	AddCardShow(cardShow cardmodel.MysqlCardShow) (error, bool)
	UpdateCardShow(cardShow cardmodel.MysqlCardShow) (error, bool)
	DeleteCardShow(showId string) (error, bool)
	DeleteCardShowByGameId(gameId string) (error, bool)
}

var CardShowDaoImplDB *CardShowDaoImpl

func init() {
	CardShowDaoImplDB = new(CardShowDaoImpl)
}

type CardShowDaoImpl struct {
}

/**
* 根据出牌id查询
 */
func (CardShowDaoImpl *CardShowDaoImpl) GetCardShow(showId string) cardmodel.MysqlCardShow {
	var cardQuery cardmodel.MysqlCardShow
	db := DB().Where("show_id = ?", showId).First(&cardQuery)
	affectRow := db.RowsAffected
	if db.Error != nil {
		return cardmodel.MysqlCardShow{}
	}
	if affectRow >= 1 {
		return cardQuery
	}
	return cardmodel.MysqlCardShow{}
}

/**
* 根据游戏id查询
 */
func (CardShowDaoImpl *CardShowDaoImpl) GetCardShowByGameId(gameId string) []cardmodel.MysqlCardShow {
	var cardQuery []cardmodel.MysqlCardShow
	db := DB().Where("game_id = ?", gameId).Find(&cardQuery)
	affectRow := db.RowsAffected
	if db.Error != nil {
		return make([]cardmodel.MysqlCardShow, 0)
	}
	if affectRow >= 1 {
		return cardQuery
	}
	return make([]cardmodel.MysqlCardShow, 0)
}

/**
* 插入出牌记录
 */
func (CardShowDaoImpl *CardShowDaoImpl) AddCardShow(cardShow cardmodel.MysqlCardShow) (error, bool) {
	db := DB().Add(&CardShowDaoImpl)
	affectRow := db.RowsAffected
	if db.Error != nil {
		return db.Error, false
	}
	if affectRow >= 1 {
		return nil, true
	}
	return nil, false
}

/**
* 更新出牌记录
 */
func (CardShowDaoImpl *CardShowDaoImpl) UpdateCardShow(cardShow cardmodel.MysqlCardShow) (error, bool) {
	db := DB().Model(&cardShow).Where("show_id = ?", cardShow.ShowId).Updates(map[string]interface{}{"game_id": cardShow.GameId,
		"show_time": cardShow.ShowTime, "show_value": cardShow.ShowValue, "card_map": cardShow.CardMap, "max_count": cardShow.MaxCount,
		"max_values": cardShow.MaxValues, "compare_value": cardShow.CompareValue, "compare_count": cardShow.CompareCount,
		"card_type_status": cardShow.CardTypeStatus})
	affectRow := db.RowsAffected
	if db.Error != nil {
		return db.Error, false
	}
	if affectRow >= 1 {
		return nil, true
	}
	return nil, false
}

/**
* 根据出牌Id删除记录
 */
func (CardShowDaoImpl *CardShowDaoImpl) DeleteCardShow(showId string) (error, bool) {
	db := DB().Unscoped().Where("show_id LIKE ?", showId).Delete(cardmodel.MysqlCardShow{})
	affectRow := db.RowsAffected
	if db.Error != nil {
		return db.Error, false
	}
	if affectRow >= 1 {
		return nil, true
	}
	return nil, false
}

/**
* 根据出牌Id删除记录
 */
func (CardShowDaoImpl *CardShowDaoImpl) DeleteCardShowByGameId(gameId string) (error, bool) {
	db := DB().Unscoped().Where("game_id LIKE ?", gameId).Delete(cardmodel.MysqlCardShow{})
	affectRow := db.RowsAffected
	if db.Error != nil {
		return db.Error, false
	}
	if affectRow >= 1 {
		return nil, true
	}
	return nil, false
}
