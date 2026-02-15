package producto

type Producto struct {
	Id           int
	Nombre       string
	Precio       float64
	Cantidad     int
	Descripcion  string
	Imagen       string
	Categoria    string
	Subcategoria string
	Proveedor    string
	Estado       string
}

func ProductoVacio() *Producto {
	return &Producto{
		Id:       0,
		Nombre:   "",
		Precio:   0,
		Cantidad: 0,
	}
}
