package main

import (
	"strconv"
	"strings"
	"sync/atomic"

	"github.com/iancoleman/strcase"
	"github.com/k3a/html2text"
	"gopkg.in/jdkato/prose.v2"

	"github.com/ncarlier/feedpushr/v3/pkg/expr"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

var filterList = map[string]string{
	"person": "PERSON",
	"gpe":    "GPE",
}

var formatList = map[string]string{
	"hashtag": "Hashtag",
	"none":    "None",
}

var separatorList = map[string]string{
	"space":          "Space ' '",
	"tab":            "Tab '\\t'",
	"returncarriage": "Return Carriage '\\n'",
	"comma":          "Comma ','",
	"semicolon":      "Semi-Colon ';'",
	"pipe":           "Pipe '|'",
}

var spec = model.Spec{
	Name: "prose",
	Desc: "Extract named entity from text.",
	PropsSpec: []model.PropSpec{
		{
			Name: "minCharLength",
			Desc: "Entity's minimum character length (default: 1)",
			Type: model.Number,
		},
		{
			Name: "maxCharLength",
			Desc: "Entity's maximum character length (default: 15)",
			Type: model.Number,
		},
		{
			Name:    "format",
			Desc:    "Formatting",
			Type:    model.Select,
			Options: formatList,
		},
		{
			Name:    "separator",
			Desc:    "Enities separator",
			Type:    model.Select,
			Options: separatorList,
		},
	},
}

func toHashtag(input string) string {
	input = strings.Replace(input, "_", " ", -1)
	input = strings.Replace(input, "-", " ", -1)
	input = "#" + strcase.ToCamel(input)
	input = strings.Replace(input, " ", "", -1)
	return input
}

func removeDuplicates(elements []string) []string {
	// Use map to record duplicates as we find them.
	encountered := map[string]bool{}
	result := []string{}
	for v := range elements {
		toLower := strings.ToLower(elements[v])
		if encountered[toLower] == true {
			// Do not add duplicate.
		} else {
			// Record this element as an encountered element.
			encountered[toLower] = true
			// Append to result slice.
			result = append(result, elements[v])
		}
	}
	// Return the new slice.
	return result
}

func safeAtoi(val string, fallback int) int {
	i, err := strconv.Atoi(val)
	if err != nil {
		return fallback
	}
	return i
}

// ProseFilterPlugin is the RAKE filter plugin
type ProseFilterPlugin struct{}

// Spec returns plugin spec
func (p *ProseFilterPlugin) Spec() model.Spec {
	return spec
}

// Build creates RAKE filter
func (p *ProseFilterPlugin) Build(def *model.FilterDef) (model.Filter, error) {
	condition, err := expr.NewConditionalExpression(def.Condition)
	if err != nil {
		return nil, err
	}
	val := def.Props.Get("minCharLength")
	minCharLength := safeAtoi(val, 1)
	val = def.Props.Get("maxCharLength")
	maxCharLength := safeAtoi(val, 15)
	format := def.Props.Get("format")
	if _, exists := formatList[format]; !exists {
		format = "hashtag"
	}
	separator := def.Props.Get("separator")
	if _, exists := separatorList[separator]; !exists {
		separator = "space"
	}

	definition := *def
	definition.Spec = spec
	definition.Props["minCharLength"] = minCharLength
	definition.Props["maxCharLength"] = maxCharLength
	definition.Props["format"] = format
	definition.Props["separator"] = separator

	return &ProseFilter{
		definition:    definition,
		condition:     condition,
		minCharLength: minCharLength,
		maxCharLength: maxCharLength,
		format:        format,
		separator:     separator,
	}, nil
}

// ProseFilter filter articles by adding extracted keywords
type ProseFilter struct {
	definition    model.FilterDef
	condition     *expr.ConditionalExpression
	minCharLength int
	maxCharLength int
	format        string
	separator     string
}

// DoFilter applies filter on the article
func (f *ProseFilter) DoFilter(article *model.Article) (bool, error) {
	plain := html2text.HTML2Text(article.Content)
	if plain == "" {
		plain = article.Text
	}

	doc, err := prose.NewDocument(plain)
	if err != nil {
		atomic.AddUint64(&f.definition.NbError, 1)
		return false, err
	}

	var entities []string
	for _, ent := range doc.Entities() {
		entity := ent.Text
		if len(entity) > f.maxCharLength {
			continue
		}
		if f.format == "hashtag" {
			entity = toHashtag(entity)
		}
		entities = append(entities, entity)
	}

	entities = removeDuplicates(entities)

	var sep string
	switch f.separator {
	case "space":
		sep = " "
	case "tab":
		sep = "\t"
	case "returncarriage":
		sep = "\n"
	case "comma":
		sep = ","
	case "semicolon":
		sep = ";"
	case "pipe":
		sep = ","
	}

	article.Meta["Entities"] = strings.Join(entities, sep)
	atomic.AddUint64(&f.definition.NbSuccess, 1)
	return true, nil
}

// GetDef return output definition
func (f *ProseFilter) GetDef() model.FilterDef {
	return f.definition
}

// Match test if article matches filter condition
func (f *ProseFilter) Match(article *model.Article) bool {
	return f.condition.Match(article)
}

// GetPluginSpec return plugin informations
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: spec,
		Type: model.FilterPluginType,
	}
}

// GetFilterPlugin returns filter plugin
func GetFilterPlugin() (op model.FilterPlugin, err error) {
	return &ProseFilterPlugin{}, nil
}
