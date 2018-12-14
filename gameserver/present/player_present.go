package present

import (
	"xiaolehuigo/gameserver/gamerpc"
	"xiaolehuigo/gameserver/model"
	"xiaolehuigo/gameserver/service"
)

/**
* 填充玩家信息
 */
func FillPlayerInfo(player *cardmodel.Player) bool {
	playerQuery := gamerpc.GetPlayerInfoRpc(player.UserId)
	if playerQuery.UserId > 0 {
		player.Money = playerQuery.Money
		player.NickName = playerQuery.NickName
		return true
	}
	return false

}

/**
* 更新玩家信息到本地
 */
func UpdateLoaclPlayerInfo(player *cardmodel.Player) bool {
	ok := FillPlayerInfo(player)
	if ok {
		_, ok1 := service.AddPlayer(player)
		return ok1
	}
	return ok
}
