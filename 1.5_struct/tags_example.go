package structs

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// ========== 结构体字段标记（Tags）示例 ==========

// 方式1：基础结构体定义
type Custom struct {
	field1, field2, field3 byte
}

// 方式2：带有字段标记的结构体
type Person struct {
	Name  string            `json:"name" gorm:"column:name"`
	Age   int               `json:"age" gorm:"column:age"`
	Call  func()            `json:"-" gorm:"column:call"`
	Map   map[string]string `json:"map" gorm:"column:map_data"`
	Ch    chan string       `json:"-" gorm:"column:channel"`
	Arr   [32]uint8         `json:"arr" gorm:"column:array_data"`
	Slice []interface{}     `json:"slice" gorm:"column:slice_data"`
	Ptr   *int              `json:"-"`
	Other Other             `json:"-"`
}

// 其他结构体类型
type Other struct {
	HiddenField string `json:"-"`
}

// ========== 字段标记的使用示例 ==========

func demonstrateStructTags() {
	fmt.Println("=== 结构体字段标记（Tags）演示 ===")
	fmt.Println()

	// 创建Person实例
	person := Person{
		Name:  "张三",
		Age:   25,
		Map:   map[string]string{"key": "value"},
		Arr:   [32]uint8{1, 2, 3},
		Slice: []interface{}{"hello", 42, true},
	}

	// 1. JSON序列化 - 使用json标签
	fmt.Println("--- 1. JSON序列化（使用json标签）---")
	jsonData, err := json.Marshal(person)
	if err != nil {
		fmt.Printf("JSON序列化错误: %v\n", err)
	} else {
		fmt.Printf("JSON输出: %s\n", string(jsonData))
	}
	fmt.Println()

	// 2. 通过反射获取字段标记
	fmt.Println("--- 2. 通过反射获取字段标记 ---")
	demonstrateFieldTags()
	fmt.Println()

	// 3. 字段标记的实际应用场景
	fmt.Println("--- 3. 字段标记的应用场景 ---")
	fmt.Println("JSON标签：用于JSON序列化/反序列化")
	fmt.Println("GORM标签：用于数据库ORM映射")
	fmt.Println("验证标签：用于数据验证")
	fmt.Println("绑定标签：用于HTTP请求绑定")
	fmt.Println("忽略标签：使用 '-' 来忽略字段")
	fmt.Println()
}

// 通过反射演示字段标记
func demonstrateFieldTags() {
	personType := reflect.TypeOf(Person{})

	fmt.Println("Person结构体的字段信息：")
	for i := 0; i < personType.NumField(); i++ {
		field := personType.Field(i)

		fmt.Printf("字段名: %s\n", field.Name)
		fmt.Printf("  类型: %s\n", field.Type)

		// 获取json标签
		jsonTag := field.Tag.Get("json")
		if jsonTag != "" {
			fmt.Printf("  JSON标签: %s\n", jsonTag)
		}

		// 获取gorm标签
		gormTag := field.Tag.Get("gorm")
		if gormTag != "" {
			fmt.Printf("  GORM标签: %s\n", gormTag)
		}

		// 显示所有标签
		if field.Tag != "" {
			fmt.Printf("  所有标签: %s\n", field.Tag)
		}

		fmt.Println()
	}
}

// ========== 其他结构体标记示例 ==========

// 数据库相关的结构体
type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"uniqueIndex;size:100" validate:"required,min=3,max=50"`
	Email    string `json:"email" gorm:"uniqueIndex" validate:"required,email"`
	Password string `json:"-" gorm:"size:255" validate:"required,min=8"`
	Created  string `json:"created_at" gorm:"autoCreateTime"`
	Updated  string `json:"updated_at" gorm:"autoUpdateTime"`
}

// HTTP请求绑定的结构体
type LoginRequest struct {
	Username string `json:"username" binding:"required" form:"username"`
	Password string `json:"password" binding:"required,min=8" form:"password"`
	Remember bool   `json:"remember" form:"remember"`
}

// 配置相关的结构体
type Config struct {
	Host     string `yaml:"host" env:"HOST" default:"localhost"`
	Port     int    `yaml:"port" env:"PORT" default:"8080"`
	Debug    bool   `yaml:"debug" env:"DEBUG" default:"false"`
	Database struct {
		Driver   string `yaml:"driver" env:"DB_DRIVER" default:"postgres"`
		Host     string `yaml:"host" env:"DB_HOST" default:"localhost"`
		Port     int    `yaml:"port" env:"DB_PORT" default:"5432"`
		Name     string `yaml:"name" env:"DB_NAME" default:"mydb"`
		Username string `yaml:"username" env:"DB_USERNAME"`
		Password string `yaml:"password" env:"DB_PASSWORD"`
	} `yaml:"database"`
}

func Demo() {
	demonstrateStructTags()

	// 演示其他标记示例
	fmt.Println("=== 其他结构体标记示例 ===")
	fmt.Println()

	fmt.Println("--- 数据库模型示例 ---")
	userType := reflect.TypeOf(User{})
	fmt.Printf("User结构体有 %d 个字段\n", userType.NumField())
	for i := 0; i < userType.NumField(); i++ {
		field := userType.Field(i)
		fmt.Printf("  %s: %s\n", field.Name, field.Tag)
	}
	fmt.Println()

	fmt.Println("--- HTTP请求绑定示例 ---")
	loginType := reflect.TypeOf(LoginRequest{})
	fmt.Printf("LoginRequest结构体有 %d 个字段\n", loginType.NumField())
	for i := 0; i < loginType.NumField(); i++ {
		field := loginType.Field(i)
		fmt.Printf("  %s: %s\n", field.Name, field.Tag)
	}
	fmt.Println()

	fmt.Println("--- 配置示例 ---")
	configType := reflect.TypeOf(Config{})
	fmt.Printf("Config结构体有 %d 个字段\n", configType.NumField())
	for i := 0; i < configType.NumField(); i++ {
		field := configType.Field(i)
		fmt.Printf("  %s: %s\n", field.Name, field.Tag)
	}
}
