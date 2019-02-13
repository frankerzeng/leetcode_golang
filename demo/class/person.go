package classTest

import (
	"fmt"
)

type Person struct {
	name string
}

func (p *Person) Init(name string) {
	p.name = name
}
func (p Person) Say() {
	fmt.Println(p.name)
}

type BeautPerson struct {
	Person
	age int
}

func (b *BeautPerson) Init(name string, age int) {
	b.Person.Init(name)
	b.age = age
}
