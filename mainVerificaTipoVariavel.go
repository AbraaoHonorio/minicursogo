package main
import "fmt"
import "reflect"
func main() {
    nome := "Abraao Allysson"
    var matricula = 11320000
    var versao = 1.1
    fmt.Println("Olá sr.", nome, "sua matricula é", matricula)
    fmt.Println("Este programa está na versão", versao)
    fmt.Println("O tipo da variável idade é", reflect.TypeOf(versao))
}