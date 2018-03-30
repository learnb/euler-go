package main
import "fmt"

// Sequential Fib function that sums all even Fibs under 4 million
func seqFibSum() uint64{
    var sum uint64
    var a uint64
    var b uint64
    var swap uint64
    a = 1
    b = 1
    swap = 0
    sum = 0
    for c := 3; b <= 4000000; c++{
        swap = b
        b += a
        a = swap
        if b%2 == 0{
            sum += b
        }
    }
    return sum
}

// Recursive Fib function
func fib(x int) int{
    if x <= 1 {
        return x
    }
    return fib(x-1) + fib(x-2)
}

// Sequential Fib function
func seqFib(x int) uint64{
    if x <= 2 {
        return 1
    }
    var a uint64
    var b uint64
    var swap uint64
    a = 1
    b = 1
    swap = 0
    for c := 3; c <= x; c++{
        swap = b
        b += a
        a = swap
    }
    return b
}

func main() {
    /*
    for i := 1; i <= 100; i++{
        //fmt.Println(fib(i))
        fmt.Printf("%d: %d\n", i, seqFib(i) )
        //fmt.Println("---")
    }
    */
    fmt.Println(seqFibSum())
}
