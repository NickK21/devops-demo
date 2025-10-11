package main

import (
	"os"
	"time"
	"log"

	"github.com/gofiber/fiber/v2"
)

type RootResponse struct {
	Message	string `json:"message"`
	Timestamp	int64 `json:"timestamp"` 
}

func handleRoot(c *fiber.Ctx) error {
	var response RootResponse
		response.Message = "My name is Nick Kaplan"
		response.Timestamp = time.Now().UnixMilli()
		c.Status(fiber.StatusOK)
	return c.JSON(response)
}

func main() {
	var app *fiber.App = fiber.New()

	app.Get("/", handleRoot)

	var port string =  os.Getenv("PORT")
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