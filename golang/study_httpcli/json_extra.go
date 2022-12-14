package main

import (
	"encoding/json"
	"fmt"
	"log"

	// jsoniter "github.com/json-iterator/go"
	"github.com/oliveagle/jsonpath"
)

func Parse(data string) {
	// var json = jsoniter.ConfigCompatibleWithStandardLibrary
	var jo any
	json.Unmarshal([]byte(data), &jo)
	// res, err := jsonpath.JsonPathLookup(jo, "$.expensive")
	// for reuse
	pat, err := jsonpath.Compile("$.expensive")
	if err != nil {
		log.Fatal("compile json path failed ", err.Error())
	}

	res, err := pat.Lookup(jo)

	// 向下取整数部分
	// res = math.Floor(res.(float64))
	res = int(res.(float64))
	if err != nil {
		log.Fatal("lookup failed ", err.Error())
	}
	fmt.Printf("%v \n", res)
}
