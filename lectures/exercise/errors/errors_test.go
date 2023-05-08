package main

import (
	"testing"
)

func TestTimeParse(t *testing.T) {
	timer := map[string]bool{
		"12:26:14": true,
		"l":        false,
		"12:45":    false,
		"45:12:13": false,
		"45:87:78": false,
		"abc:12":   false,
	}

	for k, v := range timer {
		_, err := TimeParse(k)
		if (err == nil) != v {
			t.Errorf("There is an unexpected error: %v", err)
		}
	}
}
