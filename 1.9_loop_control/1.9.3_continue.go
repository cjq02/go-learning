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
	for i := 1; i <= 2; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			if j >= 7 {
				continue
			}
			fmt.Println("不使用标记，内部循环，在continue之后执行")
		}
	}

	// 3. 嵌套循环使用标记
	fmt.Println("\n--- 3. 嵌套循环使用标记 ---")
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			if j >= 7 {
				continue outter
			}
			fmt.Println("不使用标记，内部循环，在continue之后执行")
		}
	}
}

// ContinueDemo 导出函数
func ContinueDemo() {
	demonstrateContinue()
}
