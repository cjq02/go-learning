package structs

import "fmt"

// ========== 实际项目中的多文件包示例 ==========
// 在实际项目中，同一个包内包含多个文件是非常常见的做法
// 这样做的好处：
// 1. 代码组织：将相关的功能分组到不同的文件中
// 2. 可维护性：每个文件专注于特定的功能
// 3. 可读性：文件更小，更容易理解和维护
// 4. 团队协作：不同开发者可以同时编辑不同的文件而减少冲突

// ========== 示例：用户管理系统 ==========
// 假设这是一个用户管理系统的包，包含多个文件：
// - user.go: 用户结构体定义和基本方法
// - user_service.go: 用户业务逻辑
// - user_repository.go: 用户数据访问层
// - user_validator.go: 用户数据验证
// 所有这些文件都在同一个 package structs 中

// 注意：在实际项目中，这些通常会放在不同的文件中
// 这里为了演示，我们展示它们如何在同一个包内协作

// ========== 文件 1: user.go (用户结构体定义) ==========
// 在实际项目中，这个文件通常包含：
// - 用户结构体定义
// - 用户的基本方法（Getter/Setter）
// - 用户的基础操作

// AppUser 用户结构体 - 导出的，可以被其他包使用
// 在实际项目中，这个结构体通常定义在 user.go 文件中
type AppUser struct {
	// 导出字段
	ID    int
	Name  string
	Email string
	Age   int
	// 未导出字段 - 内部使用
	password string // 密码不应该被外部直接访问
	status   string // 用户状态（active, inactive, banned等）
}

// NewAppUser 构造函数 - 创建新用户
func NewAppUser(id int, name, email string, age int) *AppUser {
	return &AppUser{
		ID:       id,
		Name:     name,
		Email:    email,
		Age:      age,
		status:   "active", // 默认状态
		password: "",       // 需要单独设置
	}
}

// GetStatus 获取用户状态 - 导出方法
func (u *AppUser) GetStatus() string {
	return u.status
}

// SetPassword 设置密码 - 导出方法，但内部处理
func (u *AppUser) SetPassword(pwd string) {
	u.password = hashPassword(pwd) // 调用同包内的未导出函数
}

// ValidatePassword 验证密码 - 导出方法
func (u *AppUser) ValidatePassword(pwd string) bool {
	return u.password == hashPassword(pwd)
}

// ========== 文件 2: user_service.go (用户业务逻辑) ==========
// 在实际项目中，这个文件通常包含：
// - 用户注册、登录等业务逻辑
// - 用户状态管理
// - 用户权限检查

// AppUserService 用户服务 - 处理用户相关的业务逻辑
type AppUserService struct {
	// 可以包含其他依赖，如数据库连接、缓存等
	users []*AppUser // 简化示例，实际项目中会使用数据库
}

// NewAppUserService 创建用户服务
func NewAppUserService() *AppUserService {
	return &AppUserService{
		users: make([]*AppUser, 0),
	}
}

// RegisterUser 注册用户 - 业务逻辑层
func (s *AppUserService) RegisterUser(name, email, password string, age int) (*AppUser, error) {
	// 调用同包内的验证函数
	if !validateEmail(email) {
		return nil, fmt.Errorf("invalid email format")
	}

	if !validateAge(age) {
		return nil, fmt.Errorf("age must be between 18 and 120")
	}

	// 检查邮箱是否已存在
	if s.emailExists(email) {
		return nil, fmt.Errorf("email already exists")
	}

	// 创建新用户
	user := NewAppUser(len(s.users)+1, name, email, age)
	user.SetPassword(password)

	// 保存用户（实际项目中会保存到数据库）
	s.users = append(s.users, user)

	return user, nil
}

// LoginUser 用户登录
func (s *AppUserService) LoginUser(email, password string) (*AppUser, error) {
	user := s.findUserByEmail(email)
	if user == nil {
		return nil, fmt.Errorf("user not found")
	}

	if !user.ValidatePassword(password) {
		return nil, fmt.Errorf("invalid password")
	}

	if user.GetStatus() != "active" {
		return nil, fmt.Errorf("user account is not active")
	}

	return user, nil
}

// emailExists 检查邮箱是否存在 - 未导出方法，内部使用
func (s *AppUserService) emailExists(email string) bool {
	return s.findUserByEmail(email) != nil
}

// findUserByEmail 根据邮箱查找用户 - 未导出方法，内部使用
func (s *AppUserService) findUserByEmail(email string) *AppUser {
	for _, user := range s.users {
		if user.Email == email {
			return user
		}
	}
	return nil
}

// ========== 文件 3: user_validator.go (数据验证) ==========
// 在实际项目中，这个文件通常包含：
// - 各种数据验证函数
// - 验证规则
// - 验证工具函数

// validateEmail 验证邮箱格式 - 未导出函数，同包内使用
func validateEmail(email string) bool {
	// 简化的邮箱验证
	return len(email) > 0 && contains(email, "@")
}

// validateAge 验证年龄 - 未导出函数，同包内使用
func validateAge(age int) bool {
	return age >= 18 && age <= 120
}

// contains 检查字符串是否包含子串 - 未导出辅助函数
func contains(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

// ========== 文件 4: user_utils.go (工具函数) ==========
// 在实际项目中，这个文件通常包含：
// - 工具函数
// - 辅助函数
// - 通用操作

// hashPassword 密码哈希 - 未导出函数，同包内使用
func hashPassword(password string) string {
	// 简化的哈希，实际项目中应该使用 bcrypt 等安全哈希
	return "hashed_" + password // 仅作演示
}

// ========== 演示：实际项目中的多文件包使用 ==========

func demonstrateRealWorldExample() {
	fmt.Println("=== 实际项目中的多文件包示例 ===")
	fmt.Println()

	fmt.Println("--- 项目结构说明 ---")
	fmt.Println("在实际项目中，同一个包可能包含多个文件：")
	fmt.Println("  user.go          - 用户结构体定义和基本方法")
	fmt.Println("  user_service.go  - 用户业务逻辑")
	fmt.Println("  user_validator.go - 用户数据验证")
	fmt.Println("  user_utils.go    - 工具函数")
	fmt.Println("  所有这些文件都在同一个 package structs 中")
	fmt.Println()

	// 创建用户服务
	service := NewAppUserService()

	fmt.Println("--- 1. 用户注册（调用多个文件中的函数）---")
	user1, err := service.RegisterUser("张三", "zhangsan@example.com", "password123", 25)
	if err != nil {
		fmt.Printf("注册失败: %v\n", err)
	} else {
		fmt.Printf("注册成功: ID=%d, Name=%s, Email=%s\n", user1.ID, user1.Name, user1.Email)
	}
	fmt.Println()

	user2, err := service.RegisterUser("李四", "lisi@example.com", "password456", 30)
	if err != nil {
		fmt.Printf("注册失败: %v\n", err)
	} else {
		fmt.Printf("注册成功: ID=%d, Name=%s, Email=%s\n", user2.ID, user2.Name, user2.Email)
	}
	fmt.Println()

	fmt.Println("--- 2. 用户登录（跨文件调用）---")
	loginUser, err := service.LoginUser("zhangsan@example.com", "password123")
	if err != nil {
		fmt.Printf("登录失败: %v\n", err)
	} else {
		fmt.Printf("登录成功: ID=%d, Name=%s, Status=%s\n", loginUser.ID, loginUser.Name, loginUser.GetStatus())
	}
	fmt.Println()

	fmt.Println("--- 3. 验证失败示例 ---")
	_, err = service.RegisterUser("王五", "invalid-email", "password789", 15)
	if err != nil {
		fmt.Printf("注册失败（预期）: %v\n", err)
	}
	fmt.Println()

	fmt.Println("--- 实际项目中的优势 ---")
	fmt.Println("✓ 代码组织清晰：每个文件专注于特定功能")
	fmt.Println("✓ 易于维护：修改某个功能时只需要关注相关文件")
	fmt.Println("✓ 团队协作：多人可以同时编辑不同文件，减少冲突")
	fmt.Println("✓ 可见性控制：未导出函数只能在同包内使用，保护内部实现")
	fmt.Println("✓ 包内共享：所有文件可以访问彼此的未导出标识符")
	fmt.Println()

	fmt.Println("--- 常见项目结构示例 ---")
	fmt.Println("实际项目中的包结构：")
	fmt.Println("  package user")
	fmt.Println("    ├── user.go          (结构体定义)")
	fmt.Println("    ├── service.go       (业务逻辑)")
	fmt.Println("    ├── repository.go    (数据访问)")
	fmt.Println("    ├── validator.go     (数据验证)")
	fmt.Println("    └── utils.go         (工具函数)")
	fmt.Println()
	fmt.Println("  package order")
	fmt.Println("    ├── order.go")
	fmt.Println("    ├── service.go")
	fmt.Println("    └── repository.go")
	fmt.Println()
	fmt.Println("所有这些文件共享同一个包的命名空间")
	fmt.Println()
}

// 主函数
// 注意：如果与同包其他文件的 main 函数冲突，可以注释掉此函数
func RealWorldExampleDemo() {
	demonstrateRealWorldExample()
}
