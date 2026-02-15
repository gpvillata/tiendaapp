package clientes

type Cliente struct {
	Nombre    string
	Correo    string
	Direccion string
	Telefono  string
}

func (c *Cliente) SetNombre(nombre string) {
	c.Nombre = nombre
}

func (c *Cliente) SetCorreo(correo string) {
	c.Correo = correo
}

func (c *Cliente) SetDireccion(direccion string) {
	c.Direccion = direccion
}

func (c *Cliente) SetTelefono(telefono string) {
	c.Telefono = telefono
}

func (c *Cliente) GetNombre() string {
	return c.Nombre
}

func (c *Cliente) GetCorreo() string {
	return c.Correo
}

func (c *Cliente) GetDireccion() string {
	return c.Direccion
}
