package main

import (
	"FaceRecognitionClient/clients/face_recognizer"
	"FaceRecognitionClient/internal/domain"
	handler "FaceRecognitionClient/internal/handler/recognition"
	"context"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	client, err := face_recognizer.NewClient()
	if err != nil {
		logrus.Panic("can't create client for face_recognizer", "err", err)
	}
	handler := handler.NewHadler(client)
	content, err := os.ReadFile("images/image.png")
	if err != nil {
		logrus.Panic("can't read file content", "err", err)
	}
	ctx := context.Background()
	start := time.Now()
	_, err = handler.Recognize(ctx, domain.RecognizeRequest{Face: content})
	dur := time.Since(start)
	if err != nil {
		logrus.Error("something went wrong", "err", err)
	}
	logrus.Println("time of recognition", "time", dur)
}
