package variablescope

import (
	"fmt"
	"time"
)

// ========== 1.11.1 局部变量 ==========
//
// 局部变量：在函数内声明的变量，作用域范围只在函数体内
// 函数的参数和返回值也是局部变量
//
// 特殊作用域：if、for、switch、select、匿名代码块中声明的变量
// 作用域范围更小，仅在小的代码块内有效

// demonstrateLocalVariableInFunction 演示函数内的局部变量
func demonstrateLocalVariableInFunction() {
	fmt.Println("=== 1. 函数内的局部变量 ===")

	// 示例函数：展示 parameter、res、decVar 都是局部变量
	localVariable := func(parameter int) (res int) {
		decVar := parameter + 10
		res = decVar
		return
	}

	result := localVariable(5)
	fmt.Printf("调用 localVariable(5) = %d\n", result)
	fmt.Println("说明：parameter、res、decVar 都是局部变量，仅在函数内有效")
	fmt.Println()
}

// demonstrateLocalVariableInControlFlow 演示控制流语句中的局部变量作用域
func demonstrateLocalVariableInControlFlow() {
	fmt.Println("=== 2. 控制流语句中的局部变量作用域 ===")

	var a int
	fmt.Printf("函数级变量 a = %d\n", a)

	// if 语句中的变量作用域
	fmt.Println("\n--- if 语句中的变量作用域 ---")
	if b := 1; b == 0 {
		fmt.Println("b == 0")
	} else {
		c := 2
		fmt.Println("declare c =", c)
		fmt.Println("b == 1")
	}
	// ⚠️ 注意：b 和 c 在 if 语句外部不可访问
	// fmt.Println(b)  // 编译错误：undefined: b
	// fmt.Println(c)  // 编译错误：undefined: c
	fmt.Println("说明：b 和 c 只在 if-else 代码块内有效")

	// switch 语句中的变量作用域
	fmt.Println("\n--- switch 语句中的变量作用域 ---")
	switch d := 3; d {
	case 1:
		e := 4
		fmt.Println("declare e =", e)
		fmt.Println("d == 1")
	case 3:
		f := 4
		fmt.Println("declare f =", f)
		fmt.Println("d == 3")
	}
	// ⚠️ 注意：d、e、f 在 switch 语句外部不可访问
	// fmt.Println(d)  // 编译错误：undefined: d
	// fmt.Println(e)  // 编译错误：undefined: e
	// fmt.Println(f)  // 编译错误：undefined: f
	fmt.Println("说明：d、e、f 只在 switch 代码块内有效")

	// for 语句中的变量作用域
	fmt.Println("\n--- for 语句中的变量作用域 ---")
	for i := 0; i < 1; i++ {
		forA := 1
		fmt.Println("forA =", forA)
	}
	// ⚠️ 注意：i 和 forA 在 for 循环外部不可访问
	// fmt.Println("i =", i)      // 编译错误：undefined: i
	// fmt.Println("forA =", forA) // 编译错误：undefined: forA
	fmt.Println("说明：i 和 forA 只在 for 循环代码块内有效")

	// select 语句中的变量作用域
	fmt.Println("\n--- select 语句中的变量作用域 ---")
	select {
	case <-time.After(time.Second):
		selectA := 1
		fmt.Println("selectA =", selectA)
	}
	// ⚠️ 注意：selectA 在 select 语句外部不可访问
	// fmt.Println("selectA =", selectA) // 编译错误：undefined: selectA
	fmt.Println("说明：selectA 只在 select case 代码块内有效")

	// 匿名代码块中的变量作用域
	fmt.Println("\n--- 匿名代码块中的变量作用域 ---")
	{
		blockA := 1
		fmt.Println("blockA =", blockA)
	}
	// ⚠️ 注意：blockA 在匿名代码块外部不可访问
	// fmt.Println("blockA =", blockA) // 编译错误：undefined: blockA
	fmt.Println("说明：blockA 只在匿名代码块内有效")

	fmt.Printf("\n函数级变量 a = %d (在整个函数内都有效)\n", a)
	fmt.Println()
}

// demonstrateVariableScopeComparison 对比不同作用域的变量
func demonstrateVariableScopeComparison() {
	fmt.Println("=== 3. 变量作用域对比 ===")

	// 函数级变量：在整个函数内都有效
	funcLevelVar := "函数级变量"
	fmt.Printf("1. %s：在整个函数内都有效\n", funcLevelVar)

	// 代码块级变量：只在代码块内有效
	{
		blockLevelVar := "代码块级变量"
		fmt.Printf("2. %s：只在当前代码块内有效\n", blockLevelVar)
	}
	// blockLevelVar 在这里不可访问

	// if 语句中的变量
	if ifVar := "if语句变量"; len(ifVar) > 0 {
		fmt.Printf("3. %s：只在 if 代码块内有效\n", ifVar)
	}
	// ifVar 在这里不可访问

	// for 循环中的变量
	for forVar := 0; forVar < 1; forVar++ {
		fmt.Printf("4. for循环变量：只在 for 代码块内有效\n")
	}
	// forVar 在这里不可访问

	fmt.Println()
}

// LocalVariableDemo 局部变量完整演示
func LocalVariableDemo() {
	fmt.Println("========== 1.11.1 局部变量 ==========")
	fmt.Println()
	fmt.Println("局部变量：在函数内声明的变量，作用域范围只在函数体内。")
	fmt.Println("函数的参数和返回值也是局部变量。")
	fmt.Println()
	fmt.Println("特殊作用域：")
	fmt.Println("  - if、for、switch、select、匿名代码块中声明的变量")
	fmt.Println("  - 作用域范围更小，仅在小的代码块内有效")
	fmt.Println()

	demonstrateLocalVariableInFunction()
	demonstrateLocalVariableInControlFlow()
	demonstrateVariableScopeComparison()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 函数内声明的变量：作用域在整个函数内")
	fmt.Println("✅ 函数参数和返回值：作用域在整个函数内")
	fmt.Println("✅ 控制流语句中的变量：作用域仅在对应的代码块内")
	fmt.Println("✅ 匿名代码块中的变量：作用域仅在代码块内")
	fmt.Println()
	fmt.Println("⚠️ 重要提示：")
	fmt.Println("   - 变量必须先声明后使用")
	fmt.Println("   - 变量不能跨作用域访问")
	fmt.Println("   - 内层作用域可以访问外层作用域的变量（变量遮蔽）")
	fmt.Println()
}
