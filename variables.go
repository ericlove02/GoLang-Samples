package main
import "fmt"

func main(){
	test := "Variables Test"
	fmt.Println(test)
	
	var a int
	a = 5
	a += 10
	fmt.Println("15 == ", a)
	
	emptyLine()
	
	for i := 6; i <= 10; i++{
		fmt.Println(i)
	}
	
	emptyLine()
	
	b := "Print Me"
	printVar(b)
}

func printVar(a string){
	fmt.Println(a)
}

func emptyLine(){
	fmt.Println()
}