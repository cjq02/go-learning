package main

import "fmt"

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

// PublicFunc 公开的函数
func PublicFunc() string {
	return "我是公开的函数"
}

// privateFunc 非公开的函数
func privateFunc() string {
	return "我是非公开的函数"
}

// (ps *PublicStruct) PublicMethod 公开的方法
func (ps *PublicStruct) PublicMethod() string {
	return "我是公开的方法"
}

// (ps *PublicStruct) privateMethod 非公开的方法
func (ps *PublicStruct) privateMethod() string {
	return "我是非公开的方法"
}

// (ps *privateStruct) method 私有结构体的方法
func (ps *privateStruct) method() string {
	return "私有结构体的方法"
}

// ========== 在同一个包内，可以访问所有标识符 ==========

func demonstrateSamePackage() {
	fmt.Println("=== 同一个包内的可见性演示 ===")

	// 可以访问公开的全局变量
	fmt.Printf("公开全局变量: %s\n", PublicVar)

	// 可以访问非公开的全局变量（在同一个包内）
	fmt.Printf("非公开全局变量: %s\n", privateVar)

	// 可以访问公开的函数
	fmt.Printf("公开函数: %s\n", PublicFunc())

	// 可以访问非公开的函数（在同一个包内）
	fmt.Printf("非公开函数: %s\n", privateFunc())

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

	// 可以调用公开方法
	fmt.Printf("公开方法: %s\n", ps.PublicMethod())

	// 可以调用私有方法（在同一个包内）
	fmt.Printf("私有方法: %s\n", ps.privateMethod())

	// 可以创建和使用私有结构体（在同一个包内）
	priv := privateStruct{field: "私有字段"}
	fmt.Printf("私有结构体: %+v\n", priv)
	fmt.Printf("私有结构体方法: %s\n", priv.method())

	fmt.Println()
}
