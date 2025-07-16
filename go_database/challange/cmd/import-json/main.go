package main

import (
	"app/internal/importer"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Usage: import-json <path-to-json-dir>")
	}
	dir := os.Args[1]

	if err := importer.Run(dir); err != nil {
		log.Fatalf("import failed: %v", err)
	}
	log.Println("✅ import completed")
}
