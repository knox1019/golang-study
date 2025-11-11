package main
import "fmt"
func main() {//读取固定长度的数组
  n := 10
  s := make([]int, n)
  for i := range n {
    fmt.Scan(&s[i])
  }
  fmt.Println(s)
}