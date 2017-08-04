package main
import "fmt"

func main() {

    nome := "Abraao Allysson"
	const PrecoCafe = 2.50
	var dinheiro float64
    fmt.Println("Olá sr.", nome)
    fmt.Println("Preço do café:", PrecoCafe)
	fmt.Scanf("%f",&dinheiro)   
	
	if dinheiro > PrecoCafe {
    	fmt.Println("Saldo insuficiente")
	}else {
		fmt.Println("Café saindo na hora...")
	}
}