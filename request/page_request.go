package request

type PageRequest struct {
	Page     int `json:"page" binding:"gte=1"`
	PageSize int `json:"page_size" binding:"gte=1,lte=100"`
}
