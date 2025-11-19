package mypackage

// ========== 公开的标识符（首字母大写） ==========

// PublicVar 公开的全局变量 - 可以被其他包访问
var PublicVar = "我是公开的全局变量"

// PublicConst 公开的常量 - 可以被其他包访问
const PublicConst = "我是公开的常量"

// PublicStruct 公开的结构体 - 可以被其他包访问
type PublicStruct struct {
	// PublicField 公开的字段 - 可以被其他包访问
	PublicField string
	// privateField 非公开的字段 - 只能在本包内访问
	privateField int
}

// PublicFunc 公开的函数 - 可以被其他包访问
func PublicFunc() string {
	return "我是公开的函数"
}

// (ps *PublicStruct) PublicMethod 公开的方法 - 可以被其他包访问
func (ps *PublicStruct) PublicMethod() string {
	return "我是公开的方法"
}

// ========== 非公开的标识符（首字母小写） ==========

// privateVar 非公开的全局变量 - 只能在本包内访问
var privateVar = "我是非公开的全局变量"

// privateConst 非公开的常量 - 只能在本包内访问
const privateConst = "我是非公开的常量"

// privateStruct 非公开的结构体 - 只能在本包内访问
type privateStruct struct {
	// field 非公开的字段 - 只能在本包内访问
	field string
}

// privateFunc 非公开的函数 - 只能在本包内访问
func privateFunc() string {
	return "我是非公开的函数"
}

// (ps *privateStruct) privateMethod 非公开的方法 - 只能在本包内访问
func (ps *privateStruct) privateMethod() string {
	return "我是非公开的方法"
}

// ========== 包内部可以访问所有标识符 ==========

// GetPrivateData 公开的函数，用于演示包内部可以访问非公开标识符
func GetPrivateData() string {
	// 在同一个包内，可以访问非公开的标识符
	result := privateVar + ", " + privateConst

	ps := privateStruct{field: "私有字段值"}
	result += ", " + ps.privateMethod()

	result += ", " + privateFunc()

	return result
}

// CreatePublicStruct 创建公开结构体实例
func CreatePublicStruct() PublicStruct {
	return PublicStruct{
		PublicField:  "公开字段值",
		privateField: 42, // 在包内部可以访问私有字段
	}
}

// GetPrivateField 获取私有字段的值（通过包内部的方法）
func (ps *PublicStruct) GetPrivateField() int {
	return ps.privateField // 在包内部可以访问私有字段
}
