package request

type RegisterRequest struct {
	Name string `json:"name" binding:"required"`      // 必填
	Age  int    `json:"age" binding:"gte=18,lte=120"` // 必填，18~120
}

type GetUserInfoRequest struct {
	ID uint `json:"id" binding:"required,gte=1"` // 必填，大于等于1
}

type UpdateUserRequest struct {
	ID   uint   `json:"id" binding:"required,gte=1"`            // 必填
	Name string `json:"name"`                                   // 可选
	Age  *int   `json:"age" binding:"omitempty,gte=18,lte=120"` // 可选 + 范围校验
}

type DeleteUserRequest struct {
	ID uint `json:"id" binding:"required,gte=1"` // 必填，大于等于1
}

type UserListRequest struct {
	PageRequest
	Keyword string `json:"keyword"`
	MinAge  int    `json:"min_age"`
	MaxAge  int    `json:"max_age"`
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
