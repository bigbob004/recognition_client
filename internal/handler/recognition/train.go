package recognition

import (
	"FaceRecognitionClient/internal/domain"
	"context"
)

func (h Handler) Train(ctx context.Context, request domain.TrainRequest) error {
	err := h.faceRecognizerAPI.Train(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
