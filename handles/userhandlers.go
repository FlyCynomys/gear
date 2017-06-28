package handles

import "github.com/gin-gonic/gin"

//"github.com/FlyCynomys/gear/service"

func HandleUserGet(c *gin.Context) {
	userCookie, _ := c.Cookie(forkcookie)

	userid := c.Param("userid")
	if userid == "" {
		c.JSON(200, "uid not input")
		return
	}
	if userCookie != userid || userCookie == "" {
		//get some user info,not the master
		c.JSON(200, "uid not equal")
		return
	}
	c.JSON(200, "hello")
	return
}

func HandleUserCreateTodoPlan(c *gin.Context) {

}

func HandleUserUpdateTodoPlan(c *gin.Context) {

}

func HandleUserDeleteTodoPlan(c *gin.Context) {

}

func HandleUserGetTodoPlanDetail(c *gin.Context) {

}

func HandleUserGetToDoPlanList(c *gin.Context) {

}

func HandleUserGetJoinedGroup(c *gin.Context) {

}
func HandleUserGetContribute(c *gin.Context) {

}
