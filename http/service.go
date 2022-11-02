package http

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	mdw "github.com/notblessy/memoriku/middleware"
	"github.com/notblessy/memoriku/model"
)

// HTTPService :nodoc:
type HTTPService struct {
	userRepo     model.UserRepository
	categoryRepo model.CategoryRepository
}

// NewHTTPService :nodoc:
func NewHTTPService() *HTTPService {
	return new(HTTPService)
}

// RegisterUserRepository :nodoc:
func (h *HTTPService) RegisterUserRepository(u model.UserRepository) {
	h.userRepo = u
}

// RegisterCategoryRepository :nodoc:
func (h *HTTPService) RegisterCategoryRepository(c model.CategoryRepository) {
	h.categoryRepo = c
}

// Routes :nodoc:
func (h *HTTPService) Routes(route *echo.Echo) {
	user := route.Group("/user")
	user.POST("/login", h.loginHandler)

	category := route.Group("/category")
	category.Use(middleware.Logger())
	category.Use(middleware.Recover())
	category.Use(middleware.JWTWithConfig(mdw.JWTConfig()))
	category.POST("/", h.createCategoryHandler)

}
