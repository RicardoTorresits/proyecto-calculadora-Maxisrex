package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Uso: calculadora <número1> <operador> <número2>")
		os.Exit(1)
	}

	num1, err1 := strconv.ParseFloat(os.Args[1], 64)
	operador := os.Args[2]
	num2, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Los números ingresados no son válidos")
		os.Exit(1)
	}

	resultado := 0.0

	switch operador {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		if num2 == 0 {
			fmt.Println("No se puede dividir por cero")
			os.Exit(1)
		}
		resultado = num1 / num2
	default:
		fmt.Println("Operador no válido. Use +, -, *, o /")
		os.Exit(1)
	}

	fmt.Printf("Resultado: %.2f\n", resultado)
}
