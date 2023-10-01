package main

import (
	"errors"
	"strings"
)

var ErrInvalidName = errors.New("invalid name")

// ParseName parses a name into first and last.
func ParseName(s string) (string, string, error) {
	parts := strings.Split(s, " ")
	if len(parts) != 2 {
		return "", "", ErrInvalidName
	}
	return parts[0], parts[1], nil
}
