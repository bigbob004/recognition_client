package main

import (
	"FaceRecognitionClient/internal/database"
	"FaceRecognitionClient/internal/database/users"
	"FaceRecognitionClient/internal/domain"
	authhHandler "FaceRecognitionClient/internal/handler/auth"
	"context"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	db, err := database.NewPGX(ctx, database.DataBaseConfig{Username: "postgres", DbName: "auth", Port: "5050", Password: "qwerty", Host: "localhost"}, 5)
	if err != nil {
		logrus.Fatal("can not connect to database")
	}
	dao := users.NewDao()
	handler := authhHandler.NewHadler(dao, db)
	err = handler.SignIn(ctx, domain.SignInRequest{domain.User{Username: "rus", Password: "qwerty", Name: "Ruslan"}})
	if err != nil {
		logrus.Fatal(err)
	}

}
