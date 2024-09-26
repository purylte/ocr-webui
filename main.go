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
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/a-h/templ"
	"github.com/alexedwards/scs/v2"
	"github.com/otiai10/gosseract/v2"
	"github.com/purylte/ocr-webui/cleaner"
	"github.com/purylte/ocr-webui/services"
	"github.com/purylte/ocr-webui/stores"
	"github.com/purylte/ocr-webui/templates"
	"github.com/purylte/ocr-webui/types"

	_ "image/png"
)

var imageService *services.ImageService
var sessionService *services.SessionService
var ocrService *services.OCRService

var tempDir string

func main() {
	gob.Register([]*types.ImageData{})
	gob.Register(types.ImageData{})

	sessionManager := scs.New()
	sessionManager.Lifetime = 24 * time.Hour
	sessionManager.IdleTimeout = 20 * time.Minute

	sessionService = services.NewSessionService(*sessionManager)

	imageService = services.NewImageService(*sessionManager)

	ocrClientStore := stores.NewOCRClientStore()
	ocrService = services.NewOCRService(ocrClientStore)
	cleaner.NewOCRClientCleaner(ocrClientStore, 5*time.Minute, 30*time.Minute).Start()

	var err error
	tempDir, err = initTempDir()
	if err != nil {
		log.Fatalf("cannot create temp image dir: %v", err)
	}

	imgCleaner := cleaner.NewFSCleaner(tempDir, 20*time.Minute, 1*time.Hour)
	imgCleaner.Start()

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
	img, err := imageService.GetCurrentImage(r.Context())
	var component templ.Component
	if err != nil {
		component = templates.MainLayout(nil)
	} else {
		component = templates.MainLayout(img)
	}
	component.Render(r.Context(), w)
}

func protectImageHandler(w http.ResponseWriter, r *http.Request) {
	imageName := r.URL.Path[len("/img/"):]
	if !imageService.ImageIsAllowed(r.Context(), imageName) {
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

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "Unable to decode image: "+err.Error(), http.StatusInternalServerError)
		return
	}
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	image := NewImage(handler.Filename, width, height)

	// Reset file pointer to beginning
	if _, err := file.Seek(0, 0); err != nil {
		http.Error(w, "Unable to process file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	dst, err := os.Create(image.FilePath)
	if err != nil {
		http.Error(w, "Unable to create file: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()
	if _, err := io.Copy(dst, file); err != nil {
		http.Error(w, "Unable to save file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	if err := imageService.AddAllowedImage(r.Context(), image); err != nil {
		http.Error(w, "Unable to save to session: "+err.Error(), http.StatusInternalServerError)
		return
	}

	imageService.SetCurrentImage(r.Context(), image)
	templates.CanvasImage(image).Render(r.Context(), w)

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

	ao, errA := strconv.Atoi(r.FormValue("a"))
	bo, errB := strconv.Atoi(r.FormValue("b"))
	xo, errX := strconv.Atoi(r.FormValue("x"))
	yo, errY := strconv.Atoi(r.FormValue("y"))
	width, errW := strconv.Atoi(r.FormValue("width"))
	height, errH := strconv.Atoi(r.FormValue("height"))

	if errA != nil || errB != nil || errX != nil || errY != nil || errW != nil || errH != nil {
		http.Error(w, "invalid form values", http.StatusBadRequest)
		return
	}

	sessionId := sessionService.GetOrGenerateId(r.Context())

	imageData, err := imageService.GetCurrentImage(r.Context())
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

	img, _, err := image.Decode(file)
	if err != nil {
		http.Error(w, "unable to decode "+err.Error(), http.StatusInternalServerError)
		return
	}
	points := unscalePoints(width, height, imageData.Width, imageData.Height, image.Point{X: ao, Y: bo}, image.Point{X: xo, Y: yo})
	croppedImage := cropImage(img, *points[0], *points[1])

	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, croppedImage, nil); err != nil {
		http.Error(w, "unable to encode image "+err.Error(), http.StatusInternalServerError)
		return
	}
	// ppImage, err := preprocessImage(buf.Bytes())
	// if err != nil {
	// 	http.Error(w, "pre process image failed "+err.Error(), http.StatusInternalServerError)
	// 	return
	// }

	client := gosseract.NewClient()
	defer client.Close()

	text, err := ocrService.OcrFromBytes(sessionId, buf.Bytes())
	if err != nil {
		http.Error(w, "OCR failed "+err.Error(), http.StatusInternalServerError)
		return
	}

	imgBase64 := base64.StdEncoding.EncodeToString(buf.Bytes())
	imgSrc := "data:image/jpeg;base64," + imgBase64
	templates.TextResult(imgSrc, text).Render(r.Context(), w)
}

func unscalePoints(width int, height int, originalWidth int, originalHeight int, points ...image.Point) []*image.Point {
	xScale := float64(originalWidth) / float64(width)
	yScale := float64(originalHeight) / float64(height)
	res := make([]*image.Point, len(points))

	for i, p := range points {
		res[i] = &image.Point{
			int(float64(p.X) * xScale),
			int(float64(p.Y) * yScale),
		}
	}
	return res
}

func cropImage(src image.Image, a, b image.Point) image.Image {
	rect := image.Rect(a.X, a.Y, b.X, b.Y)
	cropped := image.NewRGBA(rect)
	draw.Draw(cropped, rect, src, image.Point{a.X, a.Y}, draw.Src)

	return cropped
}

func NewImage(fileName string, width int, height int) *types.ImageData {
	randBytes := make([]byte, 8)
	rand.Read(randBytes)
	name := hex.EncodeToString(randBytes) + filepath.Ext(fileName)

	return &types.ImageData{
		OriginalName: fileName,
		Name:         name,
		FilePath:     filepath.Join(tempDir, name),
		WebPath:      "img/" + name,
		Width:        width,
		Height:       height,
	}

}
