package main

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"net/http"
	"regexp"
)

func main() {
	resp, err := http.Get("http://www.zhenai.com/zhenghun")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Error: status code ", resp.StatusCode)
		return
	}
	// 使用NewReader存储body,再使用Peek,不更新读取指针，最后读取reader
	reader := bufio.NewReader(resp.Body)
	peek, err := reader.Peek(2014)
	if err != nil {
		panic(err)
	}
	e := determineEncoding(peek)
	utf8code := transform.NewReader(reader, e.NewDecoder())
	all, err := ioutil.ReadAll(utf8code)
	printCityList(all)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(all))

}

func printCityList(contents []byte) {
	re := regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

	matches := re.FindAllSubmatch(contents, -1)
	for _, m := range matches {
		fmt.Printf("city:%s url:%s\n", m[2], m[1])
	}
}

func determineEncoding(body []byte) encoding.Encoding {
	e, name, determine := charset.DetermineEncoding(body, "")
	fmt.Printf("name:%s ,determine:%t\n", name, determine)
	return e
}
