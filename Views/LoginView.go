package views

import (
	"fmt"
	auth "tiendaapp/Auth"
)

func LoginView() {
	var usuario, password string

	fmt.Print("Usuario: ")
	fmt.Scanln(&usuario)

	fmt.Print("Contraseña: ")
	fmt.Scanln(&password)

	if auth.Login(usuario, password) {
		MainView()
	} else {
		fmt.Println("Credenciales incorrectas")
	}
}
