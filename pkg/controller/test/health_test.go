package ctrl_test

import (
	"context"
	"testing"

	"github.com/goadesign/goa"
	"github.com/ncarlier/feedpushr/autogen/app/test"
	"github.com/ncarlier/feedpushr/pkg/controller"
)

func TestGetHealth(t *testing.T) {
	var (
		service = goa.New("ctrl-test")
		ctrl    = controller.NewHealthController(service)
	)

	test.GetHealthOK(t, context.Background(), service, ctrl)
}
