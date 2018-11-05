package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

var (
	file = flag.String("f", "", "the path to the file to outline")
	port = flag.Int("port", 8080, "port to listen on")
)

func hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	assetData := MustAsset("content/index.html")
	w.Write(assetData)
}

func main() {
	flag.Parse()
	http.HandleFunc("/", hello)
	addr := fmt.Sprintf("127.0.0.1:%d", *port)
	log.Printf("Listening on %s\n", addr)

	r := mux.NewRouter()
	r.HandleFunc("/", hello).Methods("GET")

	srv := &http.Server{
		Handler:      r,
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
