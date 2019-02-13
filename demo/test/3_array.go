package main

import "../class"

func main() {
	s := classTest.Student{}
	s.Name = "1"
	s.Init("pirlo", 21, "cs")
	s.SayHi()

	p := classTest.Person{}
	p.Init("z")
	p.Say()

	b := classTest.BeautPerson{}
	b.Init("x", 12)
	b.Say()
	b.Person.Say()

}
