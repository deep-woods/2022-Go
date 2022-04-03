package main

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func UploadTest(t *testing.T) {
	// Read file
	assert := assert.New(t)
	path := "C:/Users/0woo0/Downloads/relaxing-gopher-patrice-schoefolt"
	file, _ := os.Open(path)
	defer file.Close()

	os.RemoveAll("./uploads")

	// Create FormFile
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf)
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))
	assert.NoError(err)

	// Pass in file into the FormFile
	io.Copy(multi, file)
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uploads", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadHandler(res, req)
	assert.Equal(http.StatusOK, res.Code)

	uploadedFilePath := "./uploads/" + filepath.Base(path)
	_, err2 := os.Stat(uploadedFilePath)
	assert.NoError(err2)

	uploadedFile, _ := os.Open(uploadedFilePath)
	originalFile, _ := os.Open(path)
	defer uploadedFile.Close()
	defer originalFile.Close()

	uploadedData := []byte{}
	originalData := []byte{}
	uploadedFile.Read(uploadedData)
	originalFile.Read(originalData)

	assert.Equal(uploadedData, originalData)
}
