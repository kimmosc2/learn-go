package main

import "testing"

func TestBasename(t *testing.T) {

	tests := []struct {
		path, name string
	}{
		{"/home/root/a.txt", "a"},
		{"/home//root/abc.t", "abc"},
		{"/home/root/a", "a"},
		{"/home/root/a.txt.back", "a.txt"},
		{"//home/root/a.txt.back", "a.txt"},
	}
	for _, test := range tests {
		if n := basename1(test.path); n != test.name {
			t.Errorf("path:%v, name:%v, program answer:%v\n", test.path, test.name, n)
		}
	}
}

func TestBasename2(t *testing.T) {

	tests := []struct {
		path, name string
	}{
		{"/home/root/a.txt", "a"},
		{"/home//root/abc.t", "abc"},
		{"/home/root/a", "a"},
		{"/home/root/a.txt.back", "a.txt"},
		{"//home/root/a.txt.back", "a.txt"},
	}
	for _, test := range tests {
		if n := basename2(test.path); n != test.name {
			t.Errorf("path:%v, name:%v, program answer:%v\n", test.path, test.name, n)
		}
	}
}
