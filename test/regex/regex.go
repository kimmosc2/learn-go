package main

import (
	"fmt"
	"regexp"
)

const text = `
My email0 is kimmosc2@163.com
email1 is butn@example.com
email2 is kimmo@qq.com.cn
@someone,q@@q, q@q@.com
`

func main() {
	re := regexp.MustCompile("([a-zA-Z0-9]+)@([a-zA-Z0-9]+)\\.([a-zA-Z0-9.]+)")
	match := re.FindAllStringSubmatch(text, -1)
	for _, m := range match {
		fmt.Println(m)
	}

}
