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

// createCategoryHandler :nodoc:
func (h *HTTPService) createCategoryHandler(c echo.Context) error {
	logger := log.WithField("context", utils.Encode(c))
	var data model.Category

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

	err := h.categoryRepo.Create(data)
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

// findCategoriesHandler :nodoc:
func (h *HTTPService) findCategoriesHandler(c echo.Context) error {
	logger := log.WithField("context", utils.Encode(c))

	name := c.QueryParam("name")

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = utils.DefaultPage
	}

	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		size = utils.DefaultSize
	}

	req := model.CategoryReqQuery{
		Name: name,
		Page: page,
		Size: size,
	}

	cat, count, err := h.categoryRepo.FindAll(req)
	if err != nil {
		logger.Error(err)
		return utils.ResponseError(c, &utils.Response{
			Message: fmt.Sprintf("%s", err),
			Data:    err,
		})
	}

	return utils.ResponseOK(c, &utils.Response{
		Data: utils.BuildPagination(cat, int(count), req.Page, req.Size),
	})
}

// updateCategoryHandler :nodoc:
func (h *HTTPService) updateCategoryHandler(c echo.Context) error {
	category, err := h.getCategoryRequestBody(c)
	if err != nil {
		return utils.ResponseBadRequest(c, &utils.Response{
			Message: fmt.Sprintf("error validate request: %s", ErrBadRequest),
			Data:    nil,
		})
	}

	logger := log.WithFields(log.Fields{
		"context": utils.Encode(c),
		"request": utils.Encode(category),
	})

	_, err = h.categoryRepo.FindByID(category.ID)
	if err != nil {
		logger.Error(err)
		return utils.ResponseNotFound(c, &utils.Response{
			Message: err.Error(),
		})
	}

	category.UpdatedAt = time.Now()

	err = h.categoryRepo.Update(category)
	if err != nil {
		logger.Error(err)
		return utils.ResponseError(c, &utils.Response{
			Message: err.Error(),
		})
	}

	return utils.ResponseCreated(c, &utils.Response{
		Data: category.ID,
	})
}

// findCategoryByIDHandler :nodoc:
func (h *HTTPService) findCategoryByIDHandler(c echo.Context) error {
	logger := log.WithField("context", utils.Encode(c))

	id := c.Param("categoryID")
	if id == "" {
		logger.Error(ErrBadRequest)
		return utils.ResponseBadRequest(c, &utils.Response{
			Message: fmt.Sprintf("%s", ErrBadRequest),
		})
	}

	cat, err := h.categoryRepo.FindByID(id)
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
		Data: cat,
	})
}

// deleteCategoryByID :nodoc:
func (h *HTTPService) deleteCategoryByID(c echo.Context) error {
	logger := log.WithField("context", utils.Encode(c))

	id := c.Param("categoryID")
	if id == "" {
		logger.Error(ErrBadRequest)
		return utils.ResponseBadRequest(c, &utils.Response{
			Message: fmt.Sprintf("%s", ErrBadRequest),
		})
	}

	err := h.categoryRepo.DeleteByID(id)
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

// findGrouppedCategoriesHandler :nodoc:
func (h *HTTPService) findGrouppedCategoriesHandler(c echo.Context) error {
	logger := log.WithField("context", utils.Encode(c))

	name := c.QueryParam("name")

	page, err := strconv.Atoi(c.QueryParam("page"))
	if err != nil {
		page = utils.DefaultPage
	}

	size, err := strconv.Atoi(c.QueryParam("size"))
	if err != nil {
		size = utils.DefaultSize
	}

	req := model.CategoryReqQuery{
		Name: name,
		Page: page,
		Size: size,
	}

	categories, _, err := h.categoryRepo.FindAll(req)
	if err != nil {
		logger.Error(err)
		return utils.ResponseError(c, &utils.Response{
			Message: fmt.Sprintf("%s", err),
			Data:    err,
		})
	}

	var keys = make(map[string]bool)
	var groupIDs []string
	for _, data := range *categories {
		if _, value := keys[data.GroupID]; !value {
			keys[data.GroupID] = true
			groupIDs = append(groupIDs, data.GroupID)
		}
	}

	var results []model.CategoryWeb
	for i := 0; i < len(groupIDs); i++ {
		var result model.CategoryWeb
		for _, data := range *categories {
			if groupIDs[i] == data.GroupID {
				result.GroupID = model.GroupCategory[data.GroupID]
				result.Categories = append(result.Categories, h.MakeCategoryValueObject(data))
			}
		}
		results = append(results, result)
	}
	return utils.ResponseOK(c, &utils.Response{
		Data: results,
	})
}

func (h *HTTPService) getCategoryRequestBody(c echo.Context) (*model.Category, error) {
	var data model.Category

	if err := c.Bind(&data); err != nil {
		return nil, err
	}
	if err := c.Validate(&data); err != nil {
		return nil, err
	}

	if c.Param("categoryID") != "" {
		data.ID = c.Param("categoryID")
	}

	return &data, nil
}

func (h *HTTPService) MakeCategoryValueObject(category model.Category) (result model.ValueObject) {
	result.Value = category.ID
	result.Label = category.Name

	return result
}
