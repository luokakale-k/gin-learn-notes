package utils

import "gorm.io/gorm"

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
