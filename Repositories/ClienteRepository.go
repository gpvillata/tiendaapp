package repositories

import (
	"fmt"
	cliente "tiendaapp/Dominio/Clientes"
)

func NuevoCliente(cli cliente.Cliente) {

	fmt.Println("Nuevo cliente:", cli)

}
