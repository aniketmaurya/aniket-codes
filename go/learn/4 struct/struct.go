package main

type person struct {
	name string
	age  int
}

func newPerson(name string, age int) person {
	p := person{name: name, age: age}
	return p
}

func (p *person) updateAge(age int) {
	p.age = age
}

func (p *person) updateName(name string) {
	p.name = name
}
