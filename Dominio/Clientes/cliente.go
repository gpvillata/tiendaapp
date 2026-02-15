package cliente

import "encoding/json"

type Cliente struct {
	Nombre      string
	RazonSocial string
	Correo      string
	Direccion   string
	Telefono    string
	TipoCliente string
}

func ClienteVacio() *Cliente {
	return &Cliente{
		Nombre:      "",
		Correo:      "",
		Direccion:   "",
		Telefono:    "",
		TipoCliente: "",
	}
}

// Constructor
func NewCliente(nombre string, correo string, direccion string, telefono string) *Cliente {
	return &Cliente{
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
