package classTest

import "fmt"

type Student struct {
	name  string
	Name  string
	age   int
	major string
}

func (s Student) SayHi() {
	fmt.Printf("Hi, I am %s aged %d, and my major is %s\n", s.name, s.age, s.major)
}

func (s *Student) Init(name string, age int, major string) {
	s.name = name
	s.age = age
	s.major = major
}
