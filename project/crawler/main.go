package main

import (
	"learn-go/project/crawler/engine"
	"learn-go/project/crawler/scheduler"
	"learn-go/project/crawler/zhenai/parser"
)

func main() {

	// engine.SimpleEngine{}.Run(engine.Request{
	// 	Url:        "http://www.zhenai.com/zhenghun/",
	// 	ParserFunc: parser.ParseCityList,
	// })
	a := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{

		},
		WorkerCount: 10,
	}
	a.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
