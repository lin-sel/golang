package main

import "fmt"

func main() {

	// Declaration of map
	var m map[string]int

	fmt.Println(m)

	// Declaration with initialize map
	var mp = make(map[int]int)

	fmt.Println(mp)

	// Declare Variable with map literal
	var ms = map[string]int{
		"para1": 12,
		"para2": 13,
		"para4": 14,
		"para5": 15,
		"para6": 16,
	}

	for key, value := range ms {
		fmt.Println(key, value)
	}

	// assign one map to another map.
	var mk = ms
	// map are work on reference. here mk and ms pointing to same map. change in one reflact on other.

	mk["para4"] = 99

	fmt.Println(ms)

	// Deleting data from map.
	delete(mk, "para4")

	fmt.Println(ms)

}
