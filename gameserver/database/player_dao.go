package dizhudatabase

import (
	"github.com/sirupsen/logrus"
	"xiaolehuigo/gameserver/model"
)

type PlayerDao interface {
	GetPlayer(userId int) cardmodel.MysqlPlayer
	GetPlayers(userId ...int) []cardmodel.MysqlPlayer
	GetPlayersByGameId(gameId ...int) []cardmodel.MysqlPlayer
	AddPlayer(player cardmodel.MysqlPlayer) (error, bool)
	UpdatePlayer(player cardmodel.MysqlPlayer) (error, bool)
	UpdatePlayers(players []cardmodel.MysqlPlayer)
	DeletePlayer(userId int) (error, bool)
}

var PlayerDaoImplDB *PlayerDaoImpl

func init() {
	PlayerDaoImplDB = new(PlayerDaoImpl)
}

type PlayerDaoImpl struct {
}

/**
* 查询玩家
 */
func (PlayerDaoImpl *PlayerDaoImpl) GetPlayer(userId int) cardmodel.MysqlPlayer {
	var playerQuery cardmodel.MysqlPlayer
	db := DB().Where("user_id = ?", userId).First(&playerQuery)
	affectRow := db.RowsAffected
	if db.Error != nil {
		return cardmodel.MysqlPlayer{}
	}
	if affectRow >= 1 {
		return playerQuery
	}
	return cardmodel.MysqlPlayer{}
}

/**
* 查询玩家
 */
func (PlayerDaoImpl *PlayerDaoImpl) GetPlayers(userIds []int) []cardmodel.MysqlPlayer {
	var playersQuery []cardmodel.MysqlPlayer
	db := DB()
	for i := 0; i < len(userIds); i++ {
		if i == 0 {
			db.Where("user_id = ?", userIds[i])
		} else {
			db.Or(cardmodel.Player{UserId: userIds[i]})
		}
	}
	db.Find(&playersQuery)
	affectRow := db.RowsAffected
	if db.Error != nil {
		return make([]cardmodel.MysqlPlayer, 0)
	}
	if affectRow >= 1 {
		return playersQuery
	}
	return make([]cardmodel.MysqlPlayer, 0)
}

func (PlayerDaoImpl *PlayerDaoImpl) GetPlayersByGameId(gameId string) []cardmodel.MysqlPlayer {
	var playerQuery []cardmodel.MysqlPlayer
	db := DB().Where("game_id = ?", gameId).Find(&playerQuery)
	affectRow := db.RowsAffected
	if db.Error != nil {
		return make([]cardmodel.MysqlPlayer, 0)
	}
	if affectRow >= 1 {
		return playerQuery
	}
	return make([]cardmodel.MysqlPlayer, 0)
}

/**
* 添加玩家
 */
func (PlayerDaoImpl *PlayerDaoImpl) AddPlayer(player cardmodel.MysqlPlayer) (error, bool) {
	db := DB().Add(&player)
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
* 更新玩家信息
 */
func (PlayerDaoImpl *PlayerDaoImpl) UpdatePlayer(player cardmodel.MysqlPlayer) (error, bool) {
	db := DB().Model(&player).Where("user_id = ?", player.UserId).Updates(map[string]interface{}{
		"session_id": player.SessionId, "game_id": player.GameId, "identify": player.Identify,
		"nick_name": player.NickName, "money": player.NickName, "alive_cards": player.AliveCards,
		"is_away": player.IsAway, "money_diff": player.MoneyDiff, "status": player.Status})
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
* 批量更新(事务)
 */
func (PlayerDaoImpl *PlayerDaoImpl) UpdatePlayers(players []cardmodel.MysqlPlayer) bool {
	isSuccess := true
	tx := DB().Begin()
	for _, player := range players {
		err := tx.Model(&player).Where("user_id = ?", player.UserId).Updates(map[string]interface{}{
			"session_id": player.SessionId, "game_id": player.GameId, "identify": player.Identify,
			"nick_name": player.NickName, "money": player.NickName, "alive_cards": player.AliveCards,
			"is_away": player.IsAway, "money_diff": player.MoneyDiff, "status": player.Status}).Error
		if err != nil {
			tx.Rollback()
			logrus.Error(err)
			isSuccess = false
			break
		}
	}
	tx.Commit()
	return isSuccess
}

/**
* 删除玩家
 */
func (PlayerDaoImpl *PlayerDaoImpl) DeletePlayer(userId int) (error, bool) {
	db := DB().Unscoped().Where("user_id LIKE ?", userId).Delete(cardmodel.Player{})
	affectRow := db.RowsAffected
	if db.Error != nil {
		return db.Error, false
	}
	if affectRow >= 1 {
		return nil, true
	}
	return nil, false
}
