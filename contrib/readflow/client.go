package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ncarlier/feedpushr/pkg/model"
)

// ArticleForm is a Readflow article form structure
type ArticleForm struct {
	Title       string     `json:"title,omitempty"`
	Text        *string    `json:"text,omitempty"`
	HTML        *string    `json:"html,omitempty"`
	URL         *string    `json:"url,omitempty"`
	Image       *string    `json:"image,omitempty"`
	PublishedAt *time.Time `json:"published_at,omitempty"`
	Tags        *string    `json:"tags,omitempty"`
}

// ArticlesResponse is the JSON response of readflow creation API
type ArticlesResponse struct {
	Articles []*Article
	Errors   []error
}

// Article structure definition
type Article struct {
	ID        uint       `json:"id,omitempty"`
	Title     string     `json:"title,omitempty"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
}

// Send article to a Readflow instance.
func sendToReadflow(url string, apiKey string, article *model.Article) (int, error) {
	result := 0
	b := new(bytes.Buffer)
	articleForm := ArticleForm{
		Title:       article.Title,
		URL:         &article.Link,
		PublishedAt: article.PublishedParsed,
	}
	if article.Content == "" {
		articleForm.HTML = &article.Description
	} else {
		articleForm.HTML = &article.Content
	}
	if len(article.Tags) > 0 {
		tags := strings.Join(article.Tags, ",")
		articleForm.Tags = &tags
	}
	if image, ok := article.Meta["image"]; ok {
		sImage := image.(string)
		articleForm.Image = &sImage
	}
	if text, ok := article.Meta["text"]; ok {
		sText := text.(string)
		articleForm.Text = &sText
	}

	json.NewEncoder(b).Encode([]ArticleForm{articleForm})
	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return result, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.SetBasicAuth("api", apiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()

	respJSON := &ArticlesResponse{}
	respText := ""
	switch resp.Header.Get("Content-type") {
	case "application/json":
		if err = json.NewDecoder(resp.Body).Decode(respJSON); err != nil {
			return result, fmt.Errorf("unable to read response: %s", err.Error())
		}
		// log.Println("respJSON", respJSON)
	default:
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return result, fmt.Errorf("unable to read response: %s", err.Error())
		}
		respText = string(data)
		// log.Println("respText", respText)
	}

	if resp.StatusCode >= 400 {
		if respJSON != nil && len(respJSON.Errors) > 0 {
			return result, respJSON.Errors[0]
		} else if respText != "" {
			return result, errors.New(respText)
		} else {
			return result, fmt.Errorf("bad status code: %d", resp.StatusCode)
		}
	}
	if resp.StatusCode < 300 && respJSON != nil {
		result = len(respJSON.Articles)
	}
	return result, nil
}
