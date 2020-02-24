package querystring

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateInstance(t *testing.T) {
	i := CreateInstance("http://google.com/?q=search&foo=bar")

	assert.Equal(t, "http://google.com/", i.Path)
	assert.Len(t, i.Data, 2)
}

func Test_SetParameter(t *testing.T) {
	i := CreateInstance("http://google.com/?q=search")
	assert.Len(t, i.Data, 1)

	i.SetParameter("foo", "bar")

	assert.Len(t, i.Data, 2)

	assert.Equal(t, "bar", i.Get("foo"))
}
