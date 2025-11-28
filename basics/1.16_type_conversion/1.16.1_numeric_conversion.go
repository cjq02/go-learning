package typeconversion

import "fmt"

// ========== 1.16.1 数字类型转换 ==========
//
// 在 Go 中，类型转换的基本格式如下：
// <type_name>(<expression>)
//
// type_name 为类型
// expression 为有返回值的类型
//
// 数字类型之间相互转换比较简单，并且位数较多的类型向位数较少的类型转换时，
// 高位数据会被直接截去。

// NumericConversionDemo 演示数字类型转换
func NumericConversionDemo() {
	fmt.Println("========== 1.16.1 数字类型转换 ==========")
	fmt.Println()
	fmt.Println("在 Go 中，类型转换的基本格式如下：")
	fmt.Println("  <type_name>(<expression>)")
	fmt.Println()
	fmt.Println("其中：")
	fmt.Println("  - type_name 为类型")
	fmt.Println("  - expression 为有返回值的类型")
	fmt.Println()
	fmt.Println("数字类型之间相互转换比较简单，")
	fmt.Println("并且位数较多的类型向位数较少的类型转换时，")
	fmt.Println("高位数据会被直接截去。")
	fmt.Println()

	demonstrateBasicNumericConversion()
	demonstrateTruncation()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 类型转换格式：<type_name>(<expression>)")
	fmt.Println("✅ 数字类型可以直接强转")
	fmt.Println("✅ 位数较多的类型向位数较少的类型转换时，高位会被截断")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 转换时要注意数据范围，避免数据丢失")
	fmt.Println("   - 高位截断可能导致意外的结果")
	fmt.Println()
}

// demonstrateBasicNumericConversion 演示基本的数字类型转换
func demonstrateBasicNumericConversion() {
	fmt.Println("=== 1.16.1.1 基本数字类型转换 ===")

	var i int32 = 17
	var b byte = 5

	// 数字类型可以直接强转
	f := float32(i) / float32(b)
	fmt.Printf("i = %d, b = %d\n", i, b)
	fmt.Printf("f = float32(i) / float32(b) = %f\n", f)
	fmt.Println()
}

// demonstrateTruncation 演示高位截断
func demonstrateTruncation() {
	fmt.Println("=== 1.16.1.2 高位截断示例 ===")

	// 当int32类型强转成byte时，高位被直接舍弃
	var i2 int32 = 256
	var b2 byte = byte(i2)
	fmt.Printf("i2 (int32) = %d\n", i2)
	fmt.Printf("b2 = byte(i2) = %d\n", b2)
	fmt.Println("说明：int32(256) 转换为 byte 时，高位被截断，结果为 0")
	fmt.Println()

	// 更多截断示例
	var i3 int32 = 257
	var b3 byte = byte(i3)
	fmt.Printf("i3 (int32) = %d\n", i3)
	fmt.Printf("b3 = byte(i3) = %d\n", b3)
	fmt.Println("说明：int32(257) 转换为 byte 时，结果为 1 (257 % 256)")
	fmt.Println()

	var i4 int32 = 1000
	var b4 byte = byte(i4)
	fmt.Printf("i4 (int32) = %d\n", i4)
	fmt.Printf("b4 = byte(i4) = %d\n", b4)
	fmt.Printf("说明：int32(1000) 转换为 byte 时，结果为 %d (1000 %% 256)\n", b4)
	fmt.Println()
}
