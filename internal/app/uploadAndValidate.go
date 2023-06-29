package app

import (
	"FaceRecognitionClient/internal/domain"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"net/http"
)

// Max file size for upload is 4 MB
const MAX_FILE_UPLOAD_SIZE = 4 << 20

// .png and .jpef type
const PNG_TYPE = "image/png"
const JPEG_TYPE = "image/jpeg"

func uploadAndValidate(c *gin.Context) ([]byte, error) {
	file, fileHeader, err := c.Request.FormFile("image")
	if err != nil {
		logrus.Errorf("app/uploadAndValidate/can not upload file: %s", err)
		return nil, errors.New("Can't upload file")
	}
	if fileHeader.Size >= MAX_FILE_UPLOAD_SIZE {
		logrus.Errorf("app/uploadAndValidate/maximum file size is %d, but current file is: %d", MAX_FILE_UPLOAD_SIZE, fileHeader.Size)
		return nil, domain.InvalidSizeOfImageError
	}
	var fileData []byte
	fileData, err = ioutil.ReadAll(file)
	if err != nil {
		logrus.Errorf("app/uploadAndValidate/err %s", err)
		return nil, errors.New("Can't upload file")
	}
	if fileType := http.DetectContentType(fileData); fileType != JPEG_TYPE && fileType != PNG_TYPE {
		logrus.Errorf("app/uploadAndValidate/incorrect file type, file type: %s", fileType)
		return nil, domain.InvalidFileTypeError
	}
	config, tip, err := image.DecodeConfig(file)
	fmt.Print(config, tip)
	return fileData, nil
}
