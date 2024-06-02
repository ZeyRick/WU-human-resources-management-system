package file

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"time"
)

func SaveFile(r *http.Request) (string, error) {
	// 2mb
	err := r.ParseMultipartForm(1048576 * 2)
	if err != nil {
		return "", err
	}

	// in your case file would be fileupload
	file, _, err := r.FormFile("file")
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Create a buffer to store the header of the file in
	fileHeader := make([]byte, 512)

	// Copy the headers into the FileHeader buffer
	if _, err := file.Read(fileHeader); err != nil {
		return "", err
	}

	fileType := http.DetectContentType(fileHeader)

	// set position back to start.
	if _, err := file.Seek(0, 0); err != nil {
		return "", err
	}

	if fileType != "image/png" && fileType != "image/jpg" && fileType != "image/jpeg" {
		return "", errors.New("Bad file type can only upload file in PNG, JPG, JPEG")
	}
	fileName := fmt.Sprintf("%s.%s", time.Now().UTC().Format("20060102150405"), strings.ReplaceAll(fileType, "image/", ""))
	fileDst := fmt.Sprintf("./public/images/employee/%s", fileName)
	dst, err := os.Create(fileDst)
	if err != nil {
		return fileName, err
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		return fileName, err
	}
	return fileName, nil
}

func GetExcelFile(r *http.Request) (string, *bytes.Buffer, error) {
	err := r.ParseMultipartForm(1048576 * 100)
	if err != nil {
		return "", bytes.NewBuffer(nil), err
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		return "", bytes.NewBuffer(nil), err
	}
	defer file.Close()

	//Create buffer to hold file
	buffer := bytes.NewBuffer(nil)

	//Copy content to buffer
	_, err = io.Copy(buffer, file)

	if err != nil {
		return "", bytes.NewBuffer(nil), err
	}
	return header.Filename, buffer, err
}
