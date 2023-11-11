package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

const target string = "http://192.168.230.121/rpc/Switch.GetStatus?id=0"

type power_data struct {
	Apower  float32 `json:"apower"`
	Voltage float32 `json:"voltage"`
	Current float32 `json:"current"`
}

func extract_power() {
	for {
		resp, err := http.Get(target)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			//defer resp.Body.Close()
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				fmt.Println("[pull (smart cables)] reading error: " + err.Error())
			} else {
				var dat power_data
				err = json.Unmarshal(b, &dat)
				if err != nil {
					fmt.Println("[pull (smart cables)] json error: " + err.Error())
				} else {
					forward_power(&dat)
				}
			}
			resp.Body.Close()
		}
		time.Sleep(1 * time.Second)
	}
}

func forward_power(dat *power_data) {
	//fmt.Println("[pull (smart cables)] watt: " + strconv.FormatFloat(float64(dat.Apower), 'f', -1, 32))
	fmt.Printf("[pull (smart cables)] %+v\n", dat)

	watt := strconv.FormatFloat(float64(dat.Apower), 'f', -1, 32)
	err := graphite_send(watt, "sprecometer.demo.building.0.entrance.chandelier")
	if err != nil {
		fmt.Println(err)
	}
}
