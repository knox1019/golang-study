package main	
import (
	"fmt"
	"unicode"
)
func main() {
	var a int = 16
	var b float64 = 16.0
	var c bool = true
	var d string = "hello"
	fmt.Printf("a=%d b=%f c=%t d=%s\n", a, b, c, d)//%d十进制整数 %f浮点数 %t布尔值 %s字符串
	fmt.Printf("a的类型是%T b的类型是%T c的类型是%T d的类型是%T\n", a, b, c, d)
	counting()

	counting2()
}

func counting() {//统计出字符串"hello沙河小王子"中汉字的数量
	var i string = "hello沙河小王子"
	count := 0
	for _, r := range i {
		if r > 255 {//rune类型，超过255的就是汉字
			count++
		}
	}
	fmt.Printf("字符串%s中汉字的数量是%d\n", i, count)
}


//改进
func counting2() {
	s := "hello沙河小王子"
	count := 0
	for _, r := range s {
		if unicode.Is(unicode.Han, r) { // 判断是否为中文字符（汉字）
			count++
		}
	}
	fmt.Printf("字符串%s中汉字的数量是%d\n", s, count)
}