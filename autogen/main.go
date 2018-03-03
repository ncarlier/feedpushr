//go:generate goagen bootstrap -d github.com/ncarlier/feedpushr/design

package main

import (
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
	"github.com/ncarlier/feedpushr/autogen/app"
)

func main() {
	// Create service
	service := goa.New("feedpushr")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Mount "feed" controller
	c := NewFeedController(service)
	app.MountFeedController(service, c)
	// Mount "health" controller
	c2 := NewHealthController(service)
	app.MountHealthController(service, c2)
	// Mount "opml" controller
	c3 := NewOpmlController(service)
	app.MountOpmlController(service, c3)
	// Mount "pshb" controller
	c4 := NewPshbController(service)
	app.MountPshbController(service, c4)
	// Mount "swagger" controller
	c5 := NewSwaggerController(service)
	app.MountSwaggerController(service, c5)
	// Mount "vars" controller
	c6 := NewVarsController(service)
	app.MountVarsController(service, c6)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}
