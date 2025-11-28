package interfaceexample

import "fmt"

// ========== 1.17.1 接口基本定义 ==========
//
// 在 Go 中接口是一种抽象类型，是一组方法的集合，
// 里面只声明方法，而没有任何数据成员。
//
// 定义一个接口：
// type <interface_name> interface {
//     <method_name>(<method_params>) [<return_type>...]
//     ...
// }

// InterfaceBasicDemo 演示接口基本定义
func InterfaceBasicDemo() {
	fmt.Println("========== 1.17.1 接口基本定义 ==========")
	fmt.Println()
	fmt.Println("在 Go 中接口是一种抽象类型，是一组方法的集合，")
	fmt.Println("里面只声明方法，而没有任何数据成员。")
	fmt.Println()
	fmt.Println("定义一个接口：")
	fmt.Println("  type <interface_name> interface {")
	fmt.Println("      <method_name>(<method_params>) [<return_type>...]")
	fmt.Println("      ...")
	fmt.Println("  }")
	fmt.Println()

	demonstrateInterfaceDefinition()
	demonstrateInterfaceUsage()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 接口是一种抽象类型，只包含方法声明，不包含数据成员")
	fmt.Println("✅ 接口定义格式：type InterfaceName interface { methods... }")
	fmt.Println("✅ 接口中的方法可以没有参数名称，只有类型")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 接口中声明的方法并不要求需要全部公开（可以是小写开头）")
	fmt.Println("   - 接口中的方法参数可以没有名称，只有类型")
	fmt.Println()
}

// Shape 接口定义
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Rectangle 矩形结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Area 计算矩形面积
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Perimeter 计算矩形周长
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// Circle 圆形结构体
type Circle struct {
	Radius float64
}

// Area 计算圆形面积
func (c Circle) Area() float64 {
	return 3.14159 * c.Radius * c.Radius
}

// Perimeter 计算圆形周长
func (c Circle) Perimeter() float64 {
	return 2 * 3.14159 * c.Radius
}

// demonstrateInterfaceDefinition 演示接口定义
func demonstrateInterfaceDefinition() {
	fmt.Println("=== 1.17.1.1 接口定义示例 ===")
	fmt.Println()

	fmt.Println("1. 定义 Shape 接口：")
	fmt.Println("  type Shape interface {")
	fmt.Println("      Area() float64      // 计算面积的方法")
	fmt.Println("      Perimeter() float64 // 计算周长的方法")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("   说明：")
	fmt.Println("   - Shape 接口声明了两个方法：Area() 和 Perimeter()")
	fmt.Println("   - 接口中只声明方法签名，不包含方法实现")
	fmt.Println("   - 接口中的方法没有方法体（大括号）")
	fmt.Println()

	fmt.Println("2. 定义 Rectangle（矩形）结构体：")
	fmt.Println("  type Rectangle struct {")
	fmt.Println("      Width  float64")
	fmt.Println("      Height float64")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("   实现 Shape 接口的方法：")
	fmt.Println("   func (r Rectangle) Area() float64 {")
	fmt.Println("       return r.Width * r.Height")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   func (r Rectangle) Perimeter() float64 {")
	fmt.Println("       return 2 * (r.Width + r.Height)")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("3. 定义 Circle（圆形）结构体：")
	fmt.Println("  type Circle struct {")
	fmt.Println("      Radius float64")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("   实现 Shape 接口的方法：")
	fmt.Println("   func (c Circle) Area() float64 {")
	fmt.Println("       return 3.14159 * c.Radius * c.Radius")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   func (c Circle) Perimeter() float64 {")
	fmt.Println("       return 2 * 3.14159 * c.Radius")
	fmt.Println("   }")
	fmt.Println()

	fmt.Println("4. 接口实现的关键点：")
	fmt.Println("   ✅ Go 中的接口实现是隐式的（implicit）")
	fmt.Println("   ✅ 不需要显式声明 implements 关键字")
	fmt.Println("   ✅ 只要结构体实现了接口中的所有方法，就自动实现了该接口")
	fmt.Println("   ✅ Rectangle 和 Circle 都实现了 Shape 接口的所有方法")
	fmt.Println("   ✅ 因此它们都可以赋值给 Shape 接口类型的变量")
	fmt.Println("   ✅ 方法签名必须完全匹配（方法名、参数、返回值）")
	fmt.Println()
}

// demonstrateInterfaceUsage 演示接口使用
func demonstrateInterfaceUsage() {
	fmt.Println("=== 1.17.1.2 接口使用示例 ===")

	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 5}

	fmt.Printf("矩形: width=%.2f, height=%.2f\n", rect.Width, rect.Height)
	fmt.Printf("  面积: %.2f\n", rect.Area())
	fmt.Printf("  周长: %.2f\n", rect.Perimeter())
	fmt.Println()

	fmt.Printf("圆形: radius=%.2f\n", circle.Radius)
	fmt.Printf("  面积: %.2f\n", circle.Area())
	fmt.Printf("  周长: %.2f\n", circle.Perimeter())
	fmt.Println()

	// 使用接口类型
	var shape1 Shape = rect
	var shape2 Shape = circle

	fmt.Println("使用接口类型变量：")
	fmt.Printf("shape1.Area() = %.2f\n", shape1.Area())
	fmt.Printf("shape2.Area() = %.2f\n", shape2.Area())
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - Rectangle 和 Circle 都实现了 Shape 接口的所有方法")
	fmt.Println("  - 可以将它们赋值给 Shape 接口类型的变量")
	fmt.Println("  - 通过接口可以调用实现的方法")
	fmt.Println()
}
