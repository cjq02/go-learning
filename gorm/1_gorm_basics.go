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
//
// GORM 标签说明：
// - column: 指定数据库列名（如果字段名和列名不同，必须指定）
// - primaryKey: 标记为主键
// - autoIncrement: 自增字段（通常配合主键使用）
// - uniqueIndex: 创建唯一索引（防止重复数据）
// - index: 创建普通索引（提升查询性能）
// - not null: 非空约束（数据库层面保证数据完整性）
// - size: 字符串字段的最大长度
// - default: 字段的默认值（如果插入时未指定，使用此值）
// - autoCreateTime: 创建记录时自动设置当前时间
// - autoUpdateTime: 更新记录时自动更新为当前时间
// - type: 指定数据库字段类型（如 decimal(10,2) 表示10位数字，2位小数）
//
// JSON 标签说明：
// - json:"-" : 序列化为 JSON 时忽略此字段（常用于敏感信息如密码）
// - json:"fieldName,omitempty": 如果字段为空值，序列化时忽略此字段
type User struct {
	ID           int        `json:"id" gorm:"column:id;primaryKey;autoIncrement"`                 // 主键，自增
	Username     string     `json:"username" gorm:"column:username;uniqueIndex;not null;size:50"` // 用户名，唯一索引，非空，最大50字符
	Email        *string    `json:"email,omitempty" gorm:"column:email;uniqueIndex;size:100"`     // 邮箱，指针类型表示可选，唯一索引
	PasswordHash string     `json:"-" gorm:"column:password_hash;not null;size:255"`              // 密码哈希，不序列化到JSON，非空
	FullName     *string    `json:"fullName,omitempty" gorm:"column:full_name;size:100"`          // 全名，可选字段
	Status       string     `json:"status" gorm:"column:status;default:'active';size:20"`         // 状态，默认值为 'active'
	CreatedAt    time.Time  `json:"createdAt" gorm:"column:created_at;autoCreateTime"`            // 创建时间，自动设置
	UpdatedAt    time.Time  `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`            // 更新时间，自动更新
	LastLoginAt  *time.Time `json:"lastLoginAt,omitempty" gorm:"column:last_login_at"`            // 最后登录时间，可选字段
}

// TableName 指定表名（fuyelead 项目使用 t_sys_user）
//
// GORM 默认使用结构体名的复数形式作为表名（如 User -> users）
// 但 fuyelead 项目使用自定义表名规范：
// - 系统表使用 t_sys_ 前缀
// - 业务表使用 t_ 前缀
// 通过实现 TableName() 方法可以自定义表名
func (User) TableName() string {
	return "t_sys_user"
}

// Order 订单模型 - 展示外键关联
//
// 外键关联说明：
// - foreignKey: 指定当前表中的外键字段名（user_id）
// - references: 指定关联表的主键字段名（User 表的 id）
// - 使用指针类型 *User 表示可选关联（如果不需要加载关联数据，可以为 nil）
// - json:"omitempty" 表示如果关联数据为空，序列化时忽略此字段
type Order struct {
	ID         int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`                     // 主键
	OrderNo    string    `json:"orderNo" gorm:"column:order_no;uniqueIndex;not null;size:50"`      // 订单号，唯一索引
	UserID     int       `json:"userID" gorm:"column:user_id;not null;index"`                      // 用户ID，外键，创建索引提升查询性能
	TotalPrice float64   `json:"totalPrice" gorm:"column:total_price;type:decimal(10,2);not null"` // 总价，使用 decimal 类型保证精度
	Status     string    `json:"status" gorm:"column:status;default:'pending';size:20"`            // 订单状态，默认 pending
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`                // 创建时间
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`                // 更新时间

	// 关联关系 - 使用 foreignKey 和 references 定义
	// foreignKey:user_id 表示 Order 表中的 user_id 字段
	// references:id 表示关联到 User 表的 id 字段
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
	fmt.Println()

	// 配置 GORM logger（开发环境显示 SQL，生产环境静默）
	// logger.Info: 显示所有 SQL 语句和执行时间（开发环境使用，方便调试）
	// logger.Silent: 不显示任何日志（生产环境使用，提升性能）
	// logger.Warn: 只显示警告和错误
	// logger.Error: 只显示错误
	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 开发环境使用 Info，生产环境使用 Silent
	}

	// 连接数据库（使用 SQLite 内存数据库，无需真实数据库）
	// ":memory:" 表示使用内存数据库，程序退出后数据会丢失
	// 实际项目中 fuyelead 使用 MySQL，连接方式类似：
	// db, err := gorm.Open(mysql.Open(dsn), gormConfig)
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
	fmt.Println()
	fmt.Println("   AutoMigrate 的作用：")
	fmt.Println("   - 如果表不存在，自动创建表")
	fmt.Println("   - 如果表存在但缺少字段，自动添加字段")
	fmt.Println("   - 如果字段类型改变，会尝试修改（可能失败，需要手动处理）")
	fmt.Println("   - 不会删除未使用的字段（需要手动删除）")
	fmt.Println()
	fmt.Println("   注意：生产环境建议使用数据库迁移工具（如 golang-migrate）")
	fmt.Println("   而不是 AutoMigrate，因为 AutoMigrate 可能造成数据丢失")
	fmt.Println()

	// 传入要迁移的模型，GORM 会自动处理表结构
	err = db.AutoMigrate(&User{}, &Order{})
	if err != nil {
		fmt.Printf("   迁移失败: %v\n", err)
		return
	}
	fmt.Println("   ✓ 表结构迁移完成")
	fmt.Println()

	// 3. 创建记录（Create）
	fmt.Println("3. 创建记录（Create）")
	fmt.Println()
	fmt.Println("   Create 方法说明：")
	fmt.Println("   - 插入一条新记录到数据库")
	fmt.Println("   - 自动设置 autoCreateTime 和 autoUpdateTime 字段")
	fmt.Println("   - 自动填充自增主键 ID")
	fmt.Println("   - 返回 result 包含错误信息和影响行数")
	fmt.Println()

	// 准备数据
	email := "user@example.com"
	fullName := "张三"
	user := &User{
		Username:     "zhangsan",
		Email:        &email, // 指针类型，可以设置为 nil 表示空值
		PasswordHash: "hashed_password_123",
		FullName:     &fullName,
		Status:       "active",
		// CreatedAt 和 UpdatedAt 会自动设置，不需要手动赋值
	}

	// 执行创建操作
	// result.Error: 如果创建失败，包含错误信息
	// result.RowsAffected: 影响的行数（通常为 1）
	// user.ID: 创建成功后，GORM 会自动填充自增主键的值
	result := db.Create(user)
	if result.Error != nil {
		fmt.Printf("   创建失败: %v\n", result.Error)
		return
	}
	fmt.Printf("   ✓ 创建用户成功，ID: %d, 影响行数: %d\n", user.ID, result.RowsAffected)
	fmt.Println()

	// 4. 查询记录（First, Find）
	fmt.Println("4. 查询记录")
	fmt.Println()
	fmt.Println("   查询方法说明：")
	fmt.Println("   - First: 查询第一条记录，如果没找到返回 ErrRecordNotFound")
	fmt.Println("   - Find: 查询多条记录，返回切片，即使没找到也不会报错（返回空切片）")
	fmt.Println("   - Where: 添加查询条件，支持链式调用")
	fmt.Println("   - 使用 ? 占位符防止 SQL 注入")
	fmt.Println()

	// 4.1 根据主键查询
	// First(&model, id) 会根据主键查询
	// 如果记录不存在，会返回 gorm.ErrRecordNotFound 错误
	var foundUser User
	err = db.First(&foundUser, user.ID).Error
	if err != nil {
		fmt.Printf("   查询失败: %v\n", err)
	} else {
		fmt.Printf("   根据 ID 查询: %s (ID: %d)\n", foundUser.Username, foundUser.ID)
	}

	// 4.2 条件查询
	// Where("username = ?", "zhangsan") 添加条件
	// ? 是占位符，后面的参数会安全地替换占位符（防止 SQL 注入）
	var userByUsername User
	db.Where("username = ?", "zhangsan").First(&userByUsername)
	fmt.Printf("   根据用户名查询: %s\n", userByUsername.Username)

	// 4.3 查询多条记录
	// Find 返回切片，即使没有找到记录也不会报错（返回空切片）
	var users []User
	db.Where("status = ?", "active").Find(&users)
	fmt.Printf("   查询活跃用户数: %d\n", len(users))
	fmt.Println()

	// 5. 更新记录（Update, Save）
	fmt.Println("5. 更新记录")
	fmt.Println()
	fmt.Println("   更新方法说明：")
	fmt.Println("   - Updates: 使用 map 更新多个字段，只更新指定的字段")
	fmt.Println("   - Update: 更新单个字段")
	fmt.Println("   - Save: 保存整个模型，会更新所有字段（包括零值）")
	fmt.Println("   - Model: 指定要更新的模型，可以配合 Where 使用")
	fmt.Println()

	// 5.1 使用 Updates 更新多个字段
	// Updates 使用 map[string]interface{} 可以只更新指定的字段
	// 优点：不会更新未指定的字段，即使字段为零值也不会被更新
	// 注意：UpdatedAt 字段会自动更新（如果设置了 autoUpdateTime）
	updates := map[string]interface{}{
		"status":     "inactive",
		"updated_at": time.Now(), // 虽然设置了 autoUpdateTime，但这里手动设置也可以
	}
	db.Model(&foundUser).Updates(updates)
	fmt.Printf("   ✓ 更新用户状态为 inactive\n")

	// 5.2 使用 Save 保存整个模型
	// Save 会保存模型的所有字段，包括零值
	// 注意：如果字段为零值（如 int 的 0，string 的 ""），也会被更新到数据库
	foundUser.Status = "active"
	db.Save(&foundUser)
	fmt.Printf("   ✓ 恢复用户状态为 active\n")
	fmt.Println()

	// 6. 删除记录（Delete）
	fmt.Println("6. 删除记录")
	fmt.Println()
	fmt.Println("   删除方法说明：")
	fmt.Println("   - Delete: 软删除（如果模型有 gorm.DeletedAt 字段）")
	fmt.Println("   - Unscoped().Delete: 硬删除（物理删除，从数据库彻底删除）")
	fmt.Println("   - 软删除：只是标记为已删除，数据还在数据库中，查询时默认不显示")
	fmt.Println("   - 硬删除：数据从数据库中彻底删除，无法恢复")
	fmt.Println()

	// 创建测试用户用于删除
	testUser := &User{
		Username:     "test_delete",
		PasswordHash: "hash",
		Status:       "active",
	}
	db.Create(testUser)

	// 软删除（如果模型有 DeletedAt 字段）
	// 如果模型包含 gorm.DeletedAt 字段，Delete 会执行软删除
	// 软删除后，使用 Find/First 查询时默认不会查询到已删除的记录
	// 需要使用 Unscoped().Find 才能查询到软删除的记录
	// db.Delete(&testUser)

	// 硬删除
	// Unscoped() 表示忽略软删除标记，执行物理删除
	// 注意：硬删除的数据无法恢复，生产环境要谨慎使用
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
	fmt.Println()
	fmt.Println("   错误处理最佳实践：")
	fmt.Println("   - First 方法如果没找到记录，会返回 gorm.ErrRecordNotFound")
	fmt.Println("   - Find 方法即使没找到也不会报错（返回空切片）")
	fmt.Println("   - 使用 errors.Is 判断错误类型（推荐方式）")
	fmt.Println("   - 区分记录不存在和其他数据库错误")
	fmt.Println()

	var notFoundUser User
	// 查询不存在的记录
	err = db.Where("username = ?", "nonexistent").First(&notFoundUser).Error
	if err == gorm.ErrRecordNotFound {
		// 记录不存在是正常情况，不是错误
		fmt.Println("   ✓ 正确处理记录不存在的情况")
	} else if err != nil {
		// 其他数据库错误（如连接失败、SQL 语法错误等）
		fmt.Printf("   其他错误: %v\n", err)
	}
	fmt.Println()
	fmt.Println("   fuyelead 项目中的标准错误处理模式：")
	fmt.Println("   ```go")
	fmt.Println("   if err := db.First(&user, id).Error; err != nil {")
	fmt.Println("       if errors.Is(err, gorm.ErrRecordNotFound) {")
	fmt.Println("           return nil, errors.New(\"用户不存在\")")
	fmt.Println("       }")
	fmt.Println("       return nil, err")
	fmt.Println("   }")
	fmt.Println("   ```")
	fmt.Println()

	// 9. 表名自定义
	fmt.Println("9. 表名自定义")
	fmt.Println("   fuyelead 项目使用 TableName() 方法自定义表名")
	fmt.Printf("   User 表名: %s\n", (&User{}).TableName())
	fmt.Printf("   Order 表名: %s\n", (&Order{}).TableName())
	fmt.Println()

	fmt.Println("=== GORM 基础用法示例完成 ===")
}
