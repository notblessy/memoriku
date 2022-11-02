package db

import (
	"fmt"
	"github.com/notblessy/memoriku/config"
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitiateMysql() *gorm.DB {
	err := config.LoadENV()
	if err != nil {
		logger.Fatal(err)
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s?charset=utf8&parseTime=true&loc=Local", config.MysqlUser(), config.MysqlPassword(), config.MysqlHost(), config.MysqlDBName())
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to connect: %s", err))
	}

	return db
}

func CloseMysql(db *gorm.DB) {
	sql, err := db.DB()
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to disconnect: %s", err))
	}

	err = sql.Close()
	if err != nil {
		logger.Fatal(fmt.Sprintf("failed to close: %s", err))
	}
}
