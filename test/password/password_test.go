package main

import (
	"testing"
)

func Test_checkPassword(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"1",args{s: "12333ababccd"},false},
		{"2",args{s: "abcdefgg"},true},
		{"3",args{s: "abccab7676"},true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkPassword(tt.args.s); got != tt.want {
				t.Errorf("checkPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}