package http

import (
	"github.com/labstack/echo/v4"
	"github.com/notblessy/memoriku/model"
)

// HTTPService :nodoc:
type HTTPService struct {
	userRepo model.UserRepository
}

// NewHTTPService :nodoc:
func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

// RegisterUserRepository :nodoc:
func (h *HTTPService) RegisterUserRepository(v model.UserRepository) {
	h.userRepo = v
}

// Routes :nodoc:
func (h *HTTPService) Routes(route *echo.Echo) {
	user := route.Group("/user")
	user.POST("/login", h.loginHandler)

}
