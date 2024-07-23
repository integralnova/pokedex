package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func getter(add int) string {
	return fmt.Sprint("https://pokeapi.co/api/v2/location-area/", add)
}

func caller(s int) string {
	call := getter(s)

	res, err := http.Get(call)
	if err != nil {
		log.Fatal(err)
	}
	body, err := io.ReadAll(res.Body)
	res.Body.Close()
	if res.StatusCode > 299 {
		log.Fatalf("Response failed with status code: %d and\nbody: %s\n", res.StatusCode, body)
	}
	if err != nil {
		log.Fatal(err)
	}
	return string(body)
}

func mapgetter(numlocations int) (response []string) {

	response = make([]string, numlocations)

	for i := 1; i < numlocations+1; i++ {
		response[i] = caller(i)
	}

	return response

}
