package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/notblessy/memoriku/config"
	"github.com/notblessy/memoriku/db"
	"github.com/notblessy/memoriku/http"
	"github.com/notblessy/memoriku/repository"
	"github.com/notblessy/memoriku/utils"
)

func main() {
	initDB := db.InitiateMysql()
	defer db.CloseMysql(initDB)

	e := echo.New()
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	videoRepo := repository.NewUserRepository(initDB)
	categoryRepo := repository.NewCategoryRepository(initDB)

	httpSvc := http.NewHTTPService()
	httpSvc.RegisterUserRepository(videoRepo)
	httpSvc.RegisterCategoryRepository(categoryRepo)

	httpSvc.Routes(e)

	log.Fatal(e.Start(":" + config.HTTPPort()))
}
