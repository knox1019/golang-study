package main
import "fmt"
func main() {
	n := 10
	fmt.Println("main：", n)
	test(&n)	
}
func test(p *int) {
	*p = 20
	fmt.Println("test：", *p)
}