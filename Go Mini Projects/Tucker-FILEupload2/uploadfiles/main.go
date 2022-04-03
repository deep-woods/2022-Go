package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

var PORT = ":1004"

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	uploadFile, header, err := r.FormFile("upload_file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	defer uploadFile.Close()
	/*
		var ErrMissingFile = errors.New("http: no such file")
		ErrMissingFile is returned by FormFile when the provided file field name is either not present in the request or not a file field.
	*/

	// Create directory
	directory := "./uploads"
	os.MkdirAll(directory, 0777) // 0777 read/write permission

	// Create file path
	filepath := fmt.Sprintf("%s/%s", directory, header.Filename)
	// Create a file
	file, err := os.Create(filepath)
	defer file.Close() // Close the file's handle resource

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, err)
		return
	}

	// Copy file
	io.Copy(file, uploadFile)    // copy file
	w.WriteHeader(http.StatusOK) // check status
	fmt.Fprint(w, filepath)      // show file path
}

func main() {
	http.HandleFunc("/uploads", uploadHandler)
	http.Handle("/", http.FileServer(http.Dir("public")))

	http.ListenAndServe(PORT, nil)
}

/*  Difference between http.Handle and http.HandleFunc?
https://stackoverflow.com/questions/21957455/difference-between-http-handle-and-http-handlefunc
*/
