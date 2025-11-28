// Package function 演示 Go 语言方法的使用
package function

import "fmt"

// ========== 1.10.3 方法 ==========

// Person 结构体用于演示方法定义
type Person struct {
	name string
	age  int
}

// Counter 结构体用于演示值接收者和指针接收者
type Counter struct {
	count int
}

// Bank 结构体用于演示方法赋值给函数变量
type Bank struct {
	balance int
}

// demonstrateBasicMethod 演示基本的方法定义和调用
func demonstrateBasicMethod() {
	fmt.Println("=== 1. 基本方法定义和调用 ===")
	fmt.Println("方法 = 函数 + 接收者")
	fmt.Println("语法: func (receiver Type) methodName(...) {...}")
	fmt.Println()

	p := Person{name: "张三", age: 25}

	// 调用方法
	result := p.GetInfo()
	fmt.Printf("调用方法: %s\n", result)

	fmt.Println()
}

// GetInfo 返回 Person 的信息（值接收者）
func (p Person) GetInfo() string {
	return fmt.Sprintf("%s 今年 %d 岁", p.name, p.age)
}

// ModifyAgeByValue 使用值接收者修改年龄（不会改变原值）
func (p Person) ModifyAgeByValue(newAge int) {
	p.age = newAge // 修改的是副本，不会影响原值
}
func demonstrateValueReceiver() {
	fmt.Println("=== 2. 值接收者（传递的是副本）===")
	fmt.Println()

	p := Person{name: "李四", age: 30}
	fmt.Printf("原始 Person: %s\n", p.GetInfo())

	// 尝试通过值接收者方法修改年龄
	// ModifyAgeByValue 是一个值接收者方法
	p.ModifyAgeByValue(50)

	fmt.Printf("调用 ModifyAgeByValue(50) 后: %s\n", p.GetInfo())
	fmt.Println("年龄仍然是 30，没有改变！")
	fmt.Println()
	fmt.Println("原因：值接收者接收的是结构体的副本")
	fmt.Println("      方法内修改的是副本，不会影响原结构体")
	fmt.Println()
}

// SetAge 修改年龄（指针接收者）
func (p *Person) SetAge(newAge int) {
	p.age = newAge
}

// demonstratePointerReceiver 演示指针接收者的特点
func demonstratePointerReceiver() {
	fmt.Println("=== 3. 指针接收者（传递的是指针）===")

	p := Person{name: "王五", age: 28}

	fmt.Printf("修改前: %s\n", p.GetInfo())

	// 指针接收者可以修改原始结构体
	p.SetAge(35)

	fmt.Printf("调用 SetAge(35) 后: %s\n", p.GetInfo())

	fmt.Println("说明: 指针接收者可以修改原结构体的值")

	fmt.Println()
}

// demonstrateMethodVsFunction 演示方法与函数的处区別
func demonstrateMethodVsFunction() {
	fmt.Println("=== 4. 方法与函数的处区別 ===")
	fmt.Println()

	fmt.Println("不同点：")
	fmt.Println()

	fmt.Println("1. 函数：没有接收者")
	fmt.Println("   func add(a int, b int) int {")
	fmt.Println("       return a + b")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("2. 方法：有接收者")
	fmt.Println("   func (p Person) GetInfo() string {")
	fmt.Println("       return p.name")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("相同点：")
	fmt.Println("- 两者都可以有参数（也可以没有）")
	fmt.Println("- 两者都可以有返回值（也可以没有）")
	fmt.Println()

	fmt.Println("方法的特点：不有接收者的不能称为方法")
	fmt.Println()

	// 实际演示
	p := Person{name: "赵六", age: 40}

	fmt.Println("实际例子：")
	fmt.Printf("直接调用 p.GetInfo()：%s\n", p.GetInfo())

	// Person 是值类型，但有值接收者方法
	// Go 会自动转换，你也可以用指针调用
	fmt.Printf("通过指针调用 (&p).GetInfo()：%s\n", (&p).GetInfo())

	fmt.Println()
	fmt.Println("总结：方法 = 函数 + 接收者")
	fmt.Println("方法必须定义在一个类型上，函数是独立的")

	fmt.Println()
}

// Increment 使用值接收者的方法（不会修改原值）
func (c Counter) Increment() int {
	c.count++
	return c.count
}

// IncrementPtr 使用指针接收者的方法（会修改原值）
func (c *Counter) IncrementPtr() int {
	c.count++
	return c.count
}

// demonstrateReceiverDifference 演示值接收者和指针接收者的重要差别
func demonstrateReceiverDifference() {
	fmt.Println("=== 5. 值接收者 vs 指针接收者 ===")

	fmt.Println("值接收者方法:")
	c1 := Counter{count: 0}
	fmt.Printf("初始值: %d\n", c1.count)
	result := c1.Increment() // 方法内修改的是副本
	fmt.Printf("调用 Increment() 返回: %d\n", result)
	fmt.Printf("实例的值: %d（没有改变！）\n", c1.count)

	fmt.Println()
	fmt.Println("指针接收者方法:")
	c2 := Counter{count: 0}
	fmt.Printf("初始值: %d\n", c2.count)
	result = c2.IncrementPtr() // 方法修改的是指向的值
	fmt.Printf("调用 IncrementPtr() 返回: %d\n", result)
	fmt.Printf("实例的值: %d（改变了！）\n", c2.count)

	fmt.Println("总结: 值接收者不能修改，指针接收者可以修改")

	fmt.Println()
}

// Withdraw 从银行账户取钱
func (b *Bank) Withdraw(amount int) int {
	if b.balance >= amount {
		b.balance -= amount
		return b.balance
	}
	fmt.Println("余额不足")
	return b.balance
}

// Deposit 往银行账户存钱
func (b *Bank) Deposit(amount int) int {
	b.balance += amount
	return b.balance
}

// GetBalance 查看银行账户余额
func (b *Bank) GetBalance() int {
	return b.balance
}

// demonstrateMethodAsVariable 演示将方法赋值给函数变量
func demonstrateMethodAsVariable() {
	fmt.Println("=== 6. 将方法赋值给函数变量 ===")

	bank := Bank{balance: 1000}

	fmt.Printf("初始余额: %d\n", bank.GetBalance())

	// 将方法赋值给函数变量
	withdraw := bank.Withdraw

	// 使用函数变量调用方法
	result := withdraw(300)
	fmt.Printf("取出 300 后，余额: %d\n", result)

	// 再次调用函数变量
	result = withdraw(200)
	fmt.Printf("再取出 200 后，余额: %d\n", result)

	fmt.Println("说明: 方法可以赋值给函数变量，调用方式类似闭包")

	fmt.Println()
}

// demonstrateMultipleMethods 演示为同一结构体定义多个方法
func demonstrateMultipleMethods() {
	fmt.Println("=== 7. 为同一结构体定义多个方法 ===")

	bank := Bank{balance: 5000}

	fmt.Printf("初始余额: %d\n", bank.GetBalance())

	// 调用不同的方法
	bank.Deposit(2000)
	fmt.Printf("存入 2000 后: %d\n", bank.GetBalance())

	bank.Withdraw(1500)
	fmt.Printf("取出 1500 后: %d\n", bank.GetBalance())

	bank.Deposit(500)
	fmt.Printf("再存入 500 后: %d\n", bank.GetBalance())

	fmt.Println("说明: 一个结构体可以定义多个方法")

	fmt.Println()
}

// demonstrateMethodReceiver 演示不同的接收者类型
func demonstrateMethodReceiver() {
	fmt.Println("=== 8. 不同的接收者类型 ===")

	fmt.Println("接收者可以是:")
	fmt.Println("1. 值类型: func (p Person) method() {...}")
	fmt.Println("   - 操作的是值的副本")
	fmt.Println("   - 不能修改原值")
	fmt.Println()

	fmt.Println("2. 指针类型: func (p *Person) method() {...}")
	fmt.Println("   - 操作的是值的指针")
	fmt.Println("   - 可以修改原值")
	fmt.Println()

	fmt.Println("3. 不能是接口类型或其他包装类型")
	fmt.Println()

	fmt.Println()
}

// demonstrateExample 演示实际例子
func demonstrateExample() {
	fmt.Println("=== 9. 实际例子：银行账户系统 ===")

	// 创建一个银行账户
	account := Bank{balance: 10000}

	// 将方法赋值给函数变量
	deposit := account.Deposit
	withdraw := account.Withdraw

	fmt.Printf("初始余额: %d\n", account.GetBalance())
	fmt.Println()

	// 通过函数变量调用方法
	fmt.Println("通过函数变量调用:")
	fmt.Printf("存入 5000，余额: %d\n", deposit(5000))
	fmt.Printf("取出 3000，余额: %d\n", withdraw(3000))
	fmt.Printf("取出 15000 (余额不足): %d\n", withdraw(15000))
	fmt.Printf("最终余额: %d\n", account.GetBalance())

	fmt.Println()
}

// demonstrateMethodChaining 演示方法的返回值
func demonstrateMethodChaining() {
	fmt.Println("=== 10. 方法的返回值 ===")

	bank := Bank{balance: 1000}

	// 方法可以返回单个值或多个值
	fmt.Printf("GetBalance() 返回: %d\n", bank.GetBalance())

	result := bank.Deposit(500)
	fmt.Printf("Deposit(500) 返回: %d\n", result)

	result = bank.Withdraw(200)
	fmt.Printf("Withdraw(200) 返回: %d\n", result)

	fmt.Println()
	fmt.Println("说明: 方法的返回值设计可以支持链式调用或获取操作结果")

	fmt.Println()
}

// MethodDemo 方法完整演示主函数
func MethodDemo() {
	fmt.Println("========== 1.10.3 方法 ==========")
	fmt.Println()
	fmt.Println("方法是一个包含接收者的函数。")
	fmt.Println("大部分情况下可以通过类型的实例调用。")
	fmt.Println()
	fmt.Println("也可以把方法赋值给一个函数变量，")
	fmt.Println("使用函数变量调用这个方法，调用方式类似闭包。")
	fmt.Println()
	fmt.Println("关键概念:")
	fmt.Println("- 值接收者：传递的是结构体的副本，不能修改原值")
	fmt.Println("- 指针接收者：传递的是结构体的指针，可以修改原值")
	fmt.Println()

	demonstrateBasicMethod()
	demonstrateValueReceiver()
	demonstratePointerReceiver()
	demonstrateMethodVsFunction()
	demonstrateReceiverDifference()
	demonstrateMethodAsVariable()
	demonstrateMultipleMethods()
	demonstrateMethodReceiver()
	demonstrateExample()
	demonstrateMethodChaining()
}
