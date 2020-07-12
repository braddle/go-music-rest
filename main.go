package main

import (
	"log"
	"os"

	"github.com/braddle/go-http-template/app"
	"github.com/gorilla/mux"
)

func main() {
	a := app.New(mux.NewRouter())
	log.Fatal(a.Run(":" + os.Getenv("PORT")))
}
