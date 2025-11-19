package main

import "fmt"

// ========== 嵌套结构体示例 ==========

// 结构体 A
type A struct {
	a string
}

// 结构体 A 的方法
func (a A) GetA() string {
	return a.a
}

func (a A) SetA(value string) {
	a.a = value
}

func (a *A) SetAPtr(value string) {
	a.a = value
}

// 结构体 B - 嵌套了 A
type B struct {
	A        // 匿名嵌套，可以直接访问 A 的字段和方法
	b string
}

// 结构体 B 的方法
func (b B) GetB() string {
	return b.b
}

// 结构体 C - 嵌套了 A 和 B，同时有自己的字段 a, b, c
type C struct {
	A        // 嵌套 A
	B        // 嵌套 B
	a string // 与 A 中的字段 a 同名
	b string // 与 B 中的字段 b 同名
	c string // C 自己的字段
}

// ========== 嵌套结构体的使用示例 ==========

func demonstrateNestedStruct() {
	fmt.Println("=== 嵌套结构体演示 ===")
	fmt.Println()

	// 1. 创建结构体 A 的实例
	fmt.Println("--- 1. 结构体 A 的使用 ---")
	a := A{a: "A的字段a"}
	fmt.Printf("A.a = %s\n", a.a)
	fmt.Printf("A.GetA() = %s\n", a.GetA())
	fmt.Println()

	// 2. 创建结构体 B 的实例（嵌套了 A）
	fmt.Println("--- 2. 结构体 B 的使用（嵌套了 A）---")
	b := B{
		A: A{a: "B中嵌套的A的字段a"},
		b: "B的字段b",
	}
	// 可以直接访问嵌套结构体 A 的字段（因为 B 中没有同名字段）
	fmt.Printf("B.a = %s (直接访问嵌套的A.a)\n", b.a)
	fmt.Printf("B.A.a = %s (通过B.A访问)\n", b.A.a)
	fmt.Printf("B.b = %s\n", b.b)
	// 可以直接调用嵌套结构体 A 的方法
	fmt.Printf("B.GetA() = %s (直接调用嵌套的A的方法)\n", b.GetA())
	fmt.Printf("B.GetB() = %s\n", b.GetB())
	fmt.Println()

	// 3. 创建结构体 C 的实例（嵌套了 A 和 B，同时有自己的字段 a, b, c）
	fmt.Println("--- 3. 结构体 C 的使用（嵌套了 A 和 B，有同名字段）---")
	c := C{
		A: A{a: "C中嵌套的A的字段a"},
		B: B{
			A: A{a: "C中嵌套的B中嵌套的A的字段a"},
			b: "C中嵌套的B的字段b",
		},
		a: "C自己的字段a",
		b: "C自己的字段b",
		c: "C的字段c",
	}

	// 访问 C 自己的字段
	fmt.Printf("C.a = %s (C自己的字段a)\n", c.a)
	fmt.Printf("C.b = %s (C自己的字段b)\n", c.b)
	fmt.Printf("C.c = %s (C的字段c)\n", c.c)

	// 访问嵌套的 A 的字段（通过 C.A）
	fmt.Printf("C.A.a = %s (通过C.A访问嵌套的A的字段a)\n", c.A.a)
	fmt.Printf("C.A.GetA() = %s (通过C.A调用嵌套的A的方法)\n", c.A.GetA())

	// 访问嵌套的 B 的字段（通过 C.B）
	fmt.Printf("C.B.a = %s (通过C.B访问嵌套的B中嵌套的A的字段a)\n", c.B.a)
	fmt.Printf("C.B.A.a = %s (通过C.B.A访问)\n", c.B.A.a)
	fmt.Printf("C.B.b = %s (通过C.B访问嵌套的B的字段b)\n", c.B.b)
	fmt.Printf("C.B.GetA() = %s (通过C.B调用嵌套的B中嵌套的A的方法)\n", c.B.GetA())
	fmt.Printf("C.B.GetB() = %s (通过C.B调用嵌套的B的方法)\n", c.B.GetB())
	fmt.Println()

	// 4. 演示方法调用
	fmt.Println("--- 4. 嵌套结构体的方法调用 ---")
	// C 可以直接调用 A 的方法（因为 C 中没有同名方法）
	fmt.Printf("C.GetA() = %s (直接调用嵌套的A的方法，因为C中没有同名方法)\n", c.GetA())
	// C 可以直接调用 B 的方法（因为 C 中没有同名方法）
	fmt.Printf("C.GetB() = %s (直接调用嵌套的B的方法，因为C中没有同名方法)\n", c.GetB())
	fmt.Println()

	// 5. 演示字段和方法名的冲突处理
	fmt.Println("--- 5. 字段和方法名的冲突处理 ---")
	fmt.Println("规则说明：")
	fmt.Println("1. 如果外层结构体没有同名字段，可以直接访问嵌套结构体的字段")
	fmt.Println("2. 如果外层结构体有同名字段，需要通过 外层.嵌套结构体名 的方式访问")
	fmt.Println("3. 方法调用规则与字段相同：没有同名方法可以直接调用，有同名方法需要通过嵌套结构体名调用")
	fmt.Println()

	// 6. 演示指针接收者的方法
	fmt.Println("--- 6. 指针接收者方法的使用 ---")
	cPtr := &c
	cPtr.A.SetAPtr("通过指针修改A的字段a")
	fmt.Printf("修改后 C.A.a = %s\n", cPtr.A.a)
	fmt.Println()

	// 7. 演示多层嵌套
	fmt.Println("--- 7. 多层嵌套结构体 ---")
	type D struct {
		C        // 嵌套 C
		d string
	}
	d := D{
		C: c,
		d: "D的字段d",
	}
	fmt.Printf("D.d = %s\n", d.d)
	fmt.Printf("D.c = %s (直接访问嵌套的C的字段c)\n", d.c)
	fmt.Printf("D.C.c = %s (通过D.C访问)\n", d.C.c)
	fmt.Printf("D.a = %s (直接访问嵌套的C的字段a)\n", d.a)
	fmt.Printf("D.C.a = %s (通过D.C访问)\n", d.C.a)
	fmt.Printf("D.A.a = %s (直接访问嵌套的C中嵌套的A的字段a)\n", d.A.a)
	fmt.Printf("D.C.A.a = %s (通过D.C.A访问)\n", d.C.A.a)
}

func main() {
	demonstrateNestedStruct()
}

