package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

//The joke api response
type JokeResponse struct {
	Error     bool            `json:"error"`
	Categorie string          `json:"categorie"`
	Type      string          `json:"type"`
	Joke      string          `json:"joke"`
	Flags     map[string]bool `json:"flags"`
	Id        int             `json:"id"`
	Safe      bool            `json:"safe"`
	Lang      string          `json:"lang"`
}

func main() {
	url := "https://v2.jokeapi.dev/joke/Programming?type=single"

	//Send Request
	resp, err := http.Get(url) 
	if err != nil {
		log.Fatalf("Error while sending request, check internet connection!")
	}
	defer resp.Body.Close()

	//Checking for Status Code
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Error: Unexpected status code: %v", resp.StatusCode)
	}

	//Decoding Response
	jokeResp := new(JokeResponse)
	if err := json.NewDecoder(resp.Body).Decode(jokeResp); err != nil {
		log.Fatalf("Error: Unexpected status code: %v", resp.StatusCode)
	}

	fmt.Println(jokeResp.Joke)
}
