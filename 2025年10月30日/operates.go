package main
import "fmt"

//go语言中的运算符
func main() {
	// 1. 算术运算符
	a := 10
	b := 20
	fmt.Printf("a + b = %d\n", a+b) // 加法
	fmt.Printf("a - b = %d\n", a-b) // 减法
	fmt.Printf("a * b = %d\n", a*b)	// 乘法
	fmt.Printf("b / a = %d\n", b/a)	// 除法
	fmt.Printf("b %% a = %d\n", b%a) // 取模
	fmt.Printf("a ++ = %d\n", a+1) // 自增	a++
	fmt.Printf("b -- = %d\n", b-1) // 自减  b--
	// 2. 关系运算符
	fmt.Printf("a == b: %t\n", a == b) // 等于
	fmt.Printf("a != b: %t\n", a != b) // 不等于
	fmt.Printf("a > b: %t\n", a > b)   // 大于
	fmt.Printf("a < b: %t\n", a < b)   // 小于
	fmt.Printf("a >= b: %t\n", a >= b) // 大于等于
	fmt.Printf("a <= b: %t\n", a <= b) // 小于等于
	// 3. 逻辑运算符
	x := true
	y := false
	fmt.Printf("x && y: %t\n", x && y) // 逻辑与
	fmt.Printf("x || y: %t\n", x || y) // 逻辑或
	fmt.Printf("!x: %t\n", !x)         // 逻辑非
	// 4. 位运算符
	m := 5  // 二进制: 0101
	n := 3  // 二进制: 0011
	fmt.Printf("m & n = %d\n", m&n)   // 按位与
	fmt.Printf("m | n = %d\n", m|n)   // 按位或
	fmt.Printf("m ^ n = %d\n", m^n)   // 按位异或
	fmt.Printf("m &^ n = %d\n", m&^n) // 按位清除
	fmt.Printf("m << 1 = %d\n", m<<1) // 左移
	fmt.Printf("n >> 1 = %d\n", n>>1) // 右移	
}