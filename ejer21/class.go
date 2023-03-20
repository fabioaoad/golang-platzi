package main

import "fmt"

type Person struct {
	name string
	age  int
}

type Employee struct {
	id int
}

type FullTimeEmployee struct {
	Person
	Employee
}

func getMessage(p Person) {
	fmt.Printf("%s with age %d\n", p.name, p.age)
}

func main() {
	ftEmployee := FullTimeEmployee{}
	ftEmployee.name = "Fabio"
	ftEmployee.age = 21
	ftEmployee.id = 5
	fmt.Printf("%v\n", ftEmployee)
	//getMessage(ftEmployee) //no permite el polimorfismo
}
