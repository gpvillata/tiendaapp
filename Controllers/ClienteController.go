package Controllers

import (
	"fmt"
	clienteservice "tiendaapp/Services"
)

func NuevoCliente() {

	fmt.Println("En cliente controller nuevo cliente")
	clienteservice.NuevoCliente()

}
func EliminarCliente(id int) {

}

func BuscarCliente(id int) {

}
