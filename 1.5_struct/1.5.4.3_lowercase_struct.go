package structs

import (
	"fmt"
	"strings"
)

// ========== 演示：小写结构体名称的可见性规则 ==========
// 说明：结构体名称的大小写也遵循可见性规则
// - 首字母大写：导出的，可以被其他包访问
// - 首字母小写：未导出的，只能在同一个包内访问

// 小写结构体 - 未导出的，只能在同一个包内使用
type lowercaseStruct struct {
	// 字段可以是小写或大写
	publicField  string // 虽然字段名是大写，但结构体本身未导出
	privateField string
}

// 小写结构体的方法 - 可以是导出或未导出的
// 导出方法（首字母大写）
func (l lowercaseStruct) GetPublicField() string {
	return l.publicField
}

// 导出方法（首字母大写）
func (l *lowercaseStruct) SetPublicField(value string) {
	l.publicField = value
}

// 未导出方法（首字母小写）
func (l lowercaseStruct) getPrivateField() string {
	return l.privateField
}

// 未导出方法（首字母小写）
func (l *lowercaseStruct) setPrivateField(value string) {
	l.privateField = value
}

// 大写结构体 - 导出的，可以被其他包访问
type UppercaseStruct struct {
	PublicField  string
	privateField string
}

// 大写结构体的方法
func (u UppercaseStruct) GetPublicField() string {
	return u.PublicField
}

func (u *UppercaseStruct) SetPublicField(value string) {
	u.PublicField = value
}

// 演示小写结构体的使用
func demonstrateLowercaseStruct() {
	fmt.Println("=== 演示：小写结构体名称的可见性规则 ===")
	fmt.Println()

	fmt.Println("--- 1. 小写结构体可以在同一个包内使用 ---")
	// 可以创建小写结构体的实例（在同一个包内）
	lower := lowercaseStruct{
		publicField:  "公共字段值",
		privateField: "私有字段值",
	}
	fmt.Printf("lower.publicField = %s\n", lower.publicField)
	fmt.Printf("lower.privateField = %s\n", lower.privateField)
	fmt.Println()

	fmt.Println("--- 2. 小写结构体可以有导出和未导出的方法 ---")
	// 可以调用导出方法
	fmt.Printf("lower.GetPublicField() = %s\n", lower.GetPublicField())
	lower.SetPublicField("新的公共字段值")
	fmt.Printf("lower.SetPublicField() 后 lower.publicField = %s\n", lower.publicField)
	// 可以调用未导出方法（同包内）
	fmt.Printf("lower.getPrivateField() = %s\n", lower.getPrivateField())
	lower.setPrivateField("新的私有字段值")
	fmt.Printf("lower.setPrivateField() 后 lower.privateField = %s\n", lower.privateField)
	fmt.Println()

	fmt.Println("--- 3. 小写结构体可以嵌套在其他结构体中 ---")
	// 小写结构体可以嵌套在导出结构体中
	type Wrapper struct {
		lowercaseStruct // 可以嵌套小写结构体（同包内）
		extraField      string
	}

	wrapper := Wrapper{
		lowercaseStruct: lower,
		extraField:      "额外字段",
	}
	fmt.Printf("wrapper.publicField = %s\n", wrapper.publicField)
	fmt.Printf("wrapper.privateField = %s\n", wrapper.privateField)
	fmt.Printf("wrapper.GetPublicField() = %s\n", wrapper.GetPublicField())
	fmt.Println()

	fmt.Println("--- 4. 小写结构体可以嵌套在导出结构体中，但外部包无法直接访问 ---")
	// 导出结构体可以包含小写结构体字段
	type PublicWrapper struct {
		Lower lowercaseStruct // 字段名大写，但类型是小写结构体
		Upper UppercaseStruct // 字段名和类型都是大写
		Extra string
	}

	publicWrapper := PublicWrapper{
		Lower: lower,
		Upper: UppercaseStruct{
			PublicField:  "大写结构体的字段",
			privateField: "大写结构体的私有字段",
		},
		Extra: "额外字段",
	}
	fmt.Printf("publicWrapper.Lower.publicField = %s\n", publicWrapper.Lower.publicField)
	fmt.Printf("publicWrapper.Upper.PublicField = %s\n", publicWrapper.Upper.PublicField)
	fmt.Println()

	fmt.Println("--- 5. 对比：大写结构体和小写结构体 ---")
	upper := UppercaseStruct{
		PublicField:  "大写结构体的公共字段",
		privateField: "大写结构体的私有字段",
	}
	fmt.Printf("upper.PublicField = %s\n", upper.PublicField)
	fmt.Printf("upper.privateField = %s (同包内可以访问)\n", upper.privateField)
	fmt.Println()

	fmt.Println("--- 总结 ---")
	fmt.Println("✓ 小写结构体（lowercaseStruct）：")
	fmt.Println("  - 只能在同一个包内使用")
	fmt.Println("  - 其他包无法创建该结构体的实例")
	fmt.Println("  - 其他包无法直接访问该结构体类型")
	fmt.Println("  - 但可以有导出方法（虽然外部包无法调用，因为无法创建实例）")
	fmt.Println()
	fmt.Println("✓ 大写结构体（UppercaseStruct）：")
	fmt.Println("  - 可以被其他包访问")
	fmt.Println("  - 其他包可以创建该结构体的实例")
	fmt.Println("  - 其他包可以访问导出的字段和方法")
	fmt.Println()
	fmt.Println("✓ 重要提示：")
	fmt.Println("  - 即使小写结构体有导出方法，外部包也无法使用，因为无法创建实例")
	fmt.Println("  - 小写结构体通常用于包内部的实现细节")
	fmt.Println("  - 如果需要在外部包使用，应该使用大写结构体")
	fmt.Println()
}

// 演示：小写结构体在不同文件中的使用
func demonstrateLowercaseStructCrossFile() {
	fmt.Println("=== 演示：小写结构体在同一个包内不同文件中的使用 ===")
	fmt.Println()

	// 可以创建在另一个文件中定义的小写结构体（如果存在）
	// 例如：如果在 1.5.4.1 中定义了小写结构体，这里也可以使用

	fmt.Println("--- 同一个包内的不同文件可以访问小写结构体 ---")
	fmt.Println("✓ 小写结构体可以在同一个包内的任何文件中使用")
	fmt.Println("✓ 文件不是可见性的边界，包才是")
	fmt.Println()
}

// 主函数
// 注意：如果与同包其他文件的 main 函数冲突，可以注释掉此函数
func LowercaseStructDemo() {
	demonstrateLowercaseStruct()
	fmt.Println()
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()
	demonstrateLowercaseStructCrossFile()
}
