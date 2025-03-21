package request

type RegisterRequest struct {
	Name string `json:"name" binding:"required"`      // 必填
	Age  int    `json:"age" binding:"gte=18,lte=120"` // 必填，18~120
}
