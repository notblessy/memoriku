package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/notblessy/memoriku/model"
	"github.com/notblessy/memoriku/utils"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"time"
)

// createMemoryHandler :nodoc:
func (h *HTTPService) createMemoryHandler(c echo.Context) error {
	logger := log.WithField("context", utils.Encode(c))
	var data model.Memory

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

	if data.ID == 0 {
		data.ID = time.Now().UnixNano() + int64(rand.Intn(10000))
	}

	err := h.memoryRepo.Create(&data)
	if err != nil {
		return utils.ResponseError(c, &utils.Response{
			Message: fmt.Sprintf("%s", err),
			Data:    err,
		})
	}

	return utils.ResponseOK(c, &utils.Response{
		Data: data.ID,
	})
}
