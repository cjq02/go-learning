package operators

import "fmt"

// OperatorsDemo 运算操作符总演示函数
func OperatorsDemo() {
	fmt.Println("========== 1.7 运算操作符 ==========")
	fmt.Println("Go 语言中的运算操作符包括：")
	fmt.Println("- 算术运算符")
	fmt.Println("- 关系运算符（比较运算符）")
	fmt.Println("- 逻辑运算符")
	fmt.Println("- 位运算符")
	fmt.Println("- 赋值运算符")
	fmt.Println("- 其他运算符")
	fmt.Println("- 运算优先级")
	fmt.Println()

	ArithmeticOperatorsDemo()
	fmt.Println()

	RelationalOperatorsDemo()
	fmt.Println()

	LogicalOperatorsDemo()
	fmt.Println()

	BitwiseOperatorsDemo()
	fmt.Println()

	AssignmentOperatorsDemo()
	fmt.Println()

	OtherOperatorsDemo()
	fmt.Println()

	OperatorPrecedenceDemo()
}
