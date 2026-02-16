package views

import (
	"fmt"
	helps "tiendaapp/Utilidades"
)

func MainView() {

	opcion := -1

	for opcion != 0 {
		fmt.Println("1. Gestionar Usuarios")
		fmt.Println("2. Gestionar Productos")
		fmt.Println("3. Gestionar Pedidos")
		fmt.Println("4. Gestionar Proveedores")
		fmt.Println("5. Gestionar Inventario")
		fmt.Println("6. Gestionar Facturas")
		fmt.Println("0. Salir")
		fmt.Print("Seleccione una opción: ")
		fmt.Scanln(&opcion)

		helps.CleanScreen()

		switch opcion {
		case 1:
			UsuarioView()
		case 2:
			ProductoView()

		case 3:
			fmt.Println("Gestionar Pedidos")
			// Lógica para gestionar pedidos
		case 4:
			fmt.Println("Gestionar Proveedores")
			// Lógica para gestionar proveedores
		case 5:
			fmt.Println("Gestionar Inventario")
			// Lógica para gestionar inventario
		case 6:
			fmt.Println("Gestionar Facturas")
			// Lógica para gestionar facturas
		case 0:
			fmt.Println("Saliendo de la aplicación...")
			return
		default:
			fmt.Println("Opción no válida, por favor intente de nuevo.")
		}
	}

}
