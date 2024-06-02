package main

import (
	"context"
	"log"
	"os"

	"roadside-hotel/templates"
)

func main() {
	f, err := os.Create("build/index.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = templates.Menu().Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}
