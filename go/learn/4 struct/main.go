package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(prompt string, r *bufio.Reader) string {
	fmt.Print(prompt)
	input, _ := r.ReadString('\n')

	return strings.TrimSpace(input)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	p := newPerson("Lucifer", 1000)
	fmt.Println(p)

	name := getInput("Enter new name: ", reader)
	p.updateName(name)

	p.updateAge(7)
	fmt.Println(p)
}
