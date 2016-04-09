package main

import "fmt"

func main() {
	
	
	p := Persona{"Pipe" , 21}
	fmt.Println(p.nombre, p.correr())

}
	

type Persona struct{
	nombre string
	estatura float64

}

func(p Persona) correr() string{
	return "corriendo"


}
