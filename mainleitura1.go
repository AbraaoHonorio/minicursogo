package main
import "fmt"
func main() {
    var nome = "Abraao Allysson"
    var matricula = 11320000
	 var idade
    fmt.Println("Olá sr.", nome, "sua matricula é", matricula)
    fmt.Println("Qual sua idade? ")
	 fmt.Scanf("%d",&idade)  
    fmt.Println("sua idade é ", idade)
}
