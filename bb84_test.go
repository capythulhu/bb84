package bb84

import (
	"crypto/rand"
	"testing"

	"github.com/tj/assert"
)

// Key length in bytes
const KEY_LENGTH = 8

func TestSafeKeyDistribution(t *testing.T) {
	aliceKey := make([]byte, KEY_LENGTH)
	rand.Read(aliceKey)

	aliceFilters := make([]byte, KEY_LENGTH)
	rand.Read(aliceFilters)

	photons := BitsToPhotons(aliceKey, aliceFilters)

	bobFilters := make([]byte, KEY_LENGTH)
	rand.Read(bobFilters)

	bobKey := PhotonsToBits(photons, bobFilters)

	aliceFinalKey, bobFinalKey := GenerateKeys(aliceKey, bobKey, aliceFilters, bobFilters)

	assert.Equal(t, aliceFinalKey, bobFinalKey)
}

func TestUnsafeKeyDistribution(t *testing.T) {
	aliceKey := make([]byte, KEY_LENGTH)
	rand.Read(aliceKey)

	aliceFilters := make([]byte, KEY_LENGTH)
	rand.Read(aliceFilters)

	photons := BitsToPhotons(aliceKey, aliceFilters)

	eveFilters := make([]byte, KEY_LENGTH)
	rand.Read(eveFilters)

	// Eavesdrop on the photons
	PhotonsToBits(photons, eveFilters)

	bobFilters := make([]byte, KEY_LENGTH)
	rand.Read(bobFilters)

	bobKey := PhotonsToBits(photons, bobFilters)

	aliceFinalKey, bobFinalKey := GenerateKeys(aliceKey, bobKey, aliceFilters, bobFilters)

	assert.NotEqual(t, aliceFinalKey, bobFinalKey)
}
