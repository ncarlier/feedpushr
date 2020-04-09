// Code generated by goagen v1.4.3, DO NOT EDIT.
//
// API "feedpushr": Application Controllers
//
// Command:
// $ goagen
// --design=github.com/ncarlier/feedpushr/v2/design
// --out=/home/nicolas/workspace/fe/feedpushr/autogen
// --version=v1.4.3

package app

import (
	"context"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/cors"
	"net/http"
)

// initService sets up the service encoders, decoders and mux.
func initService(service *goa.Service) {
	// Setup encoders and decoders
	service.Encoder.Register(goa.NewJSONEncoder, "application/json")
	service.Encoder.Register(goa.NewGobEncoder, "application/gob", "application/x-gob")
	service.Encoder.Register(goa.NewXMLEncoder, "application/xml")
	service.Decoder.Register(goa.NewJSONDecoder, "application/json")
	service.Decoder.Register(goa.NewGobDecoder, "application/gob", "application/x-gob")
	service.Decoder.Register(goa.NewXMLDecoder, "application/xml")

	// Setup default encoder and decoder
	service.Encoder.Register(goa.NewJSONEncoder, "*/*")
	service.Decoder.Register(goa.NewJSONDecoder, "*/*")
}

// ExploreController is the controller interface for the Explore actions.
type ExploreController interface {
	goa.Muxer
	Get(*GetExploreContext) error
}

// MountExploreController "mounts" a Explore resource controller on the given service.
func MountExploreController(service *goa.Service, ctrl ExploreController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v2/explore", ctrl.MuxHandler("preflight", handleExploreOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetExploreContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleExploreOrigin(h)
	service.Mux.Handle("GET", "/v2/explore", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Explore", "action", "Get", "route", "GET /v2/explore")
}

// handleExploreOrigin applies the CORS response headers corresponding to the origin.
func handleExploreOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// FeedController is the controller interface for the Feed actions.
type FeedController interface {
	goa.Muxer
	Create(*CreateFeedContext) error
	Delete(*DeleteFeedContext) error
	Get(*GetFeedContext) error
	List(*ListFeedContext) error
	Start(*StartFeedContext) error
	Stop(*StopFeedContext) error
	Update(*UpdateFeedContext) error
}

// MountFeedController "mounts" a Feed resource controller on the given service.
func MountFeedController(service *goa.Service, ctrl FeedController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v2/feeds", ctrl.MuxHandler("preflight", handleFeedOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v2/feeds/:id", ctrl.MuxHandler("preflight", handleFeedOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v2/feeds/:id/start", ctrl.MuxHandler("preflight", handleFeedOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v2/feeds/:id/stop", ctrl.MuxHandler("preflight", handleFeedOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateFeedContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Create(rctx)
	}
	h = handleFeedOrigin(h)
	service.Mux.Handle("POST", "/v2/feeds", ctrl.MuxHandler("create", h, nil))
	service.LogInfo("mount", "ctrl", "Feed", "action", "Create", "route", "POST /v2/feeds")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteFeedContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleFeedOrigin(h)
	service.Mux.Handle("DELETE", "/v2/feeds/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Feed", "action", "Delete", "route", "DELETE /v2/feeds/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetFeedContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleFeedOrigin(h)
	service.Mux.Handle("GET", "/v2/feeds/:id", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Feed", "action", "Get", "route", "GET /v2/feeds/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListFeedContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleFeedOrigin(h)
	service.Mux.Handle("GET", "/v2/feeds", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Feed", "action", "List", "route", "GET /v2/feeds")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewStartFeedContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Start(rctx)
	}
	h = handleFeedOrigin(h)
	service.Mux.Handle("POST", "/v2/feeds/:id/start", ctrl.MuxHandler("start", h, nil))
	service.LogInfo("mount", "ctrl", "Feed", "action", "Start", "route", "POST /v2/feeds/:id/start")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewStopFeedContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Stop(rctx)
	}
	h = handleFeedOrigin(h)
	service.Mux.Handle("POST", "/v2/feeds/:id/stop", ctrl.MuxHandler("stop", h, nil))
	service.LogInfo("mount", "ctrl", "Feed", "action", "Stop", "route", "POST /v2/feeds/:id/stop")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateFeedContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Update(rctx)
	}
	h = handleFeedOrigin(h)
	service.Mux.Handle("PUT", "/v2/feeds/:id", ctrl.MuxHandler("update", h, nil))
	service.LogInfo("mount", "ctrl", "Feed", "action", "Update", "route", "PUT /v2/feeds/:id")
}

// handleFeedOrigin applies the CORS response headers corresponding to the origin.
func handleFeedOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// FilterController is the controller interface for the Filter actions.
type FilterController interface {
	goa.Muxer
	Specs(*SpecsFilterContext) error
}

// MountFilterController "mounts" a Filter resource controller on the given service.
func MountFilterController(service *goa.Service, ctrl FilterController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v2/filters/_specs", ctrl.MuxHandler("preflight", handleFilterOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSpecsFilterContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Specs(rctx)
	}
	h = handleFilterOrigin(h)
	service.Mux.Handle("GET", "/v2/filters/_specs", ctrl.MuxHandler("specs", h, nil))
	service.LogInfo("mount", "ctrl", "Filter", "action", "Specs", "route", "GET /v2/filters/_specs")
}

// handleFilterOrigin applies the CORS response headers corresponding to the origin.
func handleFilterOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// HealthController is the controller interface for the Health actions.
type HealthController interface {
	goa.Muxer
	Get(*GetHealthContext) error
}

// MountHealthController "mounts" a Health resource controller on the given service.
func MountHealthController(service *goa.Service, ctrl HealthController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v2/healthz", ctrl.MuxHandler("preflight", handleHealthOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetHealthContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleHealthOrigin(h)
	service.Mux.Handle("GET", "/v2/healthz", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Health", "action", "Get", "route", "GET /v2/healthz")
}

// handleHealthOrigin applies the CORS response headers corresponding to the origin.
func handleHealthOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// IndexController is the controller interface for the Index actions.
type IndexController interface {
	goa.Muxer
	Get(*GetIndexContext) error
}

// MountIndexController "mounts" a Index resource controller on the given service.
func MountIndexController(service *goa.Service, ctrl IndexController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v2/", ctrl.MuxHandler("preflight", handleIndexOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetIndexContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleIndexOrigin(h)
	service.Mux.Handle("GET", "/v2/", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Index", "action", "Get", "route", "GET /v2/")
}

// handleIndexOrigin applies the CORS response headers corresponding to the origin.
func handleIndexOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// OpmlController is the controller interface for the Opml actions.
type OpmlController interface {
	goa.Muxer
	Get(*GetOpmlContext) error
	Status(*StatusOpmlContext) error
	Upload(*UploadOpmlContext) error
}

// MountOpmlController "mounts" a Opml resource controller on the given service.
func MountOpmlController(service *goa.Service, ctrl OpmlController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v2/opml", ctrl.MuxHandler("preflight", handleOpmlOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v2/opml/status/:id", ctrl.MuxHandler("preflight", handleOpmlOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetOpmlContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleOpmlOrigin(h)
	service.Mux.Handle("GET", "/v2/opml", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Opml", "action", "Get", "route", "GET /v2/opml")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewStatusOpmlContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Status(rctx)
	}
	h = handleOpmlOrigin(h)
	service.Mux.Handle("GET", "/v2/opml/status/:id", ctrl.MuxHandler("status", h, nil))
	service.LogInfo("mount", "ctrl", "Opml", "action", "Status", "route", "GET /v2/opml/status/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUploadOpmlContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Upload(rctx)
	}
	h = handleOpmlOrigin(h)
	service.Mux.Handle("POST", "/v2/opml", ctrl.MuxHandler("upload", h, nil))
	service.LogInfo("mount", "ctrl", "Opml", "action", "Upload", "route", "POST /v2/opml")
}

// handleOpmlOrigin applies the CORS response headers corresponding to the origin.
func handleOpmlOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// OutputController is the controller interface for the Output actions.
type OutputController interface {
	goa.Muxer
	Create(*CreateOutputContext) error
	CreateFilter(*CreateFilterOutputContext) error
	Delete(*DeleteOutputContext) error
	DeleteFilter(*DeleteFilterOutputContext) error
	Get(*GetOutputContext) error
	List(*ListOutputContext) error
	Specs(*SpecsOutputContext) error
	Update(*UpdateOutputContext) error
	UpdateFilter(*UpdateFilterOutputContext) error
}

// MountOutputController "mounts" a Output resource controller on the given service.
func MountOutputController(service *goa.Service, ctrl OutputController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v2/outputs", ctrl.MuxHandler("preflight", handleOutputOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v2/outputs/:id/filters", ctrl.MuxHandler("preflight", handleOutputOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v2/outputs/:id", ctrl.MuxHandler("preflight", handleOutputOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v2/outputs/:id/filters/:idFilter", ctrl.MuxHandler("preflight", handleOutputOrigin(cors.HandlePreflight()), nil))
	service.Mux.Handle("OPTIONS", "/v2/outputs/_specs", ctrl.MuxHandler("preflight", handleOutputOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateOutputContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateOutputPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Create(rctx)
	}
	h = handleOutputOrigin(h)
	service.Mux.Handle("POST", "/v2/outputs", ctrl.MuxHandler("create", h, unmarshalCreateOutputPayload))
	service.LogInfo("mount", "ctrl", "Output", "action", "Create", "route", "POST /v2/outputs")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewCreateFilterOutputContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*CreateFilterOutputPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.CreateFilter(rctx)
	}
	h = handleOutputOrigin(h)
	service.Mux.Handle("POST", "/v2/outputs/:id/filters", ctrl.MuxHandler("createFilter", h, unmarshalCreateFilterOutputPayload))
	service.LogInfo("mount", "ctrl", "Output", "action", "CreateFilter", "route", "POST /v2/outputs/:id/filters")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteOutputContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Delete(rctx)
	}
	h = handleOutputOrigin(h)
	service.Mux.Handle("DELETE", "/v2/outputs/:id", ctrl.MuxHandler("delete", h, nil))
	service.LogInfo("mount", "ctrl", "Output", "action", "Delete", "route", "DELETE /v2/outputs/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewDeleteFilterOutputContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.DeleteFilter(rctx)
	}
	h = handleOutputOrigin(h)
	service.Mux.Handle("DELETE", "/v2/outputs/:id/filters/:idFilter", ctrl.MuxHandler("deleteFilter", h, nil))
	service.LogInfo("mount", "ctrl", "Output", "action", "DeleteFilter", "route", "DELETE /v2/outputs/:id/filters/:idFilter")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetOutputContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleOutputOrigin(h)
	service.Mux.Handle("GET", "/v2/outputs/:id", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Output", "action", "Get", "route", "GET /v2/outputs/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewListOutputContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.List(rctx)
	}
	h = handleOutputOrigin(h)
	service.Mux.Handle("GET", "/v2/outputs", ctrl.MuxHandler("list", h, nil))
	service.LogInfo("mount", "ctrl", "Output", "action", "List", "route", "GET /v2/outputs")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSpecsOutputContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Specs(rctx)
	}
	h = handleOutputOrigin(h)
	service.Mux.Handle("GET", "/v2/outputs/_specs", ctrl.MuxHandler("specs", h, nil))
	service.LogInfo("mount", "ctrl", "Output", "action", "Specs", "route", "GET /v2/outputs/_specs")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateOutputContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateOutputPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.Update(rctx)
	}
	h = handleOutputOrigin(h)
	service.Mux.Handle("PUT", "/v2/outputs/:id", ctrl.MuxHandler("update", h, unmarshalUpdateOutputPayload))
	service.LogInfo("mount", "ctrl", "Output", "action", "Update", "route", "PUT /v2/outputs/:id")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewUpdateFilterOutputContext(ctx, req, service)
		if err != nil {
			return err
		}
		// Build the payload
		if rawPayload := goa.ContextRequest(ctx).Payload; rawPayload != nil {
			rctx.Payload = rawPayload.(*UpdateFilterOutputPayload)
		} else {
			return goa.MissingPayloadError()
		}
		return ctrl.UpdateFilter(rctx)
	}
	h = handleOutputOrigin(h)
	service.Mux.Handle("PUT", "/v2/outputs/:id/filters/:idFilter", ctrl.MuxHandler("updateFilter", h, unmarshalUpdateFilterOutputPayload))
	service.LogInfo("mount", "ctrl", "Output", "action", "UpdateFilter", "route", "PUT /v2/outputs/:id/filters/:idFilter")
}

// handleOutputOrigin applies the CORS response headers corresponding to the origin.
func handleOutputOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// unmarshalCreateOutputPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateOutputPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createOutputPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalCreateFilterOutputPayload unmarshals the request body into the context request data Payload field.
func unmarshalCreateFilterOutputPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &createFilterOutputPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	if err := payload.Validate(); err != nil {
		// Initialize payload with private data structure so it can be logged
		goa.ContextRequest(ctx).Payload = payload
		return err
	}
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateOutputPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateOutputPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateOutputPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// unmarshalUpdateFilterOutputPayload unmarshals the request body into the context request data Payload field.
func unmarshalUpdateFilterOutputPayload(ctx context.Context, service *goa.Service, req *http.Request) error {
	payload := &updateFilterOutputPayload{}
	if err := service.DecodeRequest(req, payload); err != nil {
		return err
	}
	payload.Finalize()
	goa.ContextRequest(ctx).Payload = payload.Publicize()
	return nil
}

// PshbController is the controller interface for the Pshb actions.
type PshbController interface {
	goa.Muxer
	Pub(*PubPshbContext) error
	Sub(*SubPshbContext) error
}

// MountPshbController "mounts" a Pshb resource controller on the given service.
func MountPshbController(service *goa.Service, ctrl PshbController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v2/pshb", ctrl.MuxHandler("preflight", handlePshbOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewPubPshbContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Pub(rctx)
	}
	h = handlePshbOrigin(h)
	service.Mux.Handle("POST", "/v2/pshb", ctrl.MuxHandler("pub", h, nil))
	service.LogInfo("mount", "ctrl", "Pshb", "action", "Pub", "route", "POST /v2/pshb")

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewSubPshbContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Sub(rctx)
	}
	h = handlePshbOrigin(h)
	service.Mux.Handle("GET", "/v2/pshb", ctrl.MuxHandler("sub", h, nil))
	service.LogInfo("mount", "ctrl", "Pshb", "action", "Sub", "route", "GET /v2/pshb")
}

// handlePshbOrigin applies the CORS response headers corresponding to the origin.
func handlePshbOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// SwaggerController is the controller interface for the Swagger actions.
type SwaggerController interface {
	goa.Muxer
	Get(*GetSwaggerContext) error
}

// MountSwaggerController "mounts" a Swagger resource controller on the given service.
func MountSwaggerController(service *goa.Service, ctrl SwaggerController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v2/swagger.json", ctrl.MuxHandler("preflight", handleSwaggerOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetSwaggerContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleSwaggerOrigin(h)
	service.Mux.Handle("GET", "/v2/swagger.json", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Swagger", "action", "Get", "route", "GET /v2/swagger.json")
}

// handleSwaggerOrigin applies the CORS response headers corresponding to the origin.
func handleSwaggerOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Allow-Credentials", "false")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}

// VarsController is the controller interface for the Vars actions.
type VarsController interface {
	goa.Muxer
	Get(*GetVarsContext) error
}

// MountVarsController "mounts" a Vars resource controller on the given service.
func MountVarsController(service *goa.Service, ctrl VarsController) {
	initService(service)
	var h goa.Handler
	service.Mux.Handle("OPTIONS", "/v2/vars", ctrl.MuxHandler("preflight", handleVarsOrigin(cors.HandlePreflight()), nil))

	h = func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		// Check if there was an error loading the request
		if err := goa.ContextError(ctx); err != nil {
			return err
		}
		// Build the context
		rctx, err := NewGetVarsContext(ctx, req, service)
		if err != nil {
			return err
		}
		return ctrl.Get(rctx)
	}
	h = handleVarsOrigin(h)
	service.Mux.Handle("GET", "/v2/vars", ctrl.MuxHandler("get", h, nil))
	service.LogInfo("mount", "ctrl", "Vars", "action", "Get", "route", "GET /v2/vars")
}

// handleVarsOrigin applies the CORS response headers corresponding to the origin.
func handleVarsOrigin(h goa.Handler) goa.Handler {

	return func(ctx context.Context, rw http.ResponseWriter, req *http.Request) error {
		origin := req.Header.Get("Origin")
		if origin == "" {
			// Not a CORS request
			return h(ctx, rw, req)
		}
		if cors.MatchOrigin(origin, "*") {
			ctx = goa.WithLogContext(ctx, "origin", origin)
			rw.Header().Set("Access-Control-Allow-Origin", origin)
			rw.Header().Set("Access-Control-Max-Age", "600")
			rw.Header().Set("Access-Control-Allow-Credentials", "true")
			if acrm := req.Header.Get("Access-Control-Request-Method"); acrm != "" {
				// We are handling a preflight request
				rw.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE")
				rw.Header().Set("Access-Control-Allow-Headers", "content-type")
			}
			return h(ctx, rw, req)
		}

		return h(ctx, rw, req)
	}
}
