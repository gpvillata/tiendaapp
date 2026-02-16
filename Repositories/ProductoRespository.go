package repositories

import (
	"fmt"
	producto "tiendaapp/Dominio/Productos"
	helps "tiendaapp/Utilidades"
)

func NuevoProducto(producto producto.Producto) {

	helps.CleanScreen()
	fmt.Println("Nuevo producto creado:", producto)

}
func EliminarProducto(id int) {
	helps.CleanScreen()
	fmt.Println("Eliminando el producto:", id)
}
