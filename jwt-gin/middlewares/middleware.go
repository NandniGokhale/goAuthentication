package middlewares

import (
	"fmt"
	"jwt/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.TokenValid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unanthorized")
			fmt.Println(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
