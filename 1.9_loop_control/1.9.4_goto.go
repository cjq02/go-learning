package loopcontrol

import "fmt"

// ========== 1.9.4.3 goto 语句 ==========

func demonstrateGoto() {
	fmt.Println("=== goto 语句示例 ===")

	gotoPreset := false

preset:
	a := 5
	fmt.Println("进入 preset 标签")

process:
	if a > 0 {
		a--
		fmt.Println("当前a的值为：", a)
		goto process
	} else if a <= 0 {
		// elseProcess:
		if !gotoPreset {
			gotoPreset = true
			fmt.Println("准备跳转回 preset")
			goto preset
		} else {
			fmt.Println("准备跳转到 post")
			goto post
		}
	}

post:
	fmt.Println("main将结束，当前a的值为：", a)
}

// GotoDemo 导出函数
func GotoDemo() {
	demonstrateGoto()
}
