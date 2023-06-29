package app

import (
	"FaceRecognitionClient/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *App) signUp(c *gin.Context) {
	var request domain.SignUpRequest

	err := c.BindJSON(&request)
	if err != nil {
		newErrorResponse(c, err)
		return
	}

	err = a.authHandler.SignUp(c, request)
	if err != nil {
		newErrorResponse(c, err)
		return
	}
	c.JSON(http.StatusOK, struct{}{})
	return
}
