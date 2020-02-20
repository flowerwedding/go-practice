package main

import (
	"fmt"
	"time"
)

var (myres = make(map[int]int, 20)
     ch= make(chan int,100))
func factorial(n int) {
	var res = 1
	for i := 1; i <= n; i ++ {
		res *= i
	}
	ch<-res
	myres[n] =<- ch
	fmt.Println(myres[n])
}

func main() {
	for i := 1; i <= 20 ; i ++ {
		go factorial(i)
	}
	time.Sleep(100 * time.Second)
}