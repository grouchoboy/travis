package travis

import "testing"
import "github.com/stretchr/testify/assert"

func TestSayHello(t *testing.T) {
	assert.Equal(t, "Hello Manu", SayHello("Manu"))
}
