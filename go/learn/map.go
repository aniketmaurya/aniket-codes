package main

import "fmt"

func learn_map() map[string]int {
	var phonebook = map[string]int{
		"Adam": 123456,
		"Eve":  874622,
	}
	for name, number := range phonebook {
		fmt.Println(name, number)
	}
	return phonebook

}
