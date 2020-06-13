package main

import (
	"fmt"
	"strings"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v3/pkg/format"
	"github.com/ncarlier/feedpushr/v3/pkg/format/fn"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/tebeka/selenium"
)

var browserList = map[string]string{
	"chrome":  "Chrome",
	"firefox": "Firefox",
}

var spec = model.Spec{
	Name: "twitter-selenium",
	Desc: "Send new articles to a Twitter timeline with Selenium.",
	PropsSpec: []model.PropSpec{
		{
			Name:    "browser",
			Desc:    "Browser",
			Type:    model.Select,
			Options: browserList,
		},
		{
			Name: "arguments",
			Desc: "arguments to pass to the browser",
			Type: model.Text,
		},
		{
			Name: "seleniumAddr",
			Desc: "selenium address (default: 127.0.0.1:4444)",
			Type: model.Text,
		},
		{
			Name: "username",
			Desc: "Username",
			Type: model.Text,
		},
		{
			Name: "password",
			Desc: "Password",
			Type: model.Password,
		},
		{
			Name: "format",
			Desc: "Tweet format (default: `{{tweet .Title .Link}}`)",
			Type: model.Textarea,
		},
	},
}

// TwitterSeleniumOutputPlugin is the TwitterSelenium output plugin
type TwitterSeleniumOutputPlugin struct{}

// Spec returns plugin spec
func (p *TwitterSeleniumOutputPlugin) Spec() model.Spec {
	return spec
}

// Build creates TwitterSelenium output provider instance
func (p *TwitterSeleniumOutputPlugin) Build(def *model.OutputDef) (model.Output, error) {
	// Default format
	if frmt, ok := def.Props["format"]; !ok || frmt == "" {
		def.Props["format"] = "{{tweet .Title .Link}}"
	}
	formatter, err := format.NewOutputFormatter(def)
	if err != nil {
		return nil, err
	}
	browser := def.Props.Get("browser")
	if _, exists := browserList[browser]; !exists {
		browser = "chrome"
	}
	arguments := def.Props.Get("arguments")
	args := strings.Split(arguments, " ")
	seleniumAddr := def.Props.Get("seleniumAddr")
	if seleniumAddr == "" {
		seleniumAddr = "127.0.0.1:4444"
	}
	username := def.Props.Get("username")
	if username == "" {
		return nil, fmt.Errorf("missing username key property")
	}
	password := def.Props.Get("password")
	if password == "" {
		return nil, fmt.Errorf("missing password property")
	}

	wd, err := initWebDriver(browser, args, seleniumAddr)
	if err != nil {
		return nil, fmt.Errorf("could not init selenium web driver, msg=%s", err)
	}

	definition := *def
	definition.Spec = spec

	t := &TwitterSeleniumOutputProvider{
		definition: definition,
		formatter:  formatter,
		browser:    browser,
		wd:         wd,
		username:   username,
		password:   password,
	}

	err = t.Login()
	if err != nil {
		return nil, fmt.Errorf("could not init login to twitter")
	}

	return t, nil
}

// TwitterSeleniumOutputProvider output provider to send articles to TwitterSelenium
type TwitterSeleniumOutputProvider struct {
	definition model.OutputDef
	formatter  format.Formatter
	browser    string
	username   string
	password   string
	wd         selenium.WebDriver
}

// Send sent an article as Tweet to a TwitterSelenium timeline
func (op *TwitterSeleniumOutputProvider) Send(article *model.Article) (bool, error) {
	b, err := op.formatter.Format(article)
	if err != nil {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, err
	}
	tweet := fn.Truncate(270, b.String())
	err = op.Tweet(tweet)
	if err != nil {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, nil
	}
	atomic.AddUint64(&op.definition.NbSuccess, 1)
	return true, err
}

// GetDef return filter definition
func (op *TwitterSeleniumOutputProvider) GetDef() model.OutputDef {
	return op.definition
}

// GetPluginSpec returns plugin spec
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: spec,
		Type: model.OutputPluginType,
	}
}

// GetOutputPlugin returns output plugin
func GetOutputPlugin() (op model.OutputPlugin, err error) {
	return &TwitterSeleniumOutputPlugin{}, nil
}
