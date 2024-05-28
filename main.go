package main

import (
	"context"
	"log"
	"os"

	"roadside-hotel/templates"
)

func main() {
	// app := fiber.New()

	// app.Get("/", func(c *fiber.Ctx) error {

	// 	// component := hello("Abhijit")
	// 	// component.Render(context.Background(), os.Stdout)
	// 	c.Set("Content-type", "text/html")
	// 	return hello("Abhijit").Render(context.Background(), c.Response().BodyWriter())

	// })
	// app.Get("/", adaptor.HTTPHandler(templ.Handler(templates.Hello("Abhijit"))))
	// app.Get("/hello", adaptor.HTTPHandler(templ.Handler(templates.Hello("Abhijit Kar"))))

	// app.Listen(":3000")

	f, err := os.Create("build/index.html")
	if err != nil {
		log.Fatalf("failed to create output file: %v", err)
	}

	err = templates.Menu("John").Render(context.Background(), f)
	if err != nil {
		log.Fatalf("failed to write output file: %v", err)
	}
}
