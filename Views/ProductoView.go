package views

import (
	"fmt"
	productocontroller "tiendaapp/Controllers"
	producto "tiendaapp/Dominio/Productos"
)

func ProductoView() {

	opcion := -1
	for opcion != 5 {
		fmt.Println("*****************GESTION DE PRODUCTOS **********")
		fmt.Println("1. Nuevo Producto")
		fmt.Println("2. Eliminar Producto")
		fmt.Println("3. Buscar Producto")
		fmt.Println("4. Listar los  Producto")
		fmt.Println("5. Salir de Productos")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			NuevoProductoView()
			return
		case 2:
			EliminarProductoView()
			return
		}
	}

}
func NuevoProductoView() {
	fmt.Println("******************GESTIÓN DE PRODUCTOS******************")
	fmt.Println("Ingrese los datos del producto:")
	var id int
	var nombre string
	var precio float64
	var cantidad int
	fmt.Print("ID: ")
	fmt.Scanln(&id)
	fmt.Print("Nombre: ")
	fmt.Scanln(&nombre)
	fmt.Print("Precio: ")
	fmt.Scanln(&precio)
	fmt.Print("Cantidad: ")
	fmt.Scanln(&cantidad)

	prod := producto.NewProducto(id, nombre, precio, cantidad)

	productocontroller.NuevoProducto(prod)
}

func EliminarProductoView() {
	fmt.Println("Ingrese el id del producto:")
	var id int
	fmt.Scanln(&id)
	productocontroller.EliminarProducto(id)

}
