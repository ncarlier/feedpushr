package main

import (
	"fmt"
	"net/url"
	"sync/atomic"

	"github.com/ncarlier/feedpushr/v3/pkg/format"
	"github.com/ncarlier/feedpushr/v3/pkg/format/fn"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
)

var tootVisibilities = map[string]string{
	"public":   "Public",
	"private":  "Private",
	"direct":   "Direct",
	"unlisted": "Unlisted",
}

// see https://github.com/mastodon/mastodon/blob/main/app/helpers/languages_helper.rb
var tootLanguages = map[string]string{
	"": "- server default -",
	// ISO_639_1
	"aa": "Afar",
	"ab": "Abkhaz",
	"ae": "Avestan",
	"af": "Afrikaans",
	"ak": "Akan",
	"am": "Amharic",
	"an": "Aragonese",
	"ar": "Arabic",
	"as": "Assamese",
	"av": "Avaric",
	"ay": "Aymara",
	"az": "Azerbaijani",
	"ba": "Bashkir",
	"be": "Belarusian",
	"bg": "Bulgarian",
	"bh": "Bihari",
	"bi": "Bislama",
	"bm": "Bambara",
	"bn": "Bengali",
	"bo": "Tibetan",
	"br": "Breton",
	"bs": "Bosnian",
	"ca": "Catalan",
	"ce": "Chechen",
	"ch": "Chamorro",
	"co": "Corsican",
	"cr": "Cree",
	"cs": "Czech",
	"cu": "Old Church Slavonic",
	"cv": "Chuvash",
	"cy": "Welsh",
	"da": "Danish",
	"de": "German",
	"dv": "Divehi",
	"dz": "Dzongkha",
	"ee": "Ewe",
	"el": "Greek",
	"en": "English",
	"eo": "Esperanto",
	"es": "Spanish",
	"et": "Estonian",
	"eu": "Basque",
	"fa": "Persian",
	"ff": "Fula",
	"fi": "Finnish",
	"fj": "Fijian",
	"fo": "Faroese",
	"fr": "French",
	"fy": "Western Frisian",
	"ga": "Irish",
	"gd": "Scottish Gaelic",
	"gl": "Galician",
	"gu": "Gujarati",
	"gv": "Manx",
	"ha": "Hausa",
	"he": "Hebrew",
	"hi": "Hindi",
	"ho": "Hiri Motu",
	"hr": "Croatian",
	"ht": "Haitian",
	"hu": "Hungarian",
	"hy": "Armenian",
	"hz": "Herero",
	"ia": "Interlingua",
	"id": "Indonesian",
	"ie": "Interlingue",
	"ig": "Igbo",
	"ii": "Nuosu",
	"ik": "Inupiaq",
	"io": "Ido",
	"is": "Icelandic",
	"it": "Italian",
	"iu": "Inuktitut",
	"ja": "Japanese",
	"jv": "Javanese",
	"ka": "Georgian",
	"kg": "Kongo",
	"ki": "Kikuyu",
	"kj": "Kwanyama",
	"kk": "Kazakh",
	"kl": "Kalaallisut",
	"km": "Khmer",
	"kn": "Kannada",
	"ko": "Korean",
	"kr": "Kanuri",
	"ks": "Kashmiri",
	"ku": "Kurmanji (Kurdish)",
	"kv": "Komi",
	"kw": "Cornish",
	"ky": "Kyrgyz",
	"la": "Latin",
	"lb": "Luxembourgish",
	"lg": "Ganda",
	"li": "Limburgish",
	"ln": "Lingala",
	"lo": "Lao",
	"lt": "Lithuanian",
	"lu": "Luba-Katanga",
	"lv": "Latvian",
	"mg": "Malagasy",
	"mh": "Marshallese",
	"mi": "Māori",
	"mk": "Macedonian",
	"ml": "Malayalam",
	"mn": "Mongolian",
	"mr": "Marathi",
	"ms": "Malay",
	"mt": "Maltese",
	"my": "Burmese",
	"na": "Nauru",
	"nb": "Norwegian Bokmål",
	"nd": "Northern Ndebele",
	"ne": "Nepali",
	"ng": "Ndonga",
	"nl": "Dutch",
	"nn": "Norwegian Nynorsk",
	"no": "Norwegian",
	"nr": "Southern Ndebele",
	"nv": "Navajo",
	"ny": "Chichewa",
	"oc": "Occitan",
	"oj": "Ojibwe",
	"om": "Oromo",
	"or": "Oriya",
	"os": "Ossetian",
	"pa": "Panjabi",
	"pi": "Pāli",
	"pl": "Polish",
	"ps": "Pashto",
	"pt": "Portuguese",
	"qu": "Quechua",
	"rm": "Romansh",
	"rn": "Kirundi",
	"ro": "Romanian",
	"ru": "Russian",
	"rw": "Kinyarwanda",
	"sa": "Sanskrit",
	"sc": "Sardinian",
	"sd": "Sindhi",
	"se": "Northern Sami",
	"sg": "Sango",
	"si": "Sinhala",
	"sk": "Slovak",
	"sl": "Slovenian",
	"sn": "Shona",
	"so": "Somali",
	"sq": "Albanian",
	"sr": "Serbian",
	"ss": "Swati",
	"st": "Southern Sotho",
	"su": "Sundanese",
	"sv": "Swedish",
	"sw": "Swahili",
	"ta": "Tamil",
	"te": "Telugu",
	"tg": "Tajik",
	"th": "Thai",
	"ti": "Tigrinya",
	"tk": "Turkmen",
	"tl": "Tagalog",
	"tn": "Tswana",
	"to": "Tonga",
	"tr": "Turkish",
	"ts": "Tsonga",
	"tt": "Tatar",
	"tw": "Twi",
	"ty": "Tahitian",
	"ug": "Uyghur",
	"uk": "Ukrainian",
	"ur": "Urdu",
	"uz": "Uzbek",
	"ve": "Venda",
	"vi": "Vietnamese",
	"vo": "Volapük",
	"wa": "Walloon",
	"wo": "Wolof",
	"xh": "Xhosa",
	"yi": "Yiddish",
	"yo": "Yoruba",
	"za": "Zhuang",
	"zh": "Chinese",
	"zu": "Zulu",
	// ISO_639_3
	"ast": "Asturian",
	"ckb": "Sorani (Kurdish)",
	"cnr": "Montenegrin",
	"jbo": "Lojban",
	"kab": "Kabyle",
	"kmr": "Kurmanji (Kurdish)",
	"ldn": "Láadan",
	"lfn": "Lingua Franca Nova",
	"sco": "Scots",
	"sma": "Southern Sami'",
	"smj": "Lule Sami",
	"tok": "Toki Pona'",
	"zba": "Balaibalan",
	"zgh": "Standard Moroccan Tamazight",
}

var spec = model.Spec{
	Name: "mastodon",
	Desc: "Send new articles as *Toot* to a Mastodon instance.",
	PropsSpec: []model.PropSpec{
		{
			Name: "url",
			Desc: "Target URL",
			Type: model.Text,
		},
		{
			Name: "token",
			Desc: "Access token",
			Type: model.Password,
		},
		{
			Name:    "visibility",
			Desc:    "Toot visibility",
			Type:    model.Select,
			Options: tootVisibilities,
		},
		{
			Name:    "language",
			Desc:    "Toot language (leave empty for server-side default language)",
			Type:    model.Select,
			Options: tootLanguages,
		},
		{
			Name: "format",
			Desc: "Toot format (default: `{{.Title}}\\n{{.Link}}`)",
			Type: model.Textarea,
		},
	},
}

// MastodonOutputPlugin is the Mastodon output plugin
type MastodonOutputPlugin struct{}

// Spec returns plugin spec
func (p *MastodonOutputPlugin) Spec() model.Spec {
	return spec
}

// Build creates Mastodon output provider instance
func (p *MastodonOutputPlugin) Build(def *model.OutputDef) (model.Output, error) {
	// Default format
	if frmt, ok := def.Props["format"]; !ok || frmt == "" {
		def.Props["format"] = "{{.Title}}\n{{.Link}}"
	}
	formatter, err := format.NewOutputFormatter(def)
	if err != nil {
		return nil, err
	}
	u := def.Props.Get("url")
	if u == "" {
		return nil, fmt.Errorf("missing URL property")
	}
	_url, err := url.ParseRequestURI(u)
	if err != nil {
		return nil, fmt.Errorf("invalid URL property: %s", err.Error())
	}
	_url.Path = "/api/v1/statuses"
	accessToken := def.Props.Get("token")
	if accessToken == "" {
		return nil, fmt.Errorf("missing access token property")
	}
	visibility := def.Props.Get("visibility")
	if _, exists := tootVisibilities[visibility]; !exists {
		visibility = "public"
	}

	definition := *def
	definition.Spec = spec

	return &MastodonOutputProvider{
		definition:  definition,
		formatter:   formatter,
		targetURL:   _url.String(),
		accessToken: accessToken,
		visibility:  visibility,
		language:    def.Props.Get("language"),
	}, nil
}

// MastodonOutputProvider output provider to send articles to Mastodon
type MastodonOutputProvider struct {
	definition  model.OutputDef
	formatter   format.Formatter
	targetURL   string
	accessToken string
	visibility  string
	language    string
}

// Send article to a Mastodon instance.
func (op *MastodonOutputProvider) Send(article *model.Article) (bool, error) {
	b, err := op.formatter.Format(article)
	if err != nil {
		atomic.AddUint32(&op.definition.NbError, 1)
		return false, err
	}
	toot := Toot{
		Status:     fn.Truncate(500, b.String()),
		Sensitive:  false,
		Visibility: op.visibility,
		Language:   op.language,
	}
	if err := sendToMastodon(toot, op.targetURL, op.accessToken); err != nil {
		atomic.AddUint32(&op.definition.NbError, 1)
		return false, err
	}
	atomic.AddUint32(&op.definition.NbSuccess, 1)
	return true, nil
}

// GetDef return output definition
func (op *MastodonOutputProvider) GetDef() model.OutputDef {
	return op.definition
}

// GetPluginSpec return plugin informations
func GetPluginSpec() model.PluginSpec {
	return model.PluginSpec{
		Spec: spec,
		Type: model.OutputPluginType,
	}
}

// GetOutputPlugin returns output plugin
func GetOutputPlugin() (op model.OutputPlugin, err error) {
	return &MastodonOutputPlugin{}, nil
}
