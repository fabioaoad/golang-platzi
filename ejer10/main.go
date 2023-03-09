package main

import "fmt"

type car struct {
	brand string
	year  int
}

func main() {
	mycar := car{brand: "Honda", year: 2023}
	fmt.Println(mycar)

	//Otra manera
	var otherCar car
	otherCar.brand = "Toyota"
	fmt.Println(otherCar)
}
