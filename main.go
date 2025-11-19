// 包声明
package main

// 导入包
import (
	"fmt"
	structs "lesson1/1.5_struct"
	constants "lesson1/1.6_constants_enum"
	"os"
	"reflect"
	"strings"
)

// 函数声明
func sayHello() {
	fmt.Println(str)
}

// 变量声明
var str = "Hello, World!"

// DemoRegistry 示例注册表 - 函数名到函数的映射
var demoRegistry = map[string]interface{}{
	// 常量示例
	"ConstantsDemo": constants.ConstantsDemo,

	// 结构体示例
	"AnonymousStructDemo":  structs.AnonymousStructDemo,
	"NestedStructDemo":     structs.NestedStructDemo,
	"StructMethodsDemo":    structs.StructMethodsDemo,
	"CrossFileUsageDemo":   structs.CrossFileUsageDemo,
	"LowercaseStructDemo":  structs.LowercaseStructDemo,
	"RealWorldExampleDemo": structs.RealWorldExampleDemo,
	"TagsExampleDemo":      structs.TagsExampleDemo,
	"VisibilityDemo":       structs.VisibilityDemo,

	// 基础示例（特殊处理）
	"sayHello": sayHello,

	// 反射演示示例
	"demonstrateReflection": demonstrateReflection,
	"TestSmartDemo":         TestSmartDemo,
}

// aliasRegistry 别名映射 - 用户输入名到函数名的映射
var aliasRegistry = map[string]string{
	// 常量示例别名
	"constants":       "ConstantsDemo",
	"1.6":             "ConstantsDemo",
	"1.6.1_constants": "ConstantsDemo",

	// 结构体示例别名
	"anonymous_struct":   "AnonymousStructDemo",
	"1.5.2":              "AnonymousStructDemo",
	"nested_struct":      "NestedStructDemo",
	"1.5.3":              "NestedStructDemo",
	"struct_methods":     "StructMethodsDemo",
	"1.5.4":              "StructMethodsDemo",
	"cross_file_usage":   "CrossFileUsageDemo",
	"1.5.4.2":            "CrossFileUsageDemo",
	"lowercase_struct":   "LowercaseStructDemo",
	"1.5.4.3":            "LowercaseStructDemo",
	"real_world_example": "RealWorldExampleDemo",
	"1.5.4.4":            "RealWorldExampleDemo",
	"tags_example":       "TagsExampleDemo",
	"visibility":         "VisibilityDemo",

	// 基础示例别名
	"hello": "sayHello",
	"basic": "sayHello",

	// 反射演示示例别名
	"reflection": "demonstrateReflection",
	"smart_demo": "TestSmartDemo",
	"test_smart": "TestSmartDemo",
}

// callDemoByReflection 通过反射调用示例函数
func callDemoByReflection(userInput string) error {
	var funcName string
	var demoFunc interface{}
	var exists bool

	// 步骤1: 直接查找函数名（用户可能直接输入函数名）
	demoFunc, exists = demoRegistry[userInput]
	if exists {
		funcName = userInput
	} else {
		// 步骤2: 查找别名映射
		funcName, exists = aliasRegistry[userInput]
		if exists {
			demoFunc, exists = demoRegistry[funcName]
		} else {
			// 步骤3: 智能转换 - 将输入转换为Demo函数名
			// 例如: "constants" -> "ConstantsDemo"
			funcName = toDemoFunctionName(userInput)
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

// toDemoFunctionName 将用户输入转换为Demo函数名
// 例如: "constants" -> "ConstantsDemo"
//
//	"anonymous_struct" -> "AnonymousStructDemo"
func toDemoFunctionName(input string) string {
	if input == "" {
		return "Demo"
	}

	// 处理特殊情况
	switch input {
	case "hello", "basic":
		return "sayHello"
	case "reflection":
		return "demonstrateReflection"
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

	return result.String() + "Demo"
}

// 示例函数 - 这些函数可以被反射调用
func demoFunction1() {
	fmt.Println("这是示例函数 1")
}

func demoFunction2(name string) {
	fmt.Println("这是示例函数 2，参数:", name)
}

func demoFunction3() int {
	fmt.Println("这是示例函数 3，返回值")
	return 42
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
	fmt.Println("  'hello'         → 'sayHello'")
	fmt.Println("  'reflection'    → 'demonstrateReflection'")
	fmt.Println()
	fmt.Println("使用示例:")
	fmt.Println("  go run main.go constants      # 自动调用 ConstantsDemo")
	fmt.Println("  go run main.go anonymous_struct # 自动调用 AnonymousStructDemo")
	fmt.Println("  go run main.go hello          # 自动调用 sayHello")
}

// demonstrateReflection 反射调用演示
func demonstrateReflection() {
	fmt.Println("=== 反射调用方法演示 ===")

	// 1. 调用无参数函数
	fmt.Println("1. 调用无参数函数:")
	func1 := reflect.ValueOf(demoFunction1)
	func1.Call([]reflect.Value{})

	// 2. 调用带参数函数
	fmt.Println("\n2. 调用带参数函数:")
	func2 := reflect.ValueOf(demoFunction2)
	args := []reflect.Value{reflect.ValueOf("反射调用")}
	func2.Call(args)

	// 3. 调用带返回值函数
	fmt.Println("\n3. 调用带返回值函数:")
	func3 := reflect.ValueOf(demoFunction3)
	results := func3.Call([]reflect.Value{})
	fmt.Printf("返回值: %v\n", results[0].Interface())

	fmt.Println("\n=== 动态注册和调用示例 ===")

	// 创建函数注册表
	functions := map[string]interface{}{
		"func1": demoFunction1,
		"func2": demoFunction2,
		"func3": demoFunction3,
	}

	// 动态调用
	callName := "func1"
	if fn, exists := functions[callName]; exists {
		reflect.ValueOf(fn).Call([]reflect.Value{})
	}
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
	fmt.Println("📝 命名规则: 示例名 + 'Demo' = 函数名")
	fmt.Println()
	fmt.Println("可用示例:")

	// 动态列出所有可用的示例
	fmt.Println("  常量示例:")
	fmt.Println("    constants, 1.6, 1.6.1_constants → ConstantsDemo")
	fmt.Println()
	fmt.Println("  结构体示例:")
	fmt.Println("    anonymous_struct, 1.5.2           → AnonymousStructDemo")
	fmt.Println("    nested_struct, 1.5.3              → NestedStructDemo")
	fmt.Println("    struct_methods, 1.5.4             → StructMethodsDemo")
	fmt.Println("    cross_file_usage, 1.5.4.2         → CrossFileUsageDemo")
	fmt.Println("    lowercase_struct, 1.5.4.3         → LowercaseStructDemo")
	fmt.Println("    real_world_example, 1.5.4.4       → RealWorldExampleDemo")
	fmt.Println("    tags_example                      → TagsExampleDemo")
	fmt.Println("    visibility                        → VisibilityDemo")
	fmt.Println()
	fmt.Println("  基础示例:")
	fmt.Println("    hello, basic                      → sayHello")
	fmt.Println()
	fmt.Println("  反射示例:")
	fmt.Println("    reflection                        → demonstrateReflection")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  go run main.go constants     # 自动调用 ConstantsDemo")
	fmt.Println("  go run main.go anonymous_struct # 自动调用 AnonymousStructDemo")
	fmt.Println("  go run main.go nested_struct # 自动调用 NestedStructDemo")
	fmt.Println("  go run main.go 1.5.2         # 自动调用 AnonymousStructDemo")
	fmt.Println("  go run main.go reflection    # 自动调用 demonstrateReflection")
	fmt.Println()
	fmt.Printf("当前注册了 %d 个示例函数\n", len(demoRegistry))
	fmt.Printf("支持 %d 个输入别名\n", len(aliasRegistry))
	fmt.Println("\n🚀 智能匹配: 输入名称 → 自动转换 → 调用对应Demo函数")
	fmt.Println("💡 添加新示例: 只需在 demoRegistry 中添加函数注册即可！")
}
