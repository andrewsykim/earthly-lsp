package client

import (
	"testing"
)

func Test_parseInput(t *testing.T) {
	testcases := []struct {
		name  string
		input string
		path  string
		line  int
		col   int
	}{
		{
			name:  "valid path",
			input: "/tmp/foo/bar:5:10",
			path:  "/tmp/foo/bar",
			line:  5,
			col:   10,
		},
		// TODO: add error test cases
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			path, line, col, err := parseInput(testcase.input)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if path != testcase.path {
				t.Logf("actual path: %v", path)
				t.Logf("expected path: %v", testcase.path)
				t.Error("unexpected path")
			}

			if line != testcase.line {
				t.Logf("actual line: %v", line)
				t.Logf("expected line: %v", testcase.line)
				t.Error("unexpected line")
			}

			if col != testcase.col {
				t.Logf("actual col: %v", col)
				t.Logf("expected col: %v", testcase.col)
				t.Error("unexpected col")
			}
		})
	}
}
