package handles

import (
	"github.com/gin-gonic/gin"
)

func HandleIntro(c *gin.Context) {
	c.Writer.WriteString("intro")
	return
}

func HandleAbout(c *gin.Context) {
	c.Writer.WriteString("about")
	return
}
