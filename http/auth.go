package http

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/notblessy/memoriku/config"
	"github.com/notblessy/memoriku/model"
	"github.com/notblessy/memoriku/utils"
	"net/http"
	"time"
)

// loginHandler :nodoc:
func (h *HTTPService) loginHandler(c echo.Context) error {
	var data model.User

	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	user, err := h.userRepo.FindByEmail(data.Email)
	if err != nil {
		return utils.ResponseError(c, &utils.Response{
			Data: err,
		})
	}

	if err != nil && user == nil {
		return utils.ResponseUnauthorized(c, &utils.Response{
			Data: err,
		})
	}

	if user.Email == data.Email && user.Password != data.Password {
		return utils.ResponseUnauthorized(c, &utils.Response{
			Data: err,
		})
	} else {
		token := jwt.New(jwt.SigningMethodHS256)

		claims := token.Claims.(jwt.MapClaims)
		claims["id"] = user.ID
		claims["name"] = user.Name
		claims["email"] = user.Email
		claims["expired"] = time.Now().Add(time.Hour * 72).Unix()

		t, err := token.SignedString([]byte(config.JWTSecret()))
		if err != nil {
			return err
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": t,
		})
	}
}
