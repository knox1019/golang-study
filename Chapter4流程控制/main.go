package main
import "fmt"

func main() {
	//流程控制
	//条件判断 if else
	var age int = 20
	if age >= 18 {
		fmt.Println("成年人")
	}
	else {
		fmt.Println("未成年")
	}

	//if else if else
	var score int = 85
	if score >= 90 {
		fmt.Println("优秀")
	} else if score >= 60 {
		fmt.Println("及格")
	} else {
		fmt.Println("不及格")
	}



	//循环 for
	for i := 0; i < 5; i++ {
		fmt.Println("i =", i)
	}


	//switch case
	/*
		switch 变量/表达式 {
		case 值1:
			执行语句1
		case 值2:
			执行语句2
		default:
			执行默认语句
		}

		default 可选，用于处理不匹配任何 case 的情况
		可以放在任何位置，但通常放在最后
		case后面的值可以是常量、变量或表达式，但必须是唯一的，不能重复

	*/
	var day int = 3
	switch day {//switch 后可以是变量或表达式，而且可以直接省略变量/表达式，直接使用 true 作为条件,相当于多个 if-else
		//甚至可以在switch后直接定义变量，例如：switch day := 3; day {
		//switch穿透，即执行完匹配的case后继续执行后面的case，使用 fallthrough 关键字实现,目的是为了兼容C语言的switch语法
	case 1://可以有多个匹配项
	//例如case 1, 7:
		fmt.Println("星期一")
	case 2:
		fmt.Println("星期二")
	case 3:
		fmt.Println("星期三")
	default:
		fmt.Println("其他天")
	}	


	//break 和 continue
	for i := 0; i < 10; i++ {
		if i == 5 {
			break //跳出循环
		}	
		fmt.Println("i =", i)
	}

	for j := 0; j < 10; j++ {
		if j % 2 == 0 {
			continue //跳过本次循环
		}
		fmt.Println("j =", j)
	}	

	//goto 语句
	var count int = 0
Here:
	if count < 5 {
		fmt.Println("count =", count)
		count++
		goto Here //跳转到 Here 标签
	}
	
}