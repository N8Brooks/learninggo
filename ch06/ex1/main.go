package main

type Person struct {
	firstName string
	lastName  string
	age       int
}

func makePerson(firstName string, lastName string, age int) Person {
	return Person{firstName, lastName, age}
}

func makePersonPointer(firstName string, lastName string, age int) *Person {
	return &Person{firstName, lastName, age}
}

func main() {
	p1 := makePerson("John", "Doe", 30)
	p2 := makePersonPointer("Jane", "Doe", 25)
	println(p1.firstName, p1.lastName, p1.age)
	println(p2.firstName, p2.lastName, p2.age)
}
