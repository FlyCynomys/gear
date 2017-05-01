package handles

import (
	"github.com/FlyCynomys/tools/log"
	"github.com/gin-gonic/gin"
)

func HandleTodoGet(c *gin.Context) {
	c.Writer.WriteString("todo plan")
	log.Info("todo ")
	return
}
