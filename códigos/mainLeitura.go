package main
import "fmt"

func main() {

    var nome = "Abraao Allysson"
    var matricula = 11320000
	var idade int
    fmt.Println("Olá sr.", nome, "sua  matricula é", matricula)
    fmt.Println("Qual sua idade? ")
	fmt.Scan(&idade) 

    if idade > 17 {
        fmt.Println("entrada permitida", idade)
    }   else {
         fmt.Println("entrada NÃO permitida", idade)
    }
	

	
}