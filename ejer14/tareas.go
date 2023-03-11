package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}
func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre: ", tarea.nombre)
		fmt.Println("Descripcion: ", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado {
			fmt.Println("Nombre: ", tarea.nombre)
			fmt.Println("Descripcion: ", tarea.descripcion)
		}
	}
}

func (t *task) marcarCompletado() {
	t.completado = true
}
func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t1 := &task{
		nombre:      "Completar curso de Go",
		descripcion: "Realizarlo en 5 días",
	}

	t2 := &task{
		nombre:      "Completar curso de Python",
		descripcion: "Realizarlo en 10 días",
	}

	t3 := &task{
		nombre:      "Completar curso de NodeJs",
		descripcion: "Realizarlo en 30 días",
	}
	//fmt.Printf("%+v\n", t)
	//t.marcarCompletado()
	//t.actualizarNombre("Final")
	//t.actualizarDescripcion("completar")

	//fmt.Printf("%+v\n", t)

	lista := &taskList{
		tasks: []*task{
			t1, t2,
		},
	}
	//fmt.Println(lista.tasks[0])
	lista.agregarALista(t3)
	//fmt.Println(len(lista.tasks))
	//lista.eliminarDeLista(1)
	//fmt.Println(len(lista.tasks))

	for i := 0; i < len(lista.tasks); i++ {
		fmt.Println("Index ", i, "nombre ", lista.tasks[i].nombre)
	}

	for index, tarea := range lista.tasks {
		fmt.Println("Index ", index, "nombre ", tarea.nombre)
	}

	//lista.imprimirLista()
	lista.tasks[0].marcarCompletado()
	fmt.Println("Tareas Completadas")
	lista.imprimirListaCompletados()

}
