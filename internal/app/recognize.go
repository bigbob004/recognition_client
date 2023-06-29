package app

import (
	"FaceRecognitionClient/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (a *App) recognize(ctx *gin.Context) {
	fileData, err := uploadAndValidate(ctx)
	if err != nil {
		newErrorResponse(ctx, err)
		return
	}
	resp, err := a.recognitionHandler.Recognize(ctx, domain.RecognizeRequest{Face: fileData})
	if err != nil {
		newErrorResponse(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	return
}
