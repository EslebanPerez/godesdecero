// El paquete main es el paquete principal
package main

import (
	"fmt"

	"github.com/eslebanperez/godesdecero/variables"
)

// Funcion main de nuestro paquete
func main() {
	// Estado y texto es resultado del retorno de la función ConviertoTexto
	estado, texto := variables.ConviertoTexto(22)

	fmt.Println(estado)
	fmt.Println(texto)

	/*
		Comentarios multilinea
		se utilizan para código que no se va a utilizar
		variables.Circulo()
	*/
}
