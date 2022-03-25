package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewSQLiteConnection() (db *gorm.DB, err error) {
	dsn := "file:memdb1?mode=memory&cached=shared"

	db, err = gorm.Open(sqlite.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, err
}
