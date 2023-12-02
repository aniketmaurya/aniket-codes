package main

import (
	"fmt"
)

func main() {
	p := newPerson("Lucifer", 1000)
	fmt.Println(p)
	p.updateAge(7)
	fmt.Println(p)

}
