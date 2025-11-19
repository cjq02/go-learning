package main

import "fmt"

func main() {
	fmt.Println("=== Go语言可见性规则演示 ===")
	fmt.Println()

	// ========== 演示：同一个包内的可见性 ==========

	// PublicVar 公开的全局变量（大写开头）
	var PublicVar = "我是公开的全局变量"

	// privateVar 非公开的全局变量（小写开头）
	var privateVar = "我是非公开的全局变量"

	// PublicStruct 公开的结构体
	type PublicStruct struct {
		// PublicField 公开的字段
		PublicField string
		// privateField 非公开的字段
		privateField int
	}

	// privateStruct 非公开的结构体
	type privateStruct struct {
		field string
	}

	fmt.Println("=== 同一个包内的可见性演示 ===")

	// 可以访问公开的全局变量
	fmt.Printf("公开全局变量: %s\n", PublicVar)

	// 可以访问非公开的全局变量（在同一个包内）
	fmt.Printf("非公开全局变量: %s\n", privateVar)

	// 可以创建公开结构体实例
	ps := PublicStruct{
		PublicField:  "公开字段值",
		privateField: 42, // 在同一个包内可以访问私有字段
	}

	fmt.Printf("公开结构体: %+v\n", ps)

	// 可以访问公开字段
	fmt.Printf("公开字段: %s\n", ps.PublicField)

	// 可以访问私有字段（在同一个包内）
	fmt.Printf("私有字段: %d\n", ps.privateField)

	// 可以创建和使用私有结构体（在同一个包内）
	priv := privateStruct{field: "私有字段"}
	fmt.Printf("私有结构体: %+v\n", priv)

	fmt.Println()

	// ========== 总结 ==========
	fmt.Println("=== 可见性规则总结 ===")
	fmt.Println("1. 公开标识符（首字母大写）：可以被任何包访问")
	fmt.Println("2. 非公开标识符（首字母小写）：只能在相同包内访问")
	fmt.Println("3. 包是Go语言的基本可见性边界")
	fmt.Println("4. 没有其他可见性修饰符（如public、private、protected）")
	fmt.Println("5. 同一个包内的不同文件可以相互访问对方的非公开标识符")
}