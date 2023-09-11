package model

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func New(path string) (db *gorm.DB, err error) {
	db, err = gorm.Open(sqlite.Open(filepath.Join(path, "mirror.sqlite")))
	if err != nil {
		return
	}

	infoLog := logger.New(log.New(os.Stdout, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold:             200 * time.Millisecond,
		LogLevel:                  logger.Info,
		IgnoreRecordNotFoundError: false,
		Colorful:                  true,
	})

	if v, _ := strconv.ParseBool(os.Getenv("DEBUG_DB")); v {
		db.Logger = infoLog
	}

	err = db.AutoMigrate(&Rule{})

	return
}
