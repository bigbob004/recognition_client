package app

import (
	"FaceRecognitionClient/internal/domain"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *App) train(ctx *gin.Context) {
	fileData, err := uploadAndValidate(ctx)
	if err != nil {
		newErrorResponse(ctx, err)
		return
	}
	person_name := ctx.Request.Form.Get("person_name")
	if len(person_name) == 0 {
		newErrorResponse(ctx, errors.New("Не задано имя человека на фото!"))
		return
	}
	//TODO убрать заглушки в виде пустой строки имени юзера
	err = a.recognitionHandler.Train(ctx, domain.TrainRequest{Face: fileData, PersonName: person_name})
	if err != nil {
		newErrorResponse(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, struct{}{})
	return
}
