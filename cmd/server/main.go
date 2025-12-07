package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	const (
		docsDir = "docs"
		addr    = ":8080"
	)

	if _, err := os.Stat(docsDir); err != nil {
		log.Fatalf("missing docs assets: %v", err)
	}

	http.Handle("/", http.FileServer(http.Dir(docsDir)))

	log.Printf("listening on http://localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
