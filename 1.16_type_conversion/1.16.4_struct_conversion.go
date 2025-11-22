package typeconversion

import "fmt"

// ========== 1.16.4 结构体类型转换 ==========
//
// 结构体类型之间在一定条件下也可以转换的。
//
// 当两个结构体中的字段名称以及类型都完全相同，仅结构体名称不同时，
// 这两个结构体类型即可相互转换。

// StructConversionDemo 演示结构体类型转换
func StructConversionDemo() {
	fmt.Println("========== 1.16.4 结构体类型转换 ==========")
	fmt.Println()
	fmt.Println("结构体类型之间在一定条件下也可以转换的。")
	fmt.Println()
	fmt.Println("当两个结构体中的字段名称以及类型都完全相同，")
	fmt.Println("仅结构体名称不同时，这两个结构体类型即可相互转换。")
	fmt.Println()

	demonstrateSameFieldStructConversion()
	demonstratePointerConversionLimitation()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 当两个结构体的字段名称和类型都完全相同时，可以相互转换")
	fmt.Println("✅ 转换后可以使用目标结构体的方法")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 只能结构体类型实例之间相互转换")
	fmt.Println("   - 指针类型不可以相互转换，即使字段完全相同")
	fmt.Println("   - 如果需要转换指针，需要先解引用，转换后再取地址")
	fmt.Println()
}

// SameFieldA 结构体A
type SameFieldA struct {
	name  string
	value int
}

// SameFieldB 结构体B（字段与A完全相同）
type SameFieldB struct {
	name  string
	value int
}

// getValue 为 SameFieldB 添加方法
func (s *SameFieldB) getValue() int {
	return s.value
}

// demonstrateSameFieldStructConversion 演示相同字段的结构体转换
func demonstrateSameFieldStructConversion() {
	fmt.Println("=== 1.16.4.1 相同字段的结构体转换 ===")

	a := SameFieldA{
		name:  "a",
		value: 1,
	}

	fmt.Printf("转换前 - a 的类型: %T\n", a)
	fmt.Printf("转换前 - a.name = %s, a.value = %d\n", a.name, a.value)
	fmt.Println()

	// 结构体类型实例之间相互转换
	b := SameFieldB(a)
	fmt.Printf("转换后 - b 的类型: %T\n", b)
	fmt.Printf("转换后 - b.name = %s, b.value = %d\n", b.name, b.value)
	fmt.Printf("转换后 - b.getValue() = %d\n", b.getValue())
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 当两个结构体的字段名称和类型都完全相同时，可以相互转换")
	fmt.Println("  - 转换后可以使用目标结构体的方法")
	fmt.Println()
}

// demonstratePointerConversionLimitation 演示指针类型转换的限制
func demonstratePointerConversionLimitation() {
	fmt.Println("=== 1.16.4.2 指针类型转换的限制 ===")

	a := SameFieldA{
		name:  "a",
		value: 1,
	}

	fmt.Printf("a 的类型: %T\n", a)
	fmt.Printf("&a 的类型: %T\n", &a)
	fmt.Println()

	// 结构体实例可以转换
	b := SameFieldB(a)
	fmt.Printf("SameFieldB(a) 成功，b 的类型: %T\n", b)
	fmt.Println()

	// 指针类型不能直接转换（注释掉的代码）
	fmt.Println("注意：只能结构体类型实例之间相互转换，指针不可以相互转换")
	fmt.Println("  以下代码会编译失败：")
	fmt.Println("  var c interface{} = &a")
	fmt.Println("  _, ok := c.(*SameFieldB)")
	fmt.Println("  // 即使字段相同，指针类型也不能直接转换")
	fmt.Println()

	// 演示：即使通过 interface{} 也不能转换指针类型
	var c interface{} = &a
	_, ok := c.(*SameFieldB)
	fmt.Printf("尝试将 *SameFieldA 转换为 *SameFieldB: ok = %v\n", ok)
	fmt.Println("  结果：转换失败，即使字段相同，指针类型也不能直接转换")
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 只能结构体类型实例之间相互转换")
	fmt.Println("  - 指针类型不可以相互转换，即使字段完全相同")
	fmt.Println("  - 如果需要转换指针，需要先解引用，转换后再取地址")
	fmt.Println()
}

