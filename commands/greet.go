package commands

import "fmt"

func Greet(args[] string){
	if len(args)<3{
		fmt.Println("Enter name")
		return
	}
	fmt.Println("hello",args[2])
}