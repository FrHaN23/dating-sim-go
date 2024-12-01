package db

import "gorm.io/gorm"

func Paginate(size int, page int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if size == 0 {
			size = 15
		}
		offset := (page) * size
		return db.Offset(offset).Limit(size)
	}
}
