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
