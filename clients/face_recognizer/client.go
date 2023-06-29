package face_recognizer

import (
	"FaceRecognitionClient/clients/pb/face_recognizer"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const clientName = "face_recognizer"
const recognizerHost = "localhost:6066"

type Client struct {
	grpc face_recognizer.FaceRecognizerClient
	conn *grpc.ClientConn
}

func NewClient() (*Client, error) {
	cc, err := grpc.Dial(recognizerHost, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Errorf("/clients/face_recognizer can not create client for: %s", clientName, "err", err)
		return nil, err
	}
	return &Client{
		conn: cc,
		grpc: face_recognizer.NewFaceRecognizerClient(cc),
	}, nil
}
