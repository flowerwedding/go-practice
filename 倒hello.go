package main

import "fmt"

func main(){
	var s string
    s ="hello"
    fmt.Printf("%s",daoxu(s))
}

func daoxu(s string)(ss string){
	for i:=len(s)-1;i>=0;i--{
		ss+=s[i:i+1]
	}
	return ss
}