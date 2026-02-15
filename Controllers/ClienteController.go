package Controllers

import (
	cliente "tiendaapp/Dominio/Clientes"
	services "tiendaapp/Services"
)

func NuevoCliente(cliente cliente.Cliente) {

	services.NuevoCliente(cliente)

}
func EliminarCliente(id int) {

}

func BuscarCliente(id int) {

}
