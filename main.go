package main

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
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
	mux.HandleFunc("/img/", protectImageHandler)
	mux.HandleFunc("/app", appHandler)
	mux.HandleFunc("/upload", uploadHandler)
	log.Fatal(http.ListenAndServe(":3000", sessionManager.LoadAndSave(mux)))

}

func protectImageHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Path[len("/img/"):]
	if !canAccessImage(r.Context(), imageName) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}

	// file, err := os.Open(imagePath)
	// if err != nil {
	// 	http.Error(w, "File not found", http.StatusNotFound)
	// 	return
	// }
	// defer file.Close()
	http.ServeFile(w, r, tempImageDir+"/"+imageName)
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.MainLayout()
	component.Render(r.Context(), w)
}

func canAccessImage(ctx context.Context, imageName string) bool {
	imagesJson := sessionManager.GetString(ctx, "images")
	if imagesJson == "" {
		return false
	}
	var imageNames []string
	json.Unmarshal([]byte(imagesJson), &imageNames)
	for _, name := range imageNames {
		if name == imageName {
			return true
		}
	}
	return false
}

func putAllowedImage(ctx context.Context, imageName string) error {
	imagesJson := sessionManager.GetString(ctx, "images")
	var imageNames []string
	if imagesJson != "" {
		json.Unmarshal([]byte(imagesJson), &imageNames)
	}
	imageNames = append(imageNames, imageName)

	updatedImageJson, err := json.Marshal(imageNames)
	if err != nil {
		return err
	}
	sessionManager.Put(ctx, "images", string(updatedImageJson))
	return nil
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

	if err = putAllowedImage(r.Context(), randFileName); err != nil {
		http.Error(w, "unable to save to session "+err.Error(), http.StatusInternalServerError)
		return
	}

	templates.Image("img/"+randFileName).Render(r.Context(), w)

}
