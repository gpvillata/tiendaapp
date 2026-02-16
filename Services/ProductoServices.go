package services

import (
	producto "tiendaapp/Dominio/Productos"
	productorepository "tiendaapp/Repositories"
)

func NuevoProducto(producto producto.Producto) {
	productorepository.NuevoProducto(producto)
}

func EliminarProducto(id int) {
	productorepository.EliminarProducto(id)
}
