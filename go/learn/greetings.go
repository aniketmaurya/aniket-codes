package main

import "strings"

func getInitials(name string) (string, string) {
	name = strings.ToUpper(name)
	var initials []string

	names := strings.Split(name, " ")

	for _, v := range names {
		initials = append(initials, v[:1])
	}

	if len(initials) > 1 {
		return initials[0], initials[1]
	}

	return initials[0], "_"

}
