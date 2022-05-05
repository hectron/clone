package main

import (
	"reflect"
	"testing"
)

func TestParseInput(t *testing.T) {
	testCases := []struct {
		description, input string
		want               *Repository
	}{
		{"It sets a default user if none provided", "frontend", &Repository{"hectron", "frontend"}},
		{"It sets a the proper owner if provided", "tayne/notes", &Repository{"tayne", "notes"}},
		{"It is empty on bad input", "long/time/friend/of/you", &Repository{}},
		{"It is empty if input has whitespace", "guess you should think about stuff", &Repository{}},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			got := parseInput(tc.input)

			if !reflect.DeepEqual(tc.want, got) {
				t.Errorf("want %s, got %s", tc.want, got)
			}
		})
	}
}
