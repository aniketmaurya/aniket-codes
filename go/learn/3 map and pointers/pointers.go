package main

import "strings"

func update_array(arr []int) {
	arr[0] = -1
}

func updateName(n *string) {
	*n = strings.ToUpper(*n)
}
