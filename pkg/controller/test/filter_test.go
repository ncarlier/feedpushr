package test

import (
	"context"
	"testing"

	"github.com/ncarlier/feedpushr/v3/autogen/app/test"
	"github.com/ncarlier/feedpushr/v3/pkg/assert"
	"github.com/ncarlier/feedpushr/v3/pkg/controller"
)

func TestFilterDefs(t *testing.T) {
	teardown := setup(t)
	defer teardown(t)

	ctrl := controller.NewFilterController(srv, chain)
	ctx := context.Background()

	_, specs := test.SpecsFilterOK(t, ctx, srv, ctrl)
	assert.True(t, len(specs) > 0, "")
	for _, spec := range specs {
		if spec.Name == "title" {
			assert.Equal(t, "This filter will prefix the title of the article with a given value.", spec.Desc, "")
			assert.True(t, len(spec.Props) == 1, "")
			assert.Equal(t, "prefix", spec.Props[0].Name, "")
			assert.Equal(t, "Prefix to add to the article title", spec.Props[0].Desc, "")
			assert.Equal(t, "text", spec.Props[0].Type, "")
		}
	}
}
