package handlers

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

const (
	maxUploadSize = 10 << 20
	uploadDir     = "./uploads"
)

func ServeIndexHTMLHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	http.ServeFile(w, r, "index.html")
}
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "This method in not allowed: "+err.Error(), http.StatusBadRequest)
		return
	}
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Error retrieving file form: "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "File reading error: "+err.Error(), http.StatusInternalServerError)
		log.Printf("File reading error: %v\n", err)
		return
	}
	fileContent := string(fileBytes)
	convertedContent, err := service.ProcessData(fileContent)
	if err != nil {
		http.Error(w, "File processing error: "+err.Error(), http.StatusInternalServerError)
		log.Printf("File processing error: %v\n", err)
		return
	}
	if err := os.MkdirAll(uploadDir, 0755); err != nil {
		http.Error(w, "Error creating upload dir: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error creating upload dir: %v\n", uploadDir, err)
		return
	}
	fileExtension := filepath.Ext(handler.Filename)
	newFileName := fmt.Sprintf("converted_%s%s", time.Now().UTC().Format("20060102150405"), fileExtension)
	outputPath := filepath.Join(uploadDir, newFileName)

	outputFile, err := os.Create(outputPath)
	if err != nil {
		http.Error(w, "Error creating output file: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Error creating output file: %v\n", err)
		return
	}
	defer outputFile.Close()
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf("The file has been successfully processed and saved as: %s\n", newFileName)))
	_, _ = w.Write([]byte("Conversion result:\n"))
	_, _ = w.Write([]byte(convertedContent))
}
