package recognition

import (
	"FaceRecognitionClient/internal/domain"
	"context"
)

type faceRecognizerAPI interface {
	Recognize(ctx context.Context, request domain.RecognizeRequest) (domain.RecognizeResponse, error)
	Train(ctx context.Context, request domain.TrainRequest) error
}

type Handler struct {
	faceRecognizerAPI faceRecognizerAPI
}

func NewHadler(client faceRecognizerAPI) Handler {
	return Handler{faceRecognizerAPI: client}
}
