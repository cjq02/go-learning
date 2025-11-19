package structs

import (
	"fmt"
	"strings"
)

// ========== 结构体方法定义 ==========
// 在 Go 中，结构体的方法和结构体是分开的，但是一般会在声明结构体的代码下面紧跟着声明结构体的方法。
//
// 语法格式：
// func (<ref name> <Name>) <methodName>(<param name> <ParameterType>, ...) (<return name> <return type>) {
//     ...
// }
//
// func (<ref name> *<Name>) <methodName>(<param name> <ParameterType>, ...) (<return name> <return type>) {
//     ...
// }
//
// 说明：
// 1. 与普通的函数相比，结构体的方法在 func 关键字后面多了一个 (<ref name> <Name>) 结构
// 2. 当方法中需要引用到结构体中的字段或者结构体其他方法时，需要使用这个 ref name
// 3. 这个 ref name 相当于这个方法中提前声明的一个结构体类型的局部变量
// 4. Name 则是结构体名称
// 5. Go 中没有方法重写的规则，即不允许相同名称的方法存在，即使方法的参数不同
// 6. Go 中，结构体类型的变量和结构体类型指针的变量都可以直接访问结构体中声明的字段和调用声明的方法
// 7. Go 中，没有结构体构造函数，构造函数都是自己开发者自定义的，一般是按照习惯，定义一个名称是 New<struct Name> 的函数来实例化一个比较复杂的结构体
//
// 方法名称大小写规则（与 Go 语言中所有标识符的规则相同）：
// - 首字母大写的方法名：导出的（exported/public），可以被其他包访问
//   例如：func (a A) GetValue() string { ... }  // 其他包可以调用 a.GetValue()
// - 首字母小写的方法名：未导出的（unexported/private），只能在同一个包内访问
//   例如：func (a A) getValue() string { ... }  // 只能在同一个包内调用 a.getValue()
// 这个规则同样适用于结构体字段、函数、变量、类型等所有标识符

// ========== 代码示例 1：值接收者和指针接收者 ==========

// 结构体 MethodA（用于演示方法）
type MethodA struct {
	a string
}

// 值接收者方法 - 返回字符串
func (a MethodA) string() string {
	return a.a
}

// 值接收者方法 - 返回字符串（带后缀A）
func (a MethodA) stringA() string {
	return a.a
}

// 值接收者方法 - 设置字段（注意：值接收者无法修改原结构体）
func (a MethodA) setA(v string) {
	a.a = v // 这里修改的是副本，不会影响原结构体
}

// 指针接收者方法 - 返回字符串（带后缀PA）
func (a *MethodA) stringPA() string {
	return a.a
}

// 指针接收者方法 - 设置字段（可以修改原结构体）
func (a *MethodA) setPA(v string) {
	a.a = v // 这里修改的是原结构体
}

// 结构体 MethodB - 嵌套了 MethodA
type MethodB struct {
	MethodA
	b string
}

// 值接收者方法 - 返回字符串（覆盖了 MethodA 的 string 方法）
func (b MethodB) string() string {
	return b.b
}

// 值接收者方法 - 返回字符串（带后缀B）
func (b MethodB) stringB() string {
	return b.b
}

// 结构体 MethodC - 嵌套了 MethodB（MethodB 又嵌套了 MethodA），同时有自己的字段 a, b, c, d
type MethodC struct {
	MethodB
	a string
	b string
	c string
	d []byte
}

// 值接收者方法 - 返回字符串（覆盖了 MethodB 的 string 方法）
func (c MethodC) string() string {
	return c.c
}

// 值接收者方法 - 修改切片（注意：虽然 c 是值接收者，但切片是引用类型，所以可以修改）
func (c MethodC) modityD() {
	c.d[2] = 3 // 切片是引用类型，即使 c 是值接收者，也能修改切片内容
}

// 构造函数 - 按照习惯，定义一个名称是 New<struct Name> 的函数来实例化结构体
func NewMethodC() MethodC {
	return MethodC{
		MethodB: MethodB{
			MethodA: MethodA{
				a: "ba",
			},
			b: "b",
		},
		a: "ca",
		b: "cb",
		c: "c",
		d: []byte{1, 2, 3},
	}
}

// ========== 代码示例 1 的使用演示 ==========

func demonstrateStructMethods1() {
	fmt.Println("=== 代码示例 1：值接收者和指针接收者的区别 ===")
	fmt.Println()

	c := NewMethodC()
	cp := &c

	// 演示：结构体变量和指针都可以调用方法
	fmt.Println("--- 1. 结构体变量和指针都可以调用方法 ---")
	fmt.Printf("c.string() = %s\n", c.string())   // 调用 MethodC 的 string 方法
	fmt.Printf("c.stringA() = %s\n", c.stringA()) // 调用嵌套的 MethodA 的 stringA 方法
	fmt.Printf("c.stringB() = %s\n", c.stringB()) // 调用嵌套的 MethodB 的 stringB 方法

	fmt.Printf("cp.string() = %s\n", cp.string())   // 指针也可以调用
	fmt.Printf("cp.stringA() = %s\n", cp.stringA()) // 指针也可以调用
	fmt.Printf("cp.stringB() = %s\n", cp.stringB()) // 指针也可以调用
	fmt.Println()

	// 演示：值接收者方法无法修改原结构体
	fmt.Println("--- 2. 值接收者方法无法修改原结构体 ---")
	fmt.Printf("修改前 c.MethodA.a = %s\n", c.MethodA.a)
	c.setA("1a") // 值接收者，修改的是副本
	fmt.Printf("c.setA(\"1a\") 后 c.MethodA.a = %s (未改变)\n", c.MethodA.a)
	fmt.Printf("cp.MethodA.a = %s (未改变)\n", cp.MethodA.a)
	fmt.Println()

	fmt.Printf("修改前 cp.MethodA.a = %s\n", cp.MethodA.a)
	cp.setA("2a") // 即使通过指针调用，值接收者方法仍然修改的是副本
	fmt.Printf("cp.setA(\"2a\") 后 c.MethodA.a = %s (未改变)\n", c.MethodA.a)
	fmt.Printf("cp.MethodA.a = %s (未改变)\n", cp.MethodA.a)
	fmt.Println()

	// 演示：指针接收者方法可以修改原结构体
	fmt.Println("--- 3. 指针接收者方法可以修改原结构体 ---")
	fmt.Printf("修改前 c.MethodA.a = %s\n", c.MethodA.a)
	c.setPA("3a") // 指针接收者，可以修改原结构体
	fmt.Printf("c.setPA(\"3a\") 后 c.MethodA.a = %s (已改变)\n", c.MethodA.a)
	fmt.Printf("cp.MethodA.a = %s (已改变)\n", cp.MethodA.a)
	fmt.Println()

	fmt.Printf("修改前 cp.MethodA.a = %s\n", cp.MethodA.a)
	cp.setPA("4a") // 通过指针调用指针接收者方法
	fmt.Printf("cp.setPA(\"4a\") 后 c.MethodA.a = %s (已改变)\n", c.MethodA.a)
	fmt.Printf("cp.MethodA.a = %s (已改变)\n", cp.MethodA.a)
	fmt.Println()

	// 演示：值接收者方法修改引用类型（切片）
	fmt.Println("--- 4. 值接收者方法修改引用类型（切片）---")
	fmt.Printf("修改前 c.d = %v\n", c.d)
	cp.modityD() // 虽然 modityD 是值接收者，但切片是引用类型，所以可以修改
	fmt.Printf("cp.modityD() 后 c.d = %v (已改变，因为切片是引用类型)\n", c.d)
	fmt.Printf("cp.d = %v (已改变)\n", cp.d)
	fmt.Println()
}

// ========== 代码示例 2：值接收者和指针接收者的详细对比 ==========

// 结构体 A（用于示例 2）
type A2 struct {
	a     string
	bytes [2]byte
}

// 值接收者方法 - 返回字符串
func (a A2) string() string {
	return a.a
}

// 值接收者方法 - 返回字符串（带后缀A）
func (a A2) stringA() string {
	return a.a
}

// 值接收者方法 - 设置字段（无法修改原结构体）
func (a A2) setA(v string) {
	a.a = v // 修改的是副本
}

// 指针接收者方法 - 返回字符串（带后缀PA）
func (a *A2) stringPA() string {
	return a.a
}

// 指针接收者方法 - 设置字段（可以修改原结构体）
func (a *A2) setPA(v string) {
	a.a = v // 修改的是原结构体
}

// 普通函数 - 值传递
func value(a A2, value string) {
	a.a = value // 修改的是副本
}

// 普通函数 - 指针传递
func point(a *A2, value string) {
	a.a = value // 修改的是原结构体
}

// ========== 代码示例 2 的使用演示 ==========

func demonstrateStructMethods2() {
	fmt.Println("=== 代码示例 2：值接收者和指针接收者的详细对比 ===")
	fmt.Println()

	a := A2{
		a: "a",
	}

	// 演示：普通函数的值传递和指针传递
	fmt.Println("--- 1. 普通函数的值传递和指针传递 ---")
	fmt.Printf("修改前 a.a = %s\n", a.a)
	value(a, "any") // 值传递，不会修改原结构体
	fmt.Printf("value(a, \"any\") 后 a.a = %s (未改变)\n", a.a)
	fmt.Println()

	fmt.Printf("修改前 a.a = %s\n", a.a)
	point(&a, "any") // 指针传递，会修改原结构体
	fmt.Printf("point(&a, \"any\") 后 a.a = %s (已改变)\n", a.a)
	fmt.Println()

	pa := &a

	// 演示：结构体变量可以调用值接收者和指针接收者方法
	fmt.Println("--- 2. 结构体变量可以调用值接收者和指针接收者方法 ---")
	fmt.Printf("a.string() = %s (值接收者方法)\n", a.string())
	fmt.Printf("a.stringA() = %s (值接收者方法)\n", a.stringA())
	fmt.Printf("a.stringPA() = %s (指针接收者方法，Go 自动转换)\n", a.stringPA())
	fmt.Println()

	// 演示：结构体指针也可以调用值接收者和指针接收者方法
	fmt.Println("--- 3. 结构体指针也可以调用值接收者和指针接收者方法 ---")
	fmt.Printf("pa.string() = %s (值接收者方法，Go 自动解引用)\n", pa.string())
	fmt.Printf("pa.stringA() = %s (值接收者方法，Go 自动解引用)\n", pa.stringA())
	fmt.Printf("pa.stringPA() = %s (指针接收者方法)\n", pa.stringPA())
	fmt.Println()

	// 演示：值接收者方法无法修改原结构体
	fmt.Println("--- 4. 值接收者方法无法修改原结构体 ---")
	fmt.Printf("修改前 a.a = %s\n", a.a)
	a.setA("new value")
	fmt.Printf("a.setA(\"new value\") 后 a.a = %s (未改变)\n", a.a)
	fmt.Println()

	// 演示：指针接收者方法可以修改原结构体
	fmt.Println("--- 5. 指针接收者方法可以修改原结构体 ---")
	fmt.Printf("修改前 a.a = %s\n", a.a)
	a.setPA("new value")
	fmt.Printf("a.setPA(\"new value\") 后 a.a = %s (已改变)\n", a.a)
	fmt.Println()

	// 总结说明
	fmt.Println("--- 总结 ---")
	fmt.Println("1. 值接收者方法：接收的是结构体的副本，无法修改原结构体的字段（引用类型除外）")
	fmt.Println("2. 指针接收者方法：接收的是结构体的指针，可以修改原结构体的字段")
	fmt.Println("3. Go 会自动处理：")
	fmt.Println("   - 结构体变量调用指针接收者方法时，Go 会自动取地址")
	fmt.Println("   - 结构体指针调用值接收者方法时，Go 会自动解引用")
	fmt.Println("4. 对于引用类型（切片、映射、通道），即使是值接收者，也能修改其内容")
	fmt.Println()
}

// ========== 代码示例 3：方法名称大小写规则演示 ==========

// 示例结构体 - 演示导出和未导出方法
type VisibilityDemo struct {
	// 导出字段（首字母大写）
	PublicField string
	// 未导出字段（首字母小写）
	privateField string
}

// 导出方法（首字母大写）- 其他包可以调用
func (v VisibilityDemo) GetPublicField() string {
	return v.PublicField
}

// 导出方法（首字母大写）- 其他包可以调用
func (v *VisibilityDemo) SetPublicField(value string) {
	v.PublicField = value
}

// 未导出方法（首字母小写）- 只能在同一个包内调用
func (v VisibilityDemo) getPrivateField() string {
	return v.privateField
}

// 未导出方法（首字母小写）- 只能在同一个包内调用
func (v *VisibilityDemo) setPrivateField(value string) {
	v.privateField = value
}

// 导出方法可以调用未导出方法（同一个包内）
func (v VisibilityDemo) GetPrivateFieldViaPublic() string {
	return v.getPrivateField() // 可以在同一个包内调用未导出方法
}

// 演示方法名称大小写规则
func demonstrateMethodVisibility() {
	fmt.Println("=== 代码示例 3：方法名称大小写规则 ===")
	fmt.Println()

	v := VisibilityDemo{
		PublicField:  "公共字段",
		privateField: "私有字段",
	}

	fmt.Println("--- 1. 导出方法（首字母大写）---")
	fmt.Println("导出方法可以在任何包中被调用：")
	fmt.Printf("v.GetPublicField() = %s\n", v.GetPublicField())
	v.SetPublicField("新的公共字段值")
	fmt.Printf("v.SetPublicField() 后 v.PublicField = %s\n", v.PublicField)
	fmt.Println()

	fmt.Println("--- 2. 未导出方法（首字母小写）---")
	fmt.Println("未导出方法只能在同一个包内调用：")
	fmt.Printf("v.getPrivateField() = %s (同一包内可以调用)\n", v.getPrivateField())
	v.setPrivateField("新的私有字段值")
	fmt.Printf("v.setPrivateField() 后 v.privateField = %s (同一包内可以调用)\n", v.privateField)
	fmt.Println()

	fmt.Println("--- 3. 导出方法可以调用未导出方法 ---")
	fmt.Println("在同一个包内，导出方法可以调用未导出方法：")
	fmt.Printf("v.GetPrivateFieldViaPublic() = %s\n", v.GetPrivateFieldViaPublic())
	fmt.Println()

	fmt.Println("--- 4. 字段的可见性规则 ---")
	fmt.Println("字段的可见性规则与方法相同：")
	fmt.Printf("v.PublicField = %s (导出字段，其他包可以访问)\n", v.PublicField)
	fmt.Printf("v.privateField = %s (未导出字段，只能在同一包内访问)\n", v.privateField)
	fmt.Println()

	fmt.Println("--- 总结 ---")
	fmt.Println("1. 首字母大写 = 导出（exported/public），可以被其他包访问")
	fmt.Println("2. 首字母小写 = 未导出（unexported/private），只能在同一个包内访问")
	fmt.Println("3. 这个规则适用于：方法、函数、变量、类型、结构体字段等所有标识符")
	fmt.Println("4. 如果其他包尝试访问未导出的方法或字段，编译时会报错")
	fmt.Println()
	fmt.Println("--- 重要补充：同一个包内不同文件之间的访问 ---")
	fmt.Println("✓ 同一个包（package）内的所有文件共享同一个命名空间")
	fmt.Println("✓ 同一个包内的不同文件可以互相访问未导出的（小写）方法、函数、变量等")
	fmt.Println("✓ 可见性的边界是包（package），不是文件（file）")
	fmt.Println("✓ 详见示例文件：1.5.4.1 cross_file_visibility.go 和 1.5.4.2 cross_file_usage.go")
	fmt.Println()
}

// ========== 主函数 ==========
// 注意：如果与同包其他文件的 main 函数冲突，有以下解决方案：
// 1. 单独运行此文件：go run "1.5.4 struct_methods.go"
// 2. 注释掉其他文件的 main 函数
// 3. 或者将此 main 函数重命名为其他名称（如 mainStructMethods）并手动调用

func Demo() {
	demonstrateStructMethods1()
	fmt.Println()
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()
	demonstrateStructMethods2()
	fmt.Println()
	fmt.Println(strings.Repeat("=", 60))
	fmt.Println()
	demonstrateMethodVisibility()
}

// 如果需要作为主程序运行，取消下面的注释，并注释掉上面的 mainStructMethods 函数
// func main() {
// 	mainStructMethods()
// }
