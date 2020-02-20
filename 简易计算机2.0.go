package main

import (
	"fmt"
	"strconv"
)

func main(){
	s :="111+222*(333-444)"
    ans:=kuohao(s)
    fmt.Printf("%d",ans)
}

func fd(f *[999]string,sz *[1000]int64) (ans int64){
	var sz1 [1000]int64
	var f1 [999]string
	var j int
	ans = 0
	j = 0
	for i:=0;i<len(f);i++ {
		if f[i] == "*" {
			sz[i+1] = sz[i] * sz[i+1]
		} else if f[i] == "/" {
			sz[i+1] = sz[i] / sz[i+1]
		} else if f[i]=="" {
			sz1[j]=sz[i];break
	    }else {
			sz1[j]=sz[i];f1[j]=f[i];j++}
	}
	ans+=sz1[0]
	for i:=0;i<len(f1);i++{
		if f1[i]=="+" {ans+=sz1[i+1]
		} else if f1[i]=="-" {
			ans-=sz1[i+1]
		} else { break }
	}
	return ans
}

func kuohao(s string)(ans int64){
	var f [999]string
	var sz [1000]int64
	var conten string
	var ss string
	var j,n int
	j = 0
	b:=[]byte(s)
	n=len(b)
	for i:=0;i<n;i++{
		conten = string(b[i])
		if (conten=="*")||(conten=="-")||(conten=="+")||(conten=="/") {
			f[j] = conten;
			j++;
		}else if conten=="("{
			d:=i
			for{
				d++
				if string(b[d])==")" {
					ss = string(b[i+1 : d])
					sz[j] = kuohao(ss)
					break
				}
			}
			i=d
		} else {
			t, _ :=strconv.ParseInt(conten, 0, 64);
			sz[j]=sz[j]*10+t
		}
	}
	ans=fd(&f,&sz)
	return ans
}