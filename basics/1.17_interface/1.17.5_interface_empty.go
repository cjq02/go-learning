package interfaceexample

import "fmt"

// ========== 1.17.5 空接口 interface{} ==========
//
// 如果函数参数使用 interface{} 可以接受任何类型的实参。
// 同样，可以接收任何类型的值也可以赋值给 interface{} 类型的变量。
//
// interface{} 是空接口，不包含任何方法。
// 由于所有类型都实现了空接口（因为没有方法需要实现），
// 所以 interface{} 可以表示任何类型。

// InterfaceEmptyDemo 演示空接口
func InterfaceEmptyDemo() {
	fmt.Println("========== 1.17.5 空接口 interface{} ==========")
	fmt.Println()
	fmt.Println("如果函数参数使用 interface{} 可以接受任何类型的实参。")
	fmt.Println("同样，可以接收任何类型的值也可以赋值给 interface{} 类型的变量。")
	fmt.Println()
	fmt.Println("interface{} 是空接口，不包含任何方法。")
	fmt.Println("由于所有类型都实现了空接口（因为没有方法需要实现），")
	fmt.Println("所以 interface{} 可以表示任何类型。")
	fmt.Println()

	demonstrateEmptyInterface()
	demonstrateEmptyInterfaceUsage()
	demonstrateTypeAssertion()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ interface{} 是空接口，可以表示任何类型")
	fmt.Println("✅ 任何类型的值都可以赋值给 interface{} 类型的变量")
	fmt.Println("✅ 函数参数使用 interface{} 可以接受任何类型的实参")
	fmt.Println("✅ 使用类型断言可以获取 interface{} 中的具体类型和值")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - interface{} 会丢失类型信息，使用时需要类型断言")
	fmt.Println("   - 类型断言失败会返回零值和 false")
	fmt.Println("   - 可以使用 switch v := i.(type) 进行类型判断")
	fmt.Println()
}

// PayMethodSimple 简单支付接口
type PayMethodSimple interface {
	Pay(int)
}

// CreditCardSimple 信用卡结构体
type CreditCardSimple struct {
	balance int
	limit   int
}

// Pay 支付方法
func (c *CreditCardSimple) Pay(amount int) {
	if c.balance < amount {
		fmt.Println("余额不足")
		return
	}
	c.balance -= amount
	fmt.Printf("支付成功: %d\n", amount)
}

// anyParam 接受任何类型参数的函数
func anyParam(param interface{}) {
	fmt.Printf("anyParam 接收到的参数: %v, 类型: %T\n", param, param)
}

// demonstrateEmptyInterface 演示空接口基本使用
func demonstrateEmptyInterface() {
	fmt.Println("=== 1.17.5.1 空接口基本使用 ===")

	// 任何类型都可以赋值给 interface{}
	var i interface{} = 42
	var s interface{} = "hello"
	var b interface{} = true
	var f interface{} = 3.14

	fmt.Printf("i = %v, 类型: %T\n", i, i)
	fmt.Printf("s = %v, 类型: %T\n", s, s)
	fmt.Printf("b = %v, 类型: %T\n", b, b)
	fmt.Printf("f = %v, 类型: %T\n", f, f)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 任何类型的值都可以赋值给 interface{} 类型的变量")
	fmt.Println("  - interface{} 可以表示任何类型")
	fmt.Println()
}

// demonstrateEmptyInterfaceUsage 演示空接口作为函数参数
func demonstrateEmptyInterfaceUsage() {
	fmt.Println("=== 1.17.5.2 空接口作为函数参数 ===")

	c := CreditCardSimple{balance: 100, limit: 1000}
	var a PayMethodSimple = &c

	// 可以传入任何类型
	anyParam(c)
	anyParam(1)
	anyParam("123")
	anyParam(a)
	anyParam(true)
	anyParam(3.14)
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - anyParam 函数接受 interface{} 类型参数")
	fmt.Println("  - 可以传入任何类型的值")
	fmt.Println("  - 函数内部可以通过类型断言获取具体类型")
	fmt.Println()
}

// demonstrateTypeAssertion 演示类型断言
func demonstrateTypeAssertion() {
	fmt.Println("=== 1.17.5.3 类型断言示例 ===")

	var i interface{} = 42

	// 类型断言：获取具体类型和值
	if v, ok := i.(int); ok {
		fmt.Printf("i 是 int 类型，值: %d\n", v)
	} else {
		fmt.Println("i 不是 int 类型")
	}
	fmt.Println()

	// 使用 switch 进行类型判断
	var values []interface{} = []interface{}{42, "hello", true, 3.14}

	fmt.Println("使用 switch 进行类型判断：")
	for _, val := range values {
		switch v := val.(type) {
		case int:
			fmt.Printf("  %v 是 int 类型，值: %d\n", val, v)
		case string:
			fmt.Printf("  %v 是 string 类型，值: %s\n", val, v)
		case bool:
			fmt.Printf("  %v 是 bool 类型，值: %v\n", val, v)
		case float64:
			fmt.Printf("  %v 是 float64 类型，值: %f\n", val, v)
		default:
			fmt.Printf("  %v 是未知类型: %T\n", val, val)
		}
	}
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - 使用类型断言可以获取 interface{} 中的具体类型和值")
	fmt.Println("  - 类型断言格式：value, ok := i.(Type)")
	fmt.Println("  - 使用 switch v := i.(type) 可以进行多类型判断")
	fmt.Println()
}

