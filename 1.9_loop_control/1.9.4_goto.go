package loopcontrol

import "fmt"

// ========== 1.9.4.3 goto 语句 ==========
//
// goto 语句可以无条件转移到指定 label 标出的代码处
// 语法：goto 标签名
//
// 注意事项：
// 1. 标签必须在使用 goto 的同一函数内定义
// 2. 标签定义后必须被使用，否则编译错误：label xxx defined and not used
// 3. goto 不能跳过变量声明（不能跳到变量声明之前）
// 4. 一般不推荐使用 goto，会增加代码流程的混乱，不容易理解和调试
//
// 适用场景：
// 1. 错误处理（Go 标准库中常见用法）
// 2. 跳出深层嵌套循环
// 3. 状态机实现
// 4. 资源清理

// demonstrateGoto 演示基本的 goto 用法
func demonstrateGoto() {
	fmt.Println("=== 1. 基本 goto 用法示例 ===")

	gotoPreset := false

preset:
	a := 5
	fmt.Println("【进入 preset 标签】a =", a)

process:
	if a > 0 {
		a--
		fmt.Println("【process 循环】当前a的值为：", a)
		goto process // 跳转到 process 标签，实现循环效果
	} else if a <= 0 {
		// elseProcess:  // 注释掉的标签不会被使用，不会报错
		if !gotoPreset {
			gotoPreset = true
			fmt.Println("【准备跳转】回到 preset 标签")
			goto preset // 跳转到 preset 标签
		} else {
			fmt.Println("【准备跳转】到 post 标签")
			goto post // 跳转到 post 标签
		}
	}

post:
	fmt.Println("【post 标签】main将结束，当前a的值为：", a)
	fmt.Println()
}

// demonstrateGotoErrorHandling 演示 goto 在错误处理中的应用
// 这是 Go 标准库中最常见的 goto 用法模式
func demonstrateGotoErrorHandling() {
	fmt.Println("=== 2. goto 在错误处理中的应用 ===")
	fmt.Println("（模拟文件操作：打开 -> 读取 -> 关闭）")

	// ⚠️ 重要：所有变量必须在 goto 之前声明，否则编译错误
	var fileOpened bool
	var readSuccess bool
	step := 1

	// 步骤1：打开文件
	fmt.Printf("步骤 %d: 尝试打开文件...\n", step)
	fileOpened = true // 模拟成功
	if !fileOpened {
		fmt.Println("❌ 打开文件失败")
		goto cleanup // 如果失败，直接跳到清理步骤
	}
	fmt.Println("✅ 文件打开成功")
	step++

	// 步骤2：读取文件
	fmt.Printf("步骤 %d: 尝试读取文件...\n", step)
	readSuccess = true // 模拟成功
	if !readSuccess {
		fmt.Println("❌ 读取文件失败")
		goto cleanup // 如果失败，直接跳到清理步骤
	}
	fmt.Println("✅ 文件读取成功")
	step++

	// 步骤3：处理数据
	fmt.Printf("步骤 %d: 处理数据...\n", step)
	fmt.Println("✅ 数据处理完成")

cleanup:
	// 无论成功还是失败，都要执行清理工作
	fmt.Println("【清理资源】关闭文件、释放内存等...")
	// 注意：这里可以访问 fileOpened 和 readSuccess，因为它们已经在函数开头声明了
	fmt.Println()
}

// demonstrateGotoBreakNestedLoop 演示使用 goto 跳出深层嵌套循环
// 这是 goto 的另一个常见用途
func demonstrateGotoBreakNestedLoop() {
	fmt.Println("=== 3. 使用 goto 跳出深层嵌套循环 ===")

	// 场景：在三维数组中查找特定值，找到后立即退出所有循环
	target := 7
	found := false

	fmt.Printf("在嵌套循环中查找目标值 %d...\n", target)

	for i := 0; i < 3; i++ {
		fmt.Printf("外层循环 i = %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("  中层循环 j = %d\n", j)
			for k := 0; k < 3; k++ {
				value := i*9 + j*3 + k
				fmt.Printf("    内层循环 k = %d, 值 = %d\n", k, value)

				if value == target {
					fmt.Printf("✅ 找到目标值 %d！位置: (%d, %d, %d)\n", target, i, j, k)
					found = true
					goto found // 直接跳出所有嵌套循环
				}
			}
		}
	}

found:
	if found {
		fmt.Println("查找成功，已跳出所有循环")
	} else {
		fmt.Println("未找到目标值")
	}
	fmt.Println()
}

// demonstrateGotoStateMachine 演示使用 goto 实现简单的状态机
func demonstrateGotoStateMachine() {
	fmt.Println("=== 4. 使用 goto 实现状态机 ===")
	fmt.Println("（模拟一个简单的订单处理流程）")

	orderState := "pending"
	maxSteps := 10
	step := 0

start:
	step++
	if step > maxSteps {
		fmt.Println("⚠️ 达到最大步骤数，退出")
		goto end
	}

	switch orderState {
	case "pending":
		fmt.Printf("步骤 %d: 【待处理】订单创建成功\n", step)
		orderState = "paid"
		goto start

	case "paid":
		fmt.Printf("步骤 %d: 【已支付】处理支付信息\n", step)
		orderState = "shipped"
		goto start

	case "shipped":
		fmt.Printf("步骤 %d: 【已发货】准备配送\n", step)
		orderState = "delivered"
		goto start

	case "delivered":
		fmt.Printf("步骤 %d: 【已送达】订单完成\n", step)
		goto end

	default:
		fmt.Println("未知状态")
		goto end
	}

end:
	fmt.Println("状态机执行完毕")
	fmt.Println()
}

// demonstrateGotoVariableScope 演示 goto 与变量作用域的注意事项
func demonstrateGotoVariableScope() {
	fmt.Println("=== 5. goto 与变量作用域的注意事项 ===")

	// ✅ 正确：先声明变量，再使用 goto
	x := 10
	fmt.Printf("变量 x = %d\n", x)

	if x > 5 {
		goto useX // 可以跳转，因为 x 已经声明
	}

useX:
	fmt.Printf("使用变量 x = %d\n", x)

	// ❌ 错误示例（注释掉，避免编译错误）
	// goto beforeDeclare  // 编译错误：goto beforeDeclare jumps over declaration of y
	// y := 20              // goto 不能跳过变量声明
	// beforeDeclare:
	//     fmt.Println(y)

	fmt.Println("⚠️ 注意：goto 不能跳过变量声明语句")
	fmt.Println()
}

// demonstrateGotoVsOtherMethods 对比 goto 与其他方法的优缺点
func demonstrateGotoVsOtherMethods() {
	fmt.Println("=== 6. goto vs 其他方法对比 ===")

	fmt.Println("场景：跳出嵌套循环")
	fmt.Println()

	// 方法1：使用 goto
	fmt.Println("方法1：使用 goto")
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("  找到目标，使用 goto 跳出")
				goto method1End
			}
		}
	}
method1End:
	fmt.Println("  ✅ goto 优点：简单直接，可以跳出任意层级")
	fmt.Println("  ❌ goto 缺点：代码流程不清晰，难以维护")
	fmt.Println()

	// 方法2：使用函数 + return
	fmt.Println("方法2：使用函数 + return")
	func() {
		for i := 0; i < 3; i++ {
			for j := 0; j < 3; j++ {
				if i == 1 && j == 1 {
					fmt.Println("  找到目标，使用 return 跳出函数")
					return
				}
			}
		}
	}()
	fmt.Println("  ✅ 函数优点：代码结构清晰，易于理解")
	fmt.Println("  ❌ 函数缺点：需要额外的函数调用开销")
	fmt.Println()

	// 方法3：使用标签 + break
	fmt.Println("方法3：使用标签 + break")
outer:
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if i == 1 && j == 1 {
				fmt.Println("  找到目标，使用 break outer 跳出")
				break outer
			}
		}
	}
	fmt.Println("  ✅ break 标签优点：代码清晰，Go 推荐方式")
	fmt.Println("  ✅ break 标签缺点：无（这是 Go 中跳出嵌套循环的最佳实践）")
	fmt.Println()
}

// GotoDemo 导出函数 - goto 语句完整演示
func GotoDemo() {
	fmt.Println("========== 1.9.4.3 goto 语句 ==========")
	fmt.Println()
	fmt.Println("goto 语句可以无条件转移到指定 label 标出的代码处。")
	fmt.Println("一般 goto 语句会配合条件语句使用，实现条件转移、构成循环、跳出循环的功能。")
	fmt.Println()
	fmt.Println("⚠️ 重要提示：")
	fmt.Println("   - 一般不推荐使用 goto 语句")
	fmt.Println("   - goto 语句会增加代码流程的混乱，不容易理解代码和调试程序")
	fmt.Println("   - 但在错误处理和资源清理场景中，goto 是 Go 标准库的常见用法")
	fmt.Println()

	demonstrateGoto()
	demonstrateGotoErrorHandling()
	demonstrateGotoBreakNestedLoop()
	demonstrateGotoStateMachine()
	demonstrateGotoVariableScope()
	demonstrateGotoVsOtherMethods()

	fmt.Println("=== 总结 ===")
	fmt.Println("goto 的适用场景：")
	fmt.Println("  ✅ 错误处理和资源清理（Go 标准库常见模式）")
	fmt.Println("  ✅ 跳出深层嵌套循环（但更推荐使用 break + 标签）")
	fmt.Println("  ✅ 简单的状态机实现")
	fmt.Println()
	fmt.Println("goto 的替代方案：")
	fmt.Println("  ✅ 跳出嵌套循环：使用 break + 标签（推荐）")
	fmt.Println("  ✅ 错误处理：使用 defer + 命名返回值")
	fmt.Println("  ✅ 状态机：使用 switch + 函数封装")
	fmt.Println()
}
