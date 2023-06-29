package face_recognizer

import (
	"FaceRecognitionClient/clients/pb/face_recognizer"
	"FaceRecognitionClient/internal/domain"
	"context"
	"github.com/sirupsen/logrus"
)

func (c Client) Train(ctx context.Context, request domain.TrainRequest) error {
	_, err := c.grpc.Train(ctx, &face_recognizer.TrainRequest{Face: &face_recognizer.Face{Data: request.Face}, PersonName: request.PersonName})
	if err != nil {
		logrus.Errorf("/clients/face_recognizer.c.grpc.Train", "err", err)
		return mapError(err)
	}
	return nil
}
