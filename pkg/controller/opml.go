package controller

import (
	"io"
	"io/ioutil"
	"strings"
	"time"

	"github.com/goadesign/goa"
	multierror "github.com/hashicorp/go-multierror"
	"github.com/ncarlier/feedpushr/autogen/app"
	"github.com/ncarlier/feedpushr/pkg/opml"
	"github.com/ncarlier/feedpushr/pkg/store"
)

// OpmlController implements the opml resource.
type OpmlController struct {
	*goa.Controller
	db store.DB
}

// NewOpmlController creates a opml controller.
func NewOpmlController(service *goa.Service, db store.DB) *OpmlController {
	return &OpmlController{
		Controller: service.NewController("OpmlController"),
		db:         db,
	}
}

// Get feeds as OPML file.
func (c *OpmlController) Get(ctx *app.GetOpmlContext) error {
	result := opml.NewOPML("Feedpushr exports")

	err := c.db.ForEachFeed(func(feed *app.Feed) error {
		outline := opml.Outline{}
		outline.Title = feed.Title
		if feed.Text != nil {
			outline.Text = *feed.Text
		}
		outline.Type = "rss"
		outline.XMLURL = feed.XMLURL
		if feed.HTMLURL != nil {
			outline.HTMLURL = *feed.HTMLURL
		}
		outline.Created = feed.Cdate.Format(time.RFC1123)
		result.Body.Outlines = append(result.Body.Outlines, outline)
		return nil
	})
	if err != nil {
		return goa.ErrInternal(err)
	}
	xml, err := result.XML()
	if err != nil {
		return goa.ErrInternal(err)
	}
	return ctx.OK([]byte(xml))
}

// Upload OMPL file to creates feeds.
func (c *OpmlController) Upload(ctx *app.UploadOpmlContext) error {
	reader, err := ctx.MultipartReader()
	if err != nil {
		return goa.ErrBadRequest("failed to load multipart request: %s", err)
	}
	if reader == nil {
		return goa.ErrBadRequest("not a multipart request")
	}
	for {
		p, err := reader.NextPart()
		if err == io.EOF {
			break
		}
		if err != nil {
			return goa.ErrBadRequest("failed to load part: %s", err)
		}
		b, err := ioutil.ReadAll(p)
		if err != nil {
			return ctx.BadRequest(err)
		}
		o, err := opml.NewOPMLFromBytes(b)
		if err != nil {
			return ctx.BadRequest(err)
		}

		err = opml.ImportOPMLToDB(o, c.db)
		if merr, ok := err.(*multierror.Error); ok {
			// Get error details
			l := len(merr.Errors) * 2
			metas := make([]interface{}, l, l)
			for idx, e := range merr.Errors {
				r := strings.SplitN(e.Error(), ":", 2)
				metas[idx*2] = r[0]
				metas[(idx*2)+1] = r[1]
			}
			return goa.ErrBadRequest("unable to import all OPLM items", metas...)
		} else if err != nil {
			return ctx.BadRequest(err)
		}
	}

	return ctx.Created()
}
