package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func printError(w http.ResponseWriter, status int) {
	fmt.Fprintf(w, "{status: %d}", status)
}

func writeBack(w http.ResponseWriter, jsonBody []byte) {
	var x interface{}
	if err := json.Unmarshal(jsonBody, &x); err != nil {
		printError(w, -1)
		return
	}

	json.NewEncoder(w).Encode(x)
}

func handler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		printError(w, -1)
		return
	}

	type Message struct {
		Action string
	}

	var message Message
	if err = json.Unmarshal(jsonBody, &message); err != nil {
		printError(w, -1)
		return
	}

	log.Printf("{action: %s}", message.Action)

	switch message.Action {
	case "Login":
		writeBack(w, jsonBody)
		return
	default:
		printError(w, -2)
		return
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
