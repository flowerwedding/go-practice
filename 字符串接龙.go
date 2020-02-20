package main

import(
	 "fmt"
)

func main(){
	d:=666
	ptr:=&d
	ptr1:=&ptr
	ptr2:=&ptr1
	fmt.Println(ptr2)
	fmt.Println(*ptr2)
	fmt.Println(**ptr2)
	fmt.Println(***ptr2)
}