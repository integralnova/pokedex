package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
)

var page int = 0

func setPage(pageset int) {
	page = pageset
}

func getPage() int {
	return page
}
func nextPage() {
	getMap()
	setPage(getPage() + 1)
}

func previousPage() {
	if getPage() == 0 {
		fmt.Println("No more pages")
		setPage(0)
		return
	}
	if getPage() > 0 {
		setPage(getPage() - 1)
		fmt.Println("page reduced to ", getPage())
	}
	getMap()
}

func getMap() {
	var wg sync.WaitGroup
	wg.Add(20)

	go func() {

		for i := 1; i < 21; i++ {
			defer wg.Done()
			var numlocations = (page * 20) + i

			location, ok := getLocation(numlocations)

			if ok {
				fmt.Println(location.ID, ": ", location.Name)
			}
			if !ok {
				fmt.Println(numlocations, " : NOt valid ID")
			}

		}

	}()
	wg.Wait()

}

func locationFormatter(add int) string {
	return fmt.Sprint("https://pokeapi.co/api/v2/location/", add)
}

func getLocation(numLocations int) (Locationstruct, bool) {
	var result Locationstruct
	url := locationFormatter(numLocations)

	res, err := http.Get(url)
	if err != nil {
		log.Print(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		return result, false
	}
	if err != nil {
		log.Print(err)
	}

	if err := json.Unmarshal(body, &result); err != nil {
		print(err)
	}
	return result, true
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
