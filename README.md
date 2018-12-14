### socket地址端口
```
192.168.1.22:8888
```

### 基础协议
数据格式以json格式传输
|参数名|类型|描述|
|------|------|------|
|actionCode|int|指令类型|
|userId|int|用户Id|
|data|Object|指令详情|

```
{
    actionCode : 1001,
    userId : 10003,
    data : null
}
```

### 心跳协议
服务端每秒都会检查一次所有会话最后一次收发消息的时间，60秒内无消息收发则会从服务端主动断开连接  
心跳间隔控制在60秒以内

|参数名|类型|描述|
|------|------|------|
|systemTime|long|时间戳|

#### 客户端

示例
```
{
    actionCode : 1001,
    userId : 10003,
    data : {
        systemTime: 153129123934
    }
}
```
#### 服务端

示例
```
{
    actionCode : 1002,
    userId : 10003,
    data : {
        systemTime: 153129123934
    }
}
```

![image](https://upload-images.jianshu.io/upload_images/12543273-60c77c92b8ebb222.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### actionCode映射表

|actionCode|对应指令业务类型|
|------|------|
|1001|心跳请求|
|1002|心跳响应|
|2001|玩家加入|
|2002|玩家离开|
|2003|玩家准备|
|3001|牌局开始|
|3002|牌局结束|
|3003|发牌|
|3004|身份分配广播|
|3005|叫地主/抢地主|
|3006|底牌分配|
|3007|次序广播|
|3008|出牌|
|3009|出牌校验结果|
|3010|出牌结果广播|
|3011|不要|

### 游戏时序图
![牌局时序图.png](https://upload-images.jianshu.io/upload_images/12543273-b0472df84aee8572.png?imageMogr2/auto-orient/strip%7CimageView2/2/w/1240)


### 游戏环节协议

#### 1.玩家加入

|参数名|类型|描述|
|------|------|------|


示例
```
{
    actionCode : 2001,
    userId : 10003,
    data : {
    }
}
```
#### 2.玩家离开

|参数名|类型|描述|
|------|------|------|


示例
```
{
    actionCode : 2002,
    userId : 10003,
    data : {
    }
}
```

#### 3.玩家准备

|参数名|类型|描述|
|------|------|------|



示例
```
{
    actionCode : 2003,
    userId : 10003,
    data : {
    }
}
```

#### 4.牌局开始

|参数名|类型|描述|
|------|------|------|
|gamerList|array|玩家列表|

gamerList

|参数名|类型|描述|
|------|------|------|
|userId|int|玩家用户id|
|nickName|String|玩家昵称|
|level|int|玩家等级|


示例
```
{
  actionCode: 3001,
  data: {
    gamerList: [
      {
        userId: 10001,
        nickName: 小乐惠地主,
        level: 0
      },
      {
        userId: 10002,
        nickName: 小乐惠地主1,
        level: 0
      },
      {
        userId: 10003,
        nickName: 小乐惠地主,
        level: 0
      }
    ]
  }
}
```

#### 5.牌局结束

|参数名|类型|描述|
|------|------|------|
|gamerList|array|玩家列表|

gamerList

|参数名|类型|描述|
|------|------|------|
|userId|int|玩家用户id|
|nickName|String|玩家昵称|
|level|int|玩家等级|
|moneyDiff|int|玩家金币加减|
|isAway|boolean|是否逃跑|


示例
```
{
  actionCode: 3002,
  data: {
    gamerList: [
      {
        userId: 10001,
        nickName: 小乐惠地主,
        level: 0,
        moneyDiff: +200,
        isAway: false
      },
      {
        userId: 10002,
        nickName: 小乐惠地主1,
        level: 0,
        moneyDiff: -100,
        isAway: false
      },
      {
        userId: 10003,
        nickName: 小乐惠地主,
        level: 0,
        moneyDiff: -100,
        isAway: false
      }
    ]
  }
}
```
#### 6.发牌
|参数名|类型|描述|
|------|------|------|
|cards|string[]|牌面数组|

示例
```
{
  actionCode: 3003,
  data: {
    cards:[A4, C14, B14, C4, C13, C15, D6, D14, A13, B13, D11, B4, B12, C12, B9, D8, B6]
  }
}
```

#### 7.次序广播
|参数名|类型|描述|
|------|------|------|
|orderUser|int|当前次序userId|

示例
```
{
  actionCode: 3004,
  data: {
    orderUser:10001
  }
}
```

#### 8.叫抢地主

|参数名|类型|描述|
|------|------|------|
|base|int|叫分(1:100底，2:200底，3:300底)|

示例
```
{
  actionCode: 3005,
  userId:10001,
  data: {
    base: 3
  }
}
```

#### 9.身份分配广播

|参数名|类型|描述|
|------|------|------|
|userId|int|玩家id|
|identity|int|身份类型，0: 未分配 1:贫民  2: 地主|

示例
```
{
  actionCode: 3004,
  userId: 10001,
  data: [
    {
      userId: 10001,
      identity: 1
    },
    {
      userId: 10002,
      identity: 1
    },
    {
      userId: 10003,
      identity: 2
    }
  ]
}
```

#### 10.底牌分配
|参数名|类型|描述|
|------|------|------|
|cards|string[]|牌面数组|

示例
```
{
  "actionCode": 3006,
  "userId": 10001,
  "data": {
    "cards": [A4, C14, B14]
  }
}
```

#### 11.出牌
|参数名|类型|描述|
|------|------|------|
|cards|string[]|牌面数组|

示例
```
{
  actionCode: 3008,
  userId: 10001,
  data: {
    cards: [A4, C14, B4, A9]
  }
}
```

#### 12.出牌校验结果
|参数名|类型|描述|
|------|------|------|
|isCredit|boolean|是否正确|

示例
```
{
  actionCode: 3009,
  userId: 10001,
  data: {
    isCredit: true
  }
}
```

#### 12.出牌结果广播
|参数名|类型|描述|
|------|------|------|
|showTime|long|出牌时间戳|
|showValue|string[]|牌面数组|
|compareValue|int|用于比较大小的值|
|compareCount|int|用于比较的连续次数|
|cardTypeStatus|int|牌面类型(具体详见牌面协议)|
|showPlayer|int|出牌玩家id|

示例
```
{
  actionCode: 3010,
  data: {
    showTime: 153121434532，
    showValue: [A4, C14, B4, A9],
    compareValue: 4,
    compareCount: 1
    cardTypeStatus: 3,
    showPlayer: 10001
  }
}
```

#### 13.不要

|参数名|类型|描述|
|------|------|------|

示例
```
{
  actionCode: 3011,
  userId: 10001,
  data: {
     
  }
}
```
