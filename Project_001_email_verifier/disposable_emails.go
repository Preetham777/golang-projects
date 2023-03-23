package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var disposableEmailsList []string

// Will get the list of disposable email from : https://raw.githubusercontent.com/ivolo/disposable-email-domains/master/index.json

func getUpdatedDisposableEmails() {

	response, err := http.Get("https://raw.githubusercontent.com/ivolo/disposable-email-domains/master/index.json")

	if err != nil {
		fmt.Printf("response error %v\n", err)
	} else {
		json_or_err := json.NewDecoder(response.Body).Decode(&disposableEmailsList)
		if json_or_err != nil {
			panic(json_or_err)
		}
		defer response.Body.Close()
	}

}
