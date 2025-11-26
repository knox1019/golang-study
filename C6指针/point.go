/*
目标：这一小时后，你要能自信地说：“Go 的指针比 C 简单一万倍。”
核心概念：

& (取地址)：在这个变量的脑门上贴个条，看它住哪。
* (取值)：顺着地址找过去，把里面存的数据拿出来（或者改掉）。
Go 的红线：Go 不允许你做指针运算（比如 ptr + 1），所以它是安全的，你只管传地址，不用管偏移量。
【实操任务】

验证“值传递” (15分钟)：

写一个函数 func change(num int)，在里面把 num 改成 100。
在 main 里定义 x := 0，调用 change(x)。
打印 x。你会发现 x 还是 0。
结论：不传指针，函数里改不了外面的值。
使用指针修改值 (30分钟)：

修改刚才的函数为 func change(ptr *int) (参数变成了指针)。
函数体里写 *ptr = 100 (顺着地址爬过去改值)。
调用时写 change(&x) (把 x 的地址传进去)。
打印 x。这次变成 100 了。
课后作业 (15分钟)：

写一个 swap(a, b *int) 函数，交换两个整数的值。这是 C 语言经典题，用 Go 写一遍。
*/
package main

import "fmt"

func main() {
	x := 0//这个x是main函数里的变量
	fmt.Println("x =", x)
	change(x)//把x的值传给change函数里的num变量，但是num和x是两个不同的变量，改变num不会影响x
	fmt.Println("调用change函数后x =", x)	
	changePtr(&x)//把x的地址传给changePtr函数里的ptr变量，ptr和x指向同一个地址，改变ptr指向的值会影响x
	fmt.Println("调用changePtr函数后x =", x)
	fmt.Println("-----交换前-----")
	a := 10
	b := 20
	fmt.Println("a =", a, "b =", b)
	swap(&a, &b)//把a和b的地址传给swap函数里的a和b变量，swap函数里通过地址交换值
	fmt.Println("-----交换后-----")
	fmt.Println("a =", a, "b =", b)
}

func change(num int) {
	num = 100
}

func changePtr(ptr *int) {//ptr是一个指针变量，存的是一个int类型变量的地址
	*ptr = 100
}

// func swap(a, b *int) {
// 	temp := *a
// 	*a = *b
// 	*b = temp
// }
func swap(a, b *int) {
    // Go 支持“多重赋值”，不需要 temp 变量
    *a, *b = *b, *a
}