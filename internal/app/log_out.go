package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) logOut(c *gin.Context) {
	c.SetCookie("token", "", -1, "/", "", false, true)
	c.Redirect(http.StatusFound, location)
	return
}
