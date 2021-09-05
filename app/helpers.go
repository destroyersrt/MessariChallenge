package app

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func parse(w http.ResponseWriter, r *http.Request, data interface{}) error {
	return json.NewDecoder(r.Body).Decode(data)
}

func goDotEnvVariable(key string) string {

	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func sendReponse(w http.ResponseWriter, _ *http.Request, data interface{}, status int) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	if data == nil {
		return
	}

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		log.Printf("Cannot format JSON. err=%v\n", err)
	}
}
