package constants

import (
	"fmt"
	"strings"
)

// 枚举（Enums）示例
//
// Go 中没有内置枚举类型，所以 Go 中的枚举是使用 const 来定义枚举的。
// 枚举的本质就是一系列的常量。

// ========== 1.6.2.1 定义枚举 ==========

// 1. 基础枚举定义 - 直接使用常量
const (
	Male   = "Male"
	Female = "Female"
)

// 2. 使用类型别名定义枚举 - 更直观
type Gender string

const (
	MaleGender   Gender = "Male"
	FemaleGender Gender = "Female"
)

// 3. 数值型枚举 - 使用iota
type Weekday int

const (
	Sunday    Weekday = iota // 0
	Monday                   // 1
	Tuesday                  // 2
	Wednesday                // 3
	Thursday                 // 4
	Friday                   // 5
	Saturday                 // 6
)

// 4. 自定义起始值的枚举
type Priority int

const (
	Low    Priority = iota + 1 // 1
	Medium                     // 2
	High                       // 3
	Urgent                     // 4
)

// 5. 位标志枚举（用于组合权限等）
type Permission uint8

const (
	Read    Permission = 1 << iota // 1
	Write                          // 2
	Execute                        // 4
	Delete                         // 8
)

// 6. 字符串枚举的另一种写法
type Status string

const (
	Pending   Status = "pending"
	Approved  Status = "approved"
	Rejected  Status = "rejected"
	Cancelled Status = "cancelled"
)

// ========== 枚举方法 ==========

// 为 Gender 类型添加方法
func (g Gender) String() string {
	switch g {
	case MaleGender:
		return "Male"
	case FemaleGender:
		return "Female"
	default:
		return "Unknown"
	}
}

func (g Gender) IsMale() bool {
	return g == MaleGender
}

func (g Gender) IsFemale() bool {
	return g == FemaleGender
}

// 为 Weekday 类型添加方法
func (w Weekday) String() string {
	switch w {
	case Sunday:
		return "Sunday"
	case Monday:
		return "Monday"
	case Tuesday:
		return "Tuesday"
	case Wednesday:
		return "Wednesday"
	case Thursday:
		return "Thursday"
	case Friday:
		return "Friday"
	case Saturday:
		return "Saturday"
	default:
		return "Unknown"
	}
}

func (w Weekday) IsWeekend() bool {
	return w == Sunday || w == Saturday
}

func (w Weekday) IsWeekday() bool {
	return !w.IsWeekend()
}

// 为 Permission 类型添加方法
func (p Permission) Has(perm Permission) bool {
	return p&perm == perm
}

func (p Permission) Add(perm Permission) Permission {
	return p | perm
}

func (p Permission) Remove(perm Permission) Permission {
	return p &^ perm
}

func (p Permission) String() string {
	var perms []string
	if p.Has(Read) {
		perms = append(perms, "Read")
	}
	if p.Has(Write) {
		perms = append(perms, "Write")
	}
	if p.Has(Execute) {
		perms = append(perms, "Execute")
	}
	if p.Has(Delete) {
		perms = append(perms, "Delete")
	}
	if len(perms) == 0 {
		return "None"
	}
	return strings.Join(perms, ", ")
}

// ========== 演示函数 ==========

func demonstrateBasicEnums() {
	fmt.Println("=== 基础枚举定义 ===")

	// 直接使用常量
	fmt.Printf("Male: %s\n", Male)
	fmt.Printf("Female: %s\n", Female)

	// 使用类型别名
	var gender Gender = MaleGender
	fmt.Printf("Gender: %s\n", gender)
	fmt.Printf("IsMale: %v\n", gender.IsMale())
	fmt.Printf("String(): %s\n", gender.String())

	gender = FemaleGender
	fmt.Printf("Gender: %s\n", gender)
	fmt.Printf("IsFemale: %v\n", gender.IsFemale())
	fmt.Printf("String(): %s\n", gender.String())
}

func demonstrateNumericEnums() {
	fmt.Println("\n=== 数值型枚举 ===")

	// iota 自动递增
	fmt.Printf("Sunday: %d - %s\n", Sunday, Sunday.String())
	fmt.Printf("Monday: %d - %s\n", Monday, Monday.String())
	fmt.Printf("Friday: %d - %s\n", Friday, Friday.String())
	fmt.Printf("Saturday: %d - %s\n", Saturday, Saturday.String())

	// 测试方法
	fmt.Printf("Sunday.IsWeekend(): %v\n", Sunday.IsWeekend())
	fmt.Printf("Monday.IsWeekday(): %v\n", Monday.IsWeekday())
	fmt.Printf("Saturday.IsWeekend(): %v\n", Saturday.IsWeekend())
}

func demonstrateCustomEnums() {
	fmt.Println("\n=== 自定义起始值枚举 ===")

	fmt.Printf("Low: %d\n", Low)
	fmt.Printf("Medium: %d\n", Medium)
	fmt.Printf("High: %d\n", High)
	fmt.Printf("Urgent: %d\n", Urgent)
}

func demonstrateBitFlagEnums() {
	fmt.Println("\n=== 位标志枚举 ===")

	// 单个权限
	fmt.Printf("Read: %d\n", Read)
	fmt.Printf("Write: %d\n", Write)
	fmt.Printf("Execute: %d\n", Execute)
	fmt.Printf("Delete: %d\n", Delete)

	// 组合权限
	adminPerm := Read | Write | Execute | Delete
	userPerm := Read | Write
	guestPerm := Read

	fmt.Printf("Admin permissions: %s (value: %d)\n", adminPerm.String(), adminPerm)
	fmt.Printf("User permissions: %s (value: %d)\n", userPerm.String(), userPerm)
	fmt.Printf("Guest permissions: %s (value: %d)\n", guestPerm.String(), guestPerm)

	// 测试权限检查
	fmt.Printf("Admin has Write: %v\n", adminPerm.Has(Write))
	fmt.Printf("User has Execute: %v\n", userPerm.Has(Execute))
	fmt.Printf("Guest has Read: %v\n", guestPerm.Has(Read))

	// 添加权限
	userWithExecute := userPerm.Add(Execute)
	fmt.Printf("User + Execute: %s\n", userWithExecute.String())

	// 移除权限
	adminWithoutDelete := adminPerm.Remove(Delete)
	fmt.Printf("Admin - Delete: %s\n", adminWithoutDelete.String())
}

func demonstrateStringEnums() {
	fmt.Println("\n=== 字符串枚举 ===")

	statuses := []Status{Pending, Approved, Rejected, Cancelled}

	for _, status := range statuses {
		fmt.Printf("Status: %s\n", status)
	}
}

func demonstrateEnumAsParameter(gender Gender) {
	fmt.Printf("传递枚举参数: %s (%s)\n", gender, gender.String())
}

func demonstrateEnumUsage() {
	fmt.Println("\n=== 枚举作为参数传递 ===")

	demonstrateEnumAsParameter(MaleGender)
	demonstrateEnumAsParameter(FemaleGender)
}

// EnumsDemo 枚举演示主函数
func EnumsDemo() {
	fmt.Println("========== 1.6.2 枚举 ==========")
	fmt.Println("Go 中没有内置枚举类型，所以 Go 中的枚举是使用 const 来定义枚举的。")
	fmt.Println("枚举的本质就是一系列的常量。")
	fmt.Println()

	demonstrateBasicEnums()
	demonstrateNumericEnums()
	demonstrateCustomEnums()
	demonstrateBitFlagEnums()
	demonstrateStringEnums()
	demonstrateEnumUsage()

	fmt.Println("\n=== 枚举优势总结 ===")
	fmt.Println("✅ 类型安全：使用类型别名避免传递错误的值")
	fmt.Println("✅ 可读性：枚举值有意义的名字")
	fmt.Println("✅ 方法支持：可以为枚举类型添加方法")
	fmt.Println("✅ IDE支持：代码补全和错误检查")
	fmt.Println("✅ 维护性：集中定义，易于修改和扩展")
}
