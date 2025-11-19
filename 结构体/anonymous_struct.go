package main

import (
	"encoding/json"
	"fmt"
)

// ========== 匿名结构体示例 ==========

// 方式1：函数外声明匿名结构体，赋值给全局变量
var GlobalAnonymous = struct {
	Field1 string
	Field2 int
	Field3 bool
}{}

// 方式2：函数外声明匿名结构体（完整写法），包含字段标签和初始化值
var Config = struct {
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

// ========== 匿名结构体的使用示例 ==========

func demonstrateAnonymousStruct() {
	fmt.Println("=== 匿名结构体演示 ===")
	fmt.Println()

	// 方式1：基础匿名结构体
	fmt.Println("--- 方式1：基础匿名结构体 ---")
	anonymous1 := struct {
		Name string
		Age  int
		City string
	}{}
	fmt.Printf("空匿名结构体: %+v\n", anonymous1)

	// 初始化匿名结构体
	anonymous1.Name = "张三"
	anonymous1.Age = 25
	anonymous1.City = "北京"
	fmt.Printf("初始化后: %+v\n", anonymous1)
	fmt.Println()

	// 方式2：声明时直接初始化
	fmt.Println("--- 方式2：声明时直接初始化 ---")
	anonymous2 := struct {
		Name    string `json:"name"`
		Age     int    `json:"age"`
		Email   string `json:"email"`
		Active  bool   `json:"active"`
		Address struct {
			Street  string `json:"street"`
			City    string `json:"city"`
			Country string `json:"country"`
		} `json:"address"`
	}{
		Name:   "李四",
		Age:    30,
		Email:  "lisi@example.com",
		Active: true,
		Address: struct {
			Street  string `json:"street"`
			City    string `json:"city"`
			Country string `json:"country"`
		}{
			Street:  "123 Main St",
			City:    "上海",
			Country: "中国",
		},
	}
	fmt.Printf("完整初始化: %+v\n", anonymous2)
	fmt.Println()

	// 方式3：在函数中声明和使用
	fmt.Println("--- 方式3：在函数中声明和使用 ---")
	demonstrateInFunction()
	fmt.Println()

	// JSON序列化示例
	fmt.Println("--- JSON序列化示例 ---")
	jsonData, err := json.Marshal(anonymous2)
	if err != nil {
		fmt.Printf("JSON序列化错误: %v\n", err)
	} else {
		fmt.Printf("JSON输出: %s\n", string(jsonData))
	}
	fmt.Println()
}

// 方式3：在函数中声明匿名结构体
func demonstrateInFunction() {
	// 声明并初始化匿名结构体
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

// ========== 实际应用场景 ==========

// 场景1：构建测试数据（单元测试）
func demonstrateTestData() {
	fmt.Println("=== 场景1：构建测试数据 ===")

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

// 场景2：HTTP处理函数中的JSON处理（演示，不推荐在实际项目中使用）
func demonstrateHTTPHandling() {
	fmt.Println("=== 场景2：HTTP处理中的JSON处理 ===")

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

// ========== 对比：使用命名结构体 vs 匿名结构体 ==========

// 命名结构体（推荐的方式）
type UserProfile struct {
	ID          int                    `json:"id"`
	Name        string                 `json:"name"`
	Email       string                 `json:"email"`
	Preferences map[string]interface{} `json:"preferences"`
}

// 使用命名结构体的函数
func processUserProfile(profile UserProfile) {
	fmt.Printf("处理用户资料: %+v\n", profile)
}

func demonstrateComparison() {
	fmt.Println("=== 对比：命名结构体 vs 匿名结构体 ===")

	// 使用命名结构体（推荐）
	userProfile := UserProfile{
		ID:    1,
		Name:  "王五",
		Email: "wangwu@example.com",
		Preferences: map[string]interface{}{
			"theme":    "light",
			"language": "zh-CN",
		},
	}

	fmt.Println("--- 使用命名结构体（推荐）---")
	processUserProfile(userProfile)

	// 使用匿名结构体（不推荐用于复杂数据）
	anonymousUser := struct {
		ID          int                    `json:"id"`
		Name        string                 `json:"name"`
		Email       string                 `json:"email"`
		Preferences map[string]interface{} `json:"preferences"`
	}{
		ID:    2,
		Name:  "赵六",
		Email: "zhaoliu@example.com",
		Preferences: map[string]interface{}{
			"theme":    "dark",
			"language": "en-US",
		},
	}

	fmt.Println("--- 使用匿名结构体（仅用于简单场景）---")
	fmt.Printf("匿名结构体用户: %+v\n", anonymousUser)
	fmt.Println()
}

func main() {
	demonstrateAnonymousStruct()
	demonstrateTestData()
	demonstrateHTTPHandling()
	demonstrateComparison()

	fmt.Println("=== 总结 ===")
	fmt.Println("匿名结构体适用场景：")
	fmt.Println("✅ 单元测试中的测试数据构建")
	fmt.Println("✅ 简单的临时数据结构")
	fmt.Println("✅ 一次性使用的配置")
	fmt.Println("❌ HTTP处理中的复杂JSON数据（应使用命名结构体）")
	fmt.Println("❌ 需要定义方法的数据结构")
	fmt.Println("❌ 会在多个地方复用的数据结构")
}
