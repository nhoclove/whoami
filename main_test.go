package main

import (
	"fmt"
	"net/http"
	"testing"
)

func TestGetMethod(t *testing.T) {
	testcases := map[string]struct {
		path    string
		handler http.HandlerFunc
	}{
		"simple path": {
			path: "path/test",
			handler: func(w http.ResponseWriter, r *http.Request) {
				// For test
				return
			},
		},
	}

	for name, tc := range testcases {
		mux := newMux()
		t.Run(name, func(t *testing.T) {
			mux.GET(tc.path, tc.handler)
			got, ok := mux.handlers[fmt.Sprintf("%s:%s", mGet, tc.path)]
			if !ok {
				t.Fatalf("expected: true, got:false")
			}

			if got == nil {
				t.Fatalf("wrong handler returned")
			}
		})
	}
}

func TestPostMethod(t *testing.T) {
	testcases := map[string]struct {
		path    string
		handler http.HandlerFunc
	}{
		"simple path": {
			path: "path/test",
			handler: func(w http.ResponseWriter, r *http.Request) {
				// For test
				return
			},
		},
	}

	for name, tc := range testcases {
		mux := newMux()
		t.Run(name, func(t *testing.T) {
			mux.POST(tc.path, tc.handler)
			got, ok := mux.handlers[fmt.Sprintf("%s:%s", mPost, tc.path)]
			if !ok {
				t.Fatalf("expected: true, got:false")
			}

			if got == nil {
				t.Fatalf("wrong handler returned")
			}
		})
	}
}
