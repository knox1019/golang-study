package main
import "fmt"
//可变参数函数
func test(args ...int) {//无return，就可以省略返回值类型
	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d]=%d\n", i, args[i])
	}
}//...表示可变参数，可以传入0个或多个参数


func main() {
	test()
	fmt.Println("-----")
	test(10)
	fmt.Println("-----")
	test(10, 20, 30)
}