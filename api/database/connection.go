package database

import (
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db   *gorm.DB
	once sync.Once
)

func Connect() (*gorm.DB, error) {
	var err error
	once.Do(func() {
		dsn := "host=localhost user=postgres password=sysadmin dbname=gastos port=5432 sslmode=disable TimeZone=America/Lima"
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	})
	return db, err
}
