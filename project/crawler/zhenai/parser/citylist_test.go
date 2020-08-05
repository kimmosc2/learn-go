package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCityList(t *testing.T) {
	file, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(file)
	const resultSize = 470
	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectedCities := []string{
		"阿坝", "阿克苏", "阿拉善盟",
	}
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests;but had %d", resultSize, len(result.Requests))
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s;but was %s", i, url, result.Requests[i].Url)
		}
	}
	for i, city := range expectedCities {
		if result.Item[i].(string) != city {
			t.Errorf("expected city #%d: %s;but was %s", i, city, result.Item[i].(string))
		}
	}
	if len(result.Item) != resultSize {
		t.Errorf("result should have %d item;but had %d", resultSize, len(result.Requests))
	}
}
