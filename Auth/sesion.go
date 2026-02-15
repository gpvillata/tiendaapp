package auth

import "fmt"

func Login() {
	var usuario string
	var password string
	fmt.Println("*****************INICIO DE SESION**************")
	fmt.Println("Usuario:")
	fmt.Scan(usuario)
	fmt.Println("Password:")
	fmt.Scan(password)
	if usuario != "" && password != "" {
		fmt.Println("Iniciando sesion.....", usuario)
	} else {
		fmt.Println("Error al inciar sesion")
	}

}
