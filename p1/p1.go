package main
import "fmt"

func p1(limit int) int {
    sum := 0
    for i := 1; i < limit; i++ {
        if (i%3 == 0) || (i%5 == 0) {
            sum += i
        }
    }
    return sum
}

func main() {
    fmt.Println(p1(1000))
}
