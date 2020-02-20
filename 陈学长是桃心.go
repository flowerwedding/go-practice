package main

import "fmt"

func main(){
	C()
	fmt.Printf("\n")
	Z()
	fmt.Printf("\n")
	L()
}

func C(){
	shanghuxian6()
	shuxian4()
	xiahuxian6()
}

func Z(){
  henxian7()
  xiexian7()
  henxian7()
}

func L(){
  shuxian4()
  shuxian4()
  henxian7()
}

func shuxian4(){
	for i:=1;i<=4;i++{
	fmt.Printf("**\n")
	}
}

func henxian7(){
	for i:=1;i<=7;i++{
		fmt.Printf("**")
	}
	fmt.Printf("\n")
}

func xiexian7(){
	for i:=1;i<=7;i=i+1{
		for j:=1;j<=7;j++{
			if j==8-i{
				fmt.Printf("**")
			} else{
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
}

func shanghuxian6(){
    for i:=1;i<=3;i++{
    	fmt.Printf("  ")
    	for j:=1;j<=6;j++{
    	 if i+j==4{
    	 	fmt.Printf("**")
		 }	else if j-i==3 {
		 	fmt.Printf("**")
		 } else {
		 	fmt.Printf("  ")
		 }
		}
    	fmt.Printf("\n")
	}
}

func xiahuxian6(){
	for i:=1;i<=3;i++{
		fmt.Printf("  ")
		for j:=1;j<=6;j++{
			if i==j {
				fmt.Printf("**")
			}	else if j+i==7 {
				fmt.Printf("**")
			} else {
				fmt.Printf("  ")
			}
		}
		fmt.Printf("\n")
	}
}