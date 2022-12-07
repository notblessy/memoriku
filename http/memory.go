package http

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/notblessy/memoriku/model"
	"github.com/notblessy/memoriku/utils"
	log "github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"math/rand"
	"strconv"
	"time"
)

// createMemoryHandler :nodoc:
func (h *HTTPService) createMemoryHandler(c echo.Context) error {
	logger := log.WithField("context", utils.Encode(c))
	var data model.Memory

	if err := c.Bind(&data); err != nil {
		logger.Error(ErrBadRequest)
		return utils.ResponseBadRequest(c, &utils.Response{
			Message: fmt.Sprintf("error binding: %s", ErrBadRequest),
			Data:    ErrBadRequest,
		})
	}

	if err := c.Validate(&data); err != nil {
		logger.Error(ErrBadRequest)
		return utils.ResponseBadRequest(c, &utils.Response{
			Message: fmt.Sprintf("error validate: %s", ErrBadRequest),
			Data:    ErrBadRequest,
		})
	}

	if data.ID == "" {
		randomID := time.Now().Nanosecond() + rand.Intn(10000)
		data.ID = strconv.Itoa(randomID)
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

// findMemoriesHandler :nodoc:
func (h *HTTPService) findMemoriesHandler(c echo.Context) error {
	logger := log.WithField("context", utils.Encode(c))

	title := c.QueryParam("title")
	categoryID := c.QueryParam("categoryID")

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = utils.DefaultPage
	}

	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		size = utils.DefaultSize
	}

	req := model.MemoryReqQuery{
		Title:      title,
		CategoryID: categoryID,
		Page:       page,
		Size:       size,
	}

	memories, count, err := h.memoryRepo.FindAll(req)
	if err != nil {
		logger.Error(err)
		return utils.ResponseError(c, &utils.Response{
			Message: fmt.Sprintf("%s", err),
			Data:    err,
		})
	}
	
	return utils.ResponseOK(c, &utils.Response{
		Data: utils.BuildPagination(memories, int(count), req.Page, req.Size),
	})
}

// findMemoryByIDHandler :nodoc:
func (h *HTTPService) findMemoryByIDHandler(c echo.Context) error {
	logger := log.WithField("context", utils.Encode(c))

	id := c.Param("memoryID")
	if id == "" {
		logger.Error(ErrBadRequest)
		return utils.ResponseBadRequest(c, &utils.Response{
			Message: fmt.Sprintf("%s", ErrBadRequest),
		})
	}

	memory, err := h.memoryRepo.FindByID(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			logger.Error(err)
			return utils.ResponseNotFound(c, &utils.Response{
				Message: fmt.Sprintf("%s", err),
				Data:    err,
			})
		default:
			logger.Error(err)
			return utils.ResponseError(c, &utils.Response{
				Message: fmt.Sprintf("%s", err),
				Data:    err,
			})
		}
	}

	return utils.ResponseOK(c, &utils.Response{
		Data: memory,
	})
}

// deleteMemoryByID :nodoc:
func (h *HTTPService) deleteMemoryByID(c echo.Context) error {
	logger := log.WithField("context", utils.Encode(c))

	id := c.Param("memoryID")
	if id == "" {
		logger.Error(ErrBadRequest)
		return utils.ResponseBadRequest(c, &utils.Response{
			Message: fmt.Sprintf("%s", ErrBadRequest),
		})
	}

	err := h.memoryRepo.DeleteByID(id)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			logger.Error(err)
			return utils.ResponseNotFound(c, &utils.Response{
				Message: fmt.Sprintf("%s", err),
				Data:    err,
			})
		default:
			logger.Error(err)
			return utils.ResponseError(c, &utils.Response{
				Message: fmt.Sprintf("%s", err),
				Data:    err,
			})
		}
	}

	return utils.ResponseOK(c, &utils.Response{
		Data: id,
	})
}
