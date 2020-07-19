package main
import "fmt"


func main(){

	var a int = 100
	
	var b int = 200
	
	fmt.Println("Before swap a =",a)
	
	fmt.Println("Before swap b =",b)
	
	var temp int = a
	
	a = b
	
	b = temp
	
	fmt.Println("After swap a =",a)

	fmt.Println("After swap b =",b)
}