package jwt

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"gin-gonic/pkg/e"
	"gin-gonic/pkg/util"
)

// Check if a request has token or permission needed
func Check() gin.HandlerFunc {
	return func(c *gin.Context) {
		var code int
		var token string

		code = e.SUCCESS
		rToken := c.Request.Header["Authorization"]

		if len(rToken) < 1 {
			code = e.ERROR_MISSING_TOKEN
		} else {
			token = rToken[0]
			splitToken := strings.Split(token, "Bearer")
			token = strings.TrimSpace(splitToken[1])

			claims, err := util.ParseToken(token)
			if err != nil {
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
			} else {
				if time.Now().Unix() > claims.ExpiresAt {
					code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
				} else {
					c.Set("id_user", claims.ID)
				}
			}
		}

		if code != e.SUCCESS {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
			})
			c.Abort()
			return
		}

		c.Next()

	}
}
