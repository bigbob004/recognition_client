package app

import (
	"FaceRecognitionClient/internal/domain"
	"FaceRecognitionClient/internal/handler/auth"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *App) signIn(c *gin.Context) {
	var request domain.SignInRequest

	err := c.BindJSON(&request)
	if err != nil {
		newErrorResponse(c, err)
		return
	}

	//request = getParams(c)
	token, err := a.authHandler.SignIn(c, request)
	if err != nil {
		newErrorResponse(c, err)
		return
	}
	//TODO Если очень хочется, можно зашифровать куку симметричным шифрованием

	//TODO вынести в конфиг название cookie
	c.SetCookie("token", fmt.Sprintf("Bearer %s", token), int(auth.TokenTTL.Seconds()), "/", "localhost", false, true)
	c.JSON(http.StatusOK, struct{}{})

	return
}
