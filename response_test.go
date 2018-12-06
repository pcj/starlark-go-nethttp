package nethttp

import (
	"testing"
)

func TestResponse(t *testing.T) {

	tests := []testCase{}

	for i, tc := range tests {
		runTestCase(t, i, &tc)
	}
}
