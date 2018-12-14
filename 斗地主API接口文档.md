# 斗地主接口文档

### 一、说明

本文档基本遵循RESTFUL设计规范，请注意规范格式。  
所有成功请求成功均返回Http状态码200，  
具体业务由业务code码处理。

#### 1. Host地址
```
http://host
```

#### 2. 基本响应格式
|参数名|类型|描述|
|------|------|------|
|code|int|状态码|
|message|String|状态信息|
|data|Object|实际返回内容|
```
{
  code: 1000,
  message: "success",
  data: {
    
  }
}
```

#### 3. 请求公共Header参数
|参数名|类型|描述|
|------|------|------|
|access_token|String|用户访问令牌|
#### 4.状态码对照表

|code值|描述|
|------|------|
|1000|成功|
|1400|请求参数有误|
|1403|访问令牌无效|
|1405|访问资源不存在|
|1406|未查询到结果|
|1500|服务内部错误|
|2001|用户名或密码错误|
|2002|用户手机号已被注册|
|2003|该手机号未注册|
|2011|验证码错误|



### 二、账号模块
#### 1.登录接口

##### 路径
```
/account/login
```
##### 请求类型
```
POST
```

##### 请求参数
|参数名|类型|描述|
|------|------|------|
|mobile|String|用户手机号|
|password|String|密码，MD5加密|

##### 返回参数
|参数名|类型|描述|
|------|------|------|
|userId|String|用户id|
|mobile|用户手机号|
|nickName|昵称|
|money|金币|
|token|String|用户访问令牌|

###### 示例
```
{
    "code": 1000,
    "message": "登录成功",
    "data": {
        "userId": 10000,
        "mobile": "18258461820",
        "nickName": "小乐惠",
        "money": 0,
        "token": "1CA17zrhPkDoVkbLmtbajKbxC6n"
    }
}
```

#### 2.微信登录接口

##### 路径
```
/account/wxlogin
```
##### 请求类型
```
POST
```

##### 请求参数
|参数名|类型|描述|
|------|------|------|
|unicoId|String|用户唯一标识|

##### 返回参数
|参数名|类型|描述|
|------|------|------|
|userId|String|用户id|
|mobile|String|用户手机号|
|nickName|String|昵称|
|money|int|金币|
|token|String|用户访问令牌|

###### 示例
```
{
    "code": 1000,
    "message": "登录成功",
    "data": {
        "userId": 10000,
        "mobile": "18258461820",
        "nickName": "小乐惠",
        "money": 0,
        "token": "1CA17zrhPkDoVkbLmtbajKbxC6n"
    }
}
```

#### 3.创建账号接口

##### 路径
```
/account/create
```
##### 请求类型
```
POST
```

##### 请求参数
|参数名|类型|描述|
|------|------|------|
|mobile|String|用户手机号|
|password|String|密码，MD5加密|
|nickName|String|昵称|
|captchaId|String|验证码Id|
|captchaValue|String|验证码值|


##### 返回参数
|参数名|类型|描述|
|------|------|------|
|userId|String|用户id|
|mobile|String|用户手机号|
|nickName|String|昵称|
|money|int|金币|
###### 示例
```
{
    "code": 1000,
    "message": "success",
    "data": {
        "userId": 10001,
        "mobile": "18258461821",
        "nickName": "小乐惠1",
        "money": 0
    }
}
```

#### 4.个人信息查询接口

##### 路径
```
/account/userInfo
```
##### 请求类型
```
GET
```

##### 请求参数
|参数名|类型|描述|
|------|------|------|


##### 返回参数
|参数名|类型|描述|
|------|------|------|
|userId|String|用户Id|
|mobile|String|用户手机号|
|nickName|String|用户昵称|
|money|int|用户积分|

###### 示例
```
{
    "code": 1000,
    "message": "success",
    "data": {
        "userId": 10000,
        "mobile": "18258461820",
        "nickName": "小乐惠",
        "money": 0
    }
}
```

#### 5.个人信息修改接口(增量传参)

##### 路径
```
/account/modifyUserInfo
```
##### 请求类型
```
PUT
```

##### 请求参数
|参数名|类型|描述|
|------|------|------|
|userId|String|用户Id(必传)|
|mobile|String|用户手机号|
|nickName|String|用户昵称|
|money|int|用户积分|


##### 返回参数
|参数名|类型|描述|
|------|------|------|


###### 示例
```
{
  code: 1000,
  message: "修改成功",
  data: {
  }
}
```

#### 6.密码修改接口

##### 路径
```
/account/modifyPassword
```
##### 请求类型
```
PUT
```

##### 请求参数
|参数名|类型|描述|
|------|------|------|
|mobile|String|手机号|
|newPassword|String|新密码，MD5加密|
|captchaId|String|验证码Id|
|captchaValue|String|验证码值|


##### 返回参数
|参数名|类型|描述|
|------|------|------|


###### 示例
```
{
  code: 1000,
  message: "修改成功",
  data: {
  }
}
```

#### 7.登出接口

##### 路径
```
/account/loginOut
```
##### 请求类型
```
DELETE
```

##### 请求参数
|参数名|类型|描述|
|------|------|------|


##### 返回参数
|参数名|类型|描述|
|------|------|------|


###### 示例
```
{
  code: 1000,
  message: "success",
  data: {
  }
}
```

### 三、验证码模块
#### 1.获取验证码接口

##### 路径
```
/captcha/getCaptcha
```
##### 请求类型
```
GET
```

##### 请求参数
|参数名|类型|描述|
|------|------|------|

##### 返回参数
|参数名|类型|描述|
|------|------|------|
|captchaId|String|验证码Id(Hash码)|
|imageUrl|String|验证码图片地址(不带host)|

###### 示例
```
{
    "code": 1000,
    "message": "success",
    "data": {
        "captchaId": "HBKwk8MTwexdt2nB3K9W",
        "imageUrl": "/captcha/show/HBKwk8MTwexdt2nB3K9W.png"
    }
}
```

#### 2.校验验证码接口

##### 路径
```
/captcha/verifyCaptcha
```
##### 请求类型
```
GET
```

##### 请求参数
|参数名|类型|描述|
|------|------|------|
|captchaId|String|验证码Id|
|value|String|验证码值|

##### 返回参数
|参数名|类型|描述|
|------|------|------|

###### 示例
```
{
    "code": 2011,
    "message": "验证错误"
}
```

注意事项: 生成的验证码ID一旦发生校验行为，无论通不通过，验证码ID都失效，需要重新获取验证码ID和value图片

#### 3.获取验证码图片

##### 路径
```
/show/{captchaId}.png
```
##### 请求类型
```
GET
```

##### 请求参数
|参数名|类型|描述|
|------|------|------|
|captchaId|String|验证码Id|
##### 返回参数
|参数名|类型|描述|
|------|------|------|

###### 示例
```
http://192.168.2.115:8080/captcha/show/HBKwk8MTwexdt2nB3K9W.png
```

#### 4.刷新验证码图片

##### 路径
```
/show/{captchaId}.png?reload=true
```
##### 请求类型
```
GET
```

##### 请求参数
|参数名|类型|描述|
|------|------|------|
|captchaId|String|验证码Id|
##### 返回参数
|参数名|类型|描述|
|------|------|------|

###### 示例
```
http://192.168.2.115:8080/captcha/show/HBKwk8MTwexdt2nB3K9W.png?reload=true
```