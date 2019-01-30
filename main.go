package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	fileToDownload string
)

func init() {
	flag.StringVar(&fileToDownload, "file-to-download", "", "file to download")
	flag.Parse()
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, fileToDownload)
	})
	log.Print("start http server")
	exitOnError(http.ListenAndServe(":8080", nil))
}

func exitOnError(err error) {
	if err == nil {
		return
	}

	fmt.Printf("exit with code 1: %v\n", err)
	os.Exit(1)
}
