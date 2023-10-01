package main

import (
	"math"
	"strings"
)

type ValidationClient struct{}

type ValidationClientInter interface {
	IsNumbersValid(name string) bool
	IsLengthValid(name string) bool
}

var Client ValidationClientInter = &ValidationClient{}

func main() {
}

func GetBigger(a, b float64) float64 {
	return math.Max(a, b)
}

// Valid name is between 8-16 chars and doesn't contain number
func IsValidName(name string) bool {
	numbersValid := Client.IsNumbersValid(name)
	lengthValid := Client.IsLengthValid(name)
	return numbersValid && lengthValid
}

func (c *ValidationClient) IsNumbersValid(name string) bool {
	return !strings.ContainsAny(name, "0123456789")
}

func (c *ValidationClient) IsLengthValid(name string) bool {
	length := len(name)
	return length >= 8 && length <= 16
}
