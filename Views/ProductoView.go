package views

import (
	"fmt"
	productocontroller "tiendaapp/Controllers"
	producto "tiendaapp/Dominio/Productos"
)

func ProductoView() {
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
