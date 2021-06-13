package auth

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCredentials(t *testing.T) {
	htpasswdFile, err := NewHtpasswdFromFile("test.htpasswd", "*")
	assert.Nil(t, err)
	assert.NotNil(t, htpasswdFile)
	assert.True(t, htpasswdFile.validateCredentials("foo", "bar"))
	assert.False(t, htpasswdFile.validateCredentials("foo", "bir"))
}
