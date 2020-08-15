package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type calc struct{}

func (calc) operate(input string, operate string) int {
	fmt.Println("Resultado: ")
	cleanInput := strings.Split(input, operate)
	value1 := changeType(cleanInput[0])
	value2 := changeType(cleanInput[1])
	switch operate {
	case "+":
		return value1 + value2
	case "-":
		return value1 - value2
	case "*":
		return value1 * value2
	case "/":
		return value1 / value2
	default:
		fmt.Println(operate, "No soportado")
		return 0
	}
}

func changeType(input string) int {
	value, _ := strconv.Atoi(input)
	return value
}

func readInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	return scanner.Text()
}

func main() {
	fmt.Println("-------------Calculadora-----------------")
	fmt.Println("Ingresa una 2 numero a operar ejm: 2+2 ")
	input := readInput()
	fmt.Println("Ahora ingresa el operador ejm: + ")
	operate := readInput()
	c := calc{}
	fmt.Println(c.operate(input, operate))
}
