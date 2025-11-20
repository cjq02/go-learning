package loopcontrol

import "fmt"

// ========== 1.9.2.2 continue 语句 ==========

func demonstrateContinue() {
	fmt.Println("=== continue 语句示例 ===")

	// 1. 基本使用
	fmt.Println("--- 1. 基本使用 ---")
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("第", i, "次循环")
	}

	// 2. 嵌套循环不使用标记
	fmt.Println("\n--- 2. 嵌套循环不使用标记 ---")
	fmt.Println("说明：continue 默认只跳过最内层循环的当前迭代")
	for i := 1; i <= 2; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			if j >= 7 {
				continue // ⚠️ 只跳过内层循环的当前迭代，继续内层循环的下一次
			}
			fmt.Println("不使用标记，内部循环，在continue之后执行")
		}
	}

	// 3. 嵌套循环使用标记（Label）
	fmt.Println("\n--- 3. 嵌套循环使用标记（Label）---")
	fmt.Println("说明：outter: 是一个标签（Label），标记外层循环的位置")
	fmt.Println("      continue outter 会跳过内层循环，直接开始外层循环的下一次迭代")
	// 标签语法：标签名 + 冒号，必须放在语句之前
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			if j >= 7 {
				continue outter // ✅ 跳过内层循环，直接开始外层循环的下一次迭代
				// 结果：当 j >= 7 时，内层循环被跳过，外层循环进入下一次（i++）
			}
			fmt.Println("不使用标记，内部循环，在continue之后执行")
		}
	}
}

// ContinueDemo 导出函数
func ContinueDemo() {
	demonstrateContinue()
}
