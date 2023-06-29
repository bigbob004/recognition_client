package recognition

import (
	"FaceRecognitionClient/internal/domain"
	"context"
)

func (h Handler) Recognize(ctx context.Context, request domain.RecognizeRequest) (domain.RecognizeResponse, error) {
	resp, err := h.faceRecognizerAPI.Recognize(ctx, request)
	if err != nil {
		return domain.RecognizeResponse{}, err
	}
	if resp.PersonName == "" {
		return domain.RecognizeResponse{}, domain.FaceNotFound
	}
	return resp, nil
}
