package app

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

func TestUploadTest(t *testing.T) {
	// Read file
	assert := assert.New(t)
	path := "C:/Users/youre/Downloads/gopher.png"
	file, err := os.Open(path)
	defer file.Close()

	// Create a form file
	buf := &bytes.Buffer{}
	writer := multipart.NewWriter(buf) // data is stored in the buf
	multi, err := writer.CreateFormFile("upload_file", filepath.Base(path))

	assert.NoError(err)
	io.Copy(multi, file) // data has been loaded to the buffer
	writer.Close()

	res := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/uplaods", buf)
	req.Header.Set("Content-type", writer.FormDataContentType())

	uploadsHandler(res, req) // send data
	assert.Equal(http.StatusOK, res.Code)

	uploadFilePath := "./uploads/" + filepath.Base(path)
	_, err := os.Stat() // Check if the file is there. os.Stat returns file info.
	assert.NoError(err)

	uploadedFile, err := os.Open(uploadFilePath)
	originalFile, err := os.Open(path)
	defer uploadedFile.Close()
	defer originalFile.Close()

	uploadedData := []byte{}
	originalData := []byte{}
	uploadedFile.Read(uploadedData)
	originalData.Read(originalData)

	assert.Equal(uploadedData, uploadedData)