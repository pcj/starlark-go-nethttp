package nethttp

import (
	"testing"
)

func TestPut(t *testing.T) {
	tests := []testCase{
		{
			name: "check put builtin",
			input: `
print(http.put)
`,
			output: `<built-in function http.put>`,
		},
		{
			name: "put response.status",
			input: `
print(http.put("https://httpbin.org/put").status)
`,
			output: `200 OK`,
		},
		{
			name: "put response.code",
			input: `
print(http.put("https://httpbin.org/put").code)
`,
			output: `200`,
		},
	}

	for i, tc := range tests {
		runTestCase(t, i, &tc)
	}
}
