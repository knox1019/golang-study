package main
import "fmt"
var MyName string = "jack"//公有变量
var myAge int = 22//私有变量
var x = 100//全局变量
const Pi = 3.14//公有常量
func main() {
	//变量的声明和赋值
	var a int = 10//声明变量a并赋值为10
	var b = 20//类型推导，声明变量b并赋值为20
	c := 30//简短变量声明，声明变量c并赋值为30
	fmt.Println(a, b, c)
	//多变量声明
	var x, y, z int = 1, 2, 3
	fmt.Println(x, y, z)
	//变量的零值
	var m int		//整型变量的零值为0
	var n string	//字符串变量的零值为""
	var p bool		//布尔型变量的零值为false
	fmt.Println(m, n, p)
	//常量的声明
	const pi = 3.14	//声明常量pi并赋值为3.14
	const (		//批量声明常量
		e  = 2.71
		gc = 9.8
	)
	fmt.Println(pi, e, gc)
}