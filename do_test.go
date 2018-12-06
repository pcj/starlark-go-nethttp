package nethttp

import (
	"testing"
)

func TestDo(t *testing.T) {
	tests := []testCase{
		{
			name: "check do builtin",
			input: `
print(http.do)
`,
			output: `<built-in function http.do>`,
		},
		{
			name: "do with request",
			input: `
request = http.Request("https://httpbin.org/get")
response = http.do(request)
print(response.code)
			`,
			output: `200`,
		},
	}

	for i, tc := range tests {
		runTestCase(t, i, &tc)
	}
}
