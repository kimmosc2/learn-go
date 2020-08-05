package parser

import (
	"fmt"
	"learn-go/project/crawler/engine"
	"regexp"
)

const resource = `<html>*</html>`

var ageRe = regexp.MustCompile(resource)

// Deprecated:网站有202反爬,需要解析js生成cookie
func ParseProfile(content []byte) engine.ParseResult {
	fmt.Println(string(content))
	all := ageRe.FindAll(content, -1)
	fmt.Println("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaa!!!!")
	fmt.Println(all)

	return engine.ParseResult{}
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	}
	return ""
}
