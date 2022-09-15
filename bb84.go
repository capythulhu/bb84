package bb84

import (
	"math/rand"
	"time"
)

// Photon
type Photon struct {
	bit, phase bool
}

// Create a new photon with the given bit and phase
func NewPhoton(bit, phase bool) Photon {
	return Photon{bit, phase}
}

// Polarize a photon with the given filter
func (p *Photon) Measure(base bool) bool {
	if p.phase != base {
		rand.Seed(time.Now().Unix())
		p.bit = rand.Intn(2) == 0
		p.phase = base
	}
	return p.bit
}

func BitsToPhotons(bits, phases []byte) []Photon {
	if len(bits) != len(phases) {
		panic("length of bits and phases must be equal")
	}

	photons := make([]Photon, len(bits))
	for i := range photons {
		bit := (bits[i/8] & (1 << uint(i%8))) != 0
		phase := (phases[i/8] & (1 << uint(i%8))) != 0
		photons[i] = NewPhoton(bit, phase)
	}

	return photons
}

func PhotonsToBits(photons []Photon, phases []byte) []byte {
	if len(photons) != len(phases) {
		panic("length of photons and phases must be equal")
	}

	bits := make([]byte, len(photons))
	for i := range bits {
		if photons[i].Measure(phases[i/8]&(1<<uint(i%8)) != 0) {
			bits[i/8] |= 1 << uint(i%8)
		}
	}

	return bits
}

func GenerateKeys(bits1, bits2, phases1, phases2 []byte) ([]byte, []byte) {
	if len(bits1) != len(bits2) || len(bits1) != len(phases1) || len(bits1) != len(phases2) {
		panic("length of bits and phases must be equal")
	}

	key1 := make([]byte, len(bits1))
	key2 := make([]byte, len(bits2))
	for i := range key1 {
		if phases1[i/8]&(1<<uint(i%8)) == phases2[i/8]&(1<<uint(i%8)) {
			key1[i/8] |= bits1[i/8] & (1 << uint(i%8))
			key2[i/8] |= bits2[i/8] & (1 << uint(i%8))
		}
	}

	return key1, key2
}
