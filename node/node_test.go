package node

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Node(t *testing.T) {

	v1 := "value-1"
	v2 := "value-2"
	v3 := "value-3"
	v3updated := "value-3-updated"

	n1 := New(nil, nil, v1)

	assert.IsType(t, &Node{}, n1)
	assert.Equal(t, v1, n1.GetValue())
	assert.Nil(t, n1.GetPrev())
	assert.Nil(t, n1.GetNext())

	n2 := New(n1, nil, v2)
	n1.SetNext(n2)

	assert.IsType(t, &Node{}, n1.GetNext())
	assert.IsType(t, &Node{}, n2.GetPrev())

	assert.Equal(t, v1, n2.GetPrev().GetValue())
	assert.Equal(t, v2, n1.GetNext().GetValue())

	n3 := New(nil, nil, v3)
	n3.SetPrev(n2)
	n2.SetNext(n3)

	assert.IsType(t, &Node{}, n3.GetPrev())

	n3.SetValue(v3updated)

	assert.Equal(t, v3updated, n3.GetValue())
}
