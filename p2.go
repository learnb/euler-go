package main

func p2(x int) (int, int){
    val, even := fib(x)
    return val, even
}

func fib(x int) (int, int){
    if x == 0 {
        return 1, 0
    } else if x == 1 {
        return 2, 2 
    } else {
        a, _ := fib(x-1)
        b, _ := fib(x-2)
        
        res := a+b
        even := 0
        if res%2 == 0 {
            even = res
        }
        
        return res, even
    }
}

func fibSum(x int) int{
    if x == 0 {
        return 1
    } else if x == 1 {
        return 2
    } else {
        return fibSum(x-1) + fibSum(x-2)
    }
}
