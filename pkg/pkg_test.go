package pkg

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKindOfData(t *testing.T) {
	d := 100
	d1 := &d
	d2 := &d1
	d3 := &d2
	data := KindOfData(&d3)
	assert.Equal(t, data, reflect.Int)
}
