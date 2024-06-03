package main

import (
	"context"
	"log"
	"os"

	"roadside-hotel/templates"
)

func main() {
	index_page, err := os.Create("build/index.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	scan_me_page, err := os.Create("build/scan-me.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = templates.Menu().Render(context.Background(), index_page)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
	err = templates.QRCode().Render(context.Background(), scan_me_page)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}
