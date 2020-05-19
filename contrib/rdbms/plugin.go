package main

import (
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	"github.com/abadojack/whatlanggo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/ncarlier/feedpushr/v3/pkg/format"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

type Article struct {
	gorm.Model
	Title           string
	Text            string `sql:"type:text;"`
	Content         string `sql:"type:text;"`
	Link            string
	Updated         string
	UpdatedParsed   *time.Time
	Published       string
	PublishedParsed *time.Time
	GUID            string
	Tags            []Tag `gorm:"many2many:article_tags;"`
	Language        string
	LanguageConf    float64
	Categories      string
}

type Tag struct {
	gorm.Model
	Name string `gorm:"size:255;unique"`
}

var driverList = map[string]string{
	"sqlite3":  "Sqlite3",
	"mysql":    "MySQL",
	"postgres": "PostgreSQL",
}

var outputList = map[string]string{
	"quiet":   "Quiet Mode",
	"verbose": "Verbose Mode",
}

var spec = model.Spec{
	Name: "rdbms",
	Desc: "Send new articles to a relational database managed by gorm.io.",
	PropsSpec: []model.PropSpec{
		{
			Name:    "driver",
			Desc:    "Driver",
			Type:    model.Select,
			Options: driverList,
		},
		{
			Name: "database",
			Desc: "Database Name",
			Type: model.Text,
		},
		{
			Name: "host",
			Desc: "Host",
			Type: model.Text,
		},
		{
			Name: "port",
			Desc: "Port",
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
			Name:    "verbose",
			Desc:    "Verbosity",
			Type:    model.Select,
			Options: outputList,
		},
	},
}

// GormOutputPlugin is the Twitter output plugin
type GormOutputPlugin struct{}

// Spec returns plugin spec
func (p *GormOutputPlugin) Spec() model.Spec {
	return spec
}

// Build creates Twitter output provider instance
func (p *GormOutputPlugin) Build(def *model.OutputDef) (model.Output, error) {
	driver := def.Props.Get("driver")
	if _, exists := driverList[driver]; !exists {
		driver = "sqlite3"
	}
	database := def.Props.Get("database")
	if database == "" {
		return nil, fmt.Errorf("missing database property")
	}
	var host, port, username, password string
	if driver != "sqlite3" {
		host = def.Props.Get("host")
		if host == "" {
			return nil, fmt.Errorf("missing host property")
		}
		port = def.Props.Get("port")
		if port == "" {
			return nil, fmt.Errorf("missing port property")
		}
		username = def.Props.Get("username")
		if username == "" {
			return nil, fmt.Errorf("missing username property")
		}
		password = def.Props.Get("password")
		if password == "" {
			return nil, fmt.Errorf("missing password property")
		}
	}

	var db *gorm.DB
	var err error
	switch driver {
	case "mysql":
		db, err = gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4,utf8&parseTime=True&loc=Local", username, password, host, port, database))
	case "postgres":
		db, err = gorm.Open("postgres", fmt.Sprintf("postgres://%v:%v@%v/%v?sslmode=disable", username, password, host, database))
	case "sqlite":
		db, err = gorm.Open("sqlite3", fmt.Sprintf("%v", database))
	default:
		return nil, fmt.Errorf("incorrect driver provided.")
	}
	if err != nil {
		return nil, fmt.Errorf("could not connect to the database")
	}

	verbose := def.Props.Get("verbose")
	if _, exists := outputList[verbose]; !exists {
		verbose = "quiet"
	}

	switch verbose {
	case "verbose":
		db.LogMode(true)
	case "quiet":
		fallthrough
	default:
	}

	// migrate models
	db.AutoMigrate(&Article{})
	db.AutoMigrate(&Tag{})

	definition := *def
	definition.Spec = spec

	return &GormOutputProvider{
		definition: definition,
		db:         db,
	}, nil
}

// GormOutputProvider output provider to send articles to Twitter
type GormOutputProvider struct {
	definition model.OutputDef
	formatter  format.Formatter
	db         *gorm.DB
}

// Send sent an article as Tweet to a Twitter timeline
func (op *GormOutputProvider) Send(article *model.Article) (bool, error) {

	l := whatlanggo.Detect(article.Title + " " + article.Text + " " + article.Content)

	a := &Article{
		Title:           article.Title,
		Text:            article.Text,
		Content:         article.Content,
		Link:            article.Link,
		Updated:         article.Updated,
		UpdatedParsed:   article.UpdatedParsed,
		Published:       article.Published,
		PublishedParsed: article.PublishedParsed,
		Language:        l.Lang.String(),
		LanguageConf:    l.Confidence,
		Categories:      strings.Join(article.Tags, ","),
	}

	var tags []Tag
	for _, tag := range article.Tags {
		if tag == "" {
			continue
		}
		tt, err := createOrUpdateTag(op.db, &Tag{Name: tag})
		if err != nil {
			atomic.AddUint64(&op.definition.NbError, 1)
			return false, err
		}
		if tt.Name != "" {
			tags = append(tags, *tt)
		}
	}

	if len(tags) > 0 {
		a.Tags = tags
	}

	err := op.db.Create(a).Error
	// create article
	if err != nil {
		atomic.AddUint64(&op.definition.NbError, 1)
		return false, err
	}

	atomic.AddUint64(&op.definition.NbSuccess, 1)
	return true, err
}

// GetDef return filter definition
func (op *GormOutputProvider) GetDef() model.OutputDef {
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
	return &GormOutputPlugin{}, nil
}

func createOrUpdateTag(db *gorm.DB, tag *Tag) (*Tag, error) {
	var existingTag Tag
	if db.Where("name = ?", tag.Name).First(&existingTag).RecordNotFound() {
		err := db.Create(tag).Error
		return tag, err
	}
	tag.ID = existingTag.ID
	tag.CreatedAt = existingTag.CreatedAt
	return tag, nil
}

