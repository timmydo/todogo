

.PHONY: all
all: run

.PHONY: deps
deps:
	go get -u github.com/gorilla/mux
	go get -u github.com/go-bindata/go-bindata/...

.PHONY: build-html
build-html:
	go-bindata -o todo/content.go content/...

.PHONY: build
build: build-html
	go build -o server.exe ./todo

.PHONY: run
run: build
	./server.exe -port 8080