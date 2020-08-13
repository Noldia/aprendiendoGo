package main //nombre del paquete

import "fmt" //importar librerias necesarias para ejecutar el programa

func main() {
	var mensaje string = "Hola Mundo"
	mensajeFacil := "Hola Mundo usando :="
	fmt.Println(mensaje)
	fmt.Println(mensajeFacil)
	//comentarios!!
	a := 10.
	var b float64 = 3
	fmt.Println((a / b))
	var c int = 10
	d := 3
	fmt.Println(c / d)
	x := true
	y := false
	fmt.Println((x || y))
	fmt.Println(y && x)
	fmt.Println(!x)
	privada()
	Publica()
	imprimirHola()
}

func privada() {
	fmt.Println("Ejecutar logica que no necesita ser exportada")
}

func Publica() {
	fmt.Println("Logica que quiero exportar")
}

func imprimirHola() {
	defer fmt.Println("Mundo")
	fmt.Println("Hola")
}
