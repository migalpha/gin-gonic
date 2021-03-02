package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"gin-gonic/pkg/setting"
)

// GetPage return the page number depends of pagination
func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	if page > 0 {
		result = (page - 1) * setting.PageSize
	}

	return result
}
