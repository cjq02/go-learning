// Package pointer 演示 Go 語言指针的使用
package pointer

import (
	"fmt"
	"unsafe"
)

// Person 结构体定义
type Person struct {
	Name string
	Age  int
}

// PointersDemo 指针完整示例演示
func PointersDemo() {
	// 调用基础指针示例
	basicPointerDemo()

	// 调用方式一：直接创建结构体实例并取地址
	structPointerMethod1()

	// 调用方式二：对已存在的变量取地址
	structPointerMethod2()

	// 调用普通变量取地址示例
	variablePointerDemo()

	// 调用指针的实际作用示例
	pointerRealWorldDemo()

	// 调用用户问题示例
	userQuestionDemo()

	// 调用二级指针示例
	doublePointerDemo()

	// 调用指针类型对比示例
	pointerTypesDemo()
}

// basicPointerDemo 演示基础指针操作
func basicPointerDemo() {
	fmt.Println("=== 基础指针操作 ===")
	// 1. 声明一个普通变量
	num := 42

	// 2. 声明一个指向 int 的指针变量
	var ptr *int

	// 3. 将 num 的地址赋值给指针（使用 & 取地址符）
	ptr = &num

	// 4. 通过指针访问值（使用 * 解引用）
	fmt.Println("num 的值:", num)    // 输出: 42
	fmt.Println("num 的地址:", &num)  // 输出: 0xc0000140a8 (内存地址)
	fmt.Println("ptr 存储的地址:", ptr) // 输出: 0xc0000140a8
	fmt.Println("ptr 指向的值:", *ptr) // 输出: 42

	// 5. 通过指针修改值
	*ptr = 100
	fmt.Println("修改后 num 的值:", num) // 输出: 100
	fmt.Println()
}

// structPointerMethod1 演示方式一：直接创建结构体实例并取地址
// p := &<struct type>{}
func structPointerMethod1() {
	fmt.Println("=== 方式一：直接创建结构体实例并取地址 ===")
	// 直接创建结构体实例并立即取地址，一步完成

	// 创建结构体实例并立即取地址
	p1 := &Person{
		Name: "张三",
		Age:  25,
	}
	fmt.Println("p1 的类型:", fmt.Sprintf("%T", p1)) // 输出: *main.Person
	fmt.Println("p1 指向的值:", *p1)                  // 输出: {张三 25}
	fmt.Println("p1.Name:", p1.Name)              // 输出: 张三 (可以直接访问，Go自动解引用)

	// 也可以使用空结构体
	p2 := &Person{}
	p2.Name = "李四"
	p2.Age = 30
	fmt.Println("p2 指向的值:", *p2) // 输出: {李四 30}
	fmt.Println()
}

// structPointerMethod2 演示方式二：对已存在的变量取地址
// p := &<var name>
func structPointerMethod2() {
	fmt.Println("=== 方式二：对已存在的变量取地址 ===")
	// 先声明变量，然后取地址
	var person Person
	person.Name = "王五"
	person.Age = 28

	// 对已存在的变量取地址
	p3 := &person
	fmt.Println("p3 的类型:", fmt.Sprintf("%T", p3)) // 输出: *main.Person
	fmt.Println("p3 指向的值:", *p3)                  // 输出: {王五 28}

	// 通过指针修改会影响原变量
	p3.Age = 35
	fmt.Println("通过 p3 修改后，person 的值:", person) // 输出: {王五 35}
	fmt.Println()
}

// variablePointerDemo 演示普通变量取地址
func variablePointerDemo() {
	fmt.Println("=== 普通变量取地址示例 ===")
	// 对于普通变量也可以这样
	value := 100
	p4 := &value
	fmt.Println("value 的值:", value) // 输出: 100
	fmt.Println("p4 指向的值:", *p4)    // 输出: 100
	*p4 = 200
	fmt.Println("修改后 value 的值:", value) // 输出: 200
	fmt.Println()
}

// ========== 指针的真正作用演示 ==========

// 场景1: 函数参数传递 - 值传递（无法修改原变量）
func modifyValue(x int) {
	x = 999 // 只修改了函数内部的副本
	fmt.Println("函数内部 x =", x)
}

// 场景2: 函数参数传递 - 指针传递（可以修改原变量）
func modifyValueByPointer(x *int) {
	*x = 999 // 修改了原变量的值
	fmt.Println("函数内部 *x =", *x)
}

func swapWithPointer(a, b *int) {
	// 指针传递，真正交换原变量的值
	temp := *a
	*a = *b
	*b = temp
}

// 场景4: 大结构体传递 - 值传递（复制整个结构体，性能差）
func processPersonByValue(p Person) {
	p.Age += 1 // 只修改副本
	fmt.Println("函数内 p.Age =", p.Age)
}

// 场景5: 大结构体传递 - 指针传递（只传递地址，性能好）
func processPersonByPointer(p *Person) {
	p.Age += 1 // 修改原结构体
	fmt.Println("函数内 p.Age =", p.Age)
}

// 场景6: 可选参数（nil 指针表示未提供）
func createPerson(name string, age *int) Person {
	p := Person{Name: name}
	if age != nil {
		p.Age = *age
	} else {
		p.Age = 0 // 默认值
	}
	return p
}

// pointerRealWorldDemo 演示指针的实际作用和应用场景
func pointerRealWorldDemo() {
	fmt.Println("=== 指针的真正作用和应用场景 ===")
	fmt.Println()

	// ========== 场景1: 函数中修改外部变量 ==========
	fmt.Println("--- 场景1: 函数中修改外部变量 ---")
	num1 := 42
	fmt.Println("调用前 num1 =", num1)
	modifyValue(num1)               // 值传递，无法修改原变量
	fmt.Println("调用后 num1 =", num1) // 仍然是 42，没有被修改
	fmt.Println()

	num2 := 42
	fmt.Println("调用前 num2 =", num2)
	modifyValueByPointer(&num2)     // 指针传递，可以修改原变量
	fmt.Println("调用后 num2 =", num2) // 变成了 999，被成功修改
	fmt.Println()

	// ========== 场景2: 交换两个变量的值 ==========
	fmt.Println("--- 场景2: 交换两个变量的值 ---")
	a, b := 10, 20
	fmt.Printf("交换前: a=%d, b=%d\n", a, b)
	swapWithPointer(&a, &b) // 使用指针真正交换
	fmt.Printf("交换后: a=%d, b=%d\n", a, b)
	fmt.Println()

	// ========== 场景3: 大结构体传递性能对比 ==========
	fmt.Println("--- 场景3: 大结构体传递性能 ---")
	person1 := Person{Name: "张三", Age: 25}
	fmt.Println("值传递前 person1.Age =", person1.Age)
	processPersonByValue(person1)                  // 复制整个结构体（如果结构体很大，性能差）
	fmt.Println("值传递后 person1.Age =", person1.Age) // 没有被修改
	fmt.Println()

	person2 := Person{Name: "李四", Age: 25}
	fmt.Println("指针传递前 person2.Age =", person2.Age)
	processPersonByPointer(&person2)                // 只传递地址（8字节），性能好
	fmt.Println("指针传递后 person2.Age =", person2.Age) // 被成功修改
	fmt.Println()

	// ========== 场景4: 可选参数（nil 指针） ==========
	fmt.Println("--- 场景4: 可选参数（nil 指针） ---")
	age := 30
	p1 := createPerson("王五", &age) // 提供年龄
	fmt.Printf("p1: %+v\n", p1)

	p2 := createPerson("赵六", nil) // 不提供年龄（使用默认值）
	fmt.Printf("p2: %+v\n", p2)
	fmt.Println()

	// ========== 场景5: 多个函数共享同一个变量 ==========
	fmt.Println("--- 场景5: 多个函数共享同一个变量 ---")
	counter := 0
	increment := func() { counter++ }
	decrement := func() { counter-- }
	reset := func() { counter = 0 }

	fmt.Println("初始值:", counter)
	increment()
	increment()
	fmt.Println("增加2次后:", counter)
	decrement()
	fmt.Println("减少1次后:", counter)
	reset()
	fmt.Println("重置后:", counter)
	fmt.Println()

	// ========== 场景6: 实际应用 - 修改配置 ==========
	fmt.Println("--- 场景6: 实际应用 - 修改配置 ---")
	config := Person{Name: "默认用户", Age: 18}
	fmt.Println("初始配置:", config)

	updateConfig := func(p *Person) {
		p.Name = "新用户"
		p.Age = 25
	}

	updateConfig(&config)
	fmt.Println("更新后配置:", config)
	fmt.Println()

	// ========== 总结 ==========
	fmt.Println("=== 指针的作用总结 ===")
	fmt.Println("1. 在函数中修改外部变量（值传递无法做到）")
	fmt.Println("2. 避免大对象复制，提高性能（只传递8字节地址）")
	fmt.Println("3. 实现可选参数（nil 指针）")
	fmt.Println("4. 多个函数共享同一个变量")
	fmt.Println("5. 实现真正的数据交换")
	fmt.Println()
	fmt.Println("注意：对于简单变量（如 int），直接赋值 num = 100 确实更简单")
	fmt.Println("但指针的真正价值在于函数间传递和修改数据！")
}

// userQuestionDemo 详细解释用户问题的代码
func userQuestionDemo() {
	fmt.Println("=== 代码逐行解释 ===")
	fmt.Println()

	// 第1步：声明一个指向 int 的指针变量
	var p1 *int
	fmt.Println("步骤1: var p1 *int")
	fmt.Println("  解释: 声明一个指向 int 类型的指针变量 p1")
	fmt.Println("  此时 p1 的值是 nil（空指针）")
	fmt.Printf("  p1 = %v\n", p1)
	fmt.Println()

	// 第2步：声明并初始化一个 int 变量
	i := 1
	fmt.Println("步骤2: i := 1")
	fmt.Println("  解释: 声明并初始化一个 int 变量 i，值为 1")
	fmt.Printf("  i = %d\n", i)
	fmt.Printf("  i 的地址 = %p\n", &i)
	fmt.Println()

	// 第3步：将 i 的地址赋值给指针 p1
	p1 = &i
	fmt.Println("步骤3: p1 = &i")
	fmt.Println("  解释: 使用 & 取地址符，获取变量 i 的内存地址，并赋值给指针 p1")
	fmt.Printf("  p1 现在存储的是 i 的地址: %p\n", p1)
	fmt.Printf("  p1 指向的值（*p1）: %d\n", *p1)
	fmt.Printf("  i 的值: %d\n", i)
	fmt.Println("  此时 p1 和 &i 指向同一个内存地址")
	fmt.Println()

	// 第4步：比较 *p1 和 i 的值
	fmt.Println("步骤4: fmt.Println(*p1 == i)")
	fmt.Println("  解释: *p1 是解引用，获取 p1 指向的值；i 是变量本身的值")
	fmt.Printf("  *p1 = %d\n", *p1)
	fmt.Printf("  i = %d\n", i)
	fmt.Printf("  *p1 == i 的结果: %v\n", *p1 == i)
	fmt.Println("  因为 p1 指向 i，所以 *p1 和 i 的值相等，输出 true")
	fmt.Println()

	// 第5步：通过指针修改值
	*p1 = 2
	fmt.Println("步骤5: *p1 = 2")
	fmt.Println("  解释: 通过解引用 *p1，修改 p1 指向的内存地址中的值")
	fmt.Println("  因为 p1 指向 i，所以实际上是修改了 i 的值")
	fmt.Printf("  修改后 *p1 = %d\n", *p1)
	fmt.Printf("  修改后 i = %d\n", i)
	fmt.Println("  注意：i 的值也被改变了，因为 p1 指向的就是 i")
	fmt.Println()

	// 第6步：输出 i 的值
	fmt.Println("步骤6: fmt.Println(i)")
	fmt.Printf("  输出: %d\n", i)
	fmt.Println("  解释: 因为通过 *p1 = 2 修改了 i 的值，所以 i 现在是 2")
	fmt.Println()

	// ========== 内存关系图 ==========
	fmt.Println("=== 内存关系示意图 ===")
	fmt.Println("变量 i 在内存中:")
	fmt.Println("  地址: 0x... (某个内存地址)")
	fmt.Println("  值: 2")
	fmt.Println()
	fmt.Println("指针 p1 在内存中:")
	fmt.Println("  地址: 0x... (另一个内存地址)")
	fmt.Println("  值: 0x... (存储的是变量 i 的地址)")
	fmt.Println()
	fmt.Println("关系:")
	fmt.Println("  p1 → 指向 → i 的内存地址")
	fmt.Println("  *p1 → 解引用 → 获取 i 的值")
	fmt.Println("  修改 *p1 → 实际上修改了 i 的值")
	fmt.Println()

	// ========== 完整代码演示 ==========
	fmt.Println("=== 完整代码演示 ===")
	var p2 *int
	j := 1
	p2 = &j
	fmt.Printf("初始: j = %d, *p2 = %d\n", j, *p2)
	fmt.Printf("*p2 == j: %v\n", *p2 == j)
	*p2 = 2
	fmt.Printf("修改后: j = %d\n", j)
}

// doublePointerDemo 详细解释二级指针（指向指针的指针）
func doublePointerDemo() {
	fmt.Println("=== 二级指针（指向指针的指针）详解 ===")
	fmt.Println()

	// ========== 步骤1: 声明变量 a ==========
	a := 2
	fmt.Println("步骤1: a := 2")
	fmt.Println("  解释: 声明并初始化变量 a，值为 2")
	fmt.Printf("  a = %d\n", a)
	fmt.Printf("  a 的地址 (&a) = %p\n", &a)
	fmt.Println()

	// ========== 步骤2: 声明一级指针 ==========
	var p *int
	fmt.Println("步骤2: var p *int")
	fmt.Println("  解释: 声明一个指向 int 的一级指针变量 p")
	fmt.Printf("  p 的初始值 = %v (nil)\n", p)
	fmt.Println()

	// ========== 步骤3: 打印 a 的地址 ==========
	fmt.Println("步骤3: fmt.Println(&a)")
	fmt.Printf("  输出: %p (a 的内存地址)\n", &a)
	fmt.Println()

	// ========== 步骤4: p 指向 a ==========
	p = &a
	fmt.Println("步骤4: p = &a")
	fmt.Println("  解释: 将 a 的地址赋值给一级指针 p")
	fmt.Printf("  p 现在存储的值 = %p (a 的地址)\n", p)
	fmt.Printf("  *p = %d (p 指向的值，即 a 的值)\n", *p)
	fmt.Printf("  p 的地址 (&p) = %p\n", &p)
	fmt.Println()

	// ========== 步骤5: 打印 p 和 &a ==========
	fmt.Println("步骤5: fmt.Println(p, &a)")
	fmt.Printf("  输出: %p %p\n", p, &a)
	fmt.Println("  解释: p 的值和 &a 相同，都是 a 的内存地址")
	fmt.Printf("  p == &a: %v\n", p == &a)
	fmt.Println()

	// ========== 步骤6: 声明二级指针 ==========
	var pp **int
	fmt.Println("步骤6: var pp **int")
	fmt.Println("  解释: 声明一个二级指针 pp，它指向一个 *int 类型的指针")
	fmt.Println("  **int 表示：指向（指向 int 的指针）的指针")
	fmt.Printf("  pp 的初始值 = %v (nil)\n", pp)
	fmt.Println()

	// ========== 步骤7: pp 指向 p ==========
	pp = &p
	fmt.Println("步骤7: pp = &p")
	fmt.Println("  解释: 将一级指针 p 的地址赋值给二级指针 pp")
	fmt.Printf("  pp 现在存储的值 = %p (p 的地址)\n", pp)
	fmt.Printf("  *pp = %p (pp 指向的值，即 p 的值，也就是 a 的地址)\n", *pp)
	fmt.Printf("  **pp = %d (pp 指向的指针指向的值，即 a 的值)\n", **pp)
	fmt.Println()

	// ========== 步骤8: 打印 pp 和 p ==========
	fmt.Println("步骤8: fmt.Println(pp, p)")
	fmt.Printf("  输出: %p %p\n", pp, p)
	fmt.Println("  解释:")
	fmt.Printf("    pp = %p (p 的地址)\n", pp)
	fmt.Printf("    p = %p (a 的地址)\n", p)
	fmt.Println("  注意: pp 存储的是 p 的地址，p 存储的是 a 的地址")
	fmt.Println()

	// ========== 步骤9: 通过二级指针修改值 ==========
	**pp = 3
	fmt.Println("步骤9: **pp = 3")
	fmt.Println("  解释: 通过二级指针修改值")
	fmt.Println("  **pp 的含义:")
	fmt.Println("    *pp → 获取 pp 指向的值（即 p，也就是 a 的地址）")
	fmt.Println("    **pp → 再解引用一次，获取 a 的值")
	fmt.Println("    所以 **pp = 3 实际上是修改了 a 的值")
	fmt.Printf("  修改后 a = %d\n", a)
	fmt.Printf("  修改后 *p = %d\n", *p)
	fmt.Printf("  修改后 **pp = %d\n", **pp)
	fmt.Println()

	// ========== 步骤10: 打印各种关系 ==========
	fmt.Println("步骤10: fmt.Println(pp, *pp, p)")
	fmt.Printf("  输出: %p %p %p\n", pp, *pp, p)
	fmt.Println("  解释:")
	fmt.Printf("    pp = %p (p 的地址)\n", pp)
	fmt.Printf("    *pp = %p (p 的值，即 a 的地址)\n", *pp)
	fmt.Printf("    p = %p (a 的地址)\n", p)
	fmt.Printf("    所以 *pp == p: %v\n", *pp == p)
	fmt.Println()

	// ========== 步骤11: 打印 **pp 和 *p ==========
	fmt.Println("步骤11: fmt.Println(**pp, *p)")
	fmt.Printf("  输出: %d %d\n", **pp, *p)
	fmt.Println("  解释:")
	fmt.Println("    **pp → 通过二级指针获取最终的值（a 的值）")
	fmt.Println("    *p → 通过一级指针获取值（a 的值）")
	fmt.Printf("    所以 **pp == *p: %v\n", **pp == *p)
	fmt.Println()

	// ========== 步骤12: 打印 a 和 &a ==========
	fmt.Println("步骤12: fmt.Println(a, &a)")
	fmt.Printf("  输出: %d %p\n", a, &a)
	fmt.Println("  解释:")
	fmt.Printf("    a = %d (变量的值)\n", a)
	fmt.Printf("    &a = %p (变量的地址)\n", &a)
	fmt.Println()

	// ========== 完整的内存关系图 ==========
	fmt.Println("=== 完整的内存关系示意图 ===")
	fmt.Println()
	fmt.Println("内存布局:")
	fmt.Println("┌─────────────┐")
	fmt.Println("│  变量 a     │")
	fmt.Println("│  地址: &a   │")
	fmt.Println("│  值: 3      │")
	fmt.Println("└─────────────┘")
	fmt.Println("      ↑")
	fmt.Println("      │ (p 指向这里)")
	fmt.Println("┌─────────────┐")
	fmt.Println("│  指针 p     │")
	fmt.Println("│  地址: &p   │")
	fmt.Println("│  值: &a     │")
	fmt.Println("└─────────────┘")
	fmt.Println("      ↑")
	fmt.Println("      │ (pp 指向这里)")
	fmt.Println("┌─────────────┐")
	fmt.Println("│ 二级指针 pp │")
	fmt.Println("│  地址: &pp  │")
	fmt.Println("│  值: &p     │")
	fmt.Println("└─────────────┘")
	fmt.Println()

	// ========== 关系总结 ==========
	fmt.Println("=== 关系总结 ===")
	fmt.Println("变量关系:")
	fmt.Printf("  a = %d\n", a)
	fmt.Printf("  &a = %p\n", &a)
	fmt.Println()
	fmt.Println("一级指针关系:")
	fmt.Printf("  p = %p (存储 a 的地址)\n", p)
	fmt.Printf("  *p = %d (p 指向的值，即 a 的值)\n", *p)
	fmt.Printf("  &p = %p (p 自己的地址)\n", &p)
	fmt.Println()
	fmt.Println("二级指针关系:")
	fmt.Printf("  pp = %p (存储 p 的地址)\n", pp)
	fmt.Printf("  *pp = %p (pp 指向的值，即 p 的值，也就是 &a)\n", *pp)
	fmt.Printf("  **pp = %d (pp 指向的指针指向的值，即 a 的值)\n", **pp)
	fmt.Printf("  &pp = %p (pp 自己的地址)\n", &pp)
	fmt.Println()
	fmt.Println("等价关系:")
	fmt.Printf("  *pp == p: %v\n", *pp == p)
	fmt.Printf("  **pp == *p: %v\n", **pp == *p)
	fmt.Printf("  **pp == a: %v\n", **pp == a)
	fmt.Println()

	// ========== 实际应用场景 ==========
	fmt.Println("=== 二级指针的应用场景 ===")
	fmt.Println("1. 在函数中修改指针本身（而不仅仅是指针指向的值）")
	fmt.Println("2. 动态分配内存（如 C 语言中的 malloc）")
	fmt.Println("3. 链表、树等数据结构中修改指针")
	fmt.Println("4. 函数返回指针，同时需要修改指针变量本身")
	fmt.Println()

	// ========== 实际代码演示 ==========
	fmt.Println("=== 完整代码演示 ===")
	b := 2
	var p1 *int
	fmt.Printf("初始: b = %d, &b = %p\n", b, &b)
	p1 = &b
	fmt.Printf("p1 = &b: p1 = %p, *p1 = %d\n", p1, *p1)
	var pp1 **int
	pp1 = &p1
	fmt.Printf("pp1 = &p1: pp1 = %p, *pp1 = %p, **pp1 = %d\n", pp1, *pp1, **pp1)
	**pp1 = 3
	fmt.Printf("**pp1 = 3 后: b = %d, *p1 = %d, **pp1 = %d\n", b, *p1, **pp1)
	fmt.Println()
}

// pointerTypesDemo 详细解释三种指针类型：普通指针、unsafe.Pointer、uintptr
func pointerTypesDemo() {
	fmt.Println("=== 三种指针类型详解：普通指针、unsafe.Pointer、uintptr ===")
	fmt.Println()

	// ========== 1. 普通指针 (*T) ==========
	fmt.Println("--- 1. 普通指针 (*T) ---")
	fmt.Println("特点:")
	fmt.Println("  - 类型安全，只能指向特定类型")
	fmt.Println("  - 受 Go 的垃圾回收器管理")
	fmt.Println("  - 不能进行算术运算")
	fmt.Println("  - 不能直接转换为其他类型的指针")
	fmt.Println()

	var x int = 42
	var p *int = &x
	fmt.Printf("示例: var x int = 42; var p *int = &x\n")
	fmt.Printf("  x = %d\n", x)
	fmt.Printf("  p = %p\n", p)
	fmt.Printf("  *p = %d\n", *p)
	fmt.Println()

	// 普通指针不能直接转换为其他类型
	// var p2 *float64 = (*float64)(p)  // 编译错误！

	// ========== 2. unsafe.Pointer ==========
	fmt.Println("--- 2. unsafe.Pointer ---")
	fmt.Println("特点:")
	fmt.Println("  - 通用指针类型，可以指向任何类型")
	fmt.Println("  - 可以转换为任何类型的指针")
	fmt.Println("  - 可以转换为 uintptr")
	fmt.Println("  - 仍然受垃圾回收器管理")
	fmt.Println("  - 不能进行算术运算")
	fmt.Println("  - 使用 unsafe 包，需要谨慎使用")
	fmt.Println()

	var y int = 100
	var ptr unsafe.Pointer = unsafe.Pointer(&y)
	fmt.Printf("示例: var y int = 100; var ptr unsafe.Pointer = unsafe.Pointer(&y)\n")
	fmt.Printf("  y = %d\n", y)
	fmt.Printf("  ptr = %p\n", ptr)
	fmt.Println()

	// unsafe.Pointer 可以转换为其他类型的指针
	var pFloat *float64 = (*float64)(ptr)
	fmt.Printf("转换为 *float64: %p\n", pFloat)
	var pInt *int = (*int)(ptr)
	fmt.Printf("转换回 *int: %p, 值: %d\n", pInt, *pInt)
	fmt.Println()

	// ========== 3. uintptr ==========
	fmt.Println("--- 3. uintptr ---")
	fmt.Println("特点:")
	fmt.Println("  - 整数类型，足够大以存储指针值")
	fmt.Println("  - 可以进行算术运算")
	fmt.Println("  - 不受垃圾回收器管理（危险！）")
	fmt.Println("  - 可以转换为 unsafe.Pointer")
	fmt.Println("  - 主要用于底层内存操作")
	fmt.Println()

	var z int = 200
	var addr uintptr = uintptr(unsafe.Pointer(&z))
	fmt.Printf("示例: var z int = 200; var addr uintptr = uintptr(unsafe.Pointer(&z))\n")
	fmt.Printf("  z = %d\n", z)
	fmt.Printf("  addr = 0x%x (内存地址的整数值)\n", addr)
	fmt.Println()

	// uintptr 可以进行算术运算
	var addr2 uintptr = addr + unsafe.Sizeof(z)
	fmt.Printf("地址运算: addr + sizeof(int) = 0x%x\n", addr2)
	fmt.Println()

	// ========== 三种类型的转换关系 ==========
	fmt.Println("--- 三种类型的转换关系 ---")
	fmt.Println("转换路径:")
	fmt.Println("  普通指针 (*T) → unsafe.Pointer → uintptr")
	fmt.Println("  uintptr → unsafe.Pointer → 普通指针 (*T)")
	fmt.Println()

	var num int = 999
	var normalPtr *int = &num
	var unsafePtr unsafe.Pointer = unsafe.Pointer(normalPtr)
	var uintptrVal uintptr = uintptr(unsafePtr)

	fmt.Printf("转换示例:\n")
	fmt.Printf("  普通指针: %p, 值: %d\n", normalPtr, *normalPtr)
	fmt.Printf("  unsafe.Pointer: %p\n", unsafePtr)
	fmt.Printf("  uintptr: 0x%x\n", uintptrVal)
	fmt.Println()

	// 反向转换
	// 注意：在实际使用中，要确保 uintptr 转换期间对象不会被 GC 回收
	// 这里立即转换并使用，所以是安全的
	var backUnsafe unsafe.Pointer = unsafe.Pointer(uintptrVal)
	var backNormal *int = (*int)(backUnsafe)
	fmt.Printf("反向转换:\n")
	fmt.Printf("  uintptr → unsafe.Pointer → *int\n")
	fmt.Printf("  结果: %p, 值: %d\n", backNormal, *backNormal)
	fmt.Println()

	// ========== 实际应用示例 ==========
	fmt.Println("--- 实际应用示例 ---")
	fmt.Println()

	// 示例1: 获取结构体字段的偏移量
	type Example struct {
		A int32
		B int64
		C int32
	}

	ex := Example{A: 1, B: 2, C: 3}
	fmt.Println("示例1: 获取结构体字段偏移量")
	fmt.Printf("  结构体: %+v\n", ex)
	fmt.Printf("  A 字段偏移: %d 字节\n", unsafe.Offsetof(ex.A))
	fmt.Printf("  B 字段偏移: %d 字节\n", unsafe.Offsetof(ex.B))
	fmt.Printf("  C 字段偏移: %d 字节\n", unsafe.Offsetof(ex.C))
	fmt.Printf("  结构体大小: %d 字节\n", unsafe.Sizeof(ex))
	fmt.Println()

	// 示例2: 通过偏移量访问字段
	fmt.Println("示例2: 通过偏移量访问字段")
	basePtr := unsafe.Pointer(&ex)
	bPtr := (*int64)(unsafe.Pointer(uintptr(basePtr) + unsafe.Offsetof(ex.B)))
	fmt.Printf("  通过偏移量访问 B 字段: %d\n", *bPtr)
	*bPtr = 999
	fmt.Printf("  修改后 ex.B = %d\n", ex.B)
	fmt.Println()

	// 示例3: 类型转换（危险操作）
	fmt.Println("示例3: 类型转换（需要谨慎）")
	var floatVal float64 = 3.14159
	var floatPtr *float64 = &floatVal
	var intPtr *int64 = (*int64)(unsafe.Pointer(floatPtr))
	fmt.Printf("  float64 值: %f\n", floatVal)
	fmt.Printf("  转换为 int64 的位模式: %d (这是位模式的整数表示，不是实际值！)\n", *intPtr)
	fmt.Println("  警告: 这种转换只是改变解释方式，不改变内存中的位！")
	fmt.Println()

	// ========== 安全注意事项 ==========
	fmt.Println("--- 安全注意事项 ---")
	fmt.Println("⚠️ 使用 unsafe 包的风险:")
	fmt.Println("1. uintptr 不受 GC 管理，可能导致悬空指针")
	fmt.Println("2. 类型转换可能导致内存对齐问题")
	fmt.Println("3. 指针运算可能导致访问非法内存")
	fmt.Println("4. 破坏了 Go 的类型安全保证")
	fmt.Println()
	fmt.Println("✅ 安全使用原则:")
	fmt.Println("1. 只在必要时使用 unsafe 包")
	fmt.Println("2. 确保 uintptr 在使用期间对象不会被 GC 回收")
	fmt.Println("3. 遵循 unsafe 包的文档规范")
	fmt.Println("4. 进行充分的测试")
	fmt.Println()

	// ========== 对比总结 ==========
	fmt.Println("=== 三种指针类型对比总结 ===")
	fmt.Println()
	fmt.Println("┌─────────────┬──────────────┬──────────────┬──────────────┐")
	fmt.Println("│  特性       │  普通指针    │ unsafe.Pointer│  uintptr     │")
	fmt.Println("├─────────────┼──────────────┼──────────────┼──────────────┤")
	fmt.Println("│ 类型安全     │     ✅       │     ❌       │     ❌       │")
	fmt.Println("│ GC 管理     │     ✅       │     ✅       │     ❌       │")
	fmt.Println("│ 算术运算     │     ❌       │     ❌       │     ✅       │")
	fmt.Println("│ 类型转换     │     ❌       │     ✅       │     ✅       │")
	fmt.Println("│ 使用场景     │  日常开发    │  底层操作    │  内存计算    │")
	fmt.Println("└─────────────┴──────────────┴──────────────┴──────────────┘")
	fmt.Println()

	// ========== 完整转换示例 ==========
	fmt.Println("=== 完整转换示例 ===")
	value := 12345
	fmt.Printf("原始值: %d\n", value)

	// 普通指针 → unsafe.Pointer → uintptr
	normalP := &value
	unsafeP := unsafe.Pointer(normalP)
	uintptrP := uintptr(unsafeP)
	fmt.Printf("转换链: *int → unsafe.Pointer → uintptr\n")
	fmt.Printf("  *int: %p, 值: %d\n", normalP, *normalP)
	fmt.Printf("  unsafe.Pointer: %p\n", unsafeP)
	fmt.Printf("  uintptr: 0x%x\n", uintptrP)
	fmt.Println()

	// uintptr → unsafe.Pointer → 普通指针
	// 注意：在实际使用中，要确保 uintptr 转换期间对象不会被 GC 回收
	// 这里立即转换并使用，所以是安全的
	backUnsafeP := unsafe.Pointer(uintptrP)
	backNormalP := (*int)(backUnsafeP)
	fmt.Printf("反向转换: uintptr → unsafe.Pointer → *int\n")
	fmt.Printf("  结果: %p, 值: %d\n", backNormalP, *backNormalP)
	fmt.Println()

	// 修改值验证
	*backNormalP = 54321
	fmt.Printf("通过转换后的指针修改值: %d\n", value)
	fmt.Printf("原始变量也被修改: %d\n", value)
}
