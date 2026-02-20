package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"lr1/lr1/internal/client"
)

func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment")
	}
}

func main() {
	response, err := client.Get()
	if err != nil {
		log.Fatalf("Request failed: %v", err)
	}

	fmt.Println("Success!")
	fmt.Printf("Value:  %d\n", response.WorkResult.Value)
	fmt.Printf("Binary: %s\n", response.WorkResult.StrMessage)
}
