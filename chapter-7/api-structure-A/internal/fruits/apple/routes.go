package apple

import "github.com/gin-gonic/gin"

var Path = "/fruits/apple"

func RegisterRoutes(group *gin.RouterGroup) {
	group.GET(Path, gin.HandlersChain{ProcessRequest}...)
}
