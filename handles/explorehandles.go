package handles

import (
	"github.com/gin-gonic/gin"
)

func HandleExplorePlan(c *gin.Context) {
	c.Writer.WriteString("hot plan")
}

func HandleExploreActive(c *gin.Context) {
	c.Writer.WriteString("hot active")
}
