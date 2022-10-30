package middlewares

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/notblessy/memoriku/config"
	"strings"
)

type JWTClaims struct {
	jwt.Claims
	ID             string
	OrganizationID *string `json:"idOrganization"`
	UserType       string  `json:"userType"`
}

func parseJWT(token string) (*JWTClaims, error) {
	secretKey := []byte(config.JWTSecret())

	tokenClaims, err := jwt.ParseWithClaims(
		token,
		&JWTClaims{},
		func(t *jwt.Token) (interface{}, error) {
			return secretKey, nil
		},
	)

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JWTClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}

func Authorization() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			auth := c.Request().Header.Get("Authorization")
			if auth == "" {
				return echo.ErrUnauthorized
			}

			token := strings.Split(auth, " ")

			if len(token) == 2 {
				user, err := parseJWT(token[1])

				if err == nil {
					c.Set("user", user)
					if err := next(c); err != nil {
						c.Error(err)
					}
				}
			}

			return echo.ErrUnauthorized
		}
	}
}
