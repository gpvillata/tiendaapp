package services

import (
	cliente "tiendaapp/Dominio/Clientes"
	clienterepository "tiendaapp/Repositories"
)

func NuevoCliente() {

	cliente1 := cliente.NewCliente(1, "Piero", "", "Calle Falsa 123", "555-1234")

	clienterepository.NuevoCliente(*cliente1)

}

func EliminarCliente(id int) {

}
