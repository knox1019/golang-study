// 在有大量输入需要读取的时候，就建议使用bufio.Reader来进行内容读取
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	line = strings.TrimSpace(line)
	parts := strings.Split(line, " ")
	n, _ := strconv.Atoi(parts[0])
	s := make([]int, n)
	for i := 0; i < n; i++ {
		s[i], _ = strconv.Atoi(parts[i+1])
	}
	fmt.Println(s)
}