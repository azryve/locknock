package locknock

import (
	"crypto/sha256"
	"encoding/binary"

	"github.com/valyala/fastrand"
)

type KnockGenerator struct {
	Key string
	rng *fastrand.RNG
}

// Generate next random port in a sequence
func (m *KnockGenerator) Port() uint32 {
	m.initRng()
	for {
		return m.rng.Uint32()
	}
}

func (m *KnockGenerator) initRng() {
	if m.rng == nil {
		m.rng = &fastrand.RNG{}
		seed := stringToUint32Seed(m.Key)
		m.rng.Seed(seed)
	}
}

func stringToUint32Seed(input string) uint32 {
	hash := sha256.Sum256([]byte(input))
	seed := binary.BigEndian.Uint32(hash[:4])
	return seed
}
