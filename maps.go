package main

import "fmt"

func main() {

	var M map[string]int
	fmt.Println(M == nil)

	usableMap := make(map[string]string)
	fmt.Println(usableMap)
	usableMap["mujahid"] = "cute"
	fmt.Println(usableMap)

	//intializing map with data

	newMap := map[string]int{
		"amit":    29,
		"mujahid": 21,
	}
	fmt.Println(newMap)

	//accessing element

	fmt.Println(newMap["amit"])

	value, boolean := newMap["usman"]
	fmt.Println(value, boolean)

	newMap["usman"] = 44
	fmt.Println(newMap)
	//deleting from map
	delete(newMap, "usman")
	fmt.Println(newMap)

	//iterating over map

	for key, value := range newMap {
		fmt.Println(key, "->", value)
	}

	//duplicating a map
	newmapCopy := newMap

	//altering data, will reference same data

	newmapCopy["hareem"] = 30
	fmt.Println(newMap)

}
