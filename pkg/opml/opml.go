package opml

import (
	"encoding/xml"
	"io/ioutil"
	"time"
)

// OPML is the root node of an OPML document.
type OPML struct {
	XMLName xml.Name `xml:"opml"`
	Version string   `xml:"version,attr"`
	Head    Head     `xml:"head"`
	Body    Body     `xml:"body"`
}

// Head is the header node of an OPML document.
type Head struct {
	Title        string `xml:"title"`
	DateCreated  string `xml:"dateCreated,omitempty"`
	DateModified string `xml:"dateModified,omitempty"`
}

// Body is the body node of an OPML document.
type Body struct {
	Outlines []Outline `xml:"outline"`
}

// Outline contains details about the subscription.
type Outline struct {
	Outlines    []Outline `xml:"outline"`
	Text        string    `xml:"text,attr"`
	Type        string    `xml:"type,attr,omitempty"`
	Created     string    `xml:"created,attr,omitempty"`
	Category    string    `xml:"category,attr,omitempty"`
	XMLURL      string    `xml:"xmlUrl,attr,omitempty"`
	HTMLURL     string    `xml:"htmlUrl,attr,omitempty"`
	URL         string    `xml:"url,attr,omitempty"`
	Title       string    `xml:"title,attr,omitempty"`
	Description string    `xml:"description,attr,omitempty"`
}

// NewOPML creates new empty OPML object.
func NewOPML(title string) *OPML {
	result := &OPML{}
	result.Version = "1.0"
	result.Head.Title = title
	result.Head.DateCreated = time.Now().Format(time.RFC1123)
	result.Head.DateModified = time.Now().Format(time.RFC1123)
	return result
}

// NewOPMLFromFile creates a new OPML object from a file.
func NewOPMLFromFile(filename string) (*OPML, error) {
	b, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	o, err := NewOPMLFromBytes(b)
	return o, err
}

// NewOPMLFromBytes creates a new OPML object from a byte array.
func NewOPMLFromBytes(b []byte) (*OPML, error) {
	var result OPML
	err := xml.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

// XML returns the OPML document as a XML string.
func (doc *OPML) XML() (string, error) {
	b, err := xml.MarshalIndent(doc, "", "\t")
	return xml.Header + string(b), err
}
