package main

import (
	"bytes"
	"fmt"
)

// comma
func comma(s string) string {
	if len(s) <= 3 {
		return s
	}
	return comma(s[:len(s)-3]) + "," + s[len(s)-3:]
}

func comma2(s string) string {
	var buf bytes.Buffer
	for i := 1; i < len(s); i++ {
		if i % 3 == 0 {
			buf.WriteByte(',')
		}
		fmt.Fprintf(&buf,"%d",s[len(s)-i])
	}
	return buf.String()
}
