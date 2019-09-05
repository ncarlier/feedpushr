package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

// Toot is a Mastodon status
type Toot struct {
	Status     string `json:"status"`
	Sensitive  bool   `json:"sensitive"`
	Visibility string `json:"visibility"`
}

func sendToMastodon(toot Toot, url string, accessToken string) error {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(toot)
	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	} else if resp.StatusCode >= 300 {
		return fmt.Errorf("bad status code: %d", resp.StatusCode)
	}
	return nil
}
