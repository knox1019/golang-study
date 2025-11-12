//函数也是一种数据类型
package main
import "fmt"

//定义一个函数：
func test(num int){
	fmt.Println(num)
}

func test2(num1 int, num2 float32, testFunc func(int)) {
	fmt.Println("-----test2-----")
}

	func main() {
	//函数也是一种数据类型，可以赋值给一个变量
	a := test //函数名test代表函数本身
	fmt.Printf("a的类型是%T\n", a)//func(int)
	fmt.Printf("test的类型是%T\n", test)//func(int)

	//通过该变量可以调用该函数
	a(100)//等价于 test(100)

	//调用test2函数，传入一个函数作为参数
	test2(10, 20.5, test)
	test2(30, 40.5, a)//也可以传入变量a

	//自定义一个函数类型，相当于起别名
	type myfuncType int

	num3 := myfuncType(10)
	num4 := int(20)//需要强制类型转换
	fmt.Println("num3=", num3)
	fmt.Println("num4=", num4)
	fmt.Printf("num3的类型是%T\n", num3)//main.myfuncType
	fmt.Printf("num4的类型是%T\n", num4)//int
}

//2025年11月12日19:23:57 救命 太困了