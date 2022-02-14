package helpers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHash_GetMD5(t *testing.T) {
	t.Parallel()

	actual := MD5("password")
	expected := "5f4dcc3b5aa765d61d8327deb882cf99"
	assert.Equal(t, actual, expected)
}
