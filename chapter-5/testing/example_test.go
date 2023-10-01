package main

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

func TestGetBigger(t *testing.T) {
	type args struct {
		a float64
		b float64
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "a is the bigger one",
			args: args{
				a: 9,
				b: 6,
			},
			want: 9,
		},
		{
			name: "b is the bigger one",
			args: args{
				a: 3,
				b: 7,
			},
			want: 7,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetBigger(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GetBigger() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsValidName(t *testing.T) {
	tests := []struct {
		testName string
		name     string
		want     bool
	}{
		{
			testName: "Valid name",
			name:     "Adam Adam",
			want:     true,
		},
		{
			testName: "Invalid name due to number",
			name:     "Adam Adam12",
			want:     false,
		},
		{
			testName: "Invalid name due to length",
			name:     "A",
			want:     false,
		},
	}
	mockedClient := &MockValidationClientInter{}
	mockedClient.On("IsNumbersValid", mock.Anything).Return(true)
	mockedClient.On("IsLengthValid", mock.Anything).Return(true)
	Client = mockedClient
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			if got := IsValidName(tt.name); got != tt.want {
				t.Errorf("IsValidName() = %v, want %v", got, tt.want)
			}
		})
	}
}
