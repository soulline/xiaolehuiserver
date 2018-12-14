package gamerpc

import (
	"encoding/json"
	"fmt"
	"github.com/buger/jsonparser"
	"github.com/sirupsen/logrus"
	"net"
	"net/rpc/jsonrpc"
	"time"
	"xiaolehuigo/gameserver/model"
)

/**
* 查询玩家信息
 */
func GetPlayerInfoRpc(userId int) cardmodel.Player {
	timeout := time.Second * 30
	conn, err := net.DialTimeout("tcp", ":2345", timeout)
	if err != nil {
		logrus.Fatal("dialing:", err)
		return cardmodel.Player{}
	}
	defer conn.Close()
	client := jsonrpc.NewClient(conn)
	var reply string
	err = client.Call("AccountRpc.GetAccountInfo", userId, &reply)
	if err != nil {
		logrus.Fatal("AccountRpc.GetAccountInfo error:", err)
	}
	fmt.Println("GetPlayerInfoRpc reply : ", reply)
	userIdParse, _ := jsonparser.GetInt([]byte(reply), "data", "userId")
	nickName, _ := jsonparser.GetString([]byte(reply), "data", "nickName")
	money, _ := jsonparser.GetInt([]byte(reply), "data", "money")
	player := cardmodel.Player{
		UserId:   int(userIdParse),
		NickName: nickName,
		Money:    int(money),
	}
	data, err := jsonparser.GetString([]byte(reply), "data")
	logrus.Info("getPlayer : ", data)
	if err == nil {
		json.Unmarshal([]byte(data), player)
	}
	return player
}

func UpdateUsersMoneyRpc(players []cardmodel.Player) bool {
	timeout := time.Second * 30
	conn, err := net.DialTimeout("tcp", ":2345", timeout)
	if err != nil {
		logrus.Fatal("dialing:", err)
		return false
	}
	defer conn.Close()
	client := jsonrpc.NewClient(conn)
	data, _ := json.Marshal(players)
	reply := new(string)
	logrus.Info("UpdateUsersMoneyRpc request : ", string(data))
	divCall := client.Go("AccountRpc.UpdateUsersMoney", string(data), reply, nil)
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		logrus.Fatal("AccountRpc.UpdateUserMoney error:", replyCall.Error)
	}
	code, _ := jsonparser.GetInt([]byte(*reply), "code")
	logrus.Info("UpdateUserMoneyRpc response : ", *reply)
	if code == 1000 {
		return true
	}
	return false
}

/**
* 更新玩家金币
 */
func UpdateUserMoneyRpc(player cardmodel.Player) bool {
	timeout := time.Second * 30
	conn, err := net.DialTimeout("tcp", ":2345", timeout)
	if err != nil {
		logrus.Fatal("dialing:", err)
		return false
	}
	defer conn.Close()
	client := jsonrpc.NewClient(conn)
	data, _ := json.Marshal(player)
	reply := new(string)
	logrus.Info("UpdateUserMoneyRpc request : ", string(data))
	divCall := client.Go("AccountRpc.UpdateUserMoney", string(data), reply, nil)
	replyCall := <-divCall.Done
	if replyCall.Error != nil {
		logrus.Fatal("AccountRpc.UpdateUserMoney error:", replyCall.Error)
	}
	code, _ := jsonparser.GetInt([]byte(*reply), "code")
	logrus.Info("UpdateUserMoneyRpc response : ", *reply)
	if code == 1000 {
		return true
	}
	return false
}
