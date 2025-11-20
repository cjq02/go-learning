package loopcontrol

import (
	"fmt"
	"time"
)

// ========== 1.9.2.1 break 语句 ==========

func demonstrateBreak() {
	fmt.Println("=== break 语句示例 ===")

	// 1. 中断 for 循环
	fmt.Println("--- 1. 中断 for 循环 ---")
	for i := 0; i < 5; i++ {
		if i == 3 {
			break
		}
		fmt.Println("第", i, "次循环")
	}

	// 2. 中断 switch
	fmt.Println("\n--- 2. 中断 switch ---")
	switch i := 1; i {
	case 1:
		fmt.Println("进入case 1")
		if i == 1 {
			break
		}
		fmt.Println("i等于1") // 这行不会被执行
	case 2:
		fmt.Println("i等于2")
	default:
		fmt.Println("default case")
	}

	// 3. 中断 select
	fmt.Println("\n--- 3. 中断 select ---")
	// 为了演示 select，我们需要一个通道或者使用 time.After
	select {
	case <-time.After(time.Second * 2):
		fmt.Println("过了2秒")
	case <-time.After(time.Second):
		fmt.Println("经过了1秒")
		if true {
			break
		}
		fmt.Println("break 之后") // 这行不会被执行
	}

	// 4. 嵌套循环不使用标记
	fmt.Println("\n--- 4. 嵌套循环不使用标记 ---")
	fmt.Println("说明：break 默认只跳出最内层循环")
	for i := 1; i <= 3; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			break // ⚠️ 仅跳出内层循环，外层循环继续执行
			// 结果：外部循环会执行 3 次，每次内部循环只执行一次就 break
		}
	}

	// 5. 嵌套循环使用标记（Label）
	fmt.Println("\n--- 5. 嵌套循环使用标记（Label）---")
	fmt.Println("说明：outter: 是一个标签（Label），用于标记外层循环的位置")
	fmt.Println("      break outter 可以跳出到标签指定的循环层级")
	// 标签语法：标签名 + 冒号，必须放在语句之前
	// 注意：这里用的是 outter（少一个 t），更常见的拼写是 outer，但都可以用
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			break outter // ✅ 跳出到 outter 标签标记的循环（即外层循环）
			// 结果：外部循环只执行 1 次，内部循环执行一次后，整个嵌套循环都结束了
		}
	}
}

// BreakDemo 导出函数
func BreakDemo() {
	demonstrateBreak()
}
