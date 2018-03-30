package main
import "fmt"
import "strconv"

func Max(slice []int) int{
    var m int
        m = 0
        for _, elem := range slice {
                if m < elem {
                        m = elem
                }
        }
        return m
}

func isPalindrome(x int) bool {
    //str := strconv.FormatUint(x, 10)
    str := strconv.Itoa(x)
    l := len(str)
    count := l/2
    passed := true
    //fmt.Println(l)
    //fmt.Println(str)
    //for _, v := range str {
    //    fmt.Println(v)
    //}

    for i:=0; i <= count; i++ {
        if str[0+i] != str[l-1-i] {
            passed = false
        }
    }
    return passed
}

func main() {
   //fmt.Println(isPalindrome(1343))
    list := make([]int, 0, 10)
    for i:=999; i>=100; i-- {
        for j:=999; j>=100; j-- {
            if isPalindrome(i*j) {
               list = append(list, i*j)
            }
        }
    }
    //for _, v := range list {
    //    fmt.Println(v)
    //}
	fmt.Println(Max(list))
}
