package database

import (
	"sync"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Connect() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		dsn := "root:sysadmin@tcp(127.0.0.1:3306)/gastos?charset=utf8mb4&parseTime=True&loc=Local"
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	})
	return db, err
}

// func Connect() (*gorm.DB, error) {
// 	dsn := "root:sysadmin@tcp(127.0.0.1:3306)/gastos?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return db, nil
// }
