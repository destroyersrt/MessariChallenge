package main

import (
	"log"
	"messarichallenge/app"
	"net/http"
	"os"
)

func check(e error) {
	if e != nil {
		log.Println(e)
		os.Exit(1)
	}
}

func main() {

	app := app.New()

	http.HandleFunc("/", app.Router.ServeHTTP)

	err := http.ListenAndServe(":3000", nil)
	check(err)

}
