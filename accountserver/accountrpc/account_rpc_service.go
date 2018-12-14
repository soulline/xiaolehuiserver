package accountrpc

import (
	"encoding/json"
	"github.com/buger/jsonparser"
	"github.com/sirupsen/logrus"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"xiaolehuigo/accountserver/model"
	"xiaolehuigo/accountserver/model/datamodel"
	"xiaolehuigo/accountserver/service"
	"xiaolehuigo/accountserver/util"
)

type AccountRpc int

/**
* 远程RPC调用查询用户信息
 */
func (arpc *AccountRpc) GetAccountInfo(userId *int, reply *string) error {
	logrus.Info("GetAccountInfo userId:", *userId)
	userInfo := service.QueryUserByUserId(*userId)
	baseResponse := model.NewBaseResponse()
	if userInfo.UserId <= 0 {
		baseResponse.GetFailureResponse(model.QUERY_NO_DATA)
	} else {
		baseResponse.GetSuccessResponse()
		baseResponse.Data = util.TurnUserInfoResp(userInfo)
	}
	response := InterfaceToString(baseResponse)
	logrus.Info("GetAccountInfo response:", response)
	*reply = response
	return nil
}

/**
* 校验是否登录
 */
func (arpc *AccountRpc) CheckLoginStatus(userId *int, reply *string) error {
	logrus.Info("CheckLoginStatus userId:", *userId)
	ok := service.CheckLoginStatus(*userId)
	baseResponse := model.NewBaseResponse()
	if ok {
		baseResponse.GetSuccessResponse()
	} else {
		baseResponse.GetFailureResponse(model.LOGIN_INVALID)
	}
	response := InterfaceToString(baseResponse)
	logrus.Info("CheckLoginStatus response:", response)
	*reply = response
	return nil
}

func (arpc *AccountRpc) UpdateUsersMoney(users *string, reply *string) error {
	logrus.Info("UpdateUsersMoney", *users)
	baseResponse := model.NewBaseResponse()
	var usersModel []datamodel.UserInfo
	json.Unmarshal([]byte(*users), &usersModel)
	ok := service.UpdateUsersMoney(usersModel)
	if ok {
		baseResponse.GetSuccessResponse()
	} else {
		baseResponse.GetFailureResponse(model.SYSTEM_ERROE)
	}
	response := InterfaceToString(baseResponse)
	logrus.Info("UpdateUsersMoney response:", response)
	*reply = response
	return nil
}

/**
* 远程RPC调用更新用户金币
 */
func (arpc *AccountRpc) UpdateUserMoney(userInfo *string, reply *string) error {
	logrus.Info("UpdateUserMoney", *userInfo)
	userIdParse, _ := jsonparser.GetInt([]byte(*userInfo), "userId")
	userMoney, _ := jsonparser.GetInt([]byte(*userInfo), "money")
	user := datamodel.UserInfo{
		UserId: int(userIdParse),
		Money:  int(userMoney),
	}
	baseResponse := model.NewBaseResponse()
	if user.UserId > 0 {
		err1, _ := service.UpdateMoney(user)
		if err1 == nil {
			baseResponse.GetSuccessResponse()
		} else {
			baseResponse.GetFailureResponse(model.SYSTEM_ERROE)
		}
	} else {
		baseResponse.GetFailureResponse(model.QUERY_NO_DATA)
	}
	response := InterfaceToString(baseResponse)
	logrus.Info("UpdateUserMoney response:", response)
	*reply = response
	return nil
}

func InterfaceToString(v interface{}) string {
	data, _ := json.Marshal(v)
	response := string(data)
	return response
}

func Listen(ch chan int) {
	rpc.Register(new(AccountRpc))
	l, err := net.Listen("tcp", ":2345")
	if err != nil {
		logrus.Info("listener rpc error : ", err)
	}
	logrus.Info("go jsonRpc")
	ch <- 0
	for {
		logrus.Info("wating 2345")
		conn, err := l.Accept()
		if err != nil {
			logrus.Info("accept connection err : s%\n", conn)
		}
		go jsonrpc.ServeConn(conn)
	}
}
