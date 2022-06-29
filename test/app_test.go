package test

import (
	"github.com/screw-coding/filter"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestApp(t *testing.T) {
	var (
		in  = []string{"Cat", "DOG", "fisH"}
		out = []string{"cat", "dog", "fish"}
	)
	assert.Equal(t, out, filter.Lowercase(in))
}
