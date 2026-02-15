package views

import (
	"fmt"
	clientecontroller "tiendaapp/Controllers"
	cliente "tiendaapp/Dominio/Clientes"
	helps "tiendaapp/Utilidades"
)

func UsuarioView() {
	fmt.Println("******************GESTIÓN DE USUARIOS******************")
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
			helps.CleanScreen()

			var nombre string
			var razonsocial string
			var correo string
			var direccion string
			var telefono string

			fmt.Println("Crear Cliente")
			fmt.Println("Nombre:")
			fmt.Scanln(&nombre)
			fmt.Println("Razón Social:")
			fmt.Scanln(&razonsocial)
			fmt.Println("Correo:")
			fmt.Scanln(&correo)
			fmt.Println("Dirección:")
			fmt.Scanln(&direccion)
			fmt.Println("Teléfono:")
			fmt.Scanln(&telefono)

			cliente := cliente.NewCliente(0, nombre, correo, direccion, telefono)

			clientecontroller.NuevoCliente(cliente)
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
