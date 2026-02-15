package Clientes

type Cliente struct {
	Nombre    string
	Correo    string
	Direccion string
	Telefono  string
}

// Constructor
func NewCliente(nombre string, correo string, direccion string, telefono string) *Cliente {
	return &Cliente{
		Nombre:    nombre,
		Correo:    correo,
		Direccion: direccion,
		Telefono:  telefono,
	}
}

func EliminarCliente(c *Cliente) {
	c = nil
}
