package engine

// Request 包含一个url和解析函数
type Request struct {
	Url        string
	ParserFunc func([]byte) ParseResult
}

// ParseResult 包含一个Request slice
type ParseResult struct {
	Requests []Request
	Item     []interface{}
}

func NilParser(b []byte) ParseResult {
	return ParseResult{}
}