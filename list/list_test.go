package list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_List(t *testing.T) {

	list := New()

	assert.IsType(t, &List{}, list)
}
