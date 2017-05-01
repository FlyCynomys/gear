package handles

import (
	"net/http"

	"fmt"

	"github.com/FlyCynomys/gear/service"
	"github.com/FlyCynomys/tools/log"
	"github.com/gin-gonic/gin"
)

func HandleLoginGet(c *gin.Context) {
	c.JSON(http.StatusOK, "login")
	return
}

func HandleLoginPost(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	if account == "" {
		c.JSON(http.StatusOK, "account is empty")
		return
	}
	if password == "" {
		c.JSON(http.StatusOK, "password is empty")
		return
	}
	a := new(service.AuthService)
	result := a.Login(account, password)
	if result.Status != 1 {
		c.JSON(http.StatusOK, result)
		return
	}
	c.SetCookie(forkcookie, fmt.Sprintf("%d", result.Data), 86400, "", "locahost", false, true)
	target := fmt.Sprintf("/people/%d", result.Data)
	c.Redirect(302, target)
	return
}

func HandleLogoutGet(c *gin.Context) {
	data, err := c.Cookie(forkcookie)
	if err != nil {
		c.JSON(http.StatusOK, err)
		return
	}
	log.Debug(data)
	//todo : delete cookie
}

func HandleRegisterGet(c *gin.Context) {
	c.JSON(http.StatusOK, "register")
}

func HandleRegisterPost(c *gin.Context) {
	account := c.PostForm("account")
	password := c.PostForm("password")
	nickname := c.PostForm("nickname")
	if account == "" {
		c.JSON(http.StatusOK, "account is empty")
		return
	}
	if password == "" {
		c.JSON(http.StatusOK, "password is empty")
		return
	}
	if nickname == "" {
		c.JSON(http.StatusOK, "need nickname")
		return
	}
	a := new(service.AuthService)
	result := a.Register(account, password, nickname)
	if result.Status != 1 {
		c.JSON(http.StatusOK, result)
		return
	}
	log.Debug(result)
	c.SetCookie(forkcookie, fmt.Sprintf("%d", result.Data), 86400, "", "localhost", false, true)
	target := fmt.Sprintf("/people/%d", result.Data)
	c.Redirect(302, target)
	return
}
