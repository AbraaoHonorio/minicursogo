package main

import "fmt"

func main() {
	 var numero1,numero2 float32
	  fmt.Println("digite o primeiro numero 1: ")
	  fmt.Scan(&numero1) 
	  fmt.Println("digite o primeiro numero 2: ")
	  fmt.Scan(&numero2) 
	  fmt.Println(numero1, "+",numero2 ," = ")
      fmt.Println(somar(numero1 ,numero2))
}

func somar(numero1 float32, numero2 float32) float32 {
	return numero1+numero2
}
