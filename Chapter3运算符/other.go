package main
import "fmt"
func main() {
	//定义一个变量
	var age int =18
	fmt.Println("age的地址是：", &age)//取地址符 &
	//取指针变量的值
	var p *int = &age//	*int 指针类型
	fmt.Println("p的地址是：", p)
	fmt.Println("p指针变量的值是：", *p)//通过指针变量修改值
	*p = 20//修改指针变量指向的值
	fmt.Println("age的值被修改为：", age)
}