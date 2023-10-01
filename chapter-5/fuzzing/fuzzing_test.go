package main

import (
	"errors"
	"testing"
)

func FuzzParseName(f *testing.F) {
	f.Add("John Adams")
	f.Add("George Washington")
	f.Fuzz(func(t *testing.T, s string) {
		_, _, err := ParseName(s)
		if err != nil {
			if errors.Is(err, ErrInvalidName) {
				return
			}
			t.Errorf("%v", err)
		}
	})
}
