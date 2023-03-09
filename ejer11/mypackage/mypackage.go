package mypackage

import "fmt"

// carPrivate car con acceso privado --> inicia con minuscula
type carPrivate struct {
	brand string
	year  int
}

// CarPublic car con acceso público --> inicia con mayuscula
type CarPublic struct {
	Brand string
	Year  int
}

// PrintMessage para imprimir un mensaje. la fun es pública ya que inicia con mayuscula
func PrintMessage(text string) {
	fmt.Println(text)
}
