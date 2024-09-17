dev:
	go run .

templates/*_templ.go: templates/*.templ
	templ generate templates/*.templ

tmp/main: *.go templates/*_templ.go
	go build -o tmp/main

build: tmp/main 

run: build
	tmp/main