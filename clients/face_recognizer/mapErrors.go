package face_recognizer

import (
	"FaceRecognitionClient/internal/domain"
	"google.golang.org/grpc/status"
)

// TODO Пока что будем опираться на строку ошибки от внешнего сервиса, позже можно будет сделать errors.proto
const (
	AlreadyExistFaceError string = "Данное лицо уже загружен в систему. Попробуйте загрузить другое лицо"
	CanNotDetectFaceError string = "Не удалось обнаружить лицо на изображении"
	FaceNotFound          string = "Не удалось найти лицо в системе, попробуйте выбрать другую фотографию или добавьте изображение в систему"
	InvalidCountOfFaces   string = "На изображении должно быть только одно лицо"
)

func mapError(clientError error) domain.MyError {
	if e, ok := status.FromError(clientError); ok {
		switch e.Message() {
		case AlreadyExistFaceError:
			return domain.AlreadyExistFaceError
		case CanNotDetectFaceError:
			return domain.CanNotDetectFaceError
		case FaceNotFound:
			return domain.FaceNotFound
		case InvalidCountOfFaces:
			return domain.InvalidCountOfFaces
		default:
			return domain.InternalError
		}
	} else {
		return domain.InternalError
	}
}
