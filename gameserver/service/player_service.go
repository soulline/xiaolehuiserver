package service

import (
	"fmt"
	"xiaolehuigo/gameserver/database"
	"xiaolehuigo/gameserver/model"
	"xiaolehuigo/gameserver/util"
)

/**
* 添加或更新玩家信息
 */
func AddPlayer(player *cardmodel.Player) (error, bool) {
	fmt.Println("service AddPlayer : ", *player)
	playerInfo := dizhudatabase.PlayerDaoImplDB.GetPlayer(player.UserId)
	if playerInfo.UserId > 0 {
		err, ok := dizhudatabase.PlayerDaoImplDB.UpdatePlayer(PlayerTranslateMysql(*player))
		return err, ok
	} else {
		err, ok := dizhudatabase.PlayerDaoImplDB.AddPlayer(PlayerTranslateMysql(*player))
		return err, ok
	}
}

/**
* 更新玩家
 */
func UpdatePlayer(player *cardmodel.Player) (error, bool) {
	fmt.Println("service UpdatePlayer : ", *player)
	return dizhudatabase.PlayerDaoImplDB.UpdatePlayer(PlayerTranslateMysql(*player))
}

/**
* 查询玩家信息
 */
func GetPlayer(userId int) cardmodel.Player {
	player := dizhudatabase.PlayerDaoImplDB.GetPlayer(userId)
	return PlayerTranslate(player)
}

/**
* 查询玩家信息(批量)
 */
func GetPlayers(userIds []int) []cardmodel.Player {
	mysqlPlayers := dizhudatabase.PlayerDaoImplDB.GetPlayers(userIds)
	var players []cardmodel.Player
	for _, mysqlPlayer := range mysqlPlayers {
		players = append(players, PlayerTranslate(mysqlPlayer))
	}
	return players
}

func GetPlayersByGameId(gameId string) []cardmodel.Player {
	mysqlPlayers := dizhudatabase.PlayerDaoImplDB.GetPlayersByGameId(gameId)
	var players []cardmodel.Player
	for _, mysqlPlayer := range mysqlPlayers {
		players = append(players, PlayerTranslate(mysqlPlayer))
	}
	return players
}

/**
* 删除玩家
 */
func RemovePlayer(userId int) (error, bool) {
	return dizhudatabase.PlayerDaoImplDB.DeletePlayer(userId)
}

/**
* 批量更新玩家信息
 */
func UpdatePlayers(players []cardmodel.Player) {
	var mysqlPlayers []cardmodel.MysqlPlayer
	for _, player := range players {
		mysqlPlayers = append(mysqlPlayers, PlayerTranslateMysql(player))
	}
	dizhudatabase.PlayerDaoImplDB.UpdatePlayers(mysqlPlayers)
}

func PlayerTranslate(mysqlPlayer cardmodel.MysqlPlayer) cardmodel.Player {
	player := cardmodel.Player{
		UserId:     mysqlPlayer.UserId,
		SessionId:  mysqlPlayer.SessionId,
		GameId:     mysqlPlayer.GameId,
		Identify:   mysqlPlayer.Identify,
		NickName:   mysqlPlayer.NickName,
		Money:      mysqlPlayer.Money,
		AliveCards: util.StringToArray(mysqlPlayer.AliveCards),
		IsAway:     mysqlPlayer.IsAway,
		MoneyDiff:  mysqlPlayer.MoneyDiff,
		Status:     mysqlPlayer.Status,
	}
	return player
}

func PlayerTranslateMysql(player cardmodel.Player) cardmodel.MysqlPlayer {
	mysqlPlayer := cardmodel.MysqlPlayer{
		UserId:     player.UserId,
		SessionId:  player.SessionId,
		GameId:     player.GameId,
		Identify:   player.Identify,
		NickName:   player.NickName,
		Money:      player.Money,
		AliveCards: util.ArrayToString(player.AliveCards),
		IsAway:     player.IsAway,
		MoneyDiff:  player.MoneyDiff,
		Status:     player.Status,
	}
	return mysqlPlayer
}
