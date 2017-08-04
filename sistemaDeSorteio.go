package main
import(
    "fmt"
    "math/rand"
    "time"
)

func random(min, max int) int {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(max - min) + min
}

func main() {
    myrand := random(1, 14)
    fmt.Println(myrand)
}