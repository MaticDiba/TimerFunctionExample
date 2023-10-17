package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type ResponseBody struct {
	Message string `json:"message"`
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	message := "This HTTP triggered function executed successfully. Pass a name in the query string for a personalized response.\n"

	log.Printf("Executing Go HTTP trigger function")
	log.Printf("Received request on path: %s", r.URL.Path)

	url := "https://vreme.arso.gov.si/uploads/probase/www/nowcast/inca/inca_si0zm_data.json"
	data, err := fetchData(url)
	if err != nil {
		log.Fatal(err)
	}
	data.Print()

	response := ResponseBody{Message: message}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}
	http.HandleFunc("/api/TimerTriggerExample", helloHandler)
	http.HandleFunc("/TimerTriggerExample", helloHandler)
	log.Printf("About to listen on %s. Go to https://127.0.0.1%s/", listenAddr, listenAddr)
	log.Fatal(http.ListenAndServe(listenAddr, nil))
}
