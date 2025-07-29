package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"sprint-6-final/internal/service"
	"time"
)

const maxUploadSize = 1024 * 1024 // 1MB

func Init(logger interface{}) {}

// IndexHandler обрабатывает запрос к корневому эндпоинту
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "index.html")
}

// UploadHandler обрабатывает загрузку файла и конвертацию
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Ограничиваем размер файла
	r.Body = http.MaxBytesReader(w, r.Body, maxUploadSize)
	if err := r.ParseMultipartForm(maxUploadSize); err != nil {
		http.Error(w, "File too large", http.StatusBadRequest)
		return
	}

	// Получаем файл из формы
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file upload", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Читаем содержимое файла
	fileBytes, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Failed to read file", http.StatusInternalServerError)
		return
	}

	// Конвертируем содержимое
	content := string(fileBytes)
	result, err := service.AutoConvert(content)
	if err != nil {
		http.Error(w, fmt.Sprintf("Conversion error: %v", err), http.StatusInternalServerError)
		return
	}

	// Создаем выходной файл
	outputFilename := "converted_" + time.Now().UTC().Format("20060102150405") + filepath.Ext(header.Filename)
	if err := os.WriteFile(outputFilename, []byte(result.Result), 0644); err != nil {
		http.Error(w, "Failed to save result", http.StatusInternalServerError)
		return
	}

	// Возвращаем результат
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Conversion successful!\nOriginal: %s\nConverted: %s\nSaved to: %s",
		result.Original, result.Result, outputFilename)
}
