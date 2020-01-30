package main

import "net/http"

import "log"

import "io"

import "os"

func main() {
	res, err := http.Get("http://localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		log.Fatal(res.StatusCode)
	}

	io.Copy(os.Stdout, res.Body)
}