package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

// Constructor de Employee
func NewEmployee(id int, name string, vaction bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vaction,
	}
}

func main() {
	e := Employee{}
	//fmt.Printf("%v", e)
	e.id = 1
	e.name = "Name"
	//fmt.Printf("%v", e)
	e.SetId(5)
	e.SetName("Name 2")
	//fmt.Printf("%v", e)
	fmt.Println(e.GetId())
	fmt.Println(e.GetName())

	//constructores

	//forma 1
	emp1 := Employee{}
	fmt.Printf("%v\n", emp1)

	//forma2
	emp2 := Employee{
		id:       1,
		name:     "Fabio",
		vacation: true,
	}
	fmt.Printf("%v\n", emp2)

	//forma 3
	emp3 := new(Employee)     //retorna el apuntador
	fmt.Printf("%v\n", emp3)  //muestra el puntero &, es decir, la referencia de emp3
	fmt.Printf("%v\n", *emp3) //muestra el valor con *
	emp3.id = 2
	emp3.name = "Martin"
	fmt.Printf("%v\n", *emp3)

	//forma 4 - recomendada
	emp4 := NewEmployee(3, "Aoad", true)
	fmt.Printf("%v\n", *emp4)
}
