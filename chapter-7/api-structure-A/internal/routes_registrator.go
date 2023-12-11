package internal

import (
	"github.com/OrangeAVA/Microservices-with-Go/internal/fruits/apple"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(routerGroup *gin.RouterGroup) {
	apple.RegisterRoutes(routerGroup)
}
