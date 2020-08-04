package main

import (
	"learn-go/project/crawler/engine"
	"learn-go/project/crawler/zhenai/parser"
)

func main() {
	engine.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun/",
		ParserFunc: parser.ParseCityList,
	})
}


