package repositories

import (
	"fmt"
	cliente "tiendaapp/Dominio/Clientes"
	helps "tiendaapp/Utilidades"
)

func NuevoCliente(cli cliente.Cliente) {

	helps.CleanScreen()

	fmt.Println("Nuevo cliente:", cli)

}
