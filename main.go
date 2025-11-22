// 包声明
package main

// 导入包
import (
	"fmt"
	functions "go-learning/1.10_method"
	variablescope "go-learning/1.11_variable_scope"
	array "go-learning/1.12_array"
	slice "go-learning/1.13_slice"
	mapcollection "go-learning/1.14_map"
	rangeiteration "go-learning/1.15_range"
	pointers "go-learning/1.4_pointer"
	structs "go-learning/1.5_struct"
	constants "go-learning/1.6_constants_enum"
	operators "go-learning/1.7_operators"
	controlflow "go-learning/1.8_control_flow"
	loopcontrol "go-learning/1.9_loop_control"
	"os"
	"reflect"
	"strings"
)

// 函数声明

// 变量声明

// DemoRegistry 示例注册表 - 函数名到函数的映射
var demoRegistry = map[string]interface{}{
	// 指针示例
	"PointersDemo": pointers.PointersDemo,
	// 函数示例
	"FunctionsDemo": functions.FunctionsDemo,
	"ClosureDemo":   functions.ClosureDemo,
	"MethodDemo":    functions.MethodDemo,
	// 流程控制示例
	"IfStatementDemo":     controlflow.IfStatementDemo,
	"SwitchStatementDemo": controlflow.SwitchStatementDemo,
	"ForLoopDemo":         loopcontrol.ForLoopDemo,
	"BreakDemo":           loopcontrol.BreakDemo,
	"ContinueDemo":        loopcontrol.ContinueDemo,
	"GotoDemo":            loopcontrol.GotoDemo,
	// 变量作用域示例
	"LocalVariableDemo":  variablescope.LocalVariableDemo,
	"GlobalVariableDemo": variablescope.GlobalVariableDemo,
	// 数组示例
	"ArrayDeclarationDemo":      array.ArrayDeclarationDemo,
	"ArrayAccessDemo":           array.ArrayAccessDemo,
	"MultidimensionalArrayDemo": array.MultidimensionalArrayDemo,
	"ArrayAsParameterDemo":      array.ArrayAsParameterDemo,
	// 切片示例
	"SliceDeclarationDemo":         slice.SliceDeclarationDemo,
	"SliceUsageDemo":               slice.SliceUsageDemo,
	"SliceUnderlyingPrincipleDemo": slice.SliceUnderlyingPrincipleDemo,
	// map 示例
	"MapDeclarationDemo": mapcollection.MapDeclarationDemo,
	"MapUsageDemo":       mapcollection.MapUsageDemo,
	"MapAsParameterDemo": mapcollection.MapAsParameterDemo,
	"MapConcurrentDemo":  mapcollection.MapConcurrentDemo,
	// range 迭代示例
	"RangeStringDemo":     rangeiteration.RangeStringDemo,
	"RangeArraySliceDemo": rangeiteration.RangeArraySliceDemo,
	"RangeChannelDemo":    rangeiteration.RangeChannelDemo,
	// 结构体示例
	"AnonymousStructDemo":  structs.AnonymousStructDemo,
	"NestedStructDemo":     structs.NestedStructDemo,
	"StructMethodsDemo":    structs.StructMethodsDemo,
	"CrossFileUsageDemo":   structs.CrossFileUsageDemo,
	"LowercaseStructDemo":  structs.LowercaseStructDemo,
	"RealWorldExampleDemo": structs.RealWorldExampleDemo,
	// 常量示例
	"ConstantsDemo": constants.ConstantsDemo,
	"EnumsDemo":     constants.EnumsDemo,

	// 运算符示例
	"ArithmeticOperatorsDemo": operators.ArithmeticOperatorsDemo,
	"OperatorsDemo":           operators.OperatorsDemo,
}

// aliasRegistry 别名映射 - 动态生成
var aliasRegistry = generateAliasRegistry()

// generateAliasRegistry 动态生成别名映射
// 从 demoRegistry 的 key 中去掉 "Demo" 后缀生成别名
func generateAliasRegistry() map[string]string {
	aliases := make(map[string]string)

	for funcName := range demoRegistry {
		// 去掉 "Demo" 后缀
		if strings.HasSuffix(funcName, "Demo") {
			alias := strings.TrimSuffix(funcName, "Demo")
			// 将首字母转换为小写
			if len(alias) > 0 {
				alias = strings.ToLower(alias[:1]) + alias[1:]
			}
			aliases[alias] = funcName
		}
	}

	return aliases
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
	fmt.Println("📝 命名规则: 示例名 + 'Demo' = 函数名")
	fmt.Println()
	fmt.Println("可用示例:")

	// 动态列出所有可用的示例
	fmt.Println("  可用示例:")

	// 动态显示所有别名和对应的函数
	for alias, funcName := range aliasRegistry {
		fmt.Printf("    %-15s → %s\n", alias, funcName)
	}

	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  go run main.go constants      # 自动调用 ConstantsDemo")
	fmt.Println("  go run main.go anonymousStruct # 自动调用 AnonymousStructDemo")
	fmt.Println("  go run main.go nestedStruct   # 自动调用 NestedStructDemo")
	fmt.Println("  go run main.go structMethods  # 自动调用 StructMethodsDemo")
	fmt.Println()
	fmt.Printf("当前注册了 %d 个示例函数\n", len(demoRegistry))
	fmt.Printf("支持 %d 个输入别名\n", len(aliasRegistry))
	fmt.Println("\n🚀 智能匹配: 输入名称 → 自动转换 → 调用对应Demo函数")
	fmt.Println("💡 添加新示例: 只需在 demoRegistry 中添加函数注册即可！")
}
