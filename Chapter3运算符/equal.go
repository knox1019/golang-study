package main
import ( "fmt" )
func main() {
	var num1 int = 10
	fmt.Println("num1 =", num1)

	var num2 int = (10 +20) % 3 + 3 - 7 // 先算括号，再算乘除，最后算加减
	fmt.Println("num2 =", num2)
	
	var num3 int = 10
	num3 += 5 // 等价于 num3 = num3 + 5
	fmt.Println("num3 =", num3)

	var num4 int = 10
	num4++ // 等价于 num4 = num4 + 1
	fmt.Println("num4 =", num4)

	var a int = 7
	var b int = 2
	//交换
	a, b = b, a
	fmt.Println("a =", a)
	fmt.Println("b =", b)
}