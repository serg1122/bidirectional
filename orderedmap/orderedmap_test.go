package orderedmap

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_OrderedMap(t *testing.T) {

	k1 := "key-1"
	k2 := "key-2"
	k3 := "key-3"

	expectedValues := map[string]interface{}{
		k1: "value-1",
		k2: "value-2",
		k3: "value-3",
	}

	om := New()

	assert.IsType(t, &OrderedMap{}, om)
	assert.Equal(t, 0, om.Length())
	assert.True(t, om.IsEmpty())

	assert.False(t, om.HasKey(k1))
	assert.False(t, om.HasKey(k2))
	assert.False(t, om.HasKey(k3))

	om.Set(k1, expectedValues[k1])

	assert.True(t, om.HasKey(k1))
	assert.False(t, om.HasKey(k2))
	assert.False(t, om.HasKey(k3))

	om.Remove(k1)

	assert.Equal(t, 0, om.Length())
	assert.True(t, om.IsEmpty())

	assert.False(t, om.HasKey(k1))
	assert.False(t, om.HasKey(k2))
	assert.False(t, om.HasKey(k3))

	om.Set(k1, expectedValues[k1])
	om.Set(k2, expectedValues[k2])
	om.Set(k3, expectedValues[k3])

	assert.True(t, om.HasKey(k1))
	assert.True(t, om.HasKey(k2))
	assert.True(t, om.HasKey(k3))

	expectedString := expectedValues[k1].(string) + expectedValues[k2].(string) + expectedValues[k3].(string)
	gotString := ""
	node := om.GetList().GetFirstNode()
	for node != nil {
		gotString += node.GetValue().(string)
		node = node.GetNext()
	}
	assert.Equal(t, expectedString, gotString)
}
