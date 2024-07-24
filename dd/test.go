package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println(caller(2))
}
func getter(add int) string {
	return fmt.Sprint("https://pokeapi.co/api/v2/location/", add)
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

	if res.Header.Get("Content-Type") == "application/json" {
		bodyBytes, err := io.ReadAll(res.Body)
		if err != nil {
			panic(err)
		}

		fmt.Println("Response is JSON:", string(bodyBytes))
	} else {
		fmt.Println(res.Header.Get("Content-Type"))
	}

	return string(body)
}
