package testhelpers

import "testing"

// Long ...
func Long(t *testing.T) {
	if testing.Short() {
		t.Skip("skip testing in short mode")
	}
}
