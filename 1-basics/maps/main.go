package main

import (
	"fmt"
	"maps"
)

func main() {
	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["code"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"])

	delete(myMap, "key1")
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 10
	myMap["key3"] = 11
	fmt.Println(myMap)

	// clear(myMap)
	// fmt.Println(myMap)

	value, ok := myMap["key1"]
	if ok {
		fmt.Println("a value exists with key1")
	} else {
		fmt.Println("no value exists with key1")
	}
	fmt.Println(value)
	fmt.Println("is value associate with key1: ", ok)

	myMap2 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	myMap3 := map[string]int{"a": 1, "b": 2}
	fmt.Println(myMap2)

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("mymap3 and mymap2 are equal")
	}

	for k, v := range myMap3 {
		fmt.Println(k, v)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("the map is initialized to nil value")
	} else {
		fmt.Println("the map is not initialized to nil value")
	}

	val, ok := myMap4["key"]
	fmt.Println(val, ok)

	// myMap4["key"] = "value"
	// fmt.Println(myMap4)

	myMap4 = make(map[string]string)
	myMap4["key"] = "value"
	fmt.Println(myMap4)

	fmt.Println("myMap4 length is", len(myMap4))

	myMap5 := make(map[string]map[string]string)

	myMap5["map1"] = myMap4
	fmt.Println(myMap5)
}
