package main

import (
	"os"
	"fmt"
	"time"
	"log"

	"github.com/gofiber/fiber/v2"
)

var version = "dev"

func handleRoot(c *fiber.Ctx) error {
		var message string = "My name is Nick Kaplan"
		var ts int64 = time.Now().UnixMilli()

		c.Set(fiber.HeaderContentType, "application/json")

		var body string = fmt.Sprintf(`{"message":"%s","timestamp":%d,"version":"%s"}`, message, ts, version)

		c.Status(fiber.StatusOK)

	// Send raw bytes to avoid string conversions
	return c.Send([]byte(body))
}

func main() {
	var app *fiber.App = fiber.New()

	app.Get("/", handleRoot)

	var port string = os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	var address string = ":" + port
	log.Printf("Starting server on %s ...", address)

	var err error = app.Listen(address)
	if err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}