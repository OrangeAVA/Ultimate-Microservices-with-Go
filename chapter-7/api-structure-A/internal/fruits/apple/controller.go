package apple

import "github.com/gin-gonic/gin"

func ProcessRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"color": "red",
		"date":  "2023-12-11",
		"taste": "sweet",
	})
}
