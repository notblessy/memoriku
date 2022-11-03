package utils

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type Response struct {
	Status  string      `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func ResponseError(c echo.Context, response *Response) error {
	defaultValueError(response)
	return c.JSON(http.StatusInternalServerError, response)
}

func ResponseUnauthorized(c echo.Context, response *Response) error {
	defaultValueError(response)
	return c.JSON(http.StatusUnauthorized, response)
}

func ResponseNotFound(c echo.Context, response *Response) error {
	defaultValueError(response)
	return c.JSON(http.StatusNotFound, response)
}

func ResponseBadRequest(c echo.Context, response *Response) error {
	defaultValueError(response)
	return c.JSON(http.StatusBadRequest, response)
}

func ResponseCreated(c echo.Context, response *Response) error {
	defaultValueSuccess(response)
	return c.JSON(http.StatusCreated, response)
}

func ResponseOK(c echo.Context, response *Response) error {
	defaultValueSuccess(response)
	return c.JSON(http.StatusOK, response)
}

func defaultValueSuccess(response *Response) {
	if response.Status == "" {
		response.Status = "SUCCESS"
	}

	if response.Message == "" {
		response.Message = "SUCCESS"
	}
}

func defaultValueError(response *Response) {
	if response.Status == "" {
		response.Status = "ERROR"
	}

	if response.Message == "" {
		response.Message = "ERROR"
	}
}
