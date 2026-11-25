package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "index.html")
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "error while opening file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "error while reading file", http.StatusInternalServerError)
		return
	}

	result, err := service.CheckString(string(data))
	if err != nil {
		http.Error(w, "wrong data in file", http.StatusInternalServerError)
		return
	}

	filename := time.Now().UTC().String()

	resultFile, err := os.Create(strings.ReplaceAll(filename, ":", "_") + filepath.Ext(header.Filename))
	if err != nil {
		http.Error(w, "error while creating result file", http.StatusInternalServerError)
		return
	}
	defer resultFile.Close()

	_, err = io.WriteString(resultFile, result)
	if err != nil {
		http.Error(w, "error while writing in the result file", http.StatusInternalServerError)
		return
	}

	_, err = io.WriteString(w, result)
	if err != nil {
		http.Error(w, "error while sending the result", http.StatusInternalServerError)
		return
	}
}
