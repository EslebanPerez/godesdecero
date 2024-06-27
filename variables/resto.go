package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time
var variable string

func RestoVariables() {
	Nombre = "Pedro"
	Estado = true
	Sueldo = 1577.00
	Fecha = time.Now().UTC()
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
	variable = "alguito"
	fmt.Println(variable)
}

func ConviertoTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}

func NewFunction() {
	hola, mundo := "Hola", "🌎"
	mundo, happy := "mundo", "😃"
	fmt.Println(hola, mundo, happy)
}

func Circulo() {
	const PI = 3.1416
	// PI = 3.15 Las constantes no se pueden reasignar un valor
	fmt.Println(PI)
}
