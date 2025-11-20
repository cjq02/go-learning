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
		fmt.Println("进过了1秒")
		if true {
			break
		}
		fmt.Println("break 之后") // 这行不会被执行
	}

	// 4. 嵌套循环不使用标记
	fmt.Println("\n--- 4. 嵌套循环不使用标记 ---")
	for i := 1; i <= 3; i++ {
		fmt.Printf("不使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("不使用标记,内部循环 j = %d\n", j)
			break // 仅跳出内层循环
		}
	}

	// 5. 嵌套循环使用标记
	fmt.Println("\n--- 5. 嵌套循环使用标记 ---")
outter:
	for i := 1; i <= 3; i++ {
		fmt.Printf("使用标记,外部循环, i = %d\n", i)
		for j := 5; j <= 10; j++ {
			fmt.Printf("使用标记,内部循环 j = %d\n", j)
			break outter // 跳出外层循环
		}
	}
}

// BreakDemo 导出函数
func BreakDemo() {
	demonstrateBreak()
}
