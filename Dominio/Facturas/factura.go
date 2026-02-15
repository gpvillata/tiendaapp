package factura

import (
	cliente "tiendaapp/Dominio/Clientes"
	producto "tiendaapp/Dominio/Productos"
)

type Factura struct {
	Numero      string
	Fecha       string
	Total       float64
	Estado      string
	Cliente     cliente.Cliente
	Productos   []producto.Producto
	MetodoPago  string
	RazonSocial string
}
