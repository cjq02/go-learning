package gorm

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// User 用户模型 - 展示 GORM 标签和字段映射
// 基于 fuyelead 项目的 User 模型
type User struct {
	ID           int        `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Username     string     `json:"username" gorm:"column:username;uniqueIndex;not null;size:50"`
	Email        *string    `json:"email,omitempty" gorm:"column:email;uniqueIndex;size:100"`
	PasswordHash string     `json:"-" gorm:"column:password_hash;not null;size:255"` // json:"-" 表示不序列化到 JSON
	FullName     *string    `json:"fullName,omitempty" gorm:"column:full_name;size:100"`
	Status       string     `json:"status" gorm:"column:status;default:'active';size:20"`
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
	LastLoginAt  *time.Time `json:"lastLoginAt,omitempty" gorm:"column:last_login_at"`
}

// TableName 指定表名（fuyelead 项目使用 t_sys_user）
func (User) TableName() string {
	return "t_sys_user"
}

// Order 订单模型 - 展示外键关联
type Order struct {
	ID         int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	OrderNo    string    `json:"orderNo" gorm:"column:order_no;uniqueIndex;not null;size:50"`
	UserID     int       `json:"userID" gorm:"column:user_id;not null;index"`
	TotalPrice float64   `json:"totalPrice" gorm:"column:total_price;type:decimal(10,2);not null"`
	Status     string    `json:"status" gorm:"column:status;default:'pending';size:20"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`

	// 关联关系 - 使用 foreignKey 和 references 定义
	User *User `json:"user,omitempty" gorm:"foreignKey:user_id;references:id"`
}

func (Order) TableName() string {
	return "t_order"
}

// GormBasicsDemo 展示 GORM 基础用法
func GormBasicsDemo() {
	fmt.Println("=== GORM 基础用法示例（基于 fuyelead 项目）===")
	fmt.Println()

	// 1. 数据库连接配置
	fmt.Println("1. 数据库连接配置")
	fmt.Println("   fuyelead 项目使用 MySQL，这里使用 SQLite 内存数据库演示")

	// 配置 GORM logger（开发环境显示 SQL，生产环境静默）
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 开发环境使用 Info，生产环境使用 Silent
	}

	// 连接数据库（使用 SQLite 内存数据库，无需真实数据库）
	db, err := gorm.Open(sqlite.Open(":memory:"), gormConfig)
	if err != nil {
		fmt.Printf("   连接数据库失败: %v\n", err)
		return
	}
	fmt.Println("   ✓ 数据库连接成功")
	fmt.Println()

	// 2. 自动迁移（Auto Migrate）
	fmt.Println("2. 自动迁移表结构")
	fmt.Println("   GORM 会根据模型自动创建/更新表结构")

	err = db.AutoMigrate(&User{}, &Order{})
	if err != nil {
		fmt.Printf("   迁移失败: %v\n", err)
		return
	}
	fmt.Println("   ✓ 表结构迁移完成")
	fmt.Println()

	// 3. 创建记录（Create）
	fmt.Println("3. 创建记录（Create）")

	email := "user@example.com"
	fullName := "张三"
	user := &User{
		Username:     "zhangsan",
		Email:        &email,
		PasswordHash: "hashed_password_123",
		FullName:     &fullName,
		Status:       "active",
	}

	result := db.Create(user)
	if result.Error != nil {
		fmt.Printf("   创建失败: %v\n", result.Error)
		return
	}
	fmt.Printf("   ✓ 创建用户成功，ID: %d, 影响行数: %d\n", user.ID, result.RowsAffected)
	fmt.Println()

	// 4. 查询记录（First, Find）
	fmt.Println("4. 查询记录")

	// 4.1 根据主键查询
	var foundUser User
	db.First(&foundUser, user.ID)
	fmt.Printf("   根据 ID 查询: %s (ID: %d)\n", foundUser.Username, foundUser.ID)

	// 4.2 条件查询
	var userByUsername User
	db.Where("username = ?", "zhangsan").First(&userByUsername)
	fmt.Printf("   根据用户名查询: %s\n", userByUsername.Username)

	// 4.3 查询多条记录
	var users []User
	db.Where("status = ?", "active").Find(&users)
	fmt.Printf("   查询活跃用户数: %d\n", len(users))
	fmt.Println()

	// 5. 更新记录（Update, Save）
	fmt.Println("5. 更新记录")

	// 5.1 使用 Updates 更新多个字段
	updates := map[string]interface{}{
		"status":     "inactive",
		"updated_at": time.Now(),
	}
	db.Model(&foundUser).Updates(updates)
	fmt.Printf("   ✓ 更新用户状态为 inactive\n")

	// 5.2 使用 Save 保存整个模型
	foundUser.Status = "active"
	db.Save(&foundUser)
	fmt.Printf("   ✓ 恢复用户状态为 active\n")
	fmt.Println()

	// 6. 删除记录（Delete）
	fmt.Println("6. 删除记录")

	// 创建测试用户用于删除
	testUser := &User{
		Username:     "test_delete",
		PasswordHash: "hash",
		Status:       "active",
	}
	db.Create(testUser)

	// 软删除（如果模型有 DeletedAt 字段）
	// db.Delete(&testUser)

	// 硬删除
	db.Unscoped().Delete(&testUser)
	fmt.Printf("   ✓ 删除测试用户成功\n")
	fmt.Println()

	// 7. GORM 标签说明
	fmt.Println("7. GORM 标签说明（基于 fuyelead 项目）")
	fmt.Println("   column:id              - 指定数据库列名")
	fmt.Println("   primaryKey             - 主键")
	fmt.Println("   autoIncrement          - 自增")
	fmt.Println("   uniqueIndex            - 唯一索引")
	fmt.Println("   index                  - 普通索引")
	fmt.Println("   not null               - 非空约束")
	fmt.Println("   size:50                - 字符串长度")
	fmt.Println("   default:'active'      - 默认值")
	fmt.Println("   type:decimal(10,2)    - 指定数据类型")
	fmt.Println("   autoCreateTime         - 自动设置创建时间")
	fmt.Println("   autoUpdateTime         - 自动更新修改时间")
	fmt.Println("   foreignKey:user_id     - 外键字段")
	fmt.Println("   references:id          - 引用字段")
	fmt.Println()

	// 8. 错误处理
	fmt.Println("8. 错误处理（fuyelead 项目中的常见模式）")

	var notFoundUser User
	err = db.Where("username = ?", "nonexistent").First(&notFoundUser).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("   ✓ 正确处理记录不存在的情况")
	} else if err != nil {
		fmt.Printf("   其他错误: %v\n", err)
	}
	fmt.Println()

	// 9. 表名自定义
	fmt.Println("9. 表名自定义")
	fmt.Println("   fuyelead 项目使用 TableName() 方法自定义表名")
	fmt.Printf("   User 表名: %s\n", (&User{}).TableName())
	fmt.Printf("   Order 表名: %s\n", (&Order{}).TableName())
	fmt.Println()

	fmt.Println("=== GORM 基础用法示例完成 ===")
}
