package interfaceexample

import "fmt"

// ========== 1.17.2 接口实现 ==========
//
// 在 Go 中实现一个接口也不需要显式的声明，
// 只需要其他类型实现了接口中所有的方法，就是实现了这个接口。
//
// 这是 Go 语言的"鸭子类型"（Duck Typing）：
// "如果它走起来像鸭子，叫起来像鸭子，那它就是鸭子"

// InterfaceImplementationDemo 演示接口实现
func InterfaceImplementationDemo() {
	fmt.Println("========== 1.17.2 接口实现 ==========")
	fmt.Println()
	fmt.Println("在 Go 中实现一个接口也不需要显式的声明，")
	fmt.Println("只需要其他类型实现了接口中所有的方法，就是实现了这个接口。")
	fmt.Println()
	fmt.Println("这是 Go 语言的\"鸭子类型\"（Duck Typing）：")
	fmt.Println("  \"如果它走起来像鸭子，叫起来像鸭子，那它就是鸭子\"")
	fmt.Println()
	fmt.Println("核心概念：")
	fmt.Println("  - 隐式实现：不需要像 Java 那样写 implements 关键字")
	fmt.Println("  - 自动满足：只要实现了所有方法，就自动满足接口")
	fmt.Println("  - 解耦合：接口定义和使用方解耦，更灵活")
	fmt.Println()

	demonstrateImplicitImplementation()
	demonstratePaymentExample()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ Go 中接口实现是隐式的，不需要显式声明")
	fmt.Println("✅ 只要类型实现了接口中的所有方法，就自动实现了该接口")
	fmt.Println("✅ 这是 Go 语言的鸭子类型特性")
	fmt.Println("✅ 接口可以嵌入其他接口（接口组合）")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 必须实现接口中的所有方法（包括嵌入接口的方法）")
	fmt.Println("   - 方法签名必须完全匹配（方法名、参数类型、返回值类型）")
	fmt.Println("   - 接收者类型要匹配（值接收者 vs 指针接收者）")
	fmt.Println("   - 如果接口方法使用指针接收者，实现也必须使用指针接收者")
	fmt.Println()
}

// PaymentMethod 接口定义了支付方法的基本操作
type PaymentMethod interface {
	Account
	Pay(amount int) bool
}

// Account 账户接口
type Account interface {
	GetBalance() int
}

// CreditCard 信用卡结构体实现 PaymentMethod 接口
type CreditCard struct {
	balance int
	limit   int
}

// Pay 信用卡支付
func (c *CreditCard) Pay(amount int) bool {
	if c.balance+amount <= c.limit {
		c.balance += amount
		fmt.Printf("信用卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("信用卡支付失败: 超出额度")
	return false
}

// GetBalance 获取信用卡余额
func (c *CreditCard) GetBalance() int {
	return c.balance
}

// DebitCard 借记卡结构体实现 PaymentMethod 接口
type DebitCard struct {
	balance int
}

// Pay 借记卡支付
func (d *DebitCard) Pay(amount int) bool {
	if d.balance >= amount {
		d.balance -= amount
		fmt.Printf("借记卡支付成功: %d\n", amount)
		return true
	}
	fmt.Println("借记卡支付失败: 余额不足")
	return false
}

// GetBalance 获取借记卡余额
func (d *DebitCard) GetBalance() int {
	return d.balance
}

// demonstrateImplicitImplementation 演示隐式实现
func demonstrateImplicitImplementation() {
	fmt.Println("=== 1.17.2.1 隐式实现示例 ===")
	fmt.Println()

	fmt.Println("1. 接口定义（包含接口嵌入）：")
	fmt.Println("   type Account interface {")
	fmt.Println("       GetBalance() int")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   type PaymentMethod interface {")
	fmt.Println("       Account              // 嵌入 Account 接口")
	fmt.Println("       Pay(amount int) bool // 支付方法")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - PaymentMethod 接口嵌入了 Account 接口")
	fmt.Println("   - 这意味着实现 PaymentMethod 的类型必须同时实现：")
	fmt.Println("     * Account 接口的方法：GetBalance() int")
	fmt.Println("     * PaymentMethod 自己的方法：Pay(amount int) bool")
	fmt.Println()

	fmt.Println("2. CreditCard 结构体定义：")
	fmt.Println("   type CreditCard struct {")
	fmt.Println("       balance int")
	fmt.Println("       limit   int")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   实现的方法（使用指针接收者）：")
	fmt.Println("   func (c *CreditCard) Pay(amount int) bool {")
	fmt.Println("       // 信用卡支付逻辑：增加余额，不超过额度")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   func (c *CreditCard) GetBalance() int {")
	fmt.Println("       return c.balance")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - CreditCard 实现了 Pay() 和 GetBalance() 方法")
	fmt.Println("   - 使用指针接收者 *CreditCard，可以修改结构体字段")
	fmt.Println("   - 虽然没有显式声明，但自动实现了 PaymentMethod 接口")
	fmt.Println()

	fmt.Println("3. DebitCard 结构体定义：")
	fmt.Println("   type DebitCard struct {")
	fmt.Println("       balance int")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   实现的方法（使用指针接收者）：")
	fmt.Println("   func (d *DebitCard) Pay(amount int) bool {")
	fmt.Println("       // 借记卡支付逻辑：减少余额，不能透支")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   func (d *DebitCard) GetBalance() int {")
	fmt.Println("       return d.balance")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - DebitCard 也实现了 Pay() 和 GetBalance() 方法")
	fmt.Println("   - 同样使用指针接收者，可以修改余额")
	fmt.Println("   - 自动实现了 PaymentMethod 接口")
	fmt.Println()

	fmt.Println("4. 隐式实现的关键点：")
	fmt.Println("   ✅ CreditCard 和 DebitCard 都没有显式声明实现 PaymentMethod")
	fmt.Println("   ✅ 但它们都实现了 PaymentMethod 接口要求的所有方法")
	fmt.Println("   ✅ 包括嵌入接口 Account 的方法")
	fmt.Println("   ✅ 因此它们自动实现了 PaymentMethod 接口")
	fmt.Println("   ✅ 这就是 Go 的隐式接口实现机制")
	fmt.Println()
	fmt.Println("5. 指针接收者 vs 值接收者：")
	fmt.Println("   - 当前示例使用指针接收者 (*CreditCard, *DebitCard)")
	fmt.Println("   - 指针接收者可以修改结构体的字段值")
	fmt.Println("   - 如果接口方法使用指针接收者实现，")
	fmt.Println("     赋值给接口变量时也需要使用指针：")
	fmt.Println("     var pm PaymentMethod = &CreditCard{...}  // ✅ 正确")
	fmt.Println("     var pm PaymentMethod = CreditCard{...}   // ❌ 错误")
	fmt.Println()
}

// purchaseItem 使用 PaymentMethod 接口的函数
func purchaseItem(p PaymentMethod, price int) {
	if p.Pay(price) {
		fmt.Printf("购买成功，剩余余额: %d\n", p.GetBalance())
	} else {
		fmt.Println("购买失败")
	}
}

// demonstratePaymentExample 演示支付示例
func demonstratePaymentExample() {
	fmt.Println("=== 1.17.2.2 支付接口示例 ===")
	fmt.Println()

	fmt.Println("1. purchaseItem 函数定义：")
	fmt.Println("   func purchaseItem(p PaymentMethod, price int) {")
	fmt.Println("       if p.Pay(price) {")
	fmt.Println("           fmt.Printf(\"购买成功，剩余余额: [数字]\\n\", p.GetBalance())")
	fmt.Println("       } else {")
	fmt.Println("           fmt.Println(\"购买失败\")")
	fmt.Println("       }")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 函数参数类型是 PaymentMethod 接口")
	fmt.Println("   - 可以接受任何实现了 PaymentMethod 接口的类型")
	fmt.Println("   - 这是接口的核心优势：多态性")
	fmt.Println()

	fmt.Println("2. 创建支付方式实例：")
	fmt.Println("   creditCard := &CreditCard{balance: 0, limit: 1000}")
	fmt.Println("   debitCard := &DebitCard{balance: 500}")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 使用 & 创建指针，因为方法使用指针接收者")
	fmt.Println("   - creditCard: 初始余额0，额度1000（可以透支）")
	fmt.Println("   - debitCard: 初始余额500（不能透支）")
	fmt.Println()

	fmt.Println("3. 使用接口进行支付：")
	fmt.Println()
	fmt.Println("   使用信用卡购买 800：")
	creditCard := &CreditCard{balance: 0, limit: 1000}
	debitCard := &DebitCard{balance: 500}
	purchaseItem(creditCard, 800)
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 信用卡支付：增加余额（欠款），只要不超过额度即可")
	fmt.Println("   - 800 < 1000（额度），支付成功")
	fmt.Println("   - 余额变为 800（表示欠款800）")
	fmt.Println()

	fmt.Println("   使用借记卡购买 300：")
	purchaseItem(debitCard, 300)
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 借记卡支付：减少余额，不能透支")
	fmt.Println("   - 300 <= 500（余额），支付成功")
	fmt.Println("   - 余额变为 200（500 - 300）")
	fmt.Println()

	fmt.Println("   再次使用借记卡购买 300：")
	purchaseItem(debitCard, 300)
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - 当前余额 200 < 300（支付金额），余额不足")
	fmt.Println("   - 支付失败，余额不变")
	fmt.Println()

	fmt.Println("4. 接口使用的优势：")
	fmt.Println("   ✅ 多态性：同一个函数可以处理不同类型的支付方式")
	fmt.Println("   ✅ 扩展性：添加新的支付方式（如支付宝、微信）")
	fmt.Println("     只需实现 PaymentMethod 接口，无需修改 purchaseItem 函数")
	fmt.Println("   ✅ 解耦合：函数不依赖具体类型，只依赖接口")
	fmt.Println("   ✅ 可测试性：可以轻松创建 Mock 对象进行测试")
	fmt.Println()

	fmt.Println("5. 接口类型断言（补充说明）：")
	fmt.Println("   - 可以通过类型断言获取具体类型：")
	fmt.Println("     if cc, ok := p.(*CreditCard); ok {")
	fmt.Println("         // 使用 cc 的特定方法或字段")
	fmt.Println("     }")
	fmt.Println("   - 或者使用类型开关（type switch）")
	fmt.Println()
}
