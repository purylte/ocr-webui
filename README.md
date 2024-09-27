# OCR WebUI

OCR WebUI allows users to upload images, crop specific sections, and extract text using Optical Character Recognition (OCR). The project is built in Go, leveraging gosseract for OCR and HTMX for seamless interactions while being lightweight.

# Installation
## Docker
1. Run `docker run -p 3000:3000 ghcr.io/purylte/ocr-webui:latest`
2. Open http://localhost:3000/app
## Local
1. Ensure [Tesseract](https://github.com/tesseract-ocr/tesseract) and [Leptonica](https://github.com/DanBloomberg/leptonica) is installed
2. Add required languages by placing traineddata file in your tesseract installation.
3. Run `./ocr-webui`
4. Open http://localhost:3000/app 

# Development
## Using Dev Container (VS Code)
1. Ensure [Docker](https://www.docker.com/) and [Dev Containers](https://marketplace.visualstudio.com/items?itemName=ms-vscode-remote.remote-containers) extension is installed
2. Open this project in VS Code
```bash
git clone https://github.com/purylte/ocr-webui.git
code ocr-webui
```
3. Run "Dev Containers: Reopen in Container" in VS Code
4. Run `air` to start hot reload
## Manually
1. Clone the repository
```bash
git clone https://github.com/purylte/ocr-webui.git
cd ocr-webui
```

2. Install [Tesseract](https://github.com/tesseract-ocr/tesseract) and [Leptonica](https://github.com/DanBloomberg/leptonica)

3. Install the required Go tools:
```bash
go install github.com/a-h/templ/cmd/templ@latest
go install github.com/air-verse/air@latest```
```
4. Run `air` to start hot reload

# Todo
1. Preprocess image before doing OCR using [gocv](https://github.com/hybridgroup/gocv)
3. Test
2. Better logging & error handling

# Contributing

Feel free to fork this project, submit issues, and create pull requests. Contributions are welcome!

# License

This project is licensed under the MIT License - see the LICENSE file for details.