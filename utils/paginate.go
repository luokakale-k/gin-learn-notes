package utils

import "gorm.io/gorm"

type PageData struct {
	List     interface{} `json:"list"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	PageSize int         `json:"page_size"`
}

// 泛型分页函数
func Paginate[T any](db *gorm.DB, page, pageSize int) ([]T, int64, error) {
	var list []T
	var total int64

	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10
	}

	db.Count(&total)

	err := db.Offset((page - 1) * pageSize).Limit(pageSize).Find(&list).Error
	if err != nil {
		return nil, 0, err
	}

	return list, total, nil
}
