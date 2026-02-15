package main

import (
	"fmt"
	cliente "tiendaapp/Controllers"
	utilidades "tiendaapp/Utilidades"
)

func main() {

	fmt.Println("******************BIENVENIDO A LA TIENDA APP******************")
	opcion := 0
	for opcion != 3 {
		fmt.Println("1. Iniciar sesión")
		fmt.Println("2. Registrarse")
		fmt.Println("3. Salir")
		fmt.Scanln(&opcion)

		//limpiar pantalla
		utilidades.CleanScreen()
		switch opcion {
		case 1:
			fmt.Println("Iniciando sesión...")
		case 2:
			fmt.Println("Registrando usuario...")
			cliente.NuevoCliente()
		case 3:
			fmt.Println("Saliendo...")
			return
		default:
			fmt.Println("Opción inválida")
		}
		utilidades.CleanScreen()
	}

}
