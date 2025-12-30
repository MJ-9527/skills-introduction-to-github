package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	fmt.Printf("请输入名字: ") // 添加明确的输入提示
	var name string
	fmt.Scan(&name)
	fmt.Println("你输入的名字是:", name) // 输出更友好的反馈
}
