package main
import "fmt"
func main() {
	//+加号
	//1.正数 2.相加操作 3.字符串连接
	a := 10
	b := 20
	fmt.Println(a + b) //30
	str1 := "hello,"
	str2 := "world"
	fmt.Println(str1 + str2) //hello,world

	//-减号
	//1.负数 2.相减操作
	c := 30
	d := 15
	fmt.Println(c - d) //15
	fmt.Println(-c)    //-30

	// *星号
	//1.乘法 2.指针
	e := 5
	f := 6
	fmt.Println(e * f) //30

	// /斜杠
	//1.除法
	g := 20
	h := 4
	fmt.Println(g / h) //5
	//注意：整数除法会舍弃小数部分
	i := 7
	j := 2
	fmt.Println(i / j) //3
	//如果想要得到精确的结果，需要将其中一个操作数转换为浮点数
	fmt.Println(float64(i) / float64(j)) //3.5
	//或者直接使用浮点数
	k := 7.0
	l := 2.0
	fmt.Println(k / l) //3.5

	// %百分号
	//取模（求余数）
	m := 10
	n := 3
	fmt.Println(m % n) //1

	// ++自增
	//将变量的值增加1
	o := 5
	o++
	fmt.Println(o) //6

	// --自减
	//将变量的值减少1
	p := 5
	p--
	fmt.Println(p) //4
}