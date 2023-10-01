package main

import (
	"errors"
	"fmt"
)

type ApiError struct {
	Path       string
	StatusCode int
}

func (e *ApiError) Error() string {
	return fmt.Sprintf("API error: %s returned a %d", e.Path, e.StatusCode)
}

func main() {
	err := fmt.Errorf("a height of %0.2f is invalid", -2.3333)
	fmt.Println(err.Error()) // a height of -2.33 is invalid

	err = errors.New("this is invalid height")
	fmt.Println(err.Error()) // this is invalid height

	err = &ApiError{Path: "/api/users", StatusCode: 500}
	fmt.Println(err.Error()) // API error: /api/users returned a 500
}
