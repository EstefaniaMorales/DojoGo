package main

import "fmt"

func main() {
  fmt.Print("Ingrese un texto: ")
  var input string
  fmt.Scanf("%s", &input)

  //output := input * 2

  fmt.Println(input)
}

/*EJERCICIO: Cambie el anterior programa para que en vez de capturar
	un decimal, capture un texto y lo imprima.*/