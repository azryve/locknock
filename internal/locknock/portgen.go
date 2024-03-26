package locknock

import (
	"crypto/sha256"
	"encoding/binary"

	"github.com/valyala/fastrand"
)

type PortGenerator struct {
	Key string
	rng *fastrand.RNG
}

// Generate next random port in a sequence
func (m *PortGenerator) Port() int {
	m.initRng()
	for {
		port := int(m.rng.Uint32())
		// only use client tcp ports
		port = port % 65535
		if port < 1024 {
			continue
		}
		return port
	}
}

func (m *PortGenerator) initRng() {
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
