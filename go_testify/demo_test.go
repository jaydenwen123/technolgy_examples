package go_testify

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type user struct {
	name string
}

func TestSayHello(t *testing.T) {
	// assert equality
	assert.Equal(t, 123, 123, "they should be equal")

	// assert inequality
	assert.NotEqual(t, 123, 456, "they should not be equal")

	// assert for nil (good for errors)

	assert.Nil(t, &user{})

	// assert for not nil (good when you expect something)
	u := &user{name: "hello"}
	if assert.NotNil(t, u) {
		// now we know that object isn't nil, we are safe to make
		// further assertions without causing any errors
		assert.Equal(t, "hello",u.name)

	}
}
