package main

import (
	recognition "FaceRecognitionClient"
	"FaceRecognitionClient/clients/face_recognizer"
	app "FaceRecognitionClient/internal/app"
	"FaceRecognitionClient/internal/database"
	"FaceRecognitionClient/internal/database/users"
	"FaceRecognitionClient/internal/handler/auth"
	handler "FaceRecognitionClient/internal/handler/recognition"
	"context"
	"github.com/sirupsen/logrus"
)

//TODO подумать над тем, чтобы хранить фото (БУДУЩЕЕ)
//TODO подумать над разрешением фото

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	ctx := context.Background()

	db, err := database.NewPGX(ctx, database.DataBaseConfig{Username: "postgres", DbName: "auth", Port: "5050", Password: "qwerty", Host: "localhost"}, 5)
	if err != nil {
		logrus.Fatal("can not connect to database")
	}
	dao := users.NewDao()
	authHandler := auth.NewHadler(dao, db)
	client, err := face_recognizer.NewClient()
	if err != nil {
		logrus.Panic("can't create client for face_recognizer", "err", err)
	}
	recognitionHandler := handler.NewHadler(client)
	a := app.NewApp(recognitionHandler, authHandler)
	srv := new(recognition.Server)
	if err = srv.Run("7001", a.InitRoutes()); err != nil {
		logrus.Fatalf("error occcured while rinning http server: %s", err.Error())
	}
}

//TODO создать конфиг (туда положим хосты внешних сервисов, данные бля бд)!

//TODO поставить таймаут на ручку распознавания - 10 сек!
//TODO сделать даунскэйл изображения (не более 500КБ)
