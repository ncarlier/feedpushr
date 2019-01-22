package pshb

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// GetSubscriptionDetailsURL returns URL for subscription details
func GetSubscriptionDetailsURL(hub, topic, callback string) *string {
	u, err := url.Parse(hub)
	if err != nil {
		return nil
	}
	if u.Host == "pubsubhubbub.appspot.com" {
		u.Path = "/subscription-details"
	}
	q := u.Query()
	q.Set("hub.callback", callback)
	q.Set("hub.topic", topic)
	u.RawQuery = q.Encode()
	result := u.String()
	return &result
}

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
	// Required by PuSHPress the Wordpress plugin.
	// Non standard. Don't ask me why...
	form.Add("hub.verify", "foo")

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
		defer resp.Body.Close()
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("bad subscription response: %s (%s)", body, resp.Status)
	}
	return nil
}
