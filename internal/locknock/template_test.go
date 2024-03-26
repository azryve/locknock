package locknock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDedent(t *testing.T) {
	expected := `Hello, Alice!
Hello, Bob!
Hello, Charlie!
`
	assert.Equal(t, expected, dedent(`
	Hello, Alice!
	Hello, Bob!
	Hello, Charlie!
	`))
}

func TestSum(t *testing.T) {
	assert.Equal(t, 1, sum(2, -1))
}
