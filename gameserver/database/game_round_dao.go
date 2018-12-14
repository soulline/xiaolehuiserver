package dizhudatabase

import "xiaolehuigo/gameserver/model"

type GameRoundDao interface {
	GetGameRound(gameId string) cardmodel.MysqlGameRound
	AddGameRound(gameRound cardmodel.MysqlGameRound) (error, bool)
	UpdateGameRound(gameRound cardmodel.MysqlGameRound) (error, bool)
	DeleteGameRound(gameId string) (error, bool)
}

var GameRoundDaoImplDB *GameRoundDaoImpl

func init() {
	GameRoundDaoImplDB = new(GameRoundDaoImpl)
}

type GameRoundDaoImpl struct {
}

/**
* 根据出牌id查询
 */
func (GameRoundDaoImpl *GameRoundDaoImpl) GetGameRound(gameId string) cardmodel.MysqlGameRound {
	var roundQuery cardmodel.MysqlGameRound
	db := DB().Where("game_id = ?", gameId).First(&roundQuery)
	affectRow := db.RowsAffected
	if db.Error != nil {
		return cardmodel.MysqlGameRound{}
	}
	if affectRow >= 1 {
		return roundQuery
	}
	return cardmodel.MysqlGameRound{}
}

/**
* 插入牌局记录
 */
func (GameRoundDaoImpl *GameRoundDaoImpl) AddGameRound(gameRound cardmodel.MysqlGameRound) (error, bool) {
	db := DB().Add(&gameRound)
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
func (GameRoundDaoImpl *GameRoundDaoImpl) UpdateGameRound(gameRound cardmodel.MysqlGameRound) (error, bool) {
	db := DB().Model(&gameRound).Where("game_id = ?", gameRound.GameId).Updates(map[string]interface{}{
		"gamers": gameRound.Gamers, "landlord": gameRound.Landlord, "hand_card": gameRound.HandCard,
		"current_order": gameRound.CurrentOrder, "current_order_user": gameRound.CurrentOrderUser,
		"game_status": gameRound.GameStatus, "hand_money": gameRound.HandMoney, "last_show": gameRound.LastShow,
		"show_count": gameRound.ShowCount})
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
func (GameRoundDaoImpl *GameRoundDaoImpl) DeleteGameRound(gameId string) (error, bool) {
	db := DB().Unscoped().Where("game_id LIKE ?", gameId).Delete(cardmodel.MysqlGameRound{})
	affectRow := db.RowsAffected
	if db.Error != nil {
		return db.Error, false
	}
	if affectRow >= 1 {
		return nil, true
	}
	return nil, false
}
