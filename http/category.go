package http

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/notblessy/memoriku/middleware"
	"github.com/notblessy/memoriku/model"
	"github.com/notblessy/memoriku/utils"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"time"
)

var (
	ErrBadRequest error = errors.New("bad request")
)

// createCategoryHandler :nodoc:
func (h *HTTPService) createCategoryHandler(c echo.Context) error {
	logger := log.WithField("context", c)
	var data model.Category

	if err := c.Bind(&data); err != nil {
		logger.Error(ErrBadRequest)
		return utils.ResponseBadRequest(c, &utils.Response{
			Status:  "ERROR",
			Message: fmt.Sprintf("error binding: %s", ErrBadRequest),
			Data:    ErrBadRequest,
		})
	}

	if err := c.Validate(&data); err != nil {
		logger.Error(ErrBadRequest)
		return utils.ResponseBadRequest(c, &utils.Response{
			Status:  "ERROR",
			Message: fmt.Sprintf("error validate: %s", ErrBadRequest),
			Data:    ErrBadRequest,
		})
	}

	_, err := middleware.GetSessionClaims(c)
	if err != nil {
		return utils.ResponseUnauthorized(c, &utils.Response{
			Status:  "ERROR",
			Message: fmt.Sprintf("%s", err),
			Data:    err,
		})
	}

	if data.ID == 0 {
		data.ID = time.Now().UnixNano() + int64(rand.Intn(10000))
	}

	err = h.categoryRepo.Create(data)
	if err != nil {
		return utils.ResponseError(c, &utils.Response{
			Status:  "ERROR",
			Message: fmt.Sprintf("%s", err),
			Data:    err,
		})
	}

	return c.JSON(http.StatusOK, utils.Response{
		Status:  "SUCCESS",
		Message: "SUCCESS",
		Data:    data.ID,
	})
}
