package plugins

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/ncarlier/feedpushr/v3/pkg/common"
	"github.com/ncarlier/feedpushr/v3/pkg/format/fn"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
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

	if !strings.HasSuffix(url, "/articles") {
		url = url + "/articles"
	}

	b := new(bytes.Buffer)
	// Init form
	articleForm := ArticleForm{
		Title:       article.Title,
		URL:         &article.Link,
		PublishedAt: article.PublishedParsed,
	}

	// Set content
	if article.Content == "" {
		articleForm.HTML = &article.Text
	} else {
		articleForm.HTML = &article.Content
	}

	// Set tags
	if len(article.Tags) > 0 {
		tags := strings.Join(article.Tags, ",")
		articleForm.Tags = &tags
	}

	// Set image
	if image, ok := article.Meta["image"]; ok {
		if value := image.(string); value != "" {
			articleForm.Image = &value
		}
	}

	// Set text
	text := fn.Truncate(500, article.Text)
	articleForm.Text = &text
	if excerpt, ok := article.Meta["excerpt"]; ok {
		if value := excerpt.(string); value != "" {
			articleForm.Text = &value
		}
	}

	json.NewEncoder(b).Encode([]ArticleForm{articleForm})
	req, err := http.NewRequest("POST", url, b)
	if err != nil {
		return result, err
	}
	req.Header.Set("User-Agent", common.UserAgent)
	req.Header.Set("Content-Type", common.ContentTypeJSON)
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
