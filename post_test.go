package nethttp

import (
	"testing"
)

func TestPost(t *testing.T) {

	tests := []testCase{
		{
			name:   "check post builtin",
			input:  `print(http.post)`,
			output: `<built-in function http.post>`,
		},
	}

	for i, tc := range tests {
		runTestCase(t, i, &tc)
	}
}
