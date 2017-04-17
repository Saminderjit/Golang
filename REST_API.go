package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"io/ioutil"

	"github.com/julienschmidt/httprouter"
)

type data struct {
	// first letter should be Capital for json parser to work
	Msg string `json:"msg"`
}

var serverData data
var message = "hello world"

func GetRestResponse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	/* Provides Rest response for GET calls */

	response, _ := json.Marshal(serverData)

	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(200)
	fmt.Fprintf(w, "%s", response)
}

func CreateNewResponse(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	/* Create a new response */

	// parsing the rest request
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println("error:", err)
	}

	// declaring a new variable for validation
	var newClientData data

	// Validating received data
	err = json.Unmarshal(body, &newClientData)
	if err != nil {
		fmt.Println("error:", err)
	}

	serverData.Msg = newClientData.Msg

	response, _ := json.Marshal(serverData)

	//w.Header().Set("Content-Type", "application/json")
	//w.WriteHeader(201)
	fmt.Fprintf(w, "%s", response)

}
func main() {
	serverData.Msg = message
	router := httprouter.New()
	router.GET("/A/B", GetRestResponse)
	router.POST("/A/B", CreateNewResponse)

	log.Fatal(http.ListenAndServe(":8000", router))

}
