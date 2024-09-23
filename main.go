package main

import (
	"bytes"
	"crypto/rand"
	"encoding/base64"
	"encoding/gob"
	"encoding/hex"
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/otiai10/gosseract/v2"
	"github.com/purylte/ocr-webui/templates"

	_ "image/png"
)

var sessionManager *scs.SessionManager
var tempDir string

func main() {
	sessionManager = scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	gob.Register([]*ImageData{})
	gob.Register(ImageData{})

	var err error
	tempDir, err = initTempDir()
	if err != nil {
		log.Fatalf("cannot create temp image dir: %v", err)
	}

	mux := http.NewServeMux()
	mux.HandleFunc("/img/", protectImageHandler)
	mux.HandleFunc("/app", appHandler)
	mux.HandleFunc("/upload", uploadHandler)
	mux.HandleFunc("/crop", cropHandler)
	log.Fatal(http.ListenAndServe(":3000", sessionManager.LoadAndSave(mux)))

}

func initTempDir() (string, error) {
	tempBase := os.TempDir()
	appTempDir := filepath.Join(tempBase, "ocr-img")

	if err := os.MkdirAll(appTempDir, 0700); err != nil {
		return "", fmt.Errorf("failed to create app temp directory: %w", err)
	}

	testFile, err := os.CreateTemp(appTempDir, "write-test-")
	if err != nil {
		return "", fmt.Errorf("app temp directory is not writable: %w", err)
	}

	testFile.Close()
	os.Remove(testFile.Name())

	return appTempDir, nil
}

func appHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.MainLayout()
	component.Render(r.Context(), w)
}

func protectImageHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Path[len("/img/"):]
	if !canAccessImage(r.Context(), imageName) {
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	}
	http.ServeFile(w, r, tempDir+"/"+imageName)
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

	image := NewImage(handler.Filename, handler.Header.Get("Content-Type"))

	dst, err := os.Create(image.FilePath)
	if err != nil {
		http.Error(w, "unable to create file "+err.Error(), http.StatusInternalServerError)
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

	if err = putAllowedImage(r.Context(), image); err != nil {
		http.Error(w, "unable to save to session "+err.Error(), http.StatusInternalServerError)
		return
	}
	setCurrentImage(r.Context(), image)

	templates.Image(image.WebPath).Render(r.Context(), w)

}

func cropHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "expected post", http.StatusMethodNotAllowed)
		return
	}
	if err := r.ParseForm(); err != nil {
		http.Error(w, "error parsing form "+err.Error(), http.StatusBadRequest)
		return
	}

	a, errA := strconv.Atoi(r.FormValue("a"))
	b, errB := strconv.Atoi(r.FormValue("b"))
	x, errX := strconv.Atoi(r.FormValue("x"))
	y, errY := strconv.Atoi(r.FormValue("y"))

	if errA != nil || errB != nil || errX != nil || errY != nil {
		http.Error(w, "invalid coordinate values", http.StatusBadRequest)
		return
	}

	imageData, err := getCurrentImage(r.Context())
	if err != nil {
		http.Error(w, "unable to get current image "+err.Error(), http.StatusInternalServerError)
		return
	}
	file, err := os.Open(imageData.FilePath)
	if err != nil {
		http.Error(w, "unable to open image "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer file.Close()

	image, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "unable to decode "+err.Error(), http.StatusInternalServerError)
		return
	}
	croppedImage := cropImage(image, a, b, x, y)

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, croppedImage, nil); err != nil {
		http.Error(w, "unable to encode image "+err.Error(), http.StatusInternalServerError)
		return
	}

	client := gosseract.NewClient()
	defer client.Close()

	text, err := ocrFromBytes(buf.Bytes())
	if err != nil {
		http.Error(w, "OCR failed "+err.Error(), http.StatusInternalServerError)
		return
	}

	imgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	imgSrc := "data:image/jpeg;base64," + imgBase64
	templates.TextResult(imgSrc, text).Render(r.Context(), w)
}

func cropImage(src image.Image, a, b, x, y int) image.Image {
	rect := image.Rect(a, b, x, y)
	cropped := image.NewRGBA(rect)
	draw.Draw(cropped, rect, src, image.Point{a, b}, draw.Src)

	return cropped
}

func ocrFromBytes(imageByte []byte) (string, error) {
	client := gosseract.NewClient()
	defer client.Close()

	if err := client.SetImageFromBytes(imageByte); err != nil {
		return "", err
	}

	text, err := client.Text()
	if err != nil {
		return "", err
	}
	return text, nil
}

type ImageData struct {
	OriginalName string
	Name         string
	ContentType  string
	FilePath     string
	WebPath      string
}

func NewImage(fileName string, contentType string) *ImageData {
	randBytes := make([]byte, 4)
	rand.Read(randBytes)
	name := hex.EncodeToString(randBytes) + filepath.Ext(fileName)

	return &ImageData{
		OriginalName: fileName,
		Name:         name,
		FilePath:     filepath.Join(tempDir, name),
		WebPath:      "img/" + name,
		ContentType:  contentType,
	}

}
