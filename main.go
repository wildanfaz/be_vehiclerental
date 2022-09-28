package main

import (
	"log"
	"os"

	"github.com/wildanfaz/vehicle_rental/src/config"

	_ "github.com/joho/godotenv/autoload"
)

func main() {
	if err := config.Run(os.Args[1:]); err != nil {
		log.Fatal(err)
	}
}
