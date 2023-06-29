package face_recognizer

import (
	"FaceRecognitionClient/clients/pb/face_recognizer"
	"FaceRecognitionClient/internal/domain"
	"context"
	"github.com/sirupsen/logrus"
)

func (c Client) Recognize(ctx context.Context, request domain.RecognizeRequest) (domain.RecognizeResponse, error) {
	resp, err := c.grpc.RecognizeFace(ctx, &face_recognizer.RecognizeFaceRequest{Face: &face_recognizer.Face{Data: request.Face}})
	if err != nil {
		logrus.Errorf("/clients/face_recognizer.c.grpc.RecognizerFace", "err", err)
		return domain.RecognizeResponse{}, mapError(err)
	}
	return mapToDomainResponse(resp), nil
}

func mapToDomainResponse(response *face_recognizer.RecognizeFaceResponse) domain.RecognizeResponse {
	if response.FaceLocation == nil {
		return domain.RecognizeResponse{}
	}

	return domain.RecognizeResponse{
		PersonName: response.PersonName,
		FaceLocation: domain.Location{
			Top:    response.FaceLocation.Top,
			Bottom: response.FaceLocation.Bottom,
			Right:  response.FaceLocation.Right,
			Left:   response.FaceLocation.Left,
		},
	}
}
