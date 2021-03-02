package v2

import (
	"gin-gonic/models/sql"
	"gin-gonic/pkg/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Ping health check endpoint
// @Summary health check endpoint
// @Description Return "pong" if API is alive
// @Accept  json
// @Produce  json
// @Router /ping [get]
// @Success 200 {object} sql.JSONResult
func Ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "pong",
	})
}

// Login allow to get credentials
// @Summary allow to get credentials
// @Description allow to get credentials
// @Accept  json
// @Produce  json
// @Router /login [post]
// @Success 200 {object} sql.JSONResult
func Login(c *gin.Context) {
	var user sql.User = sql.User{
		ID:        1,
		FirstName: "Eren",
		LastName:  "Jaeger",
		Email:     "eren@jaeger.com",
	}

	token, err := util.GenerateToken(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "Token generating error",
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"msg": token,
	})
}

// SecurePing health check endpoint
// @Summary health check endpoint
// @Description Return "pong" if API is alive
// @Accept  json
// @Produce  json
// @Router /ping [get]
// @Success 200 {object} sql.JSONResult
func SecurePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msg": "secure pong",
	})
}
