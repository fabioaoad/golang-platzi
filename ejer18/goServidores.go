package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	for _, servidor := range servidores {
		go revisarServidor(servidor, canal)
	}

	for i := 0; i < len(servidores); i++ {
		fmt.Println(<-canal)
	}

	tiempoPaso := time.Since(inicio)
	fmt.Printf("Tiempo de ejecución %s\n", tiempoPaso)
}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		fmt.Println(servidor, "No esta disponible :(")
		canal <- servidor + " No esta disponible :( - respuesta canal"
	} else {
		fmt.Println(servidor, "Esta disponible :)")
		canal <- servidor + "  Esta disponible :) - respuesta canal"

	}
}
