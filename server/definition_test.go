package server

import (
	"testing"

	"go.lsp.dev/uri"
)

func Test_parseTarget(t *testing.T) {
	testcases := []struct {
		name   string
		arg    string
		target string
	}{
		{
			name:   "simple target arg",
			arg:    "+base",
			target: "base",
		},
		{
			name:   "target with artifact",
			arg:    "+build/earthly",
			target: "build",
		},
		{
			name:   "target in another Earthfile",
			arg:    "./ast/parser+parser/*.go",
			target: "parser",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			target := parseTarget(testcase.arg)
			if target != testcase.target {
				t.Logf("actual target: %v", target)
				t.Logf("expected target: %v", testcase.target)
				t.Errorf("unexpected target name")
			}
		})
	}
}

func Test_parsePathFromURI(t *testing.T) {
	testcases := []struct {
		name string
		arg  string
		path string
	}{
		{
			name: "simple URI",
			arg:  "file:///tmp/foo/bar",
			path: "/tmp/foo/bar",
		},
	}

	for _, testcase := range testcases {
		t.Run(testcase.name, func(t *testing.T) {
			path := parsePathFromURI(uri.New(testcase.arg))
			if path != testcase.path {
				t.Logf("actual target path: %v", path)
				t.Logf("expected target path: %v", testcase.path)
				t.Errorf("unexpected target path")
			}
		})
	}
}
