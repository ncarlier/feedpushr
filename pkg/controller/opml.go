package controller

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/opml"
	"github.com/ncarlier/feedpushr/v3/pkg/store"
)

// OpmlController implements the opml resource.
type OpmlController struct {
	*goa.Controller
	db       store.DB
	importer *opml.Importer
}

// NewOpmlController creates a opml controller.
func NewOpmlController(service *goa.Service, db store.DB) *OpmlController {
	return &OpmlController{
		Controller: service.NewController("OpmlController"),
		db:         db,
		importer:   opml.NewOPMLImporter(db),
	}
}

// Get feeds as OPML file.
func (c *OpmlController) Get(ctx *app.GetOpmlContext) error {
	result := opml.NewOPML("Feedpushr exports")

	err := c.db.ForEachFeed(func(feed *model.FeedDef) error {
		outline := opml.Outline{}
		outline.Title = feed.Title
		outline.Type = "rss"
		outline.XMLURL = feed.XMLURL
		if feed.HTMLURL != nil {
			outline.HTMLURL = *feed.HTMLURL
		}
		outline.Category = strings.Join(feed.Tags, ",")
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

// Upload OPML file to creates feeds.
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

		job, err := c.importer.ImportOPML(o)
		if err != nil {
			return ctx.BadRequest(err)
		}
		return ctx.Accepted(&app.OPMLImportJobResponse{
			ID: fmt.Sprintf("%d", job.ID),
		})
	}
	return goa.ErrBadRequest("no multipart data")
}

// Status runs the status action.
func (c *OpmlController) Status(ctx *app.StatusOpmlContext) error {
	w := ctx.ResponseWriter
	// Check that streaming is supported
	flusher, ok := w.(http.Flusher)
	if !ok {
		return goa.ErrInternal(errors.New("streaming not supported"))
	}

	// close := ctx.Request.Context().Done()

	out, err := c.importer.Get(uint(ctx.ID))
	if err != nil {
		return ctx.NotFound(goa.ErrNotFound(err))
	}
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(200)

	for line := range out {
		w.Write([]byte("event: message\n"))
		w.Write([]byte(fmt.Sprintf("data: %s\n\n", line)))
		flusher.Flush()
	}
	return nil
}
