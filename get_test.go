package nethttp

import (
	"testing"
)

func TestGet(t *testing.T) {

	tests := []testCase{

		{
			name: "check get builtin",
			input: `
print(http.get)
`,
			output: `<built-in function http.get>`,
		},

		{
			name: "get response.status",
			input: `
print(http.get("https://httpbin.org/get").status)
`,
			output: `200 OK`,
		},
		{
			name: "get response.code",
			input: `
print(http.get("https://httpbin.org/get").code)
`,
			output: `200`,
		},
		{
			name: "get response.error",
			input: `
print(http.get("https://httpbin.org/get").error)
`,
			output: ``,
		},

		{
			name: "get 404",
			input: `
print(http.get("https://httpbin.org/idontexist").status)
`,
			output: `404 NOT FOUND`, // depends on server for response message
		},

		{
			name: "get error",
			input: `
print(http.get("ssh://httpbin.org/idontexist").error)
`,
			output: `Get ssh://httpbin.org/idontexist: unsupported protocol scheme "ssh"`,
		},
	}

	for i, tc := range tests {
		runTestCase(t, i, &tc)
	}
}
