package variablescope

import "fmt"

// ========== 1.11.2 全局变量 ==========
//
// 全局变量：在函数外声明的变量
// 作用域：可以是当前整个包甚至外部包（公开的全局变量）使用
//
// 变量遮蔽（Variable Shadowing）：
// 当全局变量和局部变量重名时，函数内会使用局部变量
// 超出局部变量作用域之后，才会重新使用全局变量

// 全局变量声明（包级别）
var a int

// 公开的全局变量（首字母大写，可以被外部包访问）
var GlobalCounter int = 100

// 私有的全局变量（首字母小写，只能在当前包内访问）
var privateVar int = 200

// demonstrateGlobalVariable 演示全局变量的基本使用
func demonstrateGlobalVariable() {
	fmt.Println("=== 1. 全局变量的基本使用 ===")

	fmt.Printf("全局变量 a = %d\n", a)
	fmt.Printf("公开的全局变量 GlobalCounter = %d\n", GlobalCounter)
	fmt.Printf("私有的全局变量 privateVar = %d\n", privateVar)

	// 修改全局变量
	a = 1
	GlobalCounter = 101
	privateVar = 201

	fmt.Printf("修改后，全局变量 a = %d\n", a)
	fmt.Printf("修改后，GlobalCounter = %d\n", GlobalCounter)
	fmt.Printf("修改后，privateVar = %d\n", privateVar)
	fmt.Println()
}

// demonstrateVariableShadowing 演示全局变量和局部变量的遮蔽
func demonstrateVariableShadowing() {
	fmt.Println("=== 2. 全局变量和局部变量的遮蔽（Variable Shadowing）===")
	fmt.Println("说明：当全局变量和局部变量重名时，局部变量会遮蔽全局变量")

	// 重置全局变量 a
	a = 0

	{
		fmt.Printf("代码块开始，全局变量 a = %d\n", a)
		a = 3
		fmt.Printf("修改全局变量后，a = %d\n", a)

		// 声明局部变量 a（遮蔽全局变量 a）
		a := 10
		fmt.Printf("声明局部变量 a 后，局部变量 a = %d\n", a)
		a--
		fmt.Printf("局部变量 a-- 后，局部变量 a = %d\n", a)
		fmt.Println("⚠️ 注意：在这个代码块内，局部变量 a 遮蔽了全局变量 a")
	}

	// 代码块结束后，局部变量 a 失效，重新使用全局变量 a
	fmt.Printf("代码块结束后，全局变量 a = %d\n", a)
	fmt.Println("说明：超出局部变量作用域后，重新使用全局变量")
	fmt.Println()
}

// demonstrateLocalVariableShadowing 演示局部变量之间的遮蔽
func demonstrateLocalVariableShadowing() {
	fmt.Println("=== 3. 局部变量之间的遮蔽 ===")
	fmt.Println("说明：作用域更小的变量会遮蔽作用域更大的变量")

	var b int = 4
	fmt.Printf("函数级局部变量 b = %d\n", b)

	if b := 3; b == 3 {
		fmt.Printf("if 语句中的局部变量 b = %d\n", b)
		b--
		fmt.Printf("if 语句中的局部变量 b-- 后，b = %d\n", b)
		fmt.Println("⚠️ 注意：在 if 语句内，if 的局部变量 b 遮蔽了函数级变量 b")
	}

	// if 语句结束后，if 的局部变量 b 失效，重新使用函数级变量 b
	fmt.Printf("if 语句结束后，函数级局部变量 b = %d\n", b)
	fmt.Println("说明：超出 if 语句作用域后，重新使用函数级变量 b")
	fmt.Println()
}

// demonstrateErrorVariableShadowing 演示 error 变量的常见遮蔽问题
func demonstrateErrorVariableShadowing() {
	fmt.Println("=== 4. error 变量的遮蔽问题（常见陷阱）===")
	fmt.Println("说明：在实际代码中，经常会有各种方法返回 error，error 会赋值给 err 变量")

	// 模拟一个返回 error 的函数
	maybeError := func() error {
		return fmt.Errorf("模拟错误")
	}

	// ❌ 错误示例：变量遮蔽导致无法检查错误
	fmt.Println("--- ❌ 错误示例：变量遮蔽 ---")
	var err error
	err = maybeError()
	if err != nil {
		// 这里如果使用 := 会创建新的 err 变量，遮蔽外层的 err
		// err := maybeError()  // 错误：遮蔽了外层的 err
		fmt.Printf("检测到错误: %v\n", err)
	}

	// ✅ 正确示例：使用 = 而不是 :=
	fmt.Println("\n--- ✅ 正确示例：使用 = 赋值 ---")
	err = maybeError()
	if err != nil {
		fmt.Printf("检测到错误: %v\n", err)
	}

	// ✅ 另一种正确方式：使用不同的变量名
	fmt.Println("\n--- ✅ 另一种方式：使用不同变量名 ---")
	if err2 := maybeError(); err2 != nil {
		fmt.Printf("检测到错误: %v\n", err2)
	}

	fmt.Println()
}

// demonstrateGlobalVariableScope 演示全局变量的作用域范围
func demonstrateGlobalVariableScope() {
	fmt.Println("=== 5. 全局变量的作用域范围 ===")

	// 全局变量可以在任何函数中访问和修改
	fmt.Printf("在函数1中访问全局变量 a = %d\n", a)
	a = 10
	fmt.Printf("在函数1中修改全局变量 a = %d\n", a)

	// 调用另一个函数，验证全局变量的共享性
	modifyGlobalVariable()
	fmt.Printf("在函数1中再次访问全局变量 a = %d\n", a)
	fmt.Println("说明：全局变量在整个包内共享")
	fmt.Println()
}

// modifyGlobalVariable 修改全局变量的辅助函数
func modifyGlobalVariable() {
	fmt.Printf("在函数2中访问全局变量 a = %d\n", a)
	a = 20
	fmt.Printf("在函数2中修改全局变量 a = %d\n", a)
}

// GlobalVariableDemo 全局变量完整演示
func GlobalVariableDemo() {
	fmt.Println("========== 1.11.2 全局变量 ==========")
	fmt.Println()
	fmt.Println("全局变量：在函数外声明的变量")
	fmt.Println("作用域：可以是当前整个包甚至外部包（公开的全局变量）使用")
	fmt.Println()
	fmt.Println("变量遮蔽（Variable Shadowing）：")
	fmt.Println("  - 当全局变量和局部变量重名时，函数内会使用局部变量")
	fmt.Println("  - 超出局部变量作用域之后，才会重新使用全局变量")
	fmt.Println("  - 这种优先使用作用域更小的变量的规则，同样适用于局部变量")
	fmt.Println()

	demonstrateGlobalVariable()
	demonstrateVariableShadowing()
	demonstrateLocalVariableShadowing()
	demonstrateErrorVariableShadowing()
	demonstrateGlobalVariableScope()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 全局变量：在函数外声明，作用域是整个包")
	fmt.Println("✅ 公开全局变量：首字母大写，可以被外部包访问")
	fmt.Println("✅ 私有全局变量：首字母小写，只能在当前包内访问")
	fmt.Println("✅ 变量遮蔽：局部变量会遮蔽同名全局变量")
	fmt.Println()
	fmt.Println("⚠️ 常见陷阱：")
	fmt.Println("   - error 变量遮蔽：使用 := 可能创建新的 err 变量")
	fmt.Println("   - 建议：在 if 语句中检查 error 时，使用 = 而不是 :=")
	fmt.Println("   - 或者：使用不同的变量名（如 err2）")
	fmt.Println()
}
