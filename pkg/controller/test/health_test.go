package test

import (
	"context"
	"testing"

	"github.com/ncarlier/feedpushr/v2/autogen/app/test"
	"github.com/ncarlier/feedpushr/v2/pkg/controller"
)

func TestGetHealth(t *testing.T) {
	ctrl := controller.NewHealthController(srv)
	test.GetHealthOK(t, context.Background(), srv, ctrl)
}
