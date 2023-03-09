package main

import "fmt"

type pc struct {
	ram   int
	disk  int
	brand string
}

func (myPC pc) ping() {
	fmt.Println(myPC.brand, "Pong")
}

func (myPC *pc) duplicateRAM() {
	myPC.ram = myPC.ram * 2
}

func main() {

	//Puntero --> usado en funcionas para poder llevar las funciones o otro paquete eficientemente
	a := 50
	b := &a // puntero de a en b
	fmt.Println(a)
	fmt.Println(b) //dir de memoria

	fmt.Println(*b) //valor de a en b

	*b = 100
	fmt.Println(*b) //100
	fmt.Println(a)  //100

	myPC := pc{ram: 20, disk: 500, brand: "ASUS"}
	fmt.Println(myPC)

	myPC.ping()

	fmt.Println(myPC) //20 de ram

	myPC.duplicateRAM()
	fmt.Println(myPC) // 40 de ram

	myPC.duplicateRAM()
	fmt.Println(myPC) //80 de ram
}
