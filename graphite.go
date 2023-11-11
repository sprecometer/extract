package main

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

const graphite_address = "host.docker.internal:2003"

func graphite_send(value, path string) error {
	valueReader := strings.NewReader(path + " " + value + " -1\n")
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodPost, "http://"+graphite_address, io.NopCloser(valueReader))
	if err != nil {
		return err
	}
	req.Header.Add("Content-Type", "text/plain")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	if resp.Body != nil {
		defer resp.Body.Close()
	}
	return nil
}
