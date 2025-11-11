package main

import (
    "cmp"
    "fmt"
    "os"
)

func main() {
    // 正确使用 cmp 包的函数
    fmt.Println(cmp.Compare(1, 2)) // 输出: -1
    fmt.Println(cmp.Less(1, 2))    // 输出: true

    os.Stdout.WriteString("hello world!\n")

    fmt.Printf("%%%s\n", "hello world") // 输出: %hello world
    fmt.Printf("%s\n", "hello world")   // 输出: hello world
    fmt.Printf("%q\n", "hello world")   // 输出: "hello world"
    fmt.Printf("%d\n", 2<<7-1)          // 输出: 255

    fmt.Printf("%f\n", 1e2) // 输出: 100.000000
    fmt.Printf("%e\n", 1e2) // 输出: 1.000000e+02
    fmt.Printf("%E\n", 1e2) // 输出: 1.000000E+02
    fmt.Printf("%g\n", 1e2) // 输出: 100

    fmt.Printf("%b\n", 2<<7-1)  // 输出: 11111111
    fmt.Printf("%#b\n", 2<<7-1) // 输出: 0b11111111
    fmt.Printf("%o\n", 2<<7-1)  // 输出: 377
    fmt.Printf("%#o\n", 2<<7-1) // 输出: 0377
    fmt.Printf("%x\n", 2<<7-1)  // 输出: ff
    fmt.Printf("%#x\n", 2<<7-1) // 输出: 0xff
    fmt.Printf("%X\n", 2<<7-1)  // 输出: FF
    fmt.Printf("%#X\n", 2<<7-1) // 输出: 0XFF

    type person struct {
        name    string
        age     int
        address string
    }
    fmt.Printf("%v\n", person{"lihua", 22, "beijing"})   // 输出: {lihua 22 beijing}
    fmt.Printf("%+v\n", person{"lihua", 22, "beijing"})  // 输出: {name:lihua age:22 address:beijing}
    fmt.Printf("%#v\n", person{"lihua", 22, "beijing"})  // 输出: main.person{name:"lihua", age:22, address:"beijing"}
    fmt.Printf("%t\n", true)                             // 输出: true
    fmt.Printf("%T\n", person{})                         // 输出: main.person
    fmt.Printf("%c%c\n", 20050, 20051)                   // 输出: 乒乒（Unicode 字符）
    fmt.Printf("%U\n", '码')                              // 输出: U+7801
    fmt.Printf("%p\n", &person{})                        // 输出: 0xc0000b8010（地址值，每次运行不同）
}