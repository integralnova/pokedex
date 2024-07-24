package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func locationFormatter(add int) string {
	return fmt.Sprint("https://pokeapi.co/api/v2/location/", add)
}

func getLocation(numLocations int) Locationstruct {
	url := locationFormatter(numLocations)

	res, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Printf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Print(err)
	}
	var result Locationstruct
	if err := json.Unmarshal(body, &result); err != nil {
		print(err)
	}
	return result
}

func maplocationFormatter(numlocations int) (response Locationstruct) {

	response = getLocation(numlocations)

	return response

}

// Holds response from https://pokeapi.co/api/v2/location/
type Locationstruct struct {
	Areas []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"areas"`
	GameIndices []struct {
		GameIndex  int `json:"game_index"`
		Generation struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"generation"`
	} `json:"game_indices"`
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	Region struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"region"`
}
