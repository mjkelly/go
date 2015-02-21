// This is an even simpler webserver that serves a given directory. No
// authentication. This version uses only
//
// I run this behind an apache reverse proxy that adds digest authentication
// and SSL. I use it to serve email attachments that I want to access remotely.
package main

import (
	"log"
	"net/http"
	"os"
	"path"
)

func main() {
	dir := path.Join(os.Getenv("HOME"), "attachments")
	port := "8080"
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	if len(os.Args) > 2 {
		port = os.Args[2]
	}

	h := http.FileServer(http.Dir(dir))
	http.Handle("/", h)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
