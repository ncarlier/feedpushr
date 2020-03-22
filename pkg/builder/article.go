package builder

import (
	"github.com/mmcdole/gofeed"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

// NewArticle creates a new article from a feed item
func NewArticle(feed *model.FeedDef, item *gofeed.Item) *model.Article {
	article := &model.Article{}
	article.Content = item.Content
	article.Text = item.Description
	article.GUID = item.GUID
	article.Link = item.Link
	article.Title = item.Title
	article.Published = item.Published
	article.PublishedParsed = item.PublishedParsed
	article.Updated = item.Updated
	article.UpdatedParsed = item.UpdatedParsed
	article.Meta = make(map[string]interface{})
	article.Tags = feed.Tags
	return article
}

// NewArticles creates a new array of articles from an array of feed item
func NewArticles(feed *model.FeedDef, items []*gofeed.Item) []*model.Article {
	result := make([]*model.Article, len(items), len(items))
	for i := range items {
		result[i] = NewArticle(feed, items[i])
	}
	return result
}
