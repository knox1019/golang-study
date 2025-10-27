package main//declare the main package
import "fmt"//import the fmt package for formatted I/O
func main() {		//start of main function
	fmt.Println("Hello, World!")//	print Hello, World! to the console
	var n int = 42	//declare an integer variable n and assign it the value 42
	fmt.Println("The answer is", n, n)//fmt.Println prints the value of n in different formats
	fmt.Printf("%d\n", n)//decimal
	fmt.Printf("%T\n", n)//type
	fmt.Printf("%b\n", n)//binary
	fmt.Printf("%o\n", n)//octal
	fmt.Printf("%x\n", n)//hexadecimal
	fmt.Printf("%v\n", n)//default format
}//end of main function
now 现在处于一个新的时代
