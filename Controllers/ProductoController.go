package Controllers

import (
	producto "tiendaapp/Dominio/Productos"
	services "tiendaapp/Services"
)

func NuevoProducto(prod producto.Producto) {
	services.NuevoProducto(prod)
}
