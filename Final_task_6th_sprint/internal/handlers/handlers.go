package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// MainHandle функция подгружает .html файл
func MainHandle(w http.ResponseWriter, r *http.Request) {
	filePath, err := filepath.Abs(filepath.Join("..", "index.html"))
	if err != nil {
		http.Error(w, "error in file location path", http.StatusInternalServerError)
		return
	}
	/*data, err := os.ReadFile(filePath)
	if err != nil {
		http.Error(w, "error loading page", http.StatusInternalServerError)
		return
	}*/
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	//w.Write(data)
	http.ServeFile(w, r, filePath)
}

// UploadHandle функция, которая принимает тестовый файл и конвертирует текст в Морзе(и наоборот)
func UploadHandle(w http.ResponseWriter, req *http.Request) {
	//парсинг файла
	if err := req.ParseMultipartForm(10 << 20); err != nil {
		http.Error(w, "Parsing error", http.StatusInternalServerError)
		return
	}

	//Обработка файла
	file, header, err := req.FormFile("myFile")
	if err != nil {
		http.Error(w, "Getting file error", http.StatusInternalServerError)
		return
	}

	//Чтение файла
	fileReaded, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Reading file error", http.StatusInternalServerError)
		return
	}
	//Закрытие файла после чтения
	defer file.Close()

	//Конвертация файла
	convertedStr, err := service.Convert(string(fileReaded))
	if err != nil {
		http.Error(w, "Convertaing file error", http.StatusInternalServerError)
		return
	}

	//Создания локального файла
	ext := filepath.Ext(header.Filename)

	time := time.Now().UTC().Format("20060102T150405Z")
	filename := time + ext

	outputFile, err := os.Create(filename)
	if err != nil {
		http.Error(w, "Creating file error", http.StatusInternalServerError)
		return
	}
	//Закрытие созданного файла после чтения
	defer outputFile.Close()

	//Запись конвертированных данных в созданный локальный файл
	if _, err = outputFile.WriteString(convertedStr); err != nil {
		http.Error(w, "Writting file error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-type", "text/plain; charset=utf-8")
	w.Write([]byte(convertedStr))
}
