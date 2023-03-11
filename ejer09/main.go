package main

import "fmt"

func main() {
	//MAPS --> más eficientes que array y slice ya que implementan concurrencia
	m := make(map[string]int)

	m["Fabio"] = 21
	m["Martin"] = 32

	fmt.Println(m)

	//Recorrer map usando range
	// al imprimir map no necesariamente muestra ordenado como se insertaron los elementos
	for i, v := range m {
		fmt.Println(i, v)
	}

	//Encontrar un valor
	value := m["Fabio"]
	fmt.Println(value)

	value2 := m["Aoad"]
	fmt.Println(value2)

	value3, ok := m["FMA"]
	fmt.Println(value3, ok)

	value4, ok := m["Martin"]
	fmt.Println(value4, ok)

	//w := make(map[int][]string)
	w := make(map[int]string)
	m[2020] = "fabio"
	fmt.Println(w)

}
