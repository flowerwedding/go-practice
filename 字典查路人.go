package main

import (
	"fmt"
	"strconv"
)

func main(){
	//生成路人
	mapCreatPeople:=make(map[string]map[string]map[string]string)
	str:="计算机"
	for n:=1;n<=3;n++{
		for i:=0;i<=9;i++ {
			for j := 0; j <= 9; j++ {
				        mapCreatPeople["2019"+"00"+strconv.Itoa(n)+strconv.Itoa(i)+strconv.Itoa(j)]=map[string]map[string]string{
					    str+"路人"+strconv.Itoa(i)+strconv.Itoa(j)+"号":{
						str + strconv.Itoa(i) + "班":str+"学院",
					},
			    }
		    }
	    }
	    str ="光电"
	    if n==2 {str="网安"}
    }
	//查询路人

	var s string
    s="201900100"
	if len(s)==9{//学号
		if v,ok:=mapCreatPeople[s];ok{
			fmt.Println(v)
		}
	}else if len(s)==13||len(s)==10 { //班级
	    str:="计算机"
		for i:=201900100;i<=201900399;i++ {
			if _, ok := mapCreatPeople[strconv.Itoa(i)][str+"路人"+strconv.Itoa(i%100)+"号"][s]; ok {
				fmt.Println(i,str+"路人"+strconv.Itoa(i%100)+"号")
			}
			str="光电"
			if i>=201900299 {str="网安"}
		}
	}else{//姓名
	 for i:=201900100;i<=201900399;i++{
	 	if v,ok:=mapCreatPeople[strconv.Itoa(i)][s];ok{
	 		fmt.Println(v)
		}
	 }
	}
}