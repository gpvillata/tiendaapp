package views

import (
	"fmt"
	clientecontroller "tiendaapp/Controllers"
)

func UsuarioView() {
	fmt.Println("UsuarioView")
	opcion := -1

	for opcion != 0 {
		fmt.Println("1. Crear Usuario")
		fmt.Println("2. Listar Usuarios")
		fmt.Println("3. Eliminar Usuario")
		fmt.Println("0. Volver al menú principal")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			fmt.Println("Crear Usuario")
			clientecontroller.NuevoCliente()
			// Lógica para crear un usuario
		case 2:
			fmt.Println("Listar Usuarios")
			// Lógica para listar usuarios
		case 3:
			fmt.Println("Eliminar Usuario")
			// Lógica para eliminar un usuario
		case 0:
			fmt.Println("Volviendo al menú principal...")
			return
		default:
			fmt.Println("Opción no válida, por favor intente de nuevo.")
		}
	}

}
