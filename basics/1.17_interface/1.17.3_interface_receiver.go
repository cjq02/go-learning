package interfaceexample

import "fmt"

// ========== 1.17.3 接口与接收者 ==========
//
// 直接用接口类型作为变量时：
// - 如果方法都是值接收者，可以用值或指针
// - 如果任何方法是指针接收者，则必须用指针

// InterfaceReceiverDemo 演示接口与接收者
func InterfaceReceiverDemo() {
	fmt.Println("========== 1.17.3 接口与接收者 ==========")
	fmt.Println()
	fmt.Println("在 Go 中，将类型赋值给接口变量时，接收者类型很重要：")
	fmt.Println()
	fmt.Println("核心规则：")
	fmt.Println("  1. 如果方法都是值接收者：")
	fmt.Println("     - 可以用值赋值：var i Interface = value")
	fmt.Println("     - 也可以用指针赋值：var i Interface = &value")
	fmt.Println("     - Go 会自动处理两种情况")
	fmt.Println()
	fmt.Println("  2. 如果任何方法是指针接收者：")
	fmt.Println("     - 必须用指针赋值：var i Interface = &value")
	fmt.Println("     - 不能用值赋值（会编译错误）")
	fmt.Println("     - 这是 Go 语言的严格规则")
	fmt.Println()
	fmt.Println("原因：")
	fmt.Println("  - 接口变量必须能够调用接口中的所有方法")
	fmt.Println("  - 指针接收者方法需要指针类型，值无法满足")
	fmt.Println("  - 值接收者方法可以接受值或指针（Go 自动解引用）")
	fmt.Println()

	demonstrateValueReceiver()
	demonstratePointerReceiver()
	demonstrateMixedReceiver()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 值接收者：可以用值或指针赋值给接口")
	fmt.Println("✅ 指针接收者：必须用指针赋值给接口")
	fmt.Println("✅ 如果接口中有任何方法是指针接收者，则必须用指针")
	fmt.Println("✅ 这是 Go 语言类型安全的要求，确保接口变量可以调用所有方法")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 值接收者方法操作的是副本，修改不影响原值")
	fmt.Println("   - 指针接收者方法操作的是原值，修改会直接影响原值")
	fmt.Println("   - 选择值接收者还是指针接收者要根据需求决定")
	fmt.Println("   - 一般规则：需要修改原值用指针接收者，只读操作用值接收者")
	fmt.Println()
}

// AccountValue 账户接口（值接收者示例）
type AccountValue interface {
	getBalance() int
}

// CreditCardValue 信用卡结构体（值接收者）
type CreditCardValue struct {
	balance int
	limit   int
}

// getBalance 值接收者方法
func (c CreditCardValue) getBalance() int {
	return c.balance
}

// demonstrateValueReceiver 演示值接收者
func demonstrateValueReceiver() {
	fmt.Println("=== 1.17.3.1 值接收者示例 ===")
	fmt.Println()

	fmt.Println("1. 结构体定义和方法实现：")
	fmt.Println("   type CreditCardValue struct {")
	fmt.Println("       balance int")
	fmt.Println("       limit   int")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   // 值接收者方法")
	fmt.Println("   func (c CreditCardValue) getBalance() int {")
	fmt.Println("       return c.balance")
	fmt.Println("   }")
	fmt.Println()

	c := CreditCardValue{balance: 100, limit: 1000}

	fmt.Println("2. 赋值给接口变量（两种方式都可以）：")
	fmt.Println("   c := CreditCardValue{balance: 100, limit: 1000}")
	fmt.Println()
	fmt.Println("   ✅ 方式1：使用值赋值")
	fmt.Println("   var a1 AccountValue = c")
	fmt.Println()
	fmt.Println("   ✅ 方式2：使用指针赋值")
	fmt.Println("   var a2 AccountValue = &c")
	fmt.Println()

	// 值接收者可以用值或指针
	var a1 AccountValue = c
	var a2 AccountValue = &c

	fmt.Printf("3. 实际运行结果：\n")
	fmt.Printf("   c = %+v\n", c)
	fmt.Printf("   a1.getBalance() = %d (使用值赋值)\n", a1.getBalance())
	fmt.Printf("   a2.getBalance() = %d (使用指针赋值)\n", a2.getBalance())
	fmt.Println()

	fmt.Println("4. 说明：")
	fmt.Println("   ✅ 值接收者方法可以用值或指针赋值给接口")
	fmt.Println("   ✅ 两种方式都可以正常工作")
	fmt.Println("   ✅ Go 会自动处理指针到值的转换")
	fmt.Println("   ✅ 这是因为值接收者方法可以接受值或指针作为接收者")
	fmt.Println()
}

// AccountPointer 账户接口（指针接收者示例）
type AccountPointer interface {
	getBalance() int
	setBalance(int)
}

// CreditCardPointer 信用卡结构体（指针接收者）
type CreditCardPointer struct {
	balance int
	limit   int
}

// getBalance 指针接收者方法
func (c *CreditCardPointer) getBalance() int {
	return c.balance
}

// setBalance 指针接收者方法
func (c *CreditCardPointer) setBalance(balance int) {
	c.balance = balance
}

// demonstratePointerReceiver 演示指针接收者
func demonstratePointerReceiver() {
	fmt.Println("=== 1.17.3.2 指针接收者示例 ===")
	fmt.Println()

	fmt.Println("1. 结构体定义和方法实现：")
	fmt.Println("   type CreditCardPointer struct {")
	fmt.Println("       balance int")
	fmt.Println("       limit   int")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   // 指针接收者方法")
	fmt.Println("   func (c *CreditCardPointer) getBalance() int {")
	fmt.Println("       return c.balance")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   func (c *CreditCardPointer) setBalance(balance int) {")
	fmt.Println("       c.balance = balance")
	fmt.Println("   }")
	fmt.Println()

	c := CreditCardPointer{balance: 100, limit: 1000}

	fmt.Println("2. 赋值给接口变量：")
	fmt.Println("   c := CreditCardPointer{balance: 100, limit: 1000}")
	fmt.Println()
	fmt.Println("   ❌ 错误示例（编译错误）：")
	fmt.Println("   var a AccountPointer = c")
	fmt.Println("   // 编译错误：CreditCardPointer does not implement AccountPointer")
	fmt.Println("   //            (getBalance method has pointer receiver)")
	fmt.Println()
	fmt.Println("   ✅ 正确示例：")
	fmt.Println("   var a AccountPointer = &c")
	fmt.Println()

	// 指针接收者必须用指针
	// var a AccountPointer = c  // 这样会编译错误
	var a AccountPointer = &c

	fmt.Printf("3. 实际运行结果：\n")
	fmt.Printf("   c = %+v\n", c)
	fmt.Printf("   a.getBalance() = %d\n", a.getBalance())

	a.setBalance(200)
	fmt.Printf("   调用 setBalance(200) 后，c.balance = %d\n", c.balance)
	fmt.Println()

	fmt.Println("4. 说明：")
	fmt.Println("   ✅ 指针接收者方法必须用指针赋值给接口")
	fmt.Println("   ✅ 使用值会编译错误，因为值无法满足指针接收者的要求")
	fmt.Println("   ✅ 指针接收者可以修改原值（如 setBalance 方法）")
	fmt.Println("   ✅ 这是 Go 语言类型安全的要求")
	fmt.Println()
	fmt.Println("5. 为什么必须用指针：")
	fmt.Println("   - 指针接收者方法需要 *CreditCardPointer 类型")
	fmt.Println("   - 如果使用值 c，Go 无法自动获取指针来调用方法")
	fmt.Println("   - 必须显式使用 &c 来获取指针")
	fmt.Println()
}

// AccountMixed 账户接口（混合接收者示例）
type AccountMixed interface {
	getBalance() int
	setBalance(int)
}

// CreditCardMixed 信用卡结构体（混合接收者）
type CreditCardMixed struct {
	balance int
	limit   int
}

// getBalance 值接收者方法
func (c CreditCardMixed) getBalance() int {
	return c.balance
}

// setBalance 指针接收者方法
func (c *CreditCardMixed) setBalance(balance int) {
	c.balance = balance
}

// demonstrateMixedReceiver 演示混合接收者
func demonstrateMixedReceiver() {
	fmt.Println("=== 1.17.3.3 混合接收者示例 ===")
	fmt.Println()

	fmt.Println("1. 结构体定义和方法实现：")
	fmt.Println("   type CreditCardMixed struct {")
	fmt.Println("       balance int")
	fmt.Println("       limit   int")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   // 值接收者方法")
	fmt.Println("   func (c CreditCardMixed) getBalance() int {")
	fmt.Println("       return c.balance")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   // 指针接收者方法")
	fmt.Println("   func (c *CreditCardMixed) setBalance(balance int) {")
	fmt.Println("       c.balance = balance")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("2. 接口定义：")
	fmt.Println("   type AccountMixed interface {")
	fmt.Println("       getBalance() int    // 值接收者实现")
	fmt.Println("       setBalance(int)     // 指针接收者实现")
	fmt.Println("   }")
	fmt.Println()

	c := CreditCardMixed{balance: 100, limit: 1000}

	fmt.Println("3. 赋值给接口变量：")
	fmt.Println("   c := CreditCardMixed{balance: 100, limit: 1000}")
	fmt.Println()
	fmt.Println("   ❌ 错误示例（编译错误）：")
	fmt.Println("   var a AccountMixed = c")
	fmt.Println("   // 编译错误：CreditCardMixed does not implement AccountMixed")
	fmt.Println("   //            (setBalance method has pointer receiver)")
	fmt.Println()
	fmt.Println("   ✅ 正确示例：")
	fmt.Println("   var a AccountMixed = &c")
	fmt.Println()

	var a AccountMixed = &c

	fmt.Printf("4. 实际运行结果：\n")
	fmt.Printf("   c = %+v\n", c)
	fmt.Printf("   a.getBalance() = %d\n", a.getBalance())

	a.setBalance(200)
	fmt.Printf("   调用 setBalance(200) 后，c.balance = %d\n", c.balance)
	fmt.Println()

	fmt.Println("5. 核心规则详解：")
	fmt.Println("   ✅ 规则：如果接口中有任何方法是指针接收者实现，")
	fmt.Println("           则赋值给接口变量时必须使用指针")
	fmt.Println()
	fmt.Println("   原因分析：")
	fmt.Println("   - 接口变量需要能够调用接口中的所有方法")
	fmt.Println("   - 如果 setBalance 是指针接收者，它需要一个 *CreditCardMixed")
	fmt.Println("   - 如果使用值 c 赋值，Go 无法自动获取指针来调用 setBalance")
	fmt.Println("   - 因此必须使用 &c（指针）来赋值")
	fmt.Println()
	fmt.Println("   对比：")
	fmt.Println("   - 如果所有方法都是值接收者：可以用值或指针（Go 会自动处理）")
	fmt.Println("   - 如果有任何方法是指针接收者：必须用指针")
	fmt.Println()
	fmt.Println("6. 实际应用场景：")
	fmt.Println("   - 当接口中既有只读方法（值接收者）又有修改方法（指针接收者）时")
	fmt.Println("   - 必须使用指针赋值，以确保可以调用所有方法")
	fmt.Println("   - 这是 Go 语言类型安全的要求")
	fmt.Println()
}
