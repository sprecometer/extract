package main

import (
	"net/http"
	"strings"
)

const graphite_address = "graphite:2003"

func graphite_send(value, path string) error {
	valueReader := strings.NewReader(path + " " + value + " -1\n")
	resp, err := http.Post("http://"+graphite_address, "text/plain", valueReader)
	if err != nil {
		return err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	return nil
}
