// 包声明
package main

// 导入包
import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

// callDemoByReflection 通过反射调用示例函数
func callDemoByReflection(userInput string) error {
	var funcName string
	var demoFunc interface{}
	var exists bool

	// 步骤1: 直接查找（支持大小写不敏感）
	demoFunc, exists = demoRegistry[userInput]
	if exists {
		funcName = userInput
	} else {
		// 步骤2: 尝试首字母大写的格式（例如: "arrayAccess" -> "ArrayAccess"）
		funcName = toPascalCase(userInput)
		demoFunc, exists = demoRegistry[funcName]
		if !exists {
			// 步骤3: 尝试智能转换（处理下划线等）
			funcName = toPascalCaseFromSnakeCase(userInput)
			demoFunc, exists = demoRegistry[funcName]
		}
	}

	if !exists {
		return fmt.Errorf("未找到示例: %s (尝试调用函数: %s)", userInput, funcName)
	}

	// 获取函数的反射值
	funcValue := reflect.ValueOf(demoFunc)
	if !funcValue.IsValid() {
		return fmt.Errorf("无效的示例函数: %s", funcName)
	}

	// 检查是否为函数
	if funcValue.Kind() != reflect.Func {
		return fmt.Errorf("%s 不是一个函数", funcName)
	}

	// 检查函数参数数量（应该为0）
	if funcValue.Type().NumIn() != 0 {
		return fmt.Errorf("示例函数 %s 不应该有参数", funcName)
	}

	// 调用函数
	fmt.Printf("运行 %s 示例 (函数: %s)...\n", userInput, funcName)
	result := funcValue.Call([]reflect.Value{})

	// 检查返回值（如果有的话）
	if len(result) > 0 {
		fmt.Printf("函数返回了 %d 个值\n", len(result))
	}

	return nil
}

// toPascalCase 将输入转换为PascalCase（首字母大写）
// 例如: "arrayAccess" -> "ArrayAccess"
func toPascalCase(input string) string {
	if input == "" {
		return ""
	}
	// 首字母大写，其余保持原样
	return strings.ToUpper(input[:1]) + input[1:]
}

// toPascalCaseFromSnakeCase 将下划线分隔的名称转换为PascalCase
// 例如: "anonymous_struct" -> "AnonymousStruct"
func toPascalCaseFromSnakeCase(input string) string {
	if input == "" {
		return ""
	}

	// 处理下划线分隔的名称
	parts := strings.Split(input, "_")
	var result strings.Builder

	for _, part := range parts {
		if part == "" {
			continue
		}
		// 首字母大写
		if len(part) > 0 {
			result.WriteString(strings.ToUpper(part[:1]))
			result.WriteString(strings.ToLower(part[1:]))
		}
	}

	return result.String()
}

// TestSmartDemo 智能Demo调用演示
func TestSmartDemo() {
	fmt.Println("=== 智能Demo调用演示 ===")
	fmt.Println("这个函数演示了如何通过输入名称自动调用对应的Demo函数")
	fmt.Println()
	fmt.Println("智能转换规则:")
	fmt.Println("  'constants'     → 'ConstantsDemo'")
	fmt.Println("  'anonymous_struct' → 'AnonymousStructDemo'")
	fmt.Println("  'nested_struct' → 'NestedStructDemo'")
	fmt.Println("  'reflection'    → 'demonstrateReflection'")
	fmt.Println()
	fmt.Println("使用示例:")
	fmt.Println("  go run main.go constants      # 自动调用 ConstantsDemo")
	fmt.Println("  go run main.go anonymous_struct # 自动调用 AnonymousStructDemo")
}

func main() {
	// 检查命令行参数
	args := os.Args[1:] // 跳过程序名

	if len(args) == 0 {
		// 默认运行
		printHelp()
		return
	}

	// 通过反射调用示例
	arg := args[0]
	if err := callDemoByReflection(arg); err != nil {
		fmt.Printf("错误: %v\n", err)
		fmt.Println()
		printHelp()
	}
}

// printHelp 打印帮助信息
func printHelp() {
	fmt.Println("=== Go 语言学习示例运行器（智能反射调用版）===")
	fmt.Println("用法: go run main.go [示例名]")
	fmt.Println()
	fmt.Println("🎯 智能识别: 输入示例名自动匹配对应的 Demo 函数！")
	fmt.Println("📝 使用方式: go run main.go <示例名>")
	fmt.Println()
	fmt.Println("可用示例:")

	// 获取所有示例名称并排序
	allDemos := make([]string, 0, len(demoRegistry))
	for funcName := range demoRegistry {
		allDemos = append(allDemos, funcName)
	}

	// 按字母顺序排序
	sort.Strings(allDemos)

	// 显示所有示例
	for _, funcName := range allDemos {
		fmt.Printf("    %s\n", funcName)
	}

	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  go run main.go ArrayAccess        # 数组访问示例")
	fmt.Println("  go run main.go BasicRoutes        # Gin基础路由")
	fmt.Println("  go run main.go Constants          # 常量示例")
	fmt.Println("  go run main.go arrayAccess         # 支持小写开头（自动转换）")
	fmt.Println()
	fmt.Printf("当前注册了 %d 个示例\n", len(demoRegistry))
	fmt.Println("\n🚀 智能匹配: 支持大小写自动转换和下划线格式")
	fmt.Println("💡 添加新示例: 只需在 demoRegistry 中添加函数注册即可！")
}
