package nethttp

import (
	"testing"
)

func TestPatch(t *testing.T) {
	tests := []testCase{
		{
			name: "check patch builtin",
			input: `
print(http.patch)
`,
			output: `<built-in function http.patch>`,
		},
		{
			name: "patch response.status",
			input: `
print(http.patch("https://httpbin.org/patch").status)
`,
			output: `200 OK`,
		},
		{
			name: "patch response.code",
			input: `
print(http.patch("https://httpbin.org/patch").code)
`,
			output: `200`,
		},
	}

	for i, tc := range tests {
		runTestCase(t, i, &tc)
	}
}
