package main

import (
	"fmt"
	"io"
	"net/http"
)

/*Creo un struct que implementa un interface*/
type escritorWeb struct {
}

func (escritorWeb) Write(p []byte) (int, error) {
	//fmt.Println(p) //imprimo un slice de bytes
	fmt.Println(string(p)) // imprime codigo html de la página web
	return len(p), nil
}

func main() {
	respuesta, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(respuesta)      //
	//fmt.Println(respuesta.Body) // apuntadores

	e := escritorWeb{}
	io.Copy(e, respuesta.Body)

}
