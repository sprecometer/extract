package main

import (
	"fmt"
	"log"
	"net/http"
)

const port string = "8280"

func main() {
	http.HandleFunc("/computer/consumption", extract_computer)

	fmt.Println("Listening on port " + port)

	go extract_power()

	log.Fatal(http.ListenAndServe(":"+port, nil))
}
