package pshb

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// Subscribe to a topic of a Hub.
func Subscribe(hub, topic, callback string) error {
	return post(hub, "subscribe", topic, callback)
}

// UnSubscribe from a topic of a Hub.
func UnSubscribe(hub, topic, callback string) error {
	return post(hub, "unsubscribe", topic, callback)
}

func post(hub, mode, topic, callback string) error {
	if mode != "subscribe" && mode != "unsubscribe" {
		return fmt.Errorf("bad subscription mode: %s", mode)
	}

	// Create request body
	form := url.Values{}
	form.Add("hub.callback", callback)
	form.Add("hub.topic", topic)
	form.Add("hub.mode", mode)
	payload := form.Encode()

	// Build the request
	client := &http.Client{}
	r, err := http.NewRequest("POST", hub, strings.NewReader(payload))
	if err != nil {
		return err
	}
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	r.Header.Add("Content-Length", strconv.Itoa(len(payload)))

	// Do the request
	resp, err := client.Do(r)
	if err != nil {
		return err
	}

	// Verify Hub response
	if resp.StatusCode != 202 {
		return fmt.Errorf("bad subscription response: %s", resp.Status)
	}
	return nil
}
