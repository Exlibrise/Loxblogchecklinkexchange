package ln_test

import (
	"testing"

	"github.com/philippgille/ln-paywall/ln"
)

// TestHashPreimage tests if the result of the HashPreimage function is correct.
func TestHashPreimage(t *testing.T) {
	// Correct preimage form, taken from a payment JSON in lnd
	preimageHex := "119969c2338798cd56708126b5d6c0