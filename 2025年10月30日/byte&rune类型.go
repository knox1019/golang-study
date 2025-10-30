/*
Go 语言的字符有以下两种：

uint8类型，或者叫 byte 型，代表一个ASCII码字符。
rune类型，代表一个 UTF-8字符。
当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。rune类型实际是一个int32。
*/
package main

import (
	"fmt"
	"math"
)
func main() {
	var b byte = 'a'
	var r rune = '国'
	fmt.Printf("b=%c r=%c\n", b, r)
	traversalString()
	changeString()
	sqrtDemo()
}

//Go 使用了特殊的 rune 类型来处理 Unicode，让基于 Unicode 的文本处理更为方便，
// 也可以使用 byte 型进行默认字符串处理，性能和扩展性都有照顾。
//在 Go 语言中，字符串是以 UTF-8 编码保存的，因此可以直接使用 byte 和 rune 类型来处理字符串中的字符。
//在处理 ASCII 字符时，byte 类型更为高效，而在处理非 ASCII 字符时，rune 类型则更为合适。
//通过使用 byte 和 rune 类型，Go 语言能够高效地处理各种字符编码的文本数据，满足不同应用场景的需求。

// 遍历字符串
func traversalString() {//字符串是由字节组成的字节切片
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte字节
		fmt.Printf("%v(%c) ", s[i], s[i])//按byte遍历，  %c字符
	}
	fmt.Println()
	for _, r := range s { //按rune符文遍历，从字符串中拿出具体的字符
		fmt.Printf("%v(%c) ", r, r)//按rune遍历，%c字符
	}
	fmt.Println()
}
/*
1. 按字节遍历（byte）：
Go 的字符串底层是 []byte，默认是 UTF-8 编码。
对于 ASCII 字符（如 "hello"），每个字符占 1 个字节。
对于非 ASCII 字符（如 "沙河"），每个汉字占 3 个字节。
所以你会看到 6 个字节代表 "沙河"，打印出来是乱码或不可读的字节。
2. 按符文遍历（rune）：
rune 是 Go 的 int32 类型，表示一个 Unicode 码点。
使用 for range 会自动按 rune 解码 UTF-8 字符串。
所以 "沙河" 会被正确识别为两个字符：沙（27801）和 河（27827）
*/

//修改字符串
//要修改字符串，需要先将其转换成[]rune或[]byte，完成后再转换为string。无论哪种转换，都会重新分配内存，并复制字节数组。
func changeString() {
	s1 := "big"
	// 强制类型转换
	byteS1 := []byte(s1)//将字符串转换为byte切片
	byteS1[0] = 'p'
	fmt.Println(string(byteS1))//将byte切片转换为字符串

	s2 := "白萝卜"
	runeS2 := []rune(s2)//将字符串转换为rune切片
	runeS2[0] = '红'
	fmt.Println(runeS2)//打印rune切片
	fmt.Println(string(runeS2))//将rune切片转换为字符串

	c1 := "红"//字符串(双引号)
	c2 := '红'//rune字符(单引号) rune只是int32类型的别名
	fmt.Printf("c1=%T c2=%T\n", c1, c2)
}

// 类型转换
// //Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。

// 强制类型转换的基本语法如下：

// T(表达式)
// 其中，T表示要转换的类型。表达式包括变量、复杂算子和函数返回值等.

// 比如计算直角三角形的斜边长时使用math包的Sqrt()函数，该函数接收的是float64类型的参数，而变量a和b都是int类型的，这个时候就需要将a和b强制类型转换为float64类型。
func sqrtDemo() {
	var a, b = 3.3, 4.4
	var c float64
	// math.Sqrt()接收的参数是float64类型，需要强制转换
	c = math.Sqrt(a*a + b*b)
	fmt.Printf("%.2f\n", c)
}
