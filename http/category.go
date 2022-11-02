package http

import (
	"errors"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/notblessy/memoriku/model"
	"github.com/notblessy/memoriku/utils"
	log "github.com/sirupsen/logrus"
	"math/rand"
	"net/http"
	"strconv"
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

	if data.ID == 0 {
		data.ID = time.Now().UnixNano() + int64(rand.Intn(10000))
	}

	err := h.categoryRepo.Create(data)
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

// findCategoriesHandler :nodoc:
func (h *HTTPService) findCategoriesHandler(c echo.Context) error {
	logger := log.WithField("context", c)

	name := c.QueryParam("name")

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = utils.DefaultPage
	}
	
	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		size = utils.DefaultSize
	}

	req := model.CategoryRequest{
		Name: name,
		Page: page,
		Size: size,
	}

	cat, count, err := h.categoryRepo.FindAll(req)
	if err != nil {
		logger.Error(err)
		return utils.ResponseError(c, &utils.Response{
			Status:  "ERROR",
			Message: fmt.Sprintf("%s", err),
			Data:    err,
		})
	}

	return c.JSON(http.StatusOK, utils.BuildPagination(cat, int(count), req.Page, req.Size))
}
