package main

import (
	"context"
	"fmt"

	"FaceRecognitionClient/internal/database"
	"FaceRecognitionClient/internal/database/users"
	"FaceRecognitionClient/internal/domain"
	authhHandler "FaceRecognitionClient/internal/handler/auth"
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
	isExist, err := handler.SignUp(ctx, domain.SignUpRequest{Username: "rus", Password: "qwerty"})
	if err != nil {
		logrus.Fatal(err)
	}
	if isExist {
		fmt.Print("user rus is founded")
	} else {
		fmt.Print("user rus is not founded")
	}
}
