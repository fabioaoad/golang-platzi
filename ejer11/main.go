package main

import (
	pk "curso-golang-platzi/ejer11/mypackage"
	"fmt"
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Honda"
	myCar.Year = 2023
	fmt.Println(myCar)

	pk.PrintMessage("Hola Fabio Martin Aoad")
}
