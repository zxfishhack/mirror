package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"path/filepath"
)

func New(path string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(filepath.Join(path, "mirror.sqlite")))
	if err != nil {
		return
	}

	err = db.AutoMigrate(&Rule{})

	return
}
