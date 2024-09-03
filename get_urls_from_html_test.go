package get_urls_from_html

import (
	"reflect"
	"testing"
)

func TestGetURLsFromHTML(t *testing.T) {
	tests := []struct {
		name      string
		inputURL  string
		inputBody string
		expected  []string // Changed from string to []string
	}{
		{
			name:     "absolute and relative URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="/path/one">
						<span>Boot.dev</span>
					</a>
					<a href="https://other.com/path/one">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{"https://blog.boot.dev/path/one", "https://other.com/path/one"},
		},
		/* {
			name:     "no URLs",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<span>Boot.dev</span>
				</body>
			</html>
			`,
			expected: []string{},
		},
		{
			name:     "no href attribute",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a>
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{},
		},
		{
			name:     "invalid href attribute",
			inputURL: "https://blog.boot.dev",
			inputBody: `
			<html>
				<body>
					<a href="invalid">
						<span>Boot.dev</span>
					</a>
				</body>
			</html>
			`,
			expected: []string{},
		},
		{
			name:     "no HTML body",
			inputURL: "https://blog.boot.dev",
			inputBody: "",
			expected: []string{},
		},
		{
			name:     "empty HTML body",
			inputURL: "https://blog.boot.dev",
			inputBody: "<html></html>",
			expected: []string{},
		},
		{
			name:     "invalid HTML body",
			inputURL: "https://blog.boot.dev",
			inputBody: "<html>invalid</html>",
			expected: []string{},
		},
		{
			name:     "invalid HTML body with href",
			inputURL: "https://blog.boot.dev",
			inputBody: "<html><a href='https://blog.boot.dev'>link</a></html>",
			expected: []string{"https://blog.boot.dev"}, */
	}

	for i, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := getURLsFromHTML(tc.inputURL, tc.inputBody)
			if err != nil {
				t.Errorf("Test %v - '%s' FAIL: unexpected error: %v", i, tc.name, err)
				return
			}
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Errorf("Test %v - %s FAIL: expected URL: %v, actual: %v", i, tc.name, tc.expected, actual)
			}
		})
	}
}
