package main

import (
	"strings"
)

func basename1(s string) string {
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	for i := len(s) - 1; i > 0; i-- {
		if s[i] == '.' {
			s = s[:i]
			break
		}
	}
	return s
}

func basename2(s string) string {
	if i := strings.LastIndex(s, "/"); i != -1 {
		s = s[i+1:]
	}
	if i := strings.LastIndex(s, "."); i != -1 {
		s = s[:i]
	}
	return s
}

