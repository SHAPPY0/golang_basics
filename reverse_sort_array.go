package main

import (
		"fmt" 
		"sort"
	)



func reverse_sort(arr []int)[]int{
	if len(arr)==0{
		return []int{}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))
	return arr
}

func main(){
	var arry = []int{50,10,23,44,22,81}
	
	res:=reverse_sort(arry)
	fmt.Println(res)
}