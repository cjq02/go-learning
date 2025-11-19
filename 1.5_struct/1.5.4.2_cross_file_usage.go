package structs

import (
	"fmt"
	"strings"
)

// ========== 演示：从另一个文件访问未导出的方法和函数 ==========
// 这个文件演示了如何从同一个包内的另一个文件访问未导出的标识符

// 演示：从文件 1.5.4.1 中访问未导出的方法和函数
func demonstrateCrossFileUsage() {
	fmt.Println("=== 演示：从另一个文件访问未导出的方法和函数 ===")
	fmt.Println()

	// 可以创建在另一个文件中定义的结构体
	demo := CrossFileDemo{
		PublicField:  "从文件 1.5.4.2 设置的公共字段",
		privateField: "从文件 1.5.4.2 设置的私有字段", // ✓ 同包内可以访问
	}

	fmt.Println("--- 1. 可以调用另一个文件中定义的未导出方法 ---")
	// 可以调用在另一个文件中定义的未导出方法
	fmt.Printf("demo.getPrivateField() = %s\n", demo.getPrivateField()) // ✓ 可以调用
	fmt.Printf("demo.GetPublicField() = %s\n", demo.GetPublicField())   // ✓ 可以调用
	fmt.Println()

	fmt.Println("--- 2. 可以调用另一个文件中定义的未导出函数 ---")
	// 可以调用在另一个文件中定义的未导出函数
	fmt.Printf("privateHelperFunction() = %s\n", privateHelperFunction()) // ✓ 可以调用
	fmt.Printf("PublicHelperFunction() = %s\n", PublicHelperFunction())   // ✓ 可以调用
	fmt.Println()

	fmt.Println("--- 3. 可以直接访问另一个文件中定义的未导出字段 ---")
	// 可以直接访问在另一个文件中定义的结构体的未导出字段
	fmt.Printf("demo.privateField = %s\n", demo.privateField) // ✓ 可以直接访问
	fmt.Println()

	// 可以调用未导出的方法修改字段
	demo.setPrivateField("通过未导出方法修改的值")
	fmt.Printf("demo.setPrivateField() 后 demo.privateField = %s\n", demo.privateField)
	fmt.Println()

	fmt.Println("--- 4. 可以定义新的结构体，嵌套另一个文件中的类型 ---")
	type Wrapper struct {
		CrossFileDemo // 可以嵌套在另一个文件中定义的类型
		extraField    string
	}

	wrapper := Wrapper{
		CrossFileDemo: demo,
		extraField:    "额外字段",
	}

	// 可以访问嵌套结构体的未导出方法和字段
	fmt.Printf("wrapper.getPrivateField() = %s\n", wrapper.getPrivateField()) // ✓ 可以调用
	fmt.Printf("wrapper.privateField = %s\n", wrapper.privateField)           // ✓ 可以直接访问
	fmt.Println()

	fmt.Println("--- 5. 可以访问另一个文件中定义的小写结构体 ---")
	// 可以创建在另一个文件中定义的小写结构体（同包内）
	lower := lowercaseStruct{
		publicField:  "从文件 1.5.4.2 设置的公共字段",
		privateField: "从文件 1.5.4.2 设置的私有字段",
	}
	fmt.Printf("lower.publicField = %s\n", lower.publicField)
	fmt.Printf("lower.privateField = %s\n", lower.privateField)
	fmt.Printf("lower.GetPublicField() = %s\n", lower.GetPublicField())
	fmt.Printf("lower.getPrivateField() = %s\n", lower.getPrivateField())
	fmt.Println()

	fmt.Println("--- 重要说明 ---")
	fmt.Println("✓ 同一个包（package）内的所有文件共享同一个命名空间")
	fmt.Println("✓ 文件只是代码的组织方式，不是可见性的边界")
	fmt.Println("✓ 可见性的边界是包（package），不是文件（file）")
	fmt.Println("✓ 因此，同一个包内的不同文件可以完全访问彼此的未导出标识符")
	fmt.Println("✓ 包括：小写结构体、小写方法、小写函数、小写变量等")
	fmt.Println()
}

// 主函数 - 演示跨文件访问
// 注意：如果与同包其他文件的 main 函数冲突，可以注释掉此函数
func CrossFileUsageDemo() {
	demonstrateCrossFileVisibility()
	fmt.Println()
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()
	demonstrateCrossFileUsage()
}
