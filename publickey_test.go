package eciesgo

import (
	"crypto/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewPublicKeyFromHex(t *testing.T) {
	_, err := NewPublicKeyFromHex(testingReceiverPubkeyHex)
	assert.NoError(t, err)
}

func TestPublicKey_Equals(t *testing.T) {
	privkey, err := GenerateKey(rand.Reader)
	if !assert.NoError(t, err) {
		return
	}

	assert.True(t, privkey.PublicKey.Equals(privkey.PublicKey))
}
