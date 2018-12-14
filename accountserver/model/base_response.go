package model

const (
	SUCCESS            = 1000
	FORBIDDEN          = 1403
	QUERY_PARAM_ERROR  = 1400
	RESOURCE_NOT_FOUNT = 1405
	QUERY_NO_DATA      = 1406
	LOGIN_INVALID      = 1407
	SYSTEM_ERROE       = 1500
	PASSWORD_ERROR     = 2001
	USER_EXIST         = 2002
	USER_NOT_EXIST     = 2003
	CAPTCHA_ERROR      = 2011
)

var respErrorMap RespErrorMap

type BaseResponse struct {
	RespError
	Data interface{} `json:"data,omitempty"` //omitempty有值就输出，没值则不输出
}

type RespErrorMap map[int]string

type RespError struct {
	Code    int    `json:"code"`
	Message string `json:"message,omitempty"`
}

func (RespError RespError) Error() string {
	return RespError.Message
}

func NewBaseResponse() BaseResponse {
	return BaseResponse{}
}

func (BaseResponse *BaseResponse) GetSuccessResponse() {
	BaseResponse.Code = SUCCESS
	BaseResponse.Message = "success"
}

func (BaseResponse *BaseResponse) GetFailureResponse(code int) {
	BaseResponse.RespError = NewRespError(code).(RespError)
}

/**
返回一个成功的响应
*/
func NewSuccessResponse() BaseResponse {
	response := BaseResponse{}
	response.Code = SUCCESS
	return response
}

/**
返回一个业务失败的响应，根据不同的code返回不同的message
*/
func NewFailureResponse(code int) BaseResponse {
	response := BaseResponse{}
	response.RespError = NewRespError(code).(RespError)
	return response
}

func NewRespError(code int) error {
	return RespError{Code: code, Message: respErrorMap[code]}
}

func init() {
	respErrorMap = RespErrorMap{
		QUERY_PARAM_ERROR:  "请求参数有误",
		RESOURCE_NOT_FOUNT: "访问资源不存在",
		QUERY_NO_DATA:      "未查询到结果",
		LOGIN_INVALID:      "账号未登录",
		SYSTEM_ERROE:       "服务内部错误",
		FORBIDDEN:          "访问令牌无效",
		PASSWORD_ERROR:     "用户名或密码错误",
		USER_EXIST:         "用户手机号已被注册",
		USER_NOT_EXIST:     "该手机号未注册",
		CAPTCHA_ERROR:      "验证码错误",
	}
}
