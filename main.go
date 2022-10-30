package main

import (
	"github.com/labstack/gommon/log"
	"github.com/notblessy/memoriku/db"
)

func main() {
	initDB := db.InitiateMysql()
	defer db.CloseMysql(initDB)

	if initDB != nil {
		log.Info("CONNECT SUCCESS")
	}
}
