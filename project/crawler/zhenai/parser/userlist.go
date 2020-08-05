package parser

import (
	"learn-go/project/crawler/engine"
	"regexp"
)

const userUrlReg = `<a href="(http://album.zhenai.com/u/[0-9]+)" target="_blank">([^<]+)</a>`

var userReg = regexp.MustCompile(userUrlReg)

func ParseUserList(contents []byte) engine.ParseResult {
	result := engine.ParseResult{}
	user := userReg.FindAllSubmatch(contents, -1)
	for _, u := range user {
		result.Item = append(result.Item, string(u[2]))
	}
	return result
}
