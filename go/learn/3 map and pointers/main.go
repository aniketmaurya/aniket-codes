package main

import "fmt"

func main() {
	learn_map()

	// Learn pointers
	var ages []int = []int{1, 2, 3, 4, 5}
	fmt.Println(ages)
	update_array(ages)
	fmt.Println(ages)
	name := "lucifer"
	updateName(&name)
	fmt.Println(name)

}
