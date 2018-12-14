package enum

const (
	HEART_HELLO           = 1001 //心跳请求
	HEART_ACK             = 1002 //心跳确认
	PLAYER_JOIN           = 2001 //玩家加入
	PLAYER_EXIT           = 2002 //玩家离开
	PLAYER_READY          = 2003 //玩家准备
	GAME_START            = 3001 //牌局开始
	GAME_OVER             = 3002 //牌局结束
	GAME_DISPATCHER_CARD  = 3003 //发牌
	GAME_APPOINT_IDENTIFY = 3004 //身份分配广播
	GAME_BET              = 3005 //叫地主/抢地主
	GAME_HAND_CARD        = 3006 //底盘分配
	GAME_ORDER            = 3007 //次序广播
	GAME_SHOW_CARD        = 3008 //出牌
	GAME_VERIFY_CARD      = 3009 //出牌校验结果
	GAME_SHOW_RESULT      = 3010 //出牌结果广播
	GAME_PASS             = 3011 //不要
)
