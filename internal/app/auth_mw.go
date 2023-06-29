package app

import (
	"FaceRecognitionClient/internal/handler/auth"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
	cookieNameWithToken = "token"
	location            = "http://localhost:7001/"
)

//TODO по-хорошему, middleware нужно ставить ну группу эндпоинтов, а не на все пути (включая те, которые формы отдают) и отдавать ошибку, если не нет токена в header, а не в cookie (тобишь cookie должен приходить с каждым запросом)

func userIdentity(c *gin.Context) {
	cookieValue, err := c.Cookie(cookieNameWithToken)
	if err != nil {
		c.Redirect(http.StatusFound, location)
		return
	}

	headerParts := strings.Split(cookieValue, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		c.Redirect(http.StatusFound, location)
		return
	}

	if len(headerParts[1]) == 0 {
		c.Redirect(http.StatusFound, location)
		return
	}

	_, err = auth.ParseToken(headerParts[1])
	if err != nil {
		c.Redirect(http.StatusMovedPermanently, location)
		return
	}
	return
}
