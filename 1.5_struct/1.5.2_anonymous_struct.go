package structs

import (
	"encoding/json"
	"fmt"
)

// ========== 1.5.2 定义匿名结构体 ==========

// 匿名结构体是没有定义名称的结构体。
// 匿名结构体无法定义自己的类型方法。

// ========== 方式1：函数外声明匿名结构体 ==========
// 仅可在函数外声明，这种方式可以看成是声明了一个匿名的结构体，实例化后赋值给了的全局变量

var GlobalAnonymous = struct {
	Field1 string
	Field2 int
	Field3 bool
}{}

// ========== 方式2：函数外声明匿名结构体（完整写法）==========
// 包含字段标签和初始化值

var AppConfig = struct {
	Host     string `json:"host" env:"HOST"`
	Port     int    `json:"port" env:"PORT"`
	Debug    bool   `json:"debug" env:"DEBUG"`
	Database struct {
		Driver string `json:"driver"`
		Name   string `json:"name"`
	} `json:"database"`
}{
	Host:  "localhost",
	Port:  8080,
	Debug: false,
	Database: struct {
		Driver string `json:"driver"`
		Name   string `json:"name"`
	}{
		Driver: "postgres",
		Name:   "mydb",
	},
}

// ========== 方式3：在函数或方法中声明匿名结构体并实例化 ==========

func demonstrateAnonymousStructInFunction() {
	// 方式3：在函数中声明匿名结构体并实例化
	user := struct {
		ID       int      `json:"id"`
		Username string   `json:"username"`
		Roles    []string `json:"roles"`
		Profile  struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Bio       string `json:"bio"`
		} `json:"profile"`
	}{
		ID:       1,
		Username: "johndoe",
		Roles:    []string{"admin", "user"},
		Profile: struct {
			FirstName string `json:"first_name"`
			LastName  string `json:"last_name"`
			Bio       string `json:"bio"`
		}{
			FirstName: "John",
			LastName:  "Doe",
			Bio:       "Software Developer",
		},
	}

	fmt.Printf("函数内匿名结构体: %+v\n", user)
}

// ========== 匿名结构体的主要适用场景 ==========

// 场景1：构建测试数据（单元测试）
// 单元测试方法中一般会直接声明一个匿名结构体的切片，通过遍历切片测试方法的各个逻辑分支。
// 示例代码可以参考：go-ethereum/core/type/transaction_test.go 的 TestYParityJSONUnmarshalling 方法。

func demonstrateTestData() {
	fmt.Println("=== 场景1：构建测试数据（单元测试）===")

	// 模拟单元测试中的测试用例数据
	testCases := []struct {
		name     string
		input    int
		expected int
	}{
		{"positive number", 5, 10},
		{"zero", 0, 0},
		{"negative number", -3, -6},
	}

	// 模拟测试函数
	double := func(x int) int {
		return x * 2
	}

	// 运行测试
	for _, tc := range testCases {
		result := double(tc.input)
		status := "PASS"
		if result != tc.expected {
			status = "FAIL"
		}
		fmt.Printf("测试: %s - %s (输入: %d, 期望: %d, 实际: %d)\n",
			tc.name, status, tc.input, tc.expected, result)
	}
	fmt.Println()
}

// 场景2：HTTP处理函数中的JSON序列化和反序列化
// 注意：不推荐这么使用，应该定义一个正式的结构体。
// 优点：相比 map[string]interface{} 无需检查类型、无需检查 key 是否存在并减少相关的代码检查。

func demonstrateHTTPHandling() {
	fmt.Println("=== 场景2：HTTP处理中的JSON处理（不推荐，仅演示）===")

	// 模拟接收到的JSON数据
	jsonInput := `{
		"user_id": 123,
		"action": "update_profile",
		"data": {
			"name": "张三",
			"email": "zhangsan@example.com",
			"preferences": {
				"theme": "dark",
				"language": "zh-CN"
			}
		}
	}`

	// 使用匿名结构体解析JSON（不推荐的方式，仅用于演示）
	// 相比 map[string]interface{} 的优点：
	// 1. 无需检查类型
	// 2. 无需检查 key 是否存在
	// 3. 减少相关的代码检查
	// 但推荐使用命名结构体，详见 anonymous_struct.go 中的说明
	var request = struct {
		UserID int    `json:"user_id"`
		Action string `json:"action"`
		Data   struct {
			Name        string `json:"name"`
			Email       string `json:"email"`
			Preferences struct {
				Theme    string `json:"theme"`
				Language string `json:"language"`
			} `json:"preferences"`
		} `json:"data"`
	}{}

	err := json.Unmarshal([]byte(jsonInput), &request)
	if err != nil {
		fmt.Printf("JSON解析错误: %v\n", err)
		return
	}

	fmt.Printf("解析后的请求: %+v\n", request)

	// 模拟响应数据
	response := struct {
		Success bool   `json:"success"`
		Message string `json:"message"`
		Data    struct {
			UserID int    `json:"user_id"`
			Name   string `json:"name"`
			Email  string `json:"email"`
		} `json:"data"`
	}{
		Success: true,
		Message: "Profile updated successfully",
		Data: struct {
			UserID int    `json:"user_id"`
			Name   string `json:"name"`
			Email  string `json:"email"`
		}{
			UserID: request.UserID,
			Name:   request.Data.Name,
			Email:  request.Data.Email,
		},
	}

	responseJSON, _ := json.Marshal(response)
	fmt.Printf("响应JSON: %s\n", string(responseJSON))
	fmt.Println()
}

// ========== 对比：匿名结构体 vs map[string]interface{} ==========

func demonstrateComparison() {
	fmt.Println("=== 对比：匿名结构体 vs map[string]interface{} ===")

	jsonInput := `{"name": "张三", "age": 25, "email": "zhangsan@example.com"}`

	// 方式1：使用匿名结构体
	fmt.Println("--- 方式1：使用匿名结构体 ---")
	var user1 = struct {
		Name  string `json:"name"`
		Age   int    `json:"age"`
		Email string `json:"email"`
	}{}

	err1 := json.Unmarshal([]byte(jsonInput), &user1)
	if err1 != nil {
		fmt.Printf("解析错误: %v\n", err1)
	} else {
		fmt.Printf("解析成功: %+v\n", user1)
		// 优点：直接访问字段，无需类型断言
		fmt.Printf("姓名: %s, 年龄: %d\n", user1.Name, user1.Age)
	}
	fmt.Println()

	// 方式2：使用 map[string]interface{}
	fmt.Println("--- 方式2：使用 map[string]interface{} ---")
	var user2 map[string]interface{}
	err2 := json.Unmarshal([]byte(jsonInput), &user2)
	if err2 != nil {
		fmt.Printf("解析错误: %v\n", err2)
	} else {
		fmt.Printf("解析成功: %+v\n", user2)
		// 缺点：需要类型断言和检查 key 是否存在
		if name, ok := user2["name"].(string); ok {
			fmt.Printf("姓名: %s\n", name)
		}
		if age, ok := user2["age"].(float64); ok {
			fmt.Printf("年龄: %.0f\n", age)
		}
	}
	fmt.Println()
}

func runAnonymousStructDemo() {
	fmt.Println("========== 1.5.2 定义匿名结构体 ==========")
	fmt.Println()

	// 演示方式1和方式2（全局变量）
	fmt.Println("--- 方式1：函数外声明（空初始化）---")
	fmt.Printf("GlobalAnonymous: %+v\n", GlobalAnonymous)
	fmt.Println()

	fmt.Println("--- 方式2：函数外声明（完整初始化）---")
	fmt.Printf("AppConfig: %+v\n", AppConfig)
	configJSON, _ := json.Marshal(AppConfig)
	fmt.Printf("AppConfig JSON: %s\n", string(configJSON))
	fmt.Println()

	// 演示方式3（函数内）
	fmt.Println("--- 方式3：函数内声明并实例化 ---")
	demonstrateAnonymousStructInFunction()
	fmt.Println()

	// 演示适用场景
	demonstrateTestData()
	demonstrateHTTPHandling()
	demonstrateComparison()

	fmt.Println("=== 总结 ===")
	fmt.Println("匿名结构体的适用场景：")
	fmt.Println("✅ 构建测试数据，单元测试方法中的测试用例")
	fmt.Println("✅ HTTP处理函数中的JSON序列化和反序列化（但不推荐，应使用命名结构体）")
	fmt.Println("✅ 相比 map[string]interface{} 的优势：")
	fmt.Println("   - 无需检查类型")
	fmt.Println("   - 无需检查 key 是否存在")
	fmt.Println("   - 减少相关的代码检查")
	fmt.Println()
	fmt.Println("匿名结构体的限制：")
	fmt.Println("❌ 无法定义自己的类型方法")
	fmt.Println("❌ 无法复用（需要在多处重复定义）")
	fmt.Println("❌ 代码可读性差（特别是嵌套结构）")
}

// 运行方式：
// 方法1：单独运行此文件（需要先临时注释掉同目录下其他文件的 main 函数）
//
//	go run "1.5.2 anonymous_struct.go"
//
// 方法2：从其他 main 函数中调用
//
//	runAnonymousStructDemo()
func AnonymousStructDemo() {
	runAnonymousStructDemo()
}
