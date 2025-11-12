package main
import "fmt"
func main() {
	fmt.Println("和是：", calsum(3, 5))
	fmt.Println("差是：", calsub(10, 4))
	fmt.Println("积是：", calmul(4, 6))
	fmt.Println("商是：", caldiv(20, 5))
	a, b, c, d := moreresult(15, 3)
	fmt.Println("多重返回值：", a, b, c, d)
}

func calsum(a int, b int) int {
	return a + b
}
func calsub(a int, b int) int {
	return a - b
}
func calmul(a, b int) int {
	return a * b
}
func caldiv(a, b int) int {
	return a / b
}
func moreresult(a, b int) (int, int, int, int) {
	return a + b, a - b, a * b, a / b
}