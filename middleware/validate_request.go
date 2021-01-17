package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gofiber/fiber/v2"
	jwtutils "github.com/hieronimusbudi/go-bookstore-utils/jwt"
	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
)

var (
	jwtSecret     = os.Getenv("JWT_SECRET")
	jwtCookieName = os.Getenv("JWT_COOKIE_NAME")
)

func ValidateRequest(c *fiber.Ctx) error {
	// Get token from cookie
	token := c.Cookies(jwtCookieName)
	if token == "" {
		restJwtErr := resterrors.NewUnauthorizedError("Unauthorized")
		return c.Status(restJwtErr.Status()).JSON(restJwtErr)
	}

	// Validate token
	tokenClaims, tokenErr := jwtutils.ValidateToken(token, jwtSecret)
	if tokenErr != nil {
		restJwtErr := resterrors.NewUnauthorizedError("Token claims not exists")
		return c.Status(restJwtErr.Status()).JSON(restJwtErr)
	}

	c.Context().SetUserValue("tokenClaims", tokenClaims)
	return c.Next()
}

func _ValidateRequest(jwtSecret string, jwtCookieName string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get token from cookie
		jwtCookie, jwtErr := c.Request.Cookie(jwtCookieName)
		if jwtErr != nil {
			restJwtErr := resterrors.NewUnauthorizedError(jwtErr.Error())
			c.JSON(restJwtErr.Status(), restJwtErr)
			c.Abort()
			return
		}

		// Validate token
		tokenClaims, tokenErr := jwtutils.ValidateToken(jwtCookie.Value, jwtSecret)
		if tokenErr != nil {
			c.JSON(tokenErr.Status(), tokenErr)
			c.Abort()
			return
		}

		c.Set("tokenClaims", tokenClaims)
		c.Next()
	}
}
