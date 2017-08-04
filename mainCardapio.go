package main

import "fmt"

func main() {
    var prato string
    fmt.Println("Digite seu prato preferido...")
    fmt.Println("P - Pizza")
    fmt.Println("L - Lasanha")
    fmt.Println("C - Cachorro-quente")
    fmt.Println("H - Hambúrguer")
    fmt.Println("B - Batata frita")
    fmt.Println("S - Salada ")
    fmt.Println("O - Outros")
    fmt.Scan(&prato)

    switch prato {
    case "P":
        fmt.Println("Calabresa ou muçarela?")
     case "L":
        fmt.Println("Carne ou Frango?")
    case "C":
        fmt.Println("linguiça ou salsicha?")
    case "H":
        fmt.Println("Com Queijo ou com Ovo?")
    case "B":
        fmt.Println("Com batatas Palito ou Noisettes?")
   case "S":
        fmt.Println("É sério isso? Tu queres salada?")
    case "O":
        fmt.Println("Não gostou de nosso cardápio?")
    default:
        fmt.Println("Não entendi seu paladar...")
    }
}