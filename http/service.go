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
	route.POST("/login", h.loginHandler)

	routes := route.Group("/cms")
	routes.Use(middleware.Logger())
	routes.Use(middleware.Recover())
	routes.Use(middleware.JWTWithConfig(mdw.JWTConfig()))

	routes.GET("/user", h.profileHandler)

	routes.POST("/category", h.createCategoryHandler)

}
