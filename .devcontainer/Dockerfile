FROM mcr.microsoft.com/devcontainers/go:1-1.23-bookworm

RUN apt-get update && \
    apt-get install -y libleptonica-dev libtesseract-dev tesseract-ocr-eng && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV TESSDATA_PREFIX=/usr/share/tesseract-ocr/5/tessdata

USER root
