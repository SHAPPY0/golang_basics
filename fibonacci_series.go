package main

import "fmt"

func fibonacci_series(n int){
	fab := make([]int,n+1)
	fab[0] = 0
	fab[1] = 1

	for i:=2;i<=n;i++{
		fab[i] = fab[i-2]+fab[i-1]
	}
	fmt.Println(fab)

}

func main(){
	var n int = 5
	fibonacci_series(n)
	
}