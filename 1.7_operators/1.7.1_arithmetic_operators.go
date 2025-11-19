package operators

import "fmt"

// ========== 1.7.1 算术运算符 ==========

// demonstrateBasicArithmetic 基本算术运算演示
func demonstrateBasicArithmetic() {
	fmt.Println("=== 基本算术运算 ===")
	fmt.Println("Go 中，两个整数计算，它们计算之后的结果也还是整数。")

	a, b := 1, 2
	sum := a + b
	sub := a - b
	mul := a * b
	div := a / b // 整数除法：1/2 = 0
	mod := a % b

	fmt.Printf("a = %d, b = %d\n", a, b)
	fmt.Printf("sum (a + b) = %d\n", sum)
	fmt.Printf("sub (a - b) = %d\n", sub)
	fmt.Printf("mul (a * b) = %d\n", mul)
	fmt.Printf("div (a / b) = %d\n", div)
	fmt.Printf("mod (a %% b) = %d\n", mod)
}

// demonstrateIncrementDecrement 自增与自减运算符演示
func demonstrateIncrementDecrement() {
	fmt.Println("\n=== 自增与自减运算符 ===")
	fmt.Println("自增与自减只能以 <var name>++ 或者 <var name>-- 的模式声明")
	fmt.Println("并且只能单独存在，不能在自增或自减的同时做加减乘除的计算")

	a := 1
	fmt.Printf("初始值 a = %d\n", a)

	// 正确写法
	a++
	fmt.Printf("a++ 后: a = %d\n", a)

	a--
	fmt.Printf("a-- 后: a = %d\n", a)

	// 演示不能在表达式中使用
	fmt.Println("\n=== 错误使用方式演示（注释掉的代码会编译错误）===")
	fmt.Println("以下代码如果取消注释，会编译错误：")
	fmt.Println("// ++a      // 前缀自增不支持")
	fmt.Println("// --a      // 前缀自减不支持")
	fmt.Println("// b := a++ + 1  // 自增不能在表达式中使用")
	fmt.Println("// c := a--      // 自减不能在表达式中使用")

	// 正确的方式
	b := a + 1
	a++
	c := a
	a--

	fmt.Printf("b = a + 1 = %d + 1 = %d\n", a, b)
	fmt.Printf("c = a (在自增后) = %d\n", c)
	fmt.Printf("最终 a = %d\n", a)
}

// demonstrateTypeConversion 不同类型混合计算演示
func demonstrateTypeConversion() {
	fmt.Println("\n=== 不同类型混合计算 ===")
	fmt.Println("当不同的数字类型混合计算时，必须先把它们转换成同一类型才可以计算")

	// 错误示例（注释掉，避免编译错误）
	fmt.Println("=== 错误示例（注释掉的代码会编译错误）===")
	fmt.Println("// a := 10 + 0.1        // int + float64 不允许")
	fmt.Println("// b := byte(1) + 1      // byte + int 不允许")

	// 正确示例
	a := 10.0 + 0.1 // 都是 float64
	b := byte(1)    // byte 类型
	c := 1          // int 类型

	fmt.Printf("a (float64) = %.1f\n", a)
	fmt.Printf("b (byte) = %d\n", b)
	fmt.Printf("c (int) = %d\n", c)

	// 类型转换后才能计算
	sum := a + float64(b)
	fmt.Printf("sum = a + float64(b) = %.1f + %.1f = %.1f\n", a, float64(b), sum)

	sub := byte(a) - b
	fmt.Printf("sub = byte(a) - b = %d - %d = %d\n", byte(a), b, sub)

	mul := a * float64(b)
	div := int(a) / c

	fmt.Printf("mul = a * float64(b) = %.1f * %.1f = %.1f\n", a, float64(b), mul)
	fmt.Printf("div = int(a) / c = %d / %d = %d\n", int(a), c, div)
}

// demonstrateFloatingPointArithmetic 浮点数算术运算演示
func demonstrateFloatingPointArithmetic() {
	fmt.Println("\n=== 浮点数算术运算 ===")

	x, y := 3.14, 2.71
	sum := x + y
	sub := x - y
	mul := x * y
	div := x / y

	fmt.Printf("x = %.2f, y = %.2f\n", x, y)
	fmt.Printf("sum (x + y) = %.4f\n", sum)
	fmt.Printf("sub (x - y) = %.4f\n", sub)
	fmt.Printf("mul (x * y) = %.4f\n", mul)
	fmt.Printf("div (x / y) = %.4f\n", div)

	// 注意浮点数的精度问题
	a := 0.1
	b := 0.2
	c := a + b

	fmt.Printf("\n浮点数精度示例:\n")
	fmt.Printf("0.1 + 0.2 = %.20f (不是精确的 0.3)\n", c)
	fmt.Printf("这是浮点数二进制表示的特性导致的")
}

// demonstrateIntegerDivision 整数除法注意事项演示
func demonstrateIntegerDivision() {
	fmt.Println("\n=== 整数除法注意事项 ===")

	a, b := 7, 3
	div := a / b
	mod := a % b

	fmt.Printf("整数除法: %d / %d = %d (向下取整)\n", a, b, div)
	fmt.Printf("取余运算: %d %% %d = %d\n", a, b, mod)

	// 负数除法
	x, y := -7, 3
	div2 := x / y
	mod2 := x % y

	fmt.Printf("负数除法: %d / %d = %d\n", x, y, div2)
	fmt.Printf("负数取余: %d %% %d = %d\n", x, y, mod2)

	fmt.Println("注意：Go 中的取余运算结果的符号与被除数相同")
}

// demonstrateOperatorPrecedence 运算符优先级演示
func demonstrateOperatorPrecedence() {
	fmt.Println("\n=== 运算符优先级 ===")

	// 演示运算符优先级
	result1 := 2 + 3*4      // 3*4 先计算，然后 +2
	result2 := (2 + 3) * 4  // 2+3 先计算，然后 *4
	result3 := 10 - 2*3     // 2*3 先计算，然后 10-
	result4 := (10 - 2) * 3 // 10-2 先计算，然后 *3

	fmt.Printf("2 + 3 * 4 = %d (乘法优先级高于加法)\n", result1)
	fmt.Printf("(2 + 3) * 4 = %d (括号改变优先级)\n", result2)
	fmt.Printf("10 - 2 * 3 = %d (乘法优先级高于减法)\n", result3)
	fmt.Printf("(10 - 2) * 3 = %d (括号改变优先级)\n", result4)

	fmt.Println("\n常用运算符优先级（从高到低）:")
	fmt.Println("1. * / % （乘除取余）")
	fmt.Println("2. + - （加减）")
	fmt.Println("3. == != < <= > >= （比较）")
	fmt.Println("4. && （逻辑与）")
	fmt.Println("5. || （逻辑或）")
	fmt.Println("使用括号可以明确控制计算顺序")
}

// demonstrateArithmeticOverflow 算术溢出演示
func demonstrateArithmeticOverflow() {
	fmt.Println("\n=== 算术溢出 ===")

	// 有符号整数溢出
	var maxInt8 int8 = 127
	fmt.Printf("maxInt8 = %d\n", maxInt8)

	// 溢出（会环绕）
	overflow := maxInt8 + 1
	fmt.Printf("maxInt8 + 1 = %d (溢出环绕到最小值)\n", overflow)

	// 无符号整数溢出
	var maxUint8 uint8 = 255
	fmt.Printf("maxUint8 = %d\n", maxUint8)

	uintOverflow := maxUint8 + 1
	fmt.Printf("maxUint8 + 1 = %d (无符号整数溢出环绕到 0)\n", uintOverflow)

	fmt.Println("注意：Go 不会在运行时检查算术溢出，这是为了性能考虑")
	fmt.Println("在需要确保数值范围的场景中，需要手动检查")
}

// ArithmeticOperatorsDemo 算术运算符演示主函数
func ArithmeticOperatorsDemo() {
	fmt.Println("========== 1.7.1 算术运算符 ==========")

	demonstrateBasicArithmetic()
	demonstrateIncrementDecrement()
	demonstrateTypeConversion()
	demonstrateFloatingPointArithmetic()
	demonstrateIntegerDivision()
	demonstrateOperatorPrecedence()
	demonstrateArithmeticOverflow()

	fmt.Println("\n=== 算术运算符总结 ===")
	fmt.Println("✅ + - * / % 基本算术运算")
	fmt.Println("✅ ++ -- 自增自减（仅支持后缀形式）")
	fmt.Println("✅ 类型转换：混合类型计算前需要转换")
	fmt.Println("✅ 整数除法：结果向下取整")
	fmt.Println("✅ 浮点数：存在精度问题")
	fmt.Println("✅ 溢出：不会运行时检查，需要手动处理")
	fmt.Println("✅ 优先级：使用括号确保计算顺序正确")
}
