package main
import "fmt"

func main() {

    var nome = "Abraao Allysson"
    var matricula = 11320000
	var idade int
    fmt.Println("Olá sr.", nome, "sua matricula é", matricula)
    fmt.Println("Qual sua idade? ")
	fmt.Scanf("%d",&idade)   
	fmt.Println("sua idade é", idade)

	if idade > 22 {
    	fmt.Println("SIM",idade," > ","22")
	}else {
		fmt.Println("NÃO",idade," < ","22")
	}
}