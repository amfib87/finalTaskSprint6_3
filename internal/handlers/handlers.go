package handlers

import (
	"bufio"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandleRoot(res http.ResponseWriter, req *http.Request) {
	abs, err := filepath.Abs("../finalTaskSprint6/index.html")
	if err != nil {
		http.Error(res, "error read file", http.StatusInternalServerError)
		return
	}

	data, err := os.ReadFile(abs)
	if err != nil {
		//http.Error(res, "error read file", http.StatusInternalServerError)
		res.Header().Set("Content-Type", abs)
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write(data)

}

func HandleUpload(res http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(10 << 20)

	file, _, err := req.FormFile("myFile")
	if err != nil {
		http.Error(res, "error open file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	var content string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		content += scanner.Text()
	}

	content, err = service.Conversion(content)
	if err != nil {
		http.Error(res, "error conversion", http.StatusInternalServerError)
		return
	}

	filename := time.Now().Format("02.01.06 15-04-05") + ".txt"
	err = os.WriteFile(filename, []byte(content), 0755)
	if err != nil {
		http.Error(res, "error write local file", http.StatusInternalServerError)
		return
	}

	res.Header().Set("Content-Type", "text/html; charset=utf-8")
	res.WriteHeader(http.StatusOK)
	res.Write([]byte(content))

}
