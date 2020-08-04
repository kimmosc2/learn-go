package parser

import (
	"learn-go/project/crawler/engine"
	"regexp"
)

// cityListRe 正则表达式,匹配城市url和城市名称
const cityListRe = `<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`

// ParseCityList 城市列表解析器,contents正则匹配cityListRe
// 返回解析后的url和空解析器并将其包装为engine.ParseResult
func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)
	result := engine.ParseResult{}
	for _, m := range matches {
		result.Item = append(result.Item, string(m[2]))
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]),
			ParserFunc: engine.NilParser,
		})
	}
	return result
}
