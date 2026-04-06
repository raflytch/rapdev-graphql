package main

import (
	"log"
	"os"

	"rapdev-graphql/internal/app"

	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	application := app.NewApp()

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	log.Printf("server starting on :%s", port)
	if err := application.Listen(":" + port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
