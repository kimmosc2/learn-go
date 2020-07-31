package main

import (
	"testing"
)

func TestComma(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		want string
	}{
		{"123456789", "123,456,789"},
		{"23456789", "23,456,789"},
		{"456789", "456,789"},
		{"123", "123"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comma(tt.name); got != tt.want {
				t.Errorf("comma() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestComma2(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{"1234567"}, "1,234,567"},
		{"1", args{"123456"}, "123,456"},
		{"1", args{"123"}, "123"},
		{"1", args{""}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := comma(tt.args.s); got != tt.want {
				t.Errorf("comma() = %v, want %v", got, tt.want)
			}
		})
	}
}
