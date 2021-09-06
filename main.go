package main

import (
	"fmt"
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

	fmt.Println("Server startig at  3000")

	err := http.ListenAndServe(":3000", nil)
	check(err)

}
