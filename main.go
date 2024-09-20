package main

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/purylte/ocr-webui/templates"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/upload", uploadHandler)
	log.Fatal(http.ListenAndServe(":3000", mux))

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.MainLayout()
	component.Render(r.Context(), w)
}

func uploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "expected post", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // 10MB max file size
	if err != nil {
		http.Error(w, "parse form failed", http.StatusBadRequest)
		return
	}

	file, header, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "fail to retrieve file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		http.Error(w, "fail to copy to buffer", http.StatusInternalServerError)
	}
	imgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	contentType := header.Header.Get("Content-Type")
	dataURL := fmt.Sprintf("data:%s;base64,%s", contentType, imgBase64)
	templates.Image(dataURL).Render(r.Context(), w)
}
