package main

import (
	"log"
	"net/http"

	"github.com/purylte/ocr-webui/templates"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)

	log.Fatal(http.ListenAndServe(":3000", mux))

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	component := templates.MainLayout()
	component.Render(r.Context(), w)
}
