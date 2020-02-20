package main

import "fmt"

func main() {
   quanpailieA(1,2,3)
   quanpailieA(1,2,4)
   quanpailieA(1,3,4)
   quanpailieA(2,3,4)
   fmt.Printf("%d",4*3*2*1)
}

func quanpailieA(x,y,z int){
	jiaohuan(x,y,z)
	jiaohuan(y,x,z)
	jiaohuan(z,x,y)
}

func jiaohuan(t,m,n int){
	fmt.Printf("%d%d%d\n",t,m,n)
	fmt.Printf("%d%d%d\n",t,n,m)
}