package main

import (
	"fmt"

	"github.com/eslebanperez/godesdecero/variables"
)

func main() {
	// variables.RestoVariables()
	estado, texto := variables.ConviertoTexto(22)

	fmt.Println(estado)
	fmt.Println(texto)

	variables.NewFunction()

}
