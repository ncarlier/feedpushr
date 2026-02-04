package handler

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/ncarlier/feedpushr/v3/pkg/api"
	"github.com/ncarlier/feedpushr/v3/pkg/model"
	"github.com/ncarlier/feedpushr/v3/pkg/opml"
)

// GetOPML handles GET /v2/opml
func (h *Handlers) GetOPML(w http.ResponseWriter, r *http.Request) {
	res := api.NewResponse(w)

	result := opml.NewOPML("Feedpushr exports")

	err := h.db.ForEachFeed(r.Context(), func(feed *model.FeedDef) error {
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
		res.InternalError(err)
		return
	}

	xml, err := result.XML()
	if err != nil {
		res.InternalError(err)
		return
	}

	res.Raw("application/xml; charset=utf-8", []byte(xml))
}

// UploadOPML handles POST /v2/opml
func (h *Handlers) UploadOPML(w http.ResponseWriter, r *http.Request) {
	res := api.NewResponse(w)

	reader, err := r.MultipartReader()
	if err != nil {
		res.BadRequest(errors.New("failed to load multipart request"))
		return
	}
	if reader == nil {
		res.BadRequest(errors.New("not a multipart request"))
		return
	}

	importer := opml.NewOPMLImporter(h.db)

	p, err := reader.NextPart()
	if err == io.EOF {
		res.BadRequest(errors.New("no multipart data"))
		return
	}
	if err != nil {
		res.BadRequest(errors.New("failed to load part"))
		return
	}

	b, err := io.ReadAll(p)
	if err != nil {
		res.BadRequest(err)
		return
	}

	o, err := opml.NewOPMLFromBytes(b)
	if err != nil {
		res.BadRequest(err)
		return
	}

	job, err := importer.ImportOPML(o)
	if err != nil {
		res.BadRequest(err)
		return
	}

	res.Accepted(map[string]interface{}{
		"id": fmt.Sprintf("%d", job.ID),
	})
}

// GetOPMLStatus handles GET /v2/opml/status/{id}
func (h *Handlers) GetOPMLStatus(w http.ResponseWriter, r *http.Request) {
	req := api.NewRequest(r)
	res := api.NewResponse(w)

	idStr := req.PathParam("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		res.BadRequest(err)
		return
	}

	// Check that streaming is supported
	flusher, ok := w.(http.Flusher)
	if !ok {
		res.InternalError(errors.New("streaming not supported"))
		return
	}

	importer := opml.NewOPMLImporter(h.db)
	out, err := importer.Get(uint(id))
	if err != nil {
		res.NotFound()
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	w.WriteHeader(http.StatusOK)

	for line := range out {
		w.Write([]byte("event: message\n"))
		w.Write([]byte(fmt.Sprintf("data: %s\n\n", line)))
		flusher.Flush()
	}
}
