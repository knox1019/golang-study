package main
import "fmt"
//进制转换
func main() {
	var a int = 255
	fmt.Println("十进制：", a)
	fmt.Printf("二进制：%b\n", a)
	fmt.Printf("八进制：%o\n", a)
	fmt.Printf("十六进制：%x\n", a)
}
