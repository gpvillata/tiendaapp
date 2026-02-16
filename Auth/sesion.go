package auth

import "fmt"

func Login(usuario, password string) bool {

	if usuario == "admin" && password == "admin123" {
		fmt.Println("Login exitoso")
		return true
	}
	fmt.Println("Login fallido")
	return false

}
