package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/ID-Samurai/go-sprint6-final/internal/service"
)

func MainHandler(res http.ResponseWriter, req *http.Request) {
	wd, err := os.Getwd()
	if err != nil {
		http.Error(res, "Repository definition error", http.StatusInternalServerError)
		return
	}

	fileName := filepath.Join(wd, "index.html")
	fmt.Printf(fileName)

	data, err := os.ReadFile(fileName)
	if err != nil {
		http.Error(res, "File reading error index.html", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html")
	res.WriteHeader(http.StatusOK)
	res.Write(data)
}
func UploadHandler(res http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		http.Error(res, "Form parsing error", http.StatusInternalServerError)
		return
	}
	file, header, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "File not found", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		http.Error(res, "File reading error", http.StatusInternalServerError)
		return
	}
	text, err := service.AutoConversion(string(content))
	if err != nil {
		http.Error(res, "File decoding error", http.StatusInternalServerError)
		return
	}
	fileName := time.Now().UTC().Format("15:04:05 02-01-2006")
	fileExt := filepath.Ext(header.Filename)
	newFileName := fmt.Sprintf("Decoded_file_%s%s", fileName, fileExt)
	newFile, err := os.Create(newFileName)
	if err != nil {
		http.Error(res, "Error creating file", http.StatusInternalServerError)
		return
	}
	defer newFile.Close()

	if _, err := newFile.WriteString(text); err != nil {
		http.Error(res, "Error recording", http.StatusInternalServerError)
		return
	}

	res.WriteHeader(http.StatusOK)

	fmt.Fprintf(res, text)

}
