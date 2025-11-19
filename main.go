// 包声明
package main

// 导入包
import (
	"fmt"
	structs "lesson1/1.5_struct"
	constants "lesson1/1.6_constants_enum"
	"os"
)

// 函数声明
func sayHello() {
	fmt.Println(str)
}

// 变量声明
var str = "Hello, World!"

func main() {
	// 检查命令行参数
	args := os.Args[1:] // 跳过程序名

	if len(args) == 0 {
		// 默认运行
		fmt.Println("=== Go 语言学习示例运行器 ===")
		fmt.Println("用法: go run main.go [示例名]")
		fmt.Println("可用示例:")
		fmt.Println("  constants - 常量示例")
		fmt.Println("  struct    - 结构体示例")
		fmt.Println("  hello     - 基础 Hello World")
		fmt.Println()
		fmt.Println("示例:")
		fmt.Println("  go run main.go constants")
		fmt.Println("  go run main.go struct")
		fmt.Println("  go run main.go hello")
		return
	}

	// 根据参数运行对应的示例
	switch args[0] {
	case "constants", "1.6":
		fmt.Println("运行常量示例...")
		constants.Demo()

	case "struct", "1.5":
		fmt.Println("运行结构体示例...")
		structs.Demo()

	case "hello", "basic":
		fmt.Println("运行基础示例...")
		// fmt.Println(os.Args)
		// fmt.Println(os.Args[1:])
		sayHello()

	default:
		fmt.Printf("未知的示例: %s\n", args[0])
		fmt.Println("可用示例: constants, struct, hello")
	}
}
