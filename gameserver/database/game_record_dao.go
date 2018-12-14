package dizhudatabase

import "xiaolehuigo/gameserver/model"

type GameRecordDao interface {
	GetGameRecord(recordId string) cardmodel.MysqlGameRecord
	GetGameRecordByGameId(gameId string) []cardmodel.MysqlGameRecord
	AddGameRecord(gameRecord cardmodel.MysqlGameRecord) (error, bool)
	UpdateGameRecord(gameRecord cardmodel.MysqlGameRecord) (error, bool)
	DeleteGameRecord(gameId string) (error, bool)
}

var GameRecordDaoImplDB *GameRecordDaoImpl

func init() {
	GameRecordDaoImplDB = new(GameRecordDaoImpl)
}

type GameRecordDaoImpl struct {
}

/**
* 根据记录id查询
 */
func (GameRecordDaoImpl *GameRecordDaoImpl) GetGameRecord(recordId string) cardmodel.MysqlGameRecord {
	var recordQuery cardmodel.MysqlGameRecord
	db := DB().Where("record_id = ?", recordId).First(&recordQuery)
	affectRow := db.RowsAffected
	if db.Error != nil {
		return cardmodel.MysqlGameRecord{}
	}
	if affectRow >= 1 {
		return recordQuery
	}
	return cardmodel.MysqlGameRecord{}
}

/**
* 根据出牌id查询
 */
func (GameRecordDaoImpl *GameRecordDaoImpl) GetGameRecordByGameId(gameId string) []cardmodel.MysqlGameRecord {
	var recordQuery []cardmodel.MysqlGameRecord
	db := DB().Where("game_id = ?", gameId).Find(&recordQuery)
	affectRow := db.RowsAffected
	if db.Error != nil {
		return make([]cardmodel.MysqlGameRecord, 0)
	}
	if affectRow >= 1 {
		return recordQuery
	}
	return make([]cardmodel.MysqlGameRecord, 0)
}

/**
* 根据userId查询
 */
func (GameRecordDaoImpl *GameRecordDaoImpl) GetGameRecordByUserId(userId int) []cardmodel.MysqlGameRecord {
	var recordQuery []cardmodel.MysqlGameRecord
	db := DB().Where("player_first = ?", userId).Or("player_second = ?", userId).Or("player_third = ?", userId).Find(&recordQuery)
	affectRow := db.RowsAffected
	if db.Error != nil {
		return make([]cardmodel.MysqlGameRecord, 0)
	}
	if affectRow >= 1 {
		return recordQuery
	}
	return make([]cardmodel.MysqlGameRecord, 0)
}

/**
* 插入牌局记录
 */
func (GameRecordDaoImpl *GameRecordDaoImpl) AddGameRecord(gameRecord cardmodel.MysqlGameRecord) (error, bool) {
	db := DB().Add(&gameRecord)
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
* 更新牌局信息
 */
func (GameRecordDaoImpl *GameRecordDaoImpl) UpdateGameRecord(gameRecord cardmodel.MysqlGameRecord) (error, bool) {
	db := DB().Model(&gameRecord).Where("record_id = ?", gameRecord.GameId).Updates(map[string]interface{}{"game_id": gameRecord.GameId,
		"player_first": gameRecord.PlayerFirst, "player_second": gameRecord.PlayerSecond, "player_third": gameRecord.PlayerThird,
		"landlord": gameRecord.Landlord, "hand_card": gameRecord.HandCard, "current_order": gameRecord.CurrentOrder, "current_order_user": gameRecord.CurrentOrderUser,
		"game_status": gameRecord.GameStatus, "hand_money": gameRecord.HandMoney, "last_show": gameRecord.LastShow,
		"show_count": gameRecord.ShowCount})
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
* 根据记录Id删除记录
 */
func (GameRecordDaoImpl *GameRecordDaoImpl) DeleteGameRecordById(recordId string) (error, bool) {
	db := DB().Unscoped().Where("record_id LIKE ?", recordId).Delete(cardmodel.MysqlGameRecord{})
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
* 根据牌局Id删除记录
 */
func (GameRecordDaoImpl *GameRecordDaoImpl) DeleteGameRecord(gameId string) (error, bool) {
	db := DB().Unscoped().Where("game_id LIKE ?", gameId).Delete(cardmodel.MysqlGameRecord{})
	affectRow := db.RowsAffected
	if db.Error != nil {
		return db.Error, false
	}
	if affectRow >= 1 {
		return nil, true
	}
	return nil, false
}
