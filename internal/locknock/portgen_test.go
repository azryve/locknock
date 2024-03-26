package locknock

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPortRandNotNil(t *testing.T) {
	pr1 := PortGenerator{Key: "Key"}
	port := pr1.Port()
	assert.NotEqual(t, port, 0)
}

func TestPortRandDifferent(t *testing.T) {
	pr1 := PortGenerator{Key: "Key"}
	port1 := pr1.Port()
	port2 := pr1.Port()
	assert.NotEqual(t, port1, port2)
}

func TestPortSameKeySamePorts(t *testing.T) {
	pr1 := PortGenerator{Key: "Key"}
	pr2 := PortGenerator{Key: "Key"}
	for i := 0; i < 100; i++ {
		assert.Equal(t, pr1.Port(), pr2.Port())
	}
}

func TestPortConsistency(t *testing.T) {
	pr1 := PortGenerator{Key: "Another key"}
	expected := []int{39055, 27040, 49792, 17220, 64850, 35430, 40149, 40663, 61631, 32234}
	actual := []int{}
	for i := 0; i < 10; i++ {
		actual = append(actual, pr1.Port())
	}
	assert.Equal(t, expected, actual)
}
