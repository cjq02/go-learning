package rangeiteration

import "fmt"

// ========== 1.15 range 迭代 ==========
//
// 在 Go 中，range 关键字用于 for 循环迭代：
//   - 字符串(string)
//   - 数组(array)
//   - 切片(slice)
//   - 通道(channel)
//   - 映射集合(map)
//
// 1.15.1 对字符串迭代
// 在 Go 中，string 类型是一个比较特殊的类型，可以与 rune 切片类型、byte 切片类型相互转换
// 同时还可以使用 range 关键字来遍历一个字符串

// demonstrateStringRangeIndexOnly 演示方式1：仅使用 range 获取下标索引
func demonstrateStringRangeIndexOnly() {
	fmt.Println("=== 1.15.1.1 方式1：仅使用 range 获取下标索引 ===")

	str1 := "abc123"
	fmt.Printf("字符串 str1 = \"%s\"\n", str1)
	fmt.Println("遍历字符串（仅获取索引）：")
	for index := range str1 {
		fmt.Printf("str1 -- index:%d, value:%d\n", index, str1[index])
	}
	fmt.Println()

	str2 := "测试中文"
	fmt.Printf("字符串 str2 = \"%s\"\n", str2)
	fmt.Println("遍历字符串（仅获取索引）：")
	for index := range str2 {
		fmt.Printf("str2 -- index:%d, value:%d\n", index, str2[index])
	}
	fmt.Printf("len(str2) = %d\n", len(str2))

	runesFromStr2 := []rune(str2)
	bytesFromStr2 := []byte(str2)
	fmt.Printf("len(runesFromStr2) = %d\n", len(runesFromStr2))
	fmt.Printf("len(bytesFromStr2) = %d\n", len(bytesFromStr2))
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - str1 有6个字符，每个字符用1个byte表示，循环6次")
	fmt.Println("  - str2 有4个中文字符，但循环4次（按rune遍历）")
	fmt.Println("  - len(str2) = 12（字节数）")
	fmt.Println("  - len(runesFromStr2) = 4（rune数量，即字符数）")
	fmt.Println("  - len(bytesFromStr2) = 12（字节数）")
	fmt.Println()
	fmt.Println("⚠️ 重要理解：")
	fmt.Println("  - 在 Go 中，所有字符串都是按照 Unicode 编码的")
	fmt.Println("  - 遍历字符串时，实际上是在遍历从字符串转换来的 rune 切片")
	fmt.Println("  - 中文字符需要多个 byte 表示，但 range 按 rune（字符）遍历")
	fmt.Println()
}

// demonstrateStringRangeIndexAndValue 演示方式2：使用 range 获取下标和字符
func demonstrateStringRangeIndexAndValue() {
	fmt.Println("=== 1.15.1.2 方式2：使用 range 获取下标和字符 ===")

	str1 := "a1中文"
	fmt.Printf("字符串 str1 = \"%s\"\n", str1)
	fmt.Println("遍历字符串（获取索引和值）：")
	for index, value := range str1 {
		fmt.Printf("str1 -- index:%d, index value:%d (byte值)\n", index, str1[index])
		fmt.Printf("str1 -- index:%d, range value:%d (rune值, Unicode码点)\n", index, value)
		fmt.Printf("str1 -- index:%d, range value:%c (字符)\n", index, value)
		fmt.Println()
	}
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - str1[index]：直接使用下标取字符串某个位置的值，取出的是 byte 值")
	fmt.Println("  - range value：使用 range 获取的值，是完整的 rune 类型的值（Unicode码点）")
	fmt.Println("  - 对于 ASCII 字符（如 'a', '1'），byte 值和 rune 值相同")
	fmt.Println("  - 对于多字节字符（如中文），byte 值只是字节，rune 值是完整的 Unicode 码点")
	fmt.Println()
}

// demonstrateStringUnicodeEncoding 演示字符串的 Unicode 编码
func demonstrateStringUnicodeEncoding() {
	fmt.Println("=== 字符串的 Unicode 编码详解 ===")

	str := "a中"
	fmt.Printf("字符串 str = \"%s\"\n", str)
	fmt.Println()

	// 字节表示
	fmt.Println("--- 字节表示（byte）---")
	bytes := []byte(str)
	fmt.Printf("[]byte(str) = %v\n", bytes)
	fmt.Printf("len([]byte(str)) = %d (字节数)\n", len(bytes))
	for i, b := range bytes {
		fmt.Printf("  bytes[%d] = %d (0x%02x)\n", i, b, b)
	}
	fmt.Println()

	// Rune 表示
	fmt.Println("--- Rune 表示（Unicode码点）---")
	runes := []rune(str)
	fmt.Printf("[]rune(str) = %v\n", runes)
	fmt.Printf("len([]rune(str)) = %d (字符数)\n", len(runes))
	for i, r := range runes {
		fmt.Printf("  runes[%d] = %d (U+%04X, 字符: %c)\n", i, r, r, r)
	}
	fmt.Println()

	// Range 遍历
	fmt.Println("--- Range 遍历 ---")
	fmt.Println("range 遍历字符串时，按 rune（字符）遍历：")
	for index, value := range str {
		fmt.Printf("  index:%d, value:%d (U+%04X, 字符: %c)\n", index, value, value, value)
		fmt.Printf("    对应的 byte: str[%d] = %d\n", index, str[index])
	}
	fmt.Println()

	fmt.Println("关键理解：")
	fmt.Println("  - 'a' 是 ASCII 字符，用1个字节表示（97 = 0x61）")
	fmt.Println("  - '中' 是中文字符，用3个字节表示（UTF-8编码）")
	fmt.Println("  - range 遍历时，index 是字节位置，value 是 rune（Unicode码点）")
	fmt.Println("  - 中文字符的 index 会跳跃（0, 3），因为每个字符占3个字节")
	fmt.Println()
}

// demonstrateStringRangeComparison 对比不同遍历方式
func demonstrateStringRangeComparison() {
	fmt.Println("=== 不同遍历方式对比 ===")

	str := "Hello世界"
	fmt.Printf("字符串 str = \"%s\"\n", str)
	fmt.Println()

	// 方式1：使用下标遍历（按字节）
	fmt.Println("--- 方式1：使用下标遍历（按字节）---")
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d] = %d (0x%02x, byte)\n", i, str[i], str[i])
	}
	fmt.Println()

	// 方式2：使用 range 遍历（按 rune）
	fmt.Println("--- 方式2：使用 range 遍历（按 rune）---")
	for index, value := range str {
		fmt.Printf("index:%d, value:%d (U+%04X, 字符: %c)\n", index, value, value, value)
	}
	fmt.Println()

	// 方式3：转换为 rune 切片后遍历
	fmt.Println("--- 方式3：转换为 rune 切片后遍历 ---")
	runes := []rune(str)
	for i, r := range runes {
		fmt.Printf("runes[%d] = %d (U+%04X, 字符: %c)\n", i, r, r, r)
	}
	fmt.Println()

	fmt.Println("对比总结：")
	fmt.Println("  - 方式1：按字节遍历，适合处理字节数据")
	fmt.Println("  - 方式2：按字符（rune）遍历，适合处理文本")
	fmt.Println("  - 方式3：先转换再遍历，可以获得字符索引（0,1,2...）")
	fmt.Println()
}

// demonstrateStringRuneByteConversion 演示字符串与 rune/byte 切片的转换
func demonstrateStringRuneByteConversion() {
	fmt.Println("=== 字符串与 rune/byte 切片的转换 ===")

	str := "测试"
	fmt.Printf("原始字符串 str = \"%s\"\n", str)
	fmt.Println()

	// 字符串转 byte 切片
	bytes := []byte(str)
	fmt.Printf("[]byte(str) = %v\n", bytes)
	fmt.Printf("len([]byte(str)) = %d\n", len(bytes))

	// 字符串转 rune 切片
	runes := []rune(str)
	fmt.Printf("[]rune(str) = %v\n", runes)
	fmt.Printf("len([]rune(str)) = %d\n", len(runes))
	fmt.Println()

	// byte 切片转字符串
	strFromBytes := string(bytes)
	fmt.Printf("string([]byte) = \"%s\"\n", strFromBytes)

	// rune 切片转字符串
	strFromRunes := string(runes)
	fmt.Printf("string([]rune) = \"%s\"\n", strFromRunes)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - string 和 []byte 可以相互转换")
	fmt.Println("  - string 和 []rune 可以相互转换")
	fmt.Println("  - []byte 和 []rune 不能直接转换（需要先转 string）")
	fmt.Println()
}

// RangeStringDemo range 迭代字符串完整演示
func RangeStringDemo() {
	fmt.Println("========== 1.15.1 对字符串迭代 ==========")
	fmt.Println()
	fmt.Println("在 Go 中，string 类型是一个比较特殊的类型，")
	fmt.Println("可以与 rune 切片类型、byte 切片类型相互转换，")
	fmt.Println("同时还可以使用 range 关键字来遍历一个字符串。")
	fmt.Println()
	fmt.Println("重要概念：")
	fmt.Println("  - 在 Go 中，所有字符串都是按照 Unicode 编码的")
	fmt.Println("  - 遍历字符串时，实际上是在遍历从字符串转换来的 rune 切片")
	fmt.Println("  - 直接使用下标取字符串某个位置的值，取出的是 byte 值")
	fmt.Println("  - 使用 range 获取的值，是完整的 rune 类型的值（Unicode码点）")
	fmt.Println()

	demonstrateStringRangeIndexOnly()
	demonstrateStringRangeIndexAndValue()
	demonstrateStringUnicodeEncoding()
	demonstrateStringRangeComparison()
	demonstrateStringRuneByteConversion()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 方式1：for index := range str（仅获取索引）")
	fmt.Println("✅ 方式2：for index, value := range str（获取索引和rune值）")
	fmt.Println("✅ str[index]：获取的是 byte 值")
	fmt.Println("✅ range value：获取的是 rune 值（Unicode码点）")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 字符串按 Unicode 编码")
	fmt.Println("   - range 遍历按 rune（字符）遍历，不是按 byte 遍历")
	fmt.Println("   - 中文字符需要多个 byte，但 range 只遍历一次")
	fmt.Println("   - 字符串长度 len(str) 返回的是字节数，不是字符数")
	fmt.Println()
}
