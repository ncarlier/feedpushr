package test

import (
	"fmt"
	"testing"

	"github.com/ncarlier/feedpushr/v2/pkg/assert"
	"github.com/ncarlier/feedpushr/v2/pkg/format"
	"github.com/ncarlier/feedpushr/v2/pkg/model"
)

func TestTemplateFormater(t *testing.T) {
	art := &model.Article{
		Title: "My God, It's Full of Stars",
		Link:  "https://en.wikipedia.org/wiki/2001:_A_Space_Odyssey_(novel)",
	}
	testCases := []struct {
		value    string
		expected string
	}{
		{"{{.Title}}", art.Title},
		{"{{truncate 6 .Title}}", "My God"},
		{"{{tweet .Title .Link}}", fmt.Sprintf("%s\n%s", art.Title, art.Link)},
	}
	for idx, tc := range testCases {
		formatter, err := format.NewTemplateFormatter(fmt.Sprintf("test-%d", idx), tc.value)
		assert.Nil(t, err, "")
		buf, err := formatter.Format(art)
		assert.Equal(t, tc.expected, buf.String(), "")
	}
}
