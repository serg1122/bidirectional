package list

import (
	"strings"
	"testing"

	"github.com/serg1122/bidirectional"
	"github.com/stretchr/testify/assert"
)

func Test_List(t *testing.T) {

	v1 := "value-1"
	v2 := "value-2"
	v3 := "value-3"

	list := New()

	assert.IsType(t, &List{}, list)
	assert.True(t, list.IsEmpty())

	list.Append(v2)

	assert.False(t, list.IsEmpty())

	node2first := list.GetFirstNode()
	assert.IsType(t, &bidirectional.Node{}, node2first)
	assert.Nil(t, node2first.GetPrev())
	assert.Nil(t, node2first.GetNext())
	assert.Equal(t, v2, node2first.GetValue())

	node2last := list.GetLastNode()
	assert.IsType(t, &bidirectional.Node{}, node2last)
	assert.Nil(t, node2last.GetPrev())
	assert.Nil(t, node2last.GetNext())
	assert.Equal(t, v2, node2last.GetValue())

	list.Prepend(v1)

	node1first := list.GetFirstNode()
	assert.IsType(t, &bidirectional.Node{}, node1first)
	assert.Nil(t, node1first.GetPrev())
	assert.IsType(t, &bidirectional.Node{}, node2last.GetNext())
	assert.Equal(t, v1, node1first.GetValue())

	assert.IsType(t, &bidirectional.Node{}, node2last)
	assert.IsType(t, &bidirectional.Node{}, node2last.GetPrev())
	assert.Nil(t, node2last.GetNext())
	assert.Equal(t, v2, node2last.GetValue())

	list.Append(v3)

	node3last := list.GetLastNode()
	assert.IsType(t, &bidirectional.Node{}, node3last)
	assert.IsType(t, &bidirectional.Node{}, node3last.GetPrev())
	assert.IsType(t, &bidirectional.Node{}, node3last.GetNext())
	assert.Equal(t, v3, node3last.GetValue())

	assert.IsType(t, &bidirectional.Node{}, node2last)
	assert.IsType(t, &bidirectional.Node{}, node2last.GetPrev())
	assert.IsType(t, &bidirectional.Node{}, node2last.GetNext())
	assert.Equal(t, v2, node2last.GetValue())

	chainValueExpected := strings.Join([]string{v1, v2, v3}, ":")
	chainValueListGot := []string{}
	node := list.GetFirstNode()
	for node != nil {
		chainValueListGot = append(chainValueListGot, node.GetValue().(string))
		node = node.GetNext()
	}
	assert.Equal(t, chainValueExpected, strings.Join(chainValueListGot, ":"))
}
