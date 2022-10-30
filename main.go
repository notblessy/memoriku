package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/notblessy/memoriku/config"
	"github.com/notblessy/memoriku/db"
	"github.com/notblessy/memoriku/http"
)

func main() {
	initDB := db.InitiateMysql()
	defer db.CloseMysql(initDB)

	if initDB != nil {
		log.Info("database connected")
	}

	e := echo.New()

	http.Routes(e)

	log.Fatal(e.Start(":" + config.HTTPPort()))
}
