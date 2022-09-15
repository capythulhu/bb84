package bb84

import (
	"math/rand"
	"testing"
	"time"

	"github.com/tj/assert"
)

// Key length in bytes
const KEY_LENGTH = 1024

// Helper function for generating random bytes
func randomBytes() []byte {
	bytes := make([]byte, KEY_LENGTH)
	rand.Read(bytes)
	return bytes
}

// Alice send bits to Bob without any eavesdropping
func TestSafeKeyDistribution(t *testing.T) {
	rand.Seed(time.Now().Unix())

	aliceKey := randomBytes()
	aliceFilters := randomBytes()

	photons := BitsToPhotons(aliceKey, aliceFilters)

	bobFilters := randomBytes()
	bobKey := PhotonsToBits(photons, bobFilters)

	aliceFinalKey, bobFinalKey := GenerateKeys(aliceKey, bobKey, aliceFilters, bobFilters)
	assert.Equal(t, aliceFinalKey, bobFinalKey)
}

// Alice send bits to Bob with eavesdropping
func TestUnsafeKeyDistribution(t *testing.T) {
	rand.Seed(time.Now().Unix())

	aliceKey := randomBytes()
	aliceFilters := randomBytes()

	photons := BitsToPhotons(aliceKey, aliceFilters)

	// Eavesdrop on the photons
	eveFilters := randomBytes()
	PhotonsToBits(photons, eveFilters)

	bobFilters := randomBytes()
	bobKey := PhotonsToBits(photons, bobFilters)

	aliceFinalKey, bobFinalKey := GenerateKeys(aliceKey, bobKey, aliceFilters, bobFilters)
	assert.NotEqual(t, aliceFinalKey, bobFinalKey)
}
