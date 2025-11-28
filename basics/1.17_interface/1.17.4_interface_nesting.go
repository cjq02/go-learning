package interfaceexample

import "fmt"

// ========== 1.17.4 接口嵌套 ==========
//
// 接口可以嵌套，一个接口可以包含另一个接口的所有方法。
//
// 接口嵌套的语法：
// type InterfaceA interface {
//     MethodA()
// }
//
// type InterfaceB interface {
//     InterfaceA  // 嵌套 InterfaceA
//     MethodB()
// }

// InterfaceNestingDemo 演示接口嵌套
func InterfaceNestingDemo() {
	fmt.Println("========== 1.17.4 接口嵌套 ==========")
	fmt.Println()
	fmt.Println("接口可以嵌套，一个接口可以包含另一个接口的所有方法。")
	fmt.Println()
	fmt.Println("接口嵌套的语法：")
	fmt.Println("  type InterfaceA interface {")
	fmt.Println("      MethodA()")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("  type InterfaceB interface {")
	fmt.Println("      InterfaceA  // 嵌套 InterfaceA")
	fmt.Println("      MethodB()")
	fmt.Println("  }")
	fmt.Println()

	demonstrateInterfaceNesting()
	demonstrateNestedInterfaceUsage()

	fmt.Println("=== 总结 ===")
	fmt.Println("✅ 接口可以嵌套，一个接口可以包含另一个接口的所有方法")
	fmt.Println("✅ 嵌套接口会自动包含被嵌套接口的所有方法")
	fmt.Println("✅ 实现嵌套接口的类型必须实现所有方法（包括嵌套接口的方法）")
	fmt.Println()
	fmt.Println("⚠️ 注意事项：")
	fmt.Println("   - 嵌套接口时，只需要写接口名，不需要写方法")
	fmt.Println("   - 嵌套接口的方法会自动包含在外部接口中")
	fmt.Println()
}

// Reader 读取接口
type Reader interface {
	Read() string
}

// Writer 写入接口
type Writer interface {
	Write(string)
}

// ReadWriter 读写接口（嵌套 Reader 和 Writer）
type ReadWriter interface {
	Reader
	Writer
}

// File 文件结构体实现 ReadWriter 接口
type File struct {
	content string
}

// Read 实现 Reader 接口
func (f *File) Read() string {
	return f.content
}

// Write 实现 Writer 接口
func (f *File) Write(s string) {
	f.content = s
}

// demonstrateInterfaceNesting 演示接口嵌套
func demonstrateInterfaceNesting() {
	fmt.Println("=== 1.17.4.1 接口嵌套示例 ===")

	fmt.Println("定义 Reader 接口：")
	fmt.Println("  type Reader interface {")
	fmt.Println("      Read() string")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("定义 Writer 接口：")
	fmt.Println("  type Writer interface {")
	fmt.Println("      Write(string)")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("定义 ReadWriter 接口（嵌套 Reader 和 Writer）：")
	fmt.Println("  type ReadWriter interface {")
	fmt.Println("      Reader  // 嵌套 Reader")
	fmt.Println("      Writer  // 嵌套 Writer")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - ReadWriter 接口包含了 Reader 和 Writer 的所有方法")
	fmt.Println("  - 实现 ReadWriter 接口的类型必须实现 Read() 和 Write() 方法")
	fmt.Println()
}

// demonstrateNestedInterfaceUsage 演示嵌套接口使用
func demonstrateNestedInterfaceUsage() {
	fmt.Println("=== 1.17.4.2 嵌套接口使用示例 ===")

	file := &File{content: "初始内容"}

	// 可以作为 ReadWriter 使用
	var rw ReadWriter = file
	fmt.Printf("ReadWriter.Read() = %s\n", rw.Read())
	rw.Write("新内容")
	fmt.Printf("ReadWriter.Read() = %s\n", rw.Read())
	fmt.Println()

	// 可以作为 Reader 使用
	var r Reader = file
	fmt.Printf("Reader.Read() = %s\n", r.Read())
	fmt.Println()

	// 可以作为 Writer 使用
	var w Writer = file
	w.Write("通过 Writer 写入")
	fmt.Printf("ReadWriter.Read() = %s\n", rw.Read())
	fmt.Println()

	fmt.Println("说明：")
	fmt.Println("  - File 实现了 ReadWriter 接口")
	fmt.Println("  - 因此 File 也可以作为 Reader 或 Writer 使用")
	fmt.Println("  - 嵌套接口提供了更灵活的类型使用方式")
	fmt.Println()
}

