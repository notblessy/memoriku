package http

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/notblessy/memoriku/config"
	"github.com/notblessy/memoriku/middleware"
	"github.com/notblessy/memoriku/model"
	"github.com/notblessy/memoriku/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"time"
)

// loginHandler :nodoc:
func (h *HTTPService) loginHandler(c echo.Context) error {
	var data model.User

	if err := c.Bind(&data); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	if err := c.Validate(&data); err != nil {
		return utils.ResponseBadRequest(c, &utils.Response{
			Message: fmt.Sprintf("error validate: %s", ErrBadRequest),
			Data:    nil,
		})
	}

	logger := log.WithFields(log.Fields{
		"context": utils.Encode(c),
		"request": utils.Encode(data),
	})

	user, err := h.userRepo.FindByEmail(data.Email)
	if err != nil {
		logger.Error(err)
		return utils.ResponseError(c, &utils.Response{
			Message: err.Error(),
		})
	}

	if err != nil && user == nil {
		logger.Error(err)
		return utils.ResponseNotFound(c, &utils.Response{
			Message: ErrNotFound.Error(),
		})
	}

	if user.Email == data.Email && user.Password != data.Password {
		logger.Error(ErrIncorrectEmailOrPassword)
		return utils.ResponseUnauthorized(c, &utils.Response{
			Message: ErrIncorrectEmailOrPassword.Error(),
		})
	} else {
		claims := &middleware.JWTClaims{
			ID: user.ID,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

		t, err := token.SignedString([]byte(config.JWTSecret()))
		if err != nil {
			return err
		}

		return utils.ResponseOK(c, &utils.Response{
			Data: map[string]interface{}{
				"user_id": claims.ID,
				"token":   t,
			},
		})
	}
}

// profileHandler :nodoc:
func (h *HTTPService) profileHandler(c echo.Context) error {
	jwtClaim, err := middleware.GetSessionClaims(c)
	if err != nil {
		return utils.ResponseUnauthorized(c, &utils.Response{
			Message: fmt.Sprintf("%s", err),
			Data:    err,
		})
	}

	user, err := h.userRepo.FindByID(jwtClaim.ID)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return utils.ResponseNotFound(c, &utils.Response{
				Message: fmt.Sprintf("%s", err),
				Data:    err,
			})
		default:
			return utils.ResponseError(c, &utils.Response{
				Message: fmt.Sprintf("%s", err),
				Data:    err,
			})
		}
	}

	return utils.ResponseCreated(c, &utils.Response{
		Data: user,
	})
}

// updateProfileHandler :nodoc:
func (h *HTTPService) updateProfileHandler(c echo.Context) error {
	user, err := h.getRequestBody(c)
	if err != nil {
		return utils.ResponseBadRequest(c, &utils.Response{
			Message: fmt.Sprintf("error validate request: %s", ErrBadRequest),
			Data:    nil,
		})
	}

	logger := log.WithFields(log.Fields{
		"context": utils.Encode(c),
		"request": utils.Encode(user),
	})

	_, err = h.userRepo.FindByID(user.ID)
	if err != nil {
		logger.Error(err)
		return utils.ResponseNotFound(c, &utils.Response{
			Message: err.Error(),
		})
	}

	user.UpdatedAt = time.Now()

	err = h.userRepo.Update(user)
	if err != nil {
		logger.Error(err)
		return utils.ResponseError(c, &utils.Response{
			Message: err.Error(),
		})
	}

	return utils.ResponseCreated(c, &utils.Response{
		Data: user.ID,
	})
}

func (h *HTTPService) getRequestBody(c echo.Context) (*model.User, error) {
	var data model.User

	if err := c.Bind(&data); err != nil {
		return nil, err
	}
	if err := c.Validate(&data); err != nil {
		return nil, err
	}

	if c.Param("userID") != "" {
		userID, err := strconv.Atoi(c.Param("userID"))
		if err != nil {
			return nil, err
		}

		data.ID = int64(userID)
	}

	return &data, nil
}
