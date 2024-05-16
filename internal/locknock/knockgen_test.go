package locknock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortRandNotNil(t *testing.T) {
	pr1 := KnockGenerator{Key: "Key"}
	port := pr1.Port()
	assert.NotEqual(t, port, 0)
}

func TestPortRandDifferent(t *testing.T) {
	pr1 := KnockGenerator{Key: "Key"}
	port1 := pr1.Port()
	port2 := pr1.Port()
	assert.NotEqual(t, port1, port2)
}

func TestPortSameKeySamePorts(t *testing.T) {
	pr1 := KnockGenerator{Key: "Key"}
	pr2 := KnockGenerator{Key: "Key"}
	for i := 0; i < 100; i++ {
		assert.Equal(t, pr1.Port(), pr2.Port())
	}
}

func TestPortConsistency(t *testing.T) {
	pr1 := KnockGenerator{Key: "Another key"}
	expected := []uint32{4291860670, 1948448125, 2066040667}
	actual := []uint32{}
	for i := 0; i < 3; i++ {
		actual = append(actual, pr1.Port())
	}
	assert.Equal(t, expected, actual)
}
