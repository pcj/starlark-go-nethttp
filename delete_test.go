package nethttp

import (
	"testing"
)

func TestDelete(t *testing.T) {
	tests := []testCase{
		{
			name: "check delete builtin",
			input: `
print(http.delete)
`,
			output: `<built-in function http.delete>`,
		},
		{
			name: "delete response.status",
			input: `
print(http.delete("https://httpbin.org/delete").status)
`,
			output: `200 OK`,
		},
		{
			name: "delete response.code",
			input: `
print(http.delete("https://httpbin.org/delete").code)
`,
			output: `200`,
		},
	}

	for i, tc := range tests {
		runTestCase(t, i, &tc)
	}
}
