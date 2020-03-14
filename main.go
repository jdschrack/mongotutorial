package main

import (
	Data "github.com/jdschrack/mongotutorial/data"
	Http "github.com/jdschrack/mongotutorial/http"
	"log"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.Println("Starting Application")

	log.Print("Setting Up Resources")
	Data.Configure()
	Http.ConfigureHttp()
}
