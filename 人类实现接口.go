package main

import (
	 "fmt"
)

type gezi interface{
	gugugu()
}

func (p *person) gugugu(){
	fmt.Println(p.name,"又鸽了")
}

type fuduji interface{
	fudu(string)
}

func (p *person) fudu(word string){
	fmt.Println(word)
	fmt.Println(word)
}

type nimengjin interface{
	nimengshu()
}

func (p *person) nimengshu(){
	fmt.Println("我酸了")
}

type zhenxiangguai interface{
	zhenxiang()
}

func (p *person) zhenxiang(){
	fmt.Println("真香啊")
}

type person struct{
	name string
	age int
	gender string
}

func main(){
	p :=&person{
		name :"syr",
		age : 19,
		gender:"nv",
	}
	p.gugugu()
	p.fudu("hello world")
	p.nimengshu()
	p.zhenxiang()
}