//go:generate goagen bootstrap -d github.com/ncarlier/feedpushr/v3/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ncarlier/feedpushr/v3/autogen/app"
)

func main() {
	// Create service
	service := goa.New("feedpushr")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "explore" controller
	c := NewExploreController(service)
	app.MountExploreController(service, c)
	// Mount "feed" controller
	c2 := NewFeedController(service)
	app.MountFeedController(service, c2)
	// Mount "filter" controller
	c3 := NewFilterController(service)
	app.MountFilterController(service, c3)
	// Mount "health" controller
	c4 := NewHealthController(service)
	app.MountHealthController(service, c4)
	// Mount "index" controller
	c5 := NewIndexController(service)
	app.MountIndexController(service, c5)
	// Mount "opml" controller
	c6 := NewOpmlController(service)
	app.MountOpmlController(service, c6)
	// Mount "output" controller
	c7 := NewOutputController(service)
	app.MountOutputController(service, c7)
	// Mount "pshb" controller
	c8 := NewPshbController(service)
	app.MountPshbController(service, c8)
	// Mount "swagger" controller
	c9 := NewSwaggerController(service)
	app.MountSwaggerController(service, c9)
	// Mount "vars" controller
	c10 := NewVarsController(service)
	app.MountVarsController(service, c10)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}
}
