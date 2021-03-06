package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	if err := http.ListenAndServe(":8080", http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		bodybytes, err := ioutil.ReadAll(req.Body)
		if err != nil {
			log.Printf("%s\n", err)
		}
		req.Body.Close()
		fmt.Printf("%s\n", string(bodybytes))
	})); err != nil {
		log.Fatal(err)
	}
}
