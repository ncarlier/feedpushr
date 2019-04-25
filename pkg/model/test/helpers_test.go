package test

import (
	"testing"

	"github.com/ncarlier/feedpushr/pkg/assert"
	"github.com/ncarlier/feedpushr/pkg/model"
)

func TestMaskSecret(t *testing.T) {
	mask := "########"
	assert.Equal(t, mask, model.MaskSecret("zzz"), "")
	assert.Equal(t, "111"+mask+"999", model.MaskSecret("111zzz999"), "")
}
