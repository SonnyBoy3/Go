
package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(`Enter the number of integers`)
	var n int
	
	if m, err := Scan(&n); m != 1 {
		panic(err)
	}
	
	fmt.Println(n)
	fmt.Println(`Enter the integers`)
	all := make([]int, n)
	ReadN(all, 0, n)
	fmt.Println(all)
}

func ReadN(all []int, i, n int) {
	if n == 0 {
		return
	}
	if m, err := Scan(&all[i]); m != 1 {
		panic(err)
	}
	ReadN(all, i+1, n-1)
}

func Scan(a *int) (int, error) {
	return fmt.Fscan(r, a)
}

var r = strings.NewReader(`5
10 20 30 50 60`)
