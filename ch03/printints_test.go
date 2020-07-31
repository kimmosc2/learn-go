package main

import (
	"testing"
)

func Test_intsToString(t *testing.T) {

	tests := []struct {
		name string
		args []int
		want string
	}{
		{"1", []int{1, 2, 3, 4}, "[1,2,3,4]"},
		{"2", []int{1, 2, 3}, "[1,2,3]"},
		{"3", []int{}, "[]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intsToString(tt.args); got != tt.want {
				t.Errorf("intsToString() = %v, want %v", got, tt.want)
			}
		})
	}
}
