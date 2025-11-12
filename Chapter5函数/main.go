package main
import "fmt"
func main() {
	fmt.Println("和是：", cal(3, 5))
}

func cal(a int, b int) int {
	return a + b
}