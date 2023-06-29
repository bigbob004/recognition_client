package app

import (
	"FaceRecognitionClient/internal/domain"
	"github.com/gin-gonic/gin"
	"net/http"
)

type apiError struct {
	Message string `json:"message"`
	Code    int
}

func newErrorResponse(c *gin.Context, businessError error) {
	apiErr := mapToApiError(c, businessError)
	//TODO мне не нравится, как тут отправляется ошибка (сначал к полю обращаемся, а потом полную стр-ру кидаем)
	c.AbortWithStatusJSON(apiErr.Code, apiErr)
	return
}

//Здесь происходит преобразование бизнесовой ошибки в трпансопртную ошибку

func mapToApiError(c *gin.Context, businessErr error) apiError {
	if c.Err() != nil {
		return apiError{Code: http.StatusInternalServerError, Message: "Превышено время ожидания"}
	}

	switch businessErr {
	case domain.InvalidFileTypeError:
		return apiError{Code: http.StatusBadRequest, Message: businessErr.Error()}
	case domain.InvalidSizeOfImageError:
		return apiError{Code: http.StatusBadRequest, Message: businessErr.Error()}
	case domain.AlreadyExistFaceError:
		return apiError{Code: http.StatusBadRequest, Message: businessErr.Error()}
	case domain.CanNotDetectFaceError:
		return apiError{Code: http.StatusBadRequest, Message: businessErr.Error()}
	case domain.FaceNotFound:
		return apiError{Code: http.StatusBadRequest, Message: businessErr.Error()}
	case domain.InvalidCountOfFaces:
		return apiError{Code: http.StatusBadRequest, Message: businessErr.Error()}
	case domain.UserNameOrPasswordInIncorrectErr:
		return apiError{Code: http.StatusBadRequest, Message: businessErr.Error()}
	case domain.UserNameAlreadyExist:
		return apiError{Code: http.StatusBadRequest, Message: businessErr.Error()}
	default:
		return apiError{Code: http.StatusInternalServerError, Message: "Внутренняя ошибка"}
	}
}

//TODO сделать нормальную обработку ошибок
//TODO завести доменные ошибки для регитсрации и аутентификации
