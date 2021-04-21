package ln_test

import (
	"testing"

	"github.com/philippgille/ln-paywall/ln"
)

// TestHashPreimage tests if the result of the HashPreimage function is correct.
func TestHashPreimage(t *testing.T) {
	// Correct preimage form, taken from a