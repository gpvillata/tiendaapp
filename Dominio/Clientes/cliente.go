package cliente

import "encoding/json"

type Cliente struct {
	Id          int
	Nombre      string
	RazonSocial string
	Correo      string
	Direccion   string
	Telefono    string
	TipoCliente string
}

func ClienteVacio() *Cliente {
	return &Cliente{
		Id:          0,
		Nombre:      "",
		Correo:      "",
		Direccion:   "",
		Telefono:    "",
		TipoCliente: "",
	}
}

// Constructor
func NewCliente(id int, nombre string, correo string, direccion string, telefono string) *Cliente {
	return &Cliente{
		Id:          id,
		Nombre:      nombre,
		Correo:      correo,
		Direccion:   direccion,
		Telefono:    telefono,
		TipoCliente: "individual",
		RazonSocial: "",
	}
}

func DeleteCliente(c *Cliente) {
	c = nil
}

func SaveCliente(c *Cliente) *Cliente {
	json.Marshal(c)
	return c
}
