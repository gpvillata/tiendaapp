package cliente

import (
	"fmt"
	cliente "tiendaapp/Dominio/Clientes"
)

func ClienteController() {
	fmt.Println("Creando cliente...")
	cliente := cliente.ClienteVacio()
	fmt.Println("Ingrese el nombre del cliente:")
	fmt.Scanln(&cliente.Nombre)
	fmt.Println("Ingrese el correo del cliente:")
	fmt.Scanln(&cliente.Correo)
	fmt.Println("Ingrese la dirección del cliente:")
	fmt.Scanln(&cliente.Direccion)
	fmt.Println("Ingrese el teléfono del cliente:")
	fmt.Scanln(&cliente.Telefono)
}
