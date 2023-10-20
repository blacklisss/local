package main

import (
	"testing"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"babad", "aba"},
		{"cbbd", "bb"},
		{"a", "a"},
		{"ac", "c"},
		{"bb", "bb"},
		{"", ""},
	}

	for _, tt := range tests {
		result := longestPalindromicSubstring(tt.input)
		if result != tt.expected {
			t.Errorf("for input %q, expected %q but got %q", tt.input, tt.expected, result)
		}
	}
}
