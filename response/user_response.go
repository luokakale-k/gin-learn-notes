package response

// LoginResult 是用户登录的返回结构
type LoginResult struct {
	Token    string `json:"token"`
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
}
