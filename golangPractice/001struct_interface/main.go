package main

import (
	"log"
	"reflect"
)

// zoneDbParam xxx
var zoneDbParam struct {
	pName string
	zName string
}

// test xxx
func test(p interface{}) {
	// pp := p.(*zoneDbParam)
	// param := reflect.New(reflect.TypeOf(p).Elem())
	param := reflect.TypeOf(p).Elem()
	log.Println(param)
	// log.Println(pp.pName, pp.zName)
}

func main() {
	log.Println("hello struct")
	// var p interface{}
	p := "{\"pName\":\"nz\",\"zName\",\"hf\"}"
	// p interface{}= "hello"
	test(p)
}
