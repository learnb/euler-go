package main
import "fmt"

func Max(slice []uint64) uint64{
    var m uint64
	m = 0
	for _, elem := range slice {
		if m < elem {
			m = elem
		}
	}
	return m
}

//Returns all prime factors of a positive intager
func primeFactorsLinear(n uint64) []uint64 {
    factors := make([]uint64, 0, 10)
    var d uint64
    d = 2
    for n > 1 {
		for n%d == 0 {
			factors = append(factors, d)
			n /= d
        }
		d += 1
    }
	return factors
}

//Returns all prime factors of a positive integer
func primeFactorsSqrt(n uint64) []uint64 {
    factors := make([]uint64, 0, 10)
    var d uint64
    d = 2
    for n > 1 {
		for n%d == 0 {
			factors = append(factors, d)
			n /= d
        }
		d += 1
        if d*d > n {
            if n > 1 {
                factors = append(factors, n)
                break
            }
        }
    }
	return factors
}

func main(){
    fmt.Println(Max(primeFactorsLinear(13195)))
    fmt.Println(Max(primeFactorsSqrt(13195)))
    fmt.Println(Max(primeFactorsLinear(600851475143)))
    fmt.Println(Max(primeFactorsSqrt(600851475143)))
}
