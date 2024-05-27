package main

import (
	"context"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hello World")

	component := hello("Abhijit")
	component.Render(context.Background(), os.Stdout)
}
