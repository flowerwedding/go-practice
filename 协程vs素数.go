package main

import  "fmt"

func main(){
	for i:=1;i<=10000;i++{
		go sushu(i)
	}
}

func sushu(i int){
	var bol int
	for j:=2;j<i;j++{
		if i%j==0 { bol=1; break }
	}
	if bol==0 {fmt.Println(i)}
}