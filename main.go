package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
	}))
	e.Validator = &utils.CustomValidator{Validator: validator.New()}

	videoRepo := repository.NewUserRepository(initDB)
	categoryRepo := repository.NewCategoryRepository(initDB)
	memoryRepo := repository.NewMemoryRepository(initDB)

	httpSvc := http.NewHTTPService()
	httpSvc.RegisterUserRepository(videoRepo)
	httpSvc.RegisterCategoryRepository(categoryRepo)
	httpSvc.RegisterMemoryRepository(memoryRepo)

	httpSvc.Routes(e)

	log.Fatal(e.Start(":" + config.HTTPPort()))
}
