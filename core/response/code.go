package response

const (
	CodeSuccess      = 0     // 成功
	ParamError       = 10001 // 参数错误
	NotFound         = 10002 // 数据不存在
	DBError          = 10003 // 数据库错误
	Unauthorized     = 10004 // 未授权/未登录
	ServerError      = 10005 // 服务内部错误
	InvalidToken     = 10006
	TokenFormatError = 10007
	TokenExpired     = 10008
)
