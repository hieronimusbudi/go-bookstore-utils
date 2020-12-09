package middleware

import (
	"github.com/gin-gonic/gin"
	myjwt "github.com/hieronimusbudi/go-bookstore-utils/my_jwt"
	resterrors "github.com/hieronimusbudi/go-bookstore-utils/rest_errors"
)

func ValidateRequest(jwtSecret string, jwtCookieName string) gin.HandlerFunc {
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
		tokenClaims, tokenErr := myjwt.ValidateToken(jwtCookie.Value, jwtSecret)
		if tokenErr != nil {
			c.JSON(tokenErr.Status(), tokenErr)
			c.Abort()
			return
		}

		c.Set("tokenClaims", tokenClaims)
		c.Next()
	}
}
