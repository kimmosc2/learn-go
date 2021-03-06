package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"strings"
)

var client = http.DefaultClient

func init() {
	jar, _ := cookiejar.New(nil)
	client.Jar = jar
}

// Fetch 会自动解码url的网页,并返回解码后的byte slice
// 如果http响应码不为200,则返回nil,error，或是读取内容
// 的时候发生错误,则返回nil,error
func Fetch(url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, strings.NewReader(""))
	if err != nil {
		log.Printf("build request error:%s", err.Error())
		return nil, err
	}
	// TODO http202反爬
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("wrong status code:%d", resp.StatusCode)
	}
	// 使用NewReader存储body,再使用Peek,不更新读取指针，最后读取reader
	reader := bufio.NewReader(resp.Body)
	peek, err := reader.Peek(1024)
	if err != nil {
		return nil, err
	}
	e := determineEncoding(peek)
	utf8code := transform.NewReader(reader, e.NewDecoder())
	return ioutil.ReadAll(utf8code)
}

// determineEncoding 猜测body的编码格式并返回
func determineEncoding(body []byte) encoding.Encoding {
	e, _, _ := charset.DetermineEncoding(body, "")
	return e
}
