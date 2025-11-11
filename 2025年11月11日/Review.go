/*可见性
名称大写字母开头，即为公有类型/变量/常量,名字小写或下划线开头，即为私有类型/变量/常量
*/
package main

import "fmt"

// 公有
const MyName = "jack"

// 私有
const mySalary = 20_000

func main() {
	fmt.Println("MyName:", MyName)
	fmt.Println("mySalary:", mySalary)
	stringLiterals()
}
/*
`abc`                // "abc"
`\n
\n`                  // "\\n\n\\n"
"\n"
"\""                 // `"`
"Hello, world!\n"
"今天天气不错"
"日本語"
"\u65e5本\U00008a9e"
"\xff\u00FF"
*/
func stringLiterals() {
	println(`abc`)                // "abc"
	println(`\n
\n`)
	println("\n")
	println("\"")					 // `"`
	println("Hello, world!\n")
	println("今天天气不错")
	println("日本語")
	println("\u65e5本\U00008a9e")
	println("\xff\u00FF")
}