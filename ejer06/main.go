package main

import "fmt"

func main() {
	//DEFER
	// buena practica para cerrar archivos o cerrrar conexiones.
	// se puede usar mas de una sentencia defer en una funcion pero se recomiendoa un uso por funcion
	defer fmt.Println("Hola") //lo ultimo que imprime es esta linea
	fmt.Println("Mundo")

	//CONTINUE y BREAK
	//continuo  se usa para continuar con el codigo a pesar de que paso un error
	//break lo contrario a continuo

	for i := 0; i < 10; i++ {
		fmt.Println(i)

		//continue
		if i == 2 {
			fmt.Println("Es 2")
			continue
		}

		//break
		if i == 8 {
			fmt.Println("Break")
			break
		}
	}
}
