package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func extract_computer(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		if r.Body == nil {
			fmt.Println("[push (computer)] ? nil body ?")
		} else {
			defer r.Body.Close()
			b, err := io.ReadAll(r.Body)
			if err == nil {
				usage := strings.Split(string(b), "=")[1]
				forward_computer(usage)
			} else {
				fmt.Println("[push (computer)] error: " + err.Error())
			}
		}
	}
}

func forward_computer(cpu string) {
	fmt.Println("[push (computer)]" + cpu)

	err := graphite_send(cpu, "sprecometer.demo.building.0.entrance.reception")
	if err != nil {
		fmt.Println(err)
	}
}
