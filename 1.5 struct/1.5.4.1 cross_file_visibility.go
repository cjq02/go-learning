package main

import "fmt"

// ========== 演示：同一个包内不同文件之间的可见性 ==========
// 说明：在 Go 中，同一个包（package）内的所有文件共享同一个命名空间
// 因此，同一个包内的不同文件可以互相访问未导出的（小写）方法、函数、变量、类型等

// 结构体 CrossFileDemo - 定义在文件 1.5.4.1 cross_file_visibility.go 中
type CrossFileDemo struct {
	// 导出字段
	PublicField string
	// 未导出字段
	privateField string
}

// 导出方法 - 其他包可以调用
func (c CrossFileDemo) GetPublicField() string {
	return c.PublicField
}

// 未导出方法 - 只能在同一个包内调用（包括其他文件）
func (c CrossFileDemo) getPrivateField() string {
	return c.privateField
}

// 未导出方法 - 设置私有字段
func (c *CrossFileDemo) setPrivateField(value string) {
	c.privateField = value
}

// 未导出函数 - 只能在同一个包内调用
func privateHelperFunction() string {
	return "这是未导出的辅助函数，可以在同包的其他文件中调用"
}

// 导出函数 - 其他包可以调用
func PublicHelperFunction() string {
	return "这是导出的辅助函数，任何包都可以调用"
}

// 演示：在同一个包内的不同文件中调用未导出的方法和函数
func demonstrateCrossFileVisibility() {
	fmt.Println("=== 演示：同一个包内不同文件之间的可见性 ===")
	fmt.Println()

	// 创建结构体实例
	demo := CrossFileDemo{
		PublicField:  "公共字段值",
		privateField: "私有字段值",
	}

	fmt.Println("--- 1. 在同一个包内可以访问未导出的方法 ---")
	fmt.Printf("demo.getPrivateField() = %s\n", demo.getPrivateField())
	fmt.Printf("demo.GetPublicField() = %s\n", demo.GetPublicField())
	fmt.Println()

	fmt.Println("--- 2. 在同一个包内可以调用未导出的函数 ---")
	fmt.Printf("privateHelperFunction() = %s\n", privateHelperFunction())
	fmt.Printf("PublicHelperFunction() = %s\n", PublicHelperFunction())
	fmt.Println()

	fmt.Println("--- 3. 在同一个包内可以访问未导出的字段 ---")
	fmt.Printf("demo.privateField = %s (同包内可以直接访问)\n", demo.privateField)
	fmt.Println()

	fmt.Println("--- 总结 ---")
	fmt.Println("✓ 同一个包内的不同文件可以互相访问：")
	fmt.Println("  - 未导出的方法（小写开头）")
	fmt.Println("  - 未导出的函数（小写开头）")
	fmt.Println("  - 未导出的变量（小写开头）")
	fmt.Println("  - 未导出的类型（小写开头）")
	fmt.Println("  - 未导出的结构体字段（小写开头）")
	fmt.Println()
	fmt.Println("✓ 只有跨包访问时，才需要首字母大写的导出标识符")
	fmt.Println()
}
