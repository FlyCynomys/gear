package handles

import (
	"github.com/gin-gonic/gin"
)

func Handle(c *gin.Context) {
	c.Writer.WriteString("hot plan")
}
