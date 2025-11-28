package operators

import "fmt"

// ========== 1.7.5 赋值运算符 ==========

// demonstrateBasicAssignment 基本赋值运算符演示
func demonstrateBasicAssignment() {
	fmt.Println("=== 基本赋值运算符 ===")

	a, b := 1, 2
	var c int

	fmt.Printf("初始值: a = %d, b = %d\n", a, b)

	// 基本赋值
	c = a + b
	fmt.Printf("c = a + b, c = %d\n", c)
}

// demonstrateCompoundAssignment 复合赋值运算符演示
func demonstrateCompoundAssignment() {
	fmt.Println("\n=== 复合赋值运算符 ===")

	// 加法赋值
	func() {
		c, a := 10, 5
		fmt.Printf("加法赋值前: c = %d, a = %d\n", c, a)
		c += a // c = c + a
		fmt.Printf("c += a, c = %d\n", c)
	}()

	// 减法赋值
	func() {
		c, a := 10, 3
		fmt.Printf("减法赋值前: c = %d, a = %d\n", c, a)
		c -= a // c = c - a
		fmt.Printf("c -= a, c = %d\n", c)
	}()

	// 乘法赋值
	func() {
		c, a := 10, 3
		fmt.Printf("乘法赋值前: c = %d, a = %d\n", c, a)
		c *= a // c = c * a
		fmt.Printf("c *= a, c = %d\n", c)
	}()

	// 除法赋值
	func() {
		c, a := 20, 4
		fmt.Printf("除法赋值前: c = %d, a = %d\n", c, a)
		c /= a // c = c / a
		fmt.Printf("c /= a, c = %d\n", c)
	}()

	// 取余赋值
	func() {
		c, a := 17, 5
		fmt.Printf("取余赋值前: c = %d, a = %d\n", c, a)
		c %= a // c = c % a
		fmt.Printf("c %%= a, c = %d\n", c)
	}()
}

// demonstrateBitwiseAssignment 位运算赋值演示
func demonstrateBitwiseAssignment() {
	fmt.Println("\n=== 位运算赋值运算符 ===")

	// 左移赋值
	func() {
		c, a := 12, 2
		fmt.Printf("左移赋值前: c = %d (%04b), a = %d\n", c, c, a)
		c <<= a // c = c << a
		fmt.Printf("c <<= a, c = %d (%04b)\n", c, c)
	}()

	// 右移赋值
	func() {
		c, a := 48, 2
		fmt.Printf("右移赋值前: c = %d (%08b), a = %d\n", c, c, a)
		c >>= a // c = c >> a
		fmt.Printf("c >>= a, c = %d (%08b)\n", c, c)
	}()

	// 按位与赋值
	func() {
		c, a := 60, 13 // 60 = 111100, 13 = 001101
		fmt.Printf("按位与赋值前: c = %d (%06b), a = %d (%06b)\n", c, c, a, a)
		c &= a // c = c & a
		fmt.Printf("c &= a, c = %d (%06b)\n", c, c)
	}()

	// 按位或赋值
	func() {
		c, a := 60, 13
		fmt.Printf("按位或赋值前: c = %d (%06b), a = %d (%06b)\n", c, c, a, a)
		c |= a // c = c | a
		fmt.Printf("c |= a, c = %d (%06b)\n", c, c)
	}()

	// 按位异或赋值
	func() {
		c, a := 60, 13
		fmt.Printf("按位异或赋值前: c = %d (%06b), a = %d (%06b)\n", c, c, a, a)
		c ^= a // c = c ^ a
		fmt.Printf("c ^= a, c = %d (%06b)\n", c, c)
	}()

	// 按位清除赋值
	func() {
		c, a := 60, 13
		fmt.Printf("按位清除赋值前: c = %d (%06b), a = %d (%06b)\n", c, c, a, a)
		c &^= a // c = c &^ a
		fmt.Printf("c &^= a, c = %d (%06b)\n", c, c)
	}()
}

// demonstrateMultipleAssignment 多重赋值演示
func demonstrateMultipleAssignment() {
	fmt.Println("\n=== 多重赋值 ===")

	// 同时给多个变量赋值
	a, b, c := 1, 2, 3
	fmt.Printf("多重赋值: a, b, c = %d, %d, %d\n", a, b, c)

	// 交换两个变量的值
	x, y := 10, 20
	fmt.Printf("交换前: x = %d, y = %d\n", x, y)
	x, y = y, x
	fmt.Printf("交换后: x = %d, y = %d\n", x, y)

	// 函数返回值赋值
	sum, diff := calculate(10, 3)
	fmt.Printf("函数返回值赋值: sum = %d, diff = %d\n", sum, diff)
}

// calculate 示例函数，返回两个值
func calculate(a, b int) (int, int) {
	return a + b, a - b
}

// demonstrateUnderscoreAssignment 下划线赋值演示
func demonstrateUnderscoreAssignment() {
	fmt.Println("\n=== 下划线赋值（忽略值）===")

	// 忽略不需要的值
	sum, _ := calculate(10, 3) // 忽略差值
	fmt.Printf("只使用和: sum = %d\n", sum)

	// 忽略索引，只关心值
	numbers := []int{10, 20, 30, 40, 50}
	fmt.Print("数组元素: ")
	for _, num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()

	// 忽略值，只关心索引
	fmt.Print("数组索引: ")
	for index, _ := range numbers {
		fmt.Printf("%d ", index)
	}
	fmt.Println()
}

// AssignmentOperatorsDemo 赋值运算符演示主函数
func AssignmentOperatorsDemo() {
	fmt.Println("========== 1.7.5 赋值运算符 ==========")

	demonstrateBasicAssignment()
	demonstrateCompoundAssignment()
	demonstrateBitwiseAssignment()
	demonstrateMultipleAssignment()
	demonstrateUnderscoreAssignment()

	fmt.Println("\n=== 赋值运算符总结 ===")
	fmt.Println("✅ = 基本赋值")
	fmt.Println("✅ += -= *= /= %= 算术赋值")
	fmt.Println("✅ <<= >>= &= |= ^= &^= 位运算赋值")
	fmt.Println("✅ 多重赋值：a, b = b, a")
	fmt.Println("✅ 下划线赋值：忽略不需要的值")
	fmt.Println("✅ 复合赋值等价于：a += b 等价于 a = a + b")
}
