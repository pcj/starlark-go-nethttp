package nethttp

import (
	"testing"
)

func TestRequest(t *testing.T) {
	tests := []testCase{
		{
			name:   "Request constructor should identify itself",
			input:  `print(http.Request)`,
			output: `<built-in function http.Request>`,
		},
		{
			name:   "Request object prints itself with string dict attrs",
			input:  `print(http.Request(url=""))`,
			output: `Request({body: "", content_length: 0, headers: {}, host: "", method: "GET", url: ""})`,
		},
		{
			name:   "Request url is captured",
			input:  `print(http.Request(url="http://example.com").url)`,
			output: `http://example.com`,
		},
		{
			name:   "Request params is captured",
			input:  `print(http.Request(url="http://example.com",params={"k":"v"}).url)`,
			output: `http://example.com?k=v`,
		},
		{
			name:   "Request headers is captured",
			input:  `print(http.Request(url="",headers={"k":"v"}).headers)`,
			output: `{"k": ["v"]}`,
		},
		{
			name:   "Request body is captured",
			input:  `print(http.Request(url="", body="foo").body)`,
			output: `foo`,
		},
	}

	for i, tc := range tests {
		runTestCase(t, i, &tc)
	}
}
