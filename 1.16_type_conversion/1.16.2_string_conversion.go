package typeconversion

import (
	"fmt"
	"strconv"
)

// ========== 1.16.2 字符串类型转换 ==========
//
// string 类型、[]byte 类型与[]rune 类型之间可以类似数字类型那样相互转换，
// 并且数据不会有任何丢失。
//
// 数字与字符串相互转换需要使用到 go 提供的标准库 strconv。
// strconv 可以把数字转成字符串，也可以把字符串转换成数字。

// StringConversionDemo 演示字符串类型转换
func StringConversionDemo() {
	fmt.Println("========== 1.16.2 字符串类型转换 ==========")
	fmt.Println()
	fmt.Println("string 类型、[]byte 类型与[]rune 类型之间可以类似数字类型那样相互转换，")
	fmt.Println("并且数据不会有任何丢失。")
	fmt.Println()
	fmt.Println("数字与字符串相互转换需要使用到 go 提供的标准库 strconv。")
	fmt.Println("strconv 可以把数字转成字符串，也可以把字符串转换成数字。")
	fmt.Println()

	demonstrateStringByteRuneConversion()
	demonstrateNumericStringConversion()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ string、[]byte、[]rune 之间可以相互转换，数据不会有任何丢失")
	fmt.Println("✅ 数字与字符串转换使用 strconv 包")
	fmt.Println("   - Atoi/Itoa：最常用的 int 与 string 转换")
	fmt.Println("   - ParseUint/FormatUint：无符号数字转换，支持指定进制")
	fmt.Println("   - ParseInt/FormatInt：有符号数字转换")
	fmt.Println("   - ParseFloat/FormatFloat：浮点数转换")
	fmt.Println("   - ParseBool/FormatBool：布尔值转换")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - ParseUint 只能转换成 uint64，需要其他类型需要再次转换")
	fmt.Println("   - ParseUint 和 FormatUint 需要提供 base 参数（进制）")
	fmt.Println("   - 转换失败会返回错误，需要检查错误")
	fmt.Println()
}

// demonstrateStringByteRuneConversion 演示 string、[]byte、[]rune 之间的转换
func demonstrateStringByteRuneConversion() {
	fmt.Println("=== 1.16.2.1 string、[]byte、[]rune 之间的转换 ===")

	str := "hello, 123, 你好"
	var bytes []byte = []byte(str)
	var runes []rune = []rune(str)

	fmt.Printf("原始字符串: %s\n", str)
	fmt.Printf("转换为 []byte: %v\n", bytes)
	fmt.Printf("转换为 []rune: %v\n", runes)
	fmt.Println()

	// 转换回字符串
	str2 := string(bytes)
	str3 := string(runes)
	fmt.Printf("[]byte 转回 string: %s\n", str2)
	fmt.Printf("[]rune 转回 string: %s\n", str3)
	fmt.Println("说明：string、[]byte、[]rune 之间可以相互转换，数据不会有任何丢失")
	fmt.Println()
}

// demonstrateNumericStringConversion 演示数字与字符串之间的转换
func demonstrateNumericStringConversion() {
	fmt.Println("=== 1.16.2.2 数字与字符串之间的转换 ===")

	demonstrateAtoiItoa()
	demonstrateParseUintFormatUint()
	demonstrateOtherNumericConversions()

	fmt.Println()
}

// demonstrateAtoiItoa 演示 Atoi 和 Itoa（最常用的转换）
func demonstrateAtoiItoa() {
	fmt.Println("--- 方式1：Atoi 和 Itoa（最常用）---")

	str := "123"
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串 \"%s\" 转换为 int: %d\n", str, num)

	str1 := strconv.Itoa(num)
	fmt.Printf("int %d 转换为字符串: %s\n", num, str1)
	fmt.Println()
}

// demonstrateParseUintFormatUint 演示 ParseUint 和 FormatUint
func demonstrateParseUintFormatUint() {
	fmt.Println("--- 方式2：ParseUint 和 FormatUint（无符号数字转换）---")

	str := "123"
	// ParseUint 需要三个参数：字符串、进制、位数
	// 10 表示十进制，32 表示位数
	ui64, err := strconv.ParseUint(str, 10, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串 \"%s\" 转换为 uint64: %d\n", str, ui64)

	// FormatUint 需要两个参数：数字、进制
	// 2 表示二进制
	str2 := strconv.FormatUint(ui64, 2)
	fmt.Printf("uint64 %d 转换为二进制字符串: %s\n", ui64, str2)

	// 转换为其他进制
	str3 := strconv.FormatUint(ui64, 8)
	fmt.Printf("uint64 %d 转换为八进制字符串: %s\n", ui64, str3)

	str4 := strconv.FormatUint(ui64, 16)
	fmt.Printf("uint64 %d 转换为十六进制字符串: %s\n", ui64, str4)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - ParseUint 方法把字符串转换成数字时，需要提供第二个参数 base（进制）")
	fmt.Println("  - FormatUint 方法把数字转换成字符串时，也需要提供第二个参数 base（进制）")
	fmt.Println("  - base 参数表示数字的进制，即标识字符串输出或输入的数字进制")
	fmt.Println("  - 当需要把字符串转换成无符号数字时，目前只能转换成 uint64 类型")
	fmt.Println("  - 需要其他位的数字类型需要从 uint64 类型转到所需的数字类型")
	fmt.Println()
}

// demonstrateOtherNumericConversions 演示其他数字类型转换
func demonstrateOtherNumericConversions() {
	fmt.Println("--- 方式3：其他数字类型转换 ---")

	// 字符串转 int64
	str := "12345"
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串 \"%s\" 转换为 int64: %d\n", str, i64)

	// int64 转字符串
	strFromI64 := strconv.FormatInt(i64, 10)
	fmt.Printf("int64 %d 转换为字符串: %s\n", i64, strFromI64)
	fmt.Println()

	// 字符串转 float64
	floatStr := "123.456"
	f64, err := strconv.ParseFloat(floatStr, 64)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串 \"%s\" 转换为 float64: %f\n", floatStr, f64)

	// float64 转字符串
	strFromFloat := strconv.FormatFloat(f64, 'f', 2, 64)
	fmt.Printf("float64 %f 转换为字符串（保留2位小数）: %s\n", f64, strFromFloat)
	fmt.Println()

	// 字符串转 bool
	boolStr := "true"
	b, err := strconv.ParseBool(boolStr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("字符串 \"%s\" 转换为 bool: %v\n", boolStr, b)

	// bool 转字符串
	strFromBool := strconv.FormatBool(b)
	fmt.Printf("bool %v 转换为字符串: %s\n", b, strFromBool)
	fmt.Println()
}
