package main

import (
	"crypto/rand"
	"encoding/hex"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/purylte/ocr-webui/templates"
)

var sessionManager *scs.SessionManager

const tempImageDir = "tmp/img"

func main() {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour

	if err := os.MkdirAll(tempImageDir, os.ModePerm); err != nil {
		log.Fatalf("cannot create temp image dir: %v", err)
	}

	mux := http.NewServeMux()
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(tempImageDir))))
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
		http.Error(w, "parse form failed "+err.Error(), http.StatusBadRequest)
		return
	}

	file, handler, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "fail to retrieve file "+err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()

	// buf := bytes.NewBuffer(nil)
	// if _, err := io.Copy(buf, file); err != nil {
	// 	http.Error(w, "fail to copy to buffer", http.StatusInternalServerError)
	// }
	// imgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	// contentType := handler.Header.Get("Content-Type")
	// dataURL := fmt.Sprintf("data:%s;base64,%s", contentType, imgBase64)
	// templates.Image(dataURL).Render(r.Context(), w)
	randBytes := make([]byte, 4)
	rand.Read(randBytes)
	randFileName := hex.EncodeToString(randBytes) + filepath.Ext(handler.Filename)
	filePath := filepath.Join(tempImageDir, randFileName)
	dst, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "unable to save file "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = file.Seek(0, 0)
	if err != nil {
		http.Error(w, "unable to process file "+err.Error(), http.StatusInternalServerError)
		return
	}
	_, err = dst.ReadFrom(file)
	if err != nil {
		http.Error(w, "unable to save file "+err.Error(), http.StatusInternalServerError)
		return
	}
	templates.Image("static/"+randFileName).Render(r.Context(), w)

}
