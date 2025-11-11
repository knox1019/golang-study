package main
import "fmt"
func main() {//逻辑运算符
	//与 && 有一个false 就是 false
	//也叫做短路与
	fmt.Println(true && false) // false
	fmt.Println(true && true)  // true
	//或 || 有一个true 就是 true
	//也叫做短路或
	fmt.Println(true || false) // true
	fmt.Println(false || false)// false
	//非 ! 取反
	fmt.Println(!true)  // false
	fmt.Println(!false) // true

}