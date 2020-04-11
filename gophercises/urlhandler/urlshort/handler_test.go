package urlshort

import (
	"testing"
)

type checkFunc func([]handle, error, *testing.T)

func hasError() checkFunc {
	return func(_ []handle, err error, t *testing.T) {
		t.Helper()
		if err == nil {
			t.Error("Expected error, got: nil")
		}
	}
}

func hasNoError() checkFunc {
	return func(_ []handle, err error, t *testing.T) {
		t.Helper()
		if err != nil {
			t.Errorf("Expected no error, got: %v", err)
		}
	}

}

func checkPath(expected string) checkFunc {
	return func(handles []handle, _ error, t *testing.T) {
		t.Helper()
		if expected != handles[0].Path {
			t.Errorf("Expected path %v, got: %v", expected, handles[0].Path)
		}
	}
}

func yamlString() []byte {
	out := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	return []byte(out)
}

func TestParseYaml(t *testing.T) {
	checks := func(cs ...checkFunc) []checkFunc { return cs }

	testCases := []struct {
		desc  string
		input []byte
		ch    []checkFunc
	}{
		{
			desc:  "invalid input",
			input: make([]byte, 2),
			ch:    checks(hasError()),
		},
		{
			desc:  "valid yaml",
			input: yamlString(),
			ch:    checks(hasNoError(), checkPath("/urlshort")),
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.desc, func(t *testing.T) {
			s, err := parseYaml(testCase.input)
			for _, check := range testCase.ch {
				check(s, err, t)
			}
		})
	}
}

func TestHandleToMap(t *testing.T) {
	handles := []handle{
		{Path: "test", URL: "testURL"},
		{Path: "test2", URL: "url2"},
	}
	handleMap := handleToMap(handles)

	if handleMap["test"] != "testURL" {
		t.Errorf("Expected key test to contain: testURL, got: %v", handleMap["test"])
	}
}
