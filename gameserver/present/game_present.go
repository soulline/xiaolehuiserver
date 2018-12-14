package present

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/sirupsen/logrus"
	"net"
	"sync"
	"xiaolehuigo/gameserver/computer"
	"xiaolehuigo/gameserver/enum"
	"xiaolehuigo/gameserver/gamerpc"
	"xiaolehuigo/gameserver/model"
	"xiaolehuigo/gameserver/service"
	"xiaolehuigo/gameserver/util"
	"xiaolehuigo/gameserver/xlhsocket"
)

var ser = xlhsocket.NewMsf(&Event{})

var waitIndexList []string //等待玩家加入牌局列表

var userMap map[int]string

var joinLock sync.RWMutex

type Event struct {
}

func Listen() {
	ser.Listen(":8888")
}

//客户端握手成功事件
func (this Event) OnHandel(sessionId string, conn net.Conn) bool {
	logrus.Println(sessionId, "OnHandel")
	return true
}

//断开连接事件
func (this Event) OnClose(sessionId string) {
	logrus.Println(sessionId, "OnClose")
}

//接收到消息事件
func (this Event) OnMessage(sessionId string, msg string) bool {
	logrus.Println("onMessage", msg)
	actionCode, err1 := jsonparser.GetInt([]byte(msg), "actionCode")
	userId, err2 := jsonparser.GetInt([]byte(msg), "userId")
	logrus.Printf("OnMessage  err1 : v%   err2 : v%", err1, err2)
	if err1 == nil && err2 == nil {
		parseActionCode(int(actionCode), int(userId), sessionId, msg)
	}
	return true
}

func init() {
	userMap = make(map[int]string)
}

/**
* 新玩家加入
 */
func NewPlayerJoin(player cardmodel.Player) {
	joinLock.Lock()
	defer joinLock.Unlock()
	if len(waitIndexList) > 0 {
		joinWait(player)
	} else {
		CreateNewGame(player)
	}
}

func joinWait(player cardmodel.Player) {
	if _, ok := userMap[player.UserId]; ok {
		logrus.Printf("玩家%d已加入过,无需重新加入", player.UserId)
		return
	}
	roundId := waitIndexList[0]
	currentGame := service.GetGameRound(roundId)
	player.GameId = currentGame.GameId
	ok := UpdateLoaclPlayerInfo(&player)
	if ok {
		currentGame.Gamers = append(currentGame.Gamers, player)
		userMap[player.UserId] = currentGame.GameId
		service.AddGameRound(&currentGame)
		logrus.Println("joinWait length : ", len(currentGame.Gamers))
		if len(currentGame.Gamers) == 3 { //加入之后若满员，则删除等待队列中的此等待局数据
			waitIndexList = append(waitIndexList[:0], waitIndexList[1:]...) //删除第一条
		}
	} else {
		logrus.Info("加入牌局失败")
	}

}

/**
* 牌局创建
 */
func CreateNewGame(player cardmodel.Player) {
	gameId := "round_" + util.GetUUID()
	round := cardmodel.GameRound{
		GameId: gameId,
	}
	player.GameId = round.GameId
	ok := UpdateLoaclPlayerInfo(&player)
	if ok {
		round.AddPlayer(player)
		userMap[player.UserId] = round.GameId
		service.AddGameRound(&round)
		waitIndexList = append(waitIndexList, round.GameId)
	} else {
		logrus.Info("创建牌局失败")
		//创建牌局失败
	}
}

/**
* 开始游戏
 */
func StartGame(round *cardmodel.GameRound) {
	cards := card.CreateNew()
	card.Shuffle(cards)
	for i := 0; i < 3; i++ {
		round.Gamers[i].AliveCards = card.Dispatcher(i, cards) //发牌给玩家
		sendHandCardToPlayer(round.Gamers[i].UserId, round.Gamers[i].SessionId, enum.GAME_DISPATCHER_CARD, round.Gamers[i].AliveCards)
	}
	round.HandCard = card.Dispatcher(3, cards) //底牌
	data1, _ := json.Marshal(round)
	logrus.Println("current", string(data1))
	round.CurrentOrder = util.Random(3)
	round.CurrentOrderUser = round.Gamers[round.CurrentOrder].UserId
	round.GameStatus = 1
	service.AddGameRound(round)
	sendOrderToUser(round.CurrentOrderUser, round.Gamers[round.CurrentOrder].SessionId)
}

/**
* 重置牌局
 */
func ResetGame(round *cardmodel.GameRound) {
	for _, gamer := range round.Gamers {
		gamer.AliveCards = []string{}
		gamer.IsAway = false
		gamer.MoneyDiff = 0
		gamer.Status = 0
	}
	round.HandCard = []string{}
	round.CurrentOrder = 0
	if len(round.Gamers) > 0 {
		service.UpdatePlayers(round.Gamers)
		round.CurrentOrderUser = round.Gamers[0].UserId
	} else {
		round.CurrentOrderUser = 0
	}
	round.Landlord = 0
	round.GameStatus = 0
	round.HandMoney = 0
	round.LastShow = cardmodel.CardShow{}
	round.ShowCount = 0
	service.AddGameRound(round)
}

/**
* 地主分配
 */
func HandDispatcher(userId int, round *cardmodel.GameRound) {
	round.Landlord = userId
	round.GameStatus = 2
	for i := 0; i < len(round.Gamers); i++ {
		if round.Gamers[i].UserId == userId {
			round.CurrentOrder = i
			round.CurrentOrderUser = userId
			for _, card := range round.HandCard {
				round.Gamers[i].AliveCards = append(round.Gamers[i].AliveCards, card)
			}
			sendHandCardToPlayer(userId, round.Gamers[i].SessionId, enum.GAME_HAND_CARD, round.HandCard) //发送底牌给玩家
			break
		}
	}
	service.AddGameRound(round)
}

func GetSession(sessionId string) *xlhsocket.Session {
	return ser.SessionMaster.GetSessionById(sessionId)
}

func sendMsg(userId int, sessionId string, msg string) {
	logrus.Printf("sendMsg msg : %v  sessionId : %v", msg, sessionId)
	session := GetSession(sessionId)
	sendMessage := msg + "\n" //客户端读取是按行读取，每条消息换行符隔开
	if session == nil {       //session为空说明已掉线
		ExitPlayer(userId)
	} else {
		session.Write(sendMessage) //发送消息
	}
}

func GetRoundByUserId(userId int) cardmodel.GameRound {

	roundId := userMap[userId]
	round := service.GetGameRound(roundId)
	return round
}

/**
* 心跳确认
 */
func ackHeart(userId int, sessionId string) {
	ack := cardmodel.GetBaseMode(enum.HEART_ACK, userId)
	data, _ := json.Marshal(ack)
	sendMsg(userId, sessionId, string(data))
}

/**
* 发牌
 */
func sendHandCardToPlayer(userId int, sessionId string, sendType int, cards []string) {
	response := cardmodel.GetBaseMode(sendType, userId)
	handCards := cardmodel.HandCards{
		Cards: cards,
	}
	response.Data = handCards
	data, _ := json.Marshal(response)
	sendMsg(userId, sessionId, string(data))
}

/**
* 玩家退出
 */
func ExitPlayer(userId int) {
	round := GetRoundByUserId(userId)
	if round.GameStatus == 2 { //若玩家在游戏中退出，则调用游戏结束
		ComputerGamerResult(&round, userId, false)
		sendGameStatus(round, enum.GAME_OVER)
	}
	playerExit := service.GetPlayer(userId)
	playerExit.AliveCards = []string{}
	playerExit.IsAway = false
	playerExit.MoneyDiff = 0
	playerExit.Status = 0
	playerExit.GameId = ""
	service.UpdatePlayer(&playerExit)
	if len(round.GameId) > 0 {
		player := cardmodel.Player{
			UserId: userId,
		}
		round.RemovePlayer(player) //删除
	}
	service.AddGameRound(&round)
	service.AddGameRecordByRound(round)
	ResetGame(&round)
}

/**
* 计算游戏结果
 */
func ComputerGamerResult(round *cardmodel.GameRound, awayUser int, isFarmerWin bool) {
	var newGamers []cardmodel.Player
	if awayUser > 0 {
		for _, gamer := range round.Gamers {
			if gamer.UserId == awayUser {
				gamer.IsAway = true
				gamer.MoneyDiff = -round.HandMoney * 2
			} else {
				gamer.IsAway = false
				gamer.MoneyDiff = +round.HandMoney
			}
			gamer.Money += gamer.MoneyDiff
			newGamers = append(newGamers, gamer)
		}
	} else {
		for _, gamer := range round.Gamers {
			gamer.IsAway = false
			if round.Landlord == gamer.UserId && isFarmerWin == false {
				gamer.MoneyDiff = +round.HandMoney * 2
			} else if round.Landlord == gamer.UserId && isFarmerWin == true {
				gamer.MoneyDiff = -round.HandMoney * 2
			} else if round.Landlord != gamer.UserId && isFarmerWin == false {
				gamer.MoneyDiff = -round.HandMoney
			} else if round.Landlord != gamer.UserId && isFarmerWin == true {
				gamer.MoneyDiff = +round.HandMoney
			}
			gamer.Money += gamer.MoneyDiff
			newGamers = append(newGamers, gamer)
		}
	}
	round.Gamers = newGamers
	gamerpc.UpdateUsersMoneyRpc(round.Gamers)
}

/**
* 发送游戏结束通知
 */
func sendGameStatus(round cardmodel.GameRound, status int) {
	response := cardmodel.GetBaseModeByCode(status)
	gameStatus := cardmodel.GameStatus{
		GamerList: round.Gamers,
	}

	response.Data = gameStatus
	data, _ := json.Marshal(response)
	for _, gamer := range gameStatus.GamerList {
		sendMsg(gamer.UserId, gamer.SessionId, string(data))
	}
}

/**
* 玩家下注
 */
func addBet(userId int, msg string) {
	round := GetRoundByUserId(userId)
	base, err := jsonparser.GetInt([]byte(msg), "data", "base")
	if err == nil {
		round.HandMoney = int(base) * 100
		if round.HandMoney == 300 {
			HandDispatcher(userId, &round)
			sendIdentifyToAll(&round, userId)
			return
		}
	} else {
		logrus.Println("addBet error ", err)
	}
	RoundNextOrder(&round)
	sendOrderToUser(round.CurrentOrderUser, round.Gamers[round.CurrentOrder].SessionId)
}

/**
* 次序轮换
 */
func RoundNextOrder(round *cardmodel.GameRound) {
	orderIndex := round.NextOrder()
	round.CurrentOrder = orderIndex
	round.CurrentOrderUser = round.Gamers[orderIndex].UserId
}

/**
* 身份确认广播
 */
func sendIdentifyToAll(round *cardmodel.GameRound, landlord int) {
	response := cardmodel.GetBaseModeByCode(enum.GAME_APPOINT_IDENTIFY)
	identifyArray := make([]cardmodel.Identify, 3)
	for i := 0; i < len(round.Gamers); i++ {
		identify := cardmodel.Identify{
			UserId: round.Gamers[i].UserId,
		}
		if landlord == identify.UserId {
			identify.Identity = 2
		} else {
			identify.Identity = 1
		}
		identifyArray[i] = identify
	}
	response.Data = identifyArray
	data, _ := json.Marshal(response)
	for _, gamer := range round.Gamers {
		sendMsg(gamer.UserId, gamer.SessionId, string(data))
	}
}

/**
* 给玩家发送次序通知
 */
func sendOrderToUser(userId int, sessionId string) {
	response := cardmodel.GetBaseMode(enum.GAME_ORDER, userId)
	order := cardmodel.Order{
		OrderUser: userId,
	}
	response.Data = order
	data, _ := json.Marshal(&response)
	sendMsg(userId, sessionId, string(data))
}

/**
* 玩家出牌
 */
func showCards(userId int, sessionId string, msg string) {
	round := GetRoundByUserId(userId)
	if userId != round.CurrentOrderUser {
		//出牌玩家校验
		return
	}
	cards := GetCardsFromMsg(msg)
	cardShow := card.ParseCardsInSize(cards)
	if round.ShowCount > 0 { //所有玩家出牌总次数大于0
		if cardShow.CardTypeStatus == round.LastShow.CardTypeStatus &&
			cardShow.CompareValue > round.LastShow.CompareValue &&
			cardShow.CompareCount == round.LastShow.CompareCount { //牌面比对：类型一致，连续次数一致，比较值大于上一次出牌
			round.LastShow = cardShow
			round.ShowCount += 1
			if cardShow.CardTypeStatus == enum.BOMB { //普通炸弹翻2倍
				round.HandMoney = round.HandMoney * 2
			} else if cardShow.CardTypeStatus == enum.KING_BOMB { //王炸翻3倍
				round.HandMoney = round.HandMoney * 3
			}
			aliveLength := RemoveShowFromAlive(userId, cards, &round)
			if aliveLength == 0 {
				var isFarmerWin bool
				if userId == round.Landlord {
					isFarmerWin = false
				} else {
					isFarmerWin = true
				}
				ComputerGamerResult(&round, 0, isFarmerWin)
				service.AddGameRound(&round)
				sendGameStatus(round, enum.GAME_OVER)
				ResetGame(&round)
				return
			} else if aliveLength > 0 {
				RoundNextOrder(&round)
			}
		} else {
			sendVerifyResult(-1, false, userId, sessionId)
			return
		}
	} else { //第一次出牌
		round.LastShow = cardShow
		round.ShowCount += 1
		RoundNextOrder(&round)
	}
	sendOrderToUser(round.CurrentOrderUser, round.Gamers[round.CurrentOrder].SessionId)
	sendShowResult(cardShow, userId, round, false)
}

/**
* 将已出的牌移除
 */
func RemoveShowFromAlive(userId int, showCards []string, round *cardmodel.GameRound) int {
	for _, gamer := range round.Gamers {
		if userId == gamer.UserId {
			gamer.AliveCards = util.FilterArray(gamer.AliveCards, showCards)
			service.AddPlayer(&gamer)
			return len(gamer.AliveCards)
		}
	}
	return -1
}

/**
* 发送牌面结果通知
 */
func sendShowResult(cardShow cardmodel.CardShow, userId int, round cardmodel.GameRound, isPass bool) {
	response := cardmodel.GetBaseMode(enum.GAME_SHOW_RESULT, userId)
	showResult := cardmodel.ShowResult{
		ShowPlayer: userId,
		IsPass:     isPass,
	}
	if isPass == false {
		showResult.ShowTime = int(util.GetNowTime().UnixNano() / 1000)
		showResult.ShowValue = cardShow.ShowValue
		showResult.CompareValue = cardShow.CompareValue
		showResult.CompareCount = cardShow.CompareCount
		showResult.CardTypeStatus = cardShow.CardTypeStatus
	}
	response.Data = showResult
	data, _ := json.Marshal(response)
	for _, gamer := range round.Gamers {
		sendMsg(gamer.UserId, gamer.SessionId, string(data))
	}
}

/**
* 发送牌面校验结果
 */
func sendVerifyResult(cardType enum.CardTypeStatus, isCredit bool, userId int, sessionId string) {
	response := cardmodel.GetBaseMode(enum.GAME_VERIFY_CARD, userId)
	result := cardmodel.VerifyResult{
		IsCredit: isCredit,
		CardType: cardType,
	}
	response.Data = result
	data, _ := json.Marshal(response)
	sendMsg(userId, sessionId, string(data))
}

/**
* 获取牌面
 */
func GetCardsFromMsg(msg string) []string {
	handCards := cardmodel.HandCards{}
	cards, _, _, err := jsonparser.Get([]byte(msg), "data")
	if err == nil {
		er := json.Unmarshal(cards, &handCards)
		if er == nil {
			logrus.Println(handCards)
		} else {
			logrus.Println(er)
		}
	}
	return handCards.Cards
}

/**
* 玩家pass
 */
func playerPass(useId int) {
	round := GetRoundByUserId(useId)
	RoundNextOrder(&round)
	service.AddGameRound(&round)
	sendOrderToUser(round.CurrentOrderUser, round.Gamers[round.CurrentOrder].SessionId)
	for _, gamer := range round.Gamers {
		sendShowResult(cardmodel.CardShow{}, gamer.UserId, round, true)
	}
}

/**
* 玩家准备
 */
func playerReady(userId int) {
	round := GetRoundByUserId(userId)
	readyCount := 0
	for _, gamer := range round.Gamers {
		logrus.Printf("UserId %d readyStatus %d ", gamer.UserId, gamer.Status)
		if gamer.UserId == userId {
			gamer.Status = 1
			readyCount += 1
			service.AddPlayer(&gamer)
		} else if gamer.Status == 1 {
			readyCount += 1
		}
	}
	logrus.Println("readyCount : ", readyCount)
	if readyCount == 3 { //玩家都已准备，游戏开始
		sendGameStatus(round, enum.GAME_START)
		StartGame(&round)
	}
}

/**
* 根据ActionCode处理业务
 */
func parseActionCode(actionCode int, userId int, sessionId string, msg string) {
	switch actionCode {

	case enum.HEART_HELLO:
		ackHeart(userId, sessionId)
		break
	case enum.PLAYER_JOIN:
		player := cardmodel.Player{
			UserId:    userId,
			SessionId: sessionId,
		}
		NewPlayerJoin(player)
		break
	case enum.PLAYER_EXIT:
		ExitPlayer(userId)
		break

	case enum.PLAYER_READY:
		playerReady(userId)
		break
	case enum.GAME_BET:
		addBet(userId, msg)
		break
	case enum.GAME_SHOW_CARD:
		showCards(userId, sessionId, msg)
		break
	case enum.GAME_PASS:
		playerPass(userId)
		break
	}

}
