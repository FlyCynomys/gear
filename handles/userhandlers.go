package handles

import (
	"github.com/FlyCynomys/tools/log"
	"github.com/gin-gonic/gin"
)

func HandleUserGet(c *gin.Context) {
	data, _ := c.Cookie(forkcookie)
	log.Info(data)
	c.JSON(200, data)
}
