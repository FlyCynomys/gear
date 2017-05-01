package handles

import (
	"github.com/gin-gonic/gin"
)

func HandleGroup(c *gin.Context) {
	c.Writer.WriteString("group")
}
