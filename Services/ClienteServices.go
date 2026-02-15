package services

import (
	cliente "tiendaapp/Dominio/Clientes"
	clienterepository "tiendaapp/Repositories"
)

func NuevoCliente(cliente cliente.Cliente) {

	clienterepository.NuevoCliente(cliente)
}

func EliminarCliente(id int) {

}
