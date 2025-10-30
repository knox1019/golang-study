package main
import "fmt"
import "strings"
func main() {
	// fmt.Println("Hello, World!")
	// var n int = 42
	// fmt.Println("The answer is", n, n)
	// fmt.Printf("%d\n", n)
	// fmt.Printf("%T\n", n)
	// fmt.Printf("%b\n", n)
	// fmt.Printf("%o\n", n)
	// fmt.Printf("%x\n", n)
	// fmt.Printf("%v\n", n)
	// fmt.Println("现在我开始随便改代码了")
	
	//现在要打印Windows的路径
	fmt.Println("C:\\Program Files\\Go\\bin")
	fmt.Println(`C:\Program Files\Go\bin`)

	//字符串的长度
	fmt.Println(len("Hello, world")) // 12

	//字符串的遍历
	s := "Hello, world"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c\n", s[i])
	}
	//字符串的拼接
	s1 := "Hello, "
	s2 := "World!"
	s3 := s1 + s2
	fmt.Println("\n" + s3)
	//字符串的分隔
	s4 := "Hello,World,Go,Language"
	ret := strings.Split(s4, ",")
	fmt.Println(ret)
	//包含？
	fmt.Println(strings.Contains(s4, "World")) // true
	//前缀
	fmt.Println(strings.HasPrefix(s4, "Hello")) // true
	//后缀
	fmt.Println(strings.HasSuffix(s4, "Language")) // true	

	//子串索引
	fmt.Println(strings.Index(s4, "Go")) // 12  index从前往后找
	fmt.Println(strings.Index(s4, "Java")) // -1 java不存在

	//子串最后一次出现的索引
	fmt.Println(strings.LastIndex(s4, "o")) // 13 lastindex从后往前找
	fmt.Println(strings.LastIndex(s4, "Java")) // -1 java不存在

	//字符串替换
	s5 := "Hello, World, World, World"	// Hello, World, Go, World
	s6 := strings.ReplaceAll(s5, "World", "Go")// ReplaceAll替换所有
	fmt.Println(s6) // Hello, Go, Go, Go
	//Replace替换指定数量
	//s6 := strings.Replace(s5, "World", "Go", 1) // Hello, Go, World, World
	//s6 := strings.Replace(s5, "World", "Go", 2) // Hello, Go, Go, World
	
	//join拼接
	s7 := []string{"Hello", "World", "Go", "Language"}
	s8 := strings.Join(s7, "-")
	fmt.Println(s8) // Hello-World-Go-Language

	}