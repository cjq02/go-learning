package gorm

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// OfferingCategory 服务分类模型（基于 fuyelead 项目）
type OfferingCategory struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	Name      string    `json:"name" gorm:"column:category_name;not null;size:100"`
	Sequence  int       `json:"sequence" gorm:"column:sequence;default:0"`
	Status    string    `json:"status" gorm:"column:status;default:'active';size:20"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`

	// 一对多关联：一个分类有多个服务
	Offerings []Offering `json:"offerings,omitempty" gorm:"foreignKey:category_id;references:id"`
}

func (OfferingCategory) TableName() string {
	return "t_order_offering_cate"
}

// Offering 服务模型（基于 fuyelead 项目）
type Offering struct {
	ID         int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	CategoryID int       `json:"categoryID" gorm:"column:category_id;not null;index"`
	Name       string    `json:"name" gorm:"column:offering_name;not null;size:100"`
	UnitPrice  float64   `json:"unitPrice" gorm:"column:unit_price;type:decimal(10,2);not null"`
	Status     string    `json:"status" gorm:"column:status;default:'active';size:20"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`

	// 多对一关联：多个服务属于一个分类
	Category *OfferingCategory `json:"category,omitempty" gorm:"foreignKey:category_id;references:id"`
}

func (Offering) TableName() string {
	return "t_order_offering"
}

// OrderWithRelations 带关联的订单模型
type OrderWithRelations struct {
	ID         int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	OrderNo    string    `json:"orderNo" gorm:"column:order_no;uniqueIndex;not null;size:50"`
	UserID     int       `json:"userID" gorm:"column:user_id;not null;index"`
	OfferingID int       `json:"offeringID" gorm:"column:offering_id;not null;index"`
	TotalPrice float64   `json:"totalPrice" gorm:"column:total_price;type:decimal(10,2);not null"`
	Status     string    `json:"status" gorm:"column:status;default:'pending';size:20"`
	CreatedAt  time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`

	// 多个关联关系
	User     *User     `json:"user,omitempty" gorm:"foreignKey:user_id;references:id"`
	Offering *Offering `json:"offering,omitempty" gorm:"foreignKey:offering_id;references:id"`
}

func (OrderWithRelations) TableName() string {
	return "t_order_with_relations"
}

// GormRelationshipsDemo 展示 GORM 关联查询（基于 fuyelead 项目）
func GormRelationshipsDemo() {
	fmt.Println("=== GORM 关联查询示例（基于 fuyelead 项目）===")
	fmt.Println()

	// 连接数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("连接数据库失败: %v\n", err)
		return
	}

	// 自动迁移
	err = db.AutoMigrate(&User{}, &OfferingCategory{}, &Offering{}, &OrderWithRelations{})
	if err != nil {
		fmt.Printf("迁移失败: %v\n", err)
		return
	}
	fmt.Println("✓ 数据库表创建成功")
	fmt.Println()

	// 1. 创建测试数据
	fmt.Println("1. 创建测试数据")

	// 创建用户
	email := "user@example.com"
	user := &User{
		Username:     "testuser",
		Email:        &email,
		PasswordHash: "hash123",
		Status:       "active",
	}
	db.Create(user)
	fmt.Printf("   ✓ 创建用户: %s (ID: %d)\n", user.Username, user.ID)

	// 创建分类
	category := &OfferingCategory{
		Name:     "基础服务",
		Sequence: 1,
		Status:   "active",
	}
	db.Create(category)
	fmt.Printf("   ✓ 创建分类: %s (ID: %d)\n", category.Name, category.ID)

	// 创建服务
	offering := &Offering{
		CategoryID: category.ID,
		Name:       "测试服务",
		UnitPrice:  99.99,
		Status:     "active",
	}
	db.Create(offering)
	fmt.Printf("   ✓ 创建服务: %s (ID: %d)\n", offering.Name, offering.ID)

	// 创建订单
	order := &OrderWithRelations{
		OrderNo:    "ORD202312251430",
		UserID:     user.ID,
		OfferingID: offering.ID,
		TotalPrice: 99.99,
		Status:     "pending",
	}
	db.Create(order)
	fmt.Printf("   ✓ 创建订单: %s (ID: %d)\n", order.OrderNo, order.ID)
	fmt.Println()

	// 2. Preload 预加载关联（fuyelead 项目中的核心用法）
	fmt.Println("2. Preload 预加载关联")
	fmt.Println("   Preload 可以避免 N+1 查询问题，一次性加载所有关联数据")
	fmt.Println()
	fmt.Println("   N+1 问题说明：")
	fmt.Println("   - 不使用 Preload：查询 N 条订单，然后循环查询每个订单的用户")
	fmt.Println("   - 结果：1 次查询订单 + N 次查询用户 = N+1 次查询")
	fmt.Println("   - 使用 Preload：1 次查询订单 + 1 次批量查询所有用户 = 2 次查询")
	fmt.Println("   - 性能提升：从 N+1 次查询减少到 2 次查询")
	fmt.Println()

	// 2.1 预加载单个关联
	fmt.Println("   2.1 预加载单个关联（订单 -> 用户）")
	fmt.Println("   语法：Preload(\"关联字段名\")")
	fmt.Println("   执行流程：")
	fmt.Println("   1. 查询订单数据")
	fmt.Println("   2. 收集所有 UserID")
	fmt.Println("   3. 一次性查询所有相关用户")
	fmt.Println("   4. 自动关联到对应的订单")
	fmt.Println()
	var orderWithUser OrderWithRelations
	// Preload("User") 表示预加载 Order 模型中的 User 关联字段
	// GORM 会自动根据 foreignKey 和 references 配置进行关联
	db.Preload("User").First(&orderWithUser, order.ID)
	if orderWithUser.User != nil {
		fmt.Printf("      订单 %s 的用户: %s\n", orderWithUser.OrderNo, orderWithUser.User.Username)
	}
	fmt.Println()

	// 2.2 预加载多个关联
	fmt.Println("   2.2 预加载多个关联（订单 -> 用户 + 服务）")
	fmt.Println("   可以链式调用多个 Preload，一次性加载所有关联")
	fmt.Println("   执行流程：")
	fmt.Println("   1. 查询订单")
	fmt.Println("   2. 批量查询所有用户")
	fmt.Println("   3. 批量查询所有服务")
	fmt.Println("   4. 自动关联到对应的订单")
	fmt.Println()
	var orderWithAll OrderWithRelations
	// 链式调用多个 Preload，GORM 会优化查询，不会产生 N+1 问题
	db.Preload("User").Preload("Offering").First(&orderWithAll, order.ID)
	if orderWithAll.User != nil && orderWithAll.Offering != nil {
		fmt.Printf("      订单: %s\n", orderWithAll.OrderNo)
		fmt.Printf("      用户: %s\n", orderWithAll.User.Username)
		fmt.Printf("      服务: %s (价格: %.2f)\n", orderWithAll.Offering.Name, orderWithAll.Offering.UnitPrice)
	}
	fmt.Println()

	// 2.3 预加载嵌套关联（fuyelead 项目中的常见模式）
	fmt.Println("   2.3 预加载嵌套关联（订单 -> 服务 -> 分类）")
	fmt.Println("   使用点号(.)连接多层关联，一次性加载所有层级")
	fmt.Println("   语法：Preload(\"一级关联.二级关联\")")
	fmt.Println("   执行流程：")
	fmt.Println("   1. 查询订单")
	fmt.Println("   2. 批量查询服务（订单的关联）")
	fmt.Println("   3. 批量查询分类（服务的关联）")
	fmt.Println("   4. 自动关联到对应的服务，再关联到订单")
	fmt.Println()
	var orderWithNested OrderWithRelations
	// Preload("Offering.Category") 表示：
	// 1. 先预加载 Order 的 Offering 关联
	// 2. 再预加载 Offering 的 Category 关联
	// 这样一次性加载三层关联，避免多次查询
	db.Preload("Offering.Category").First(&orderWithNested, order.ID)
	if orderWithNested.Offering != nil && orderWithNested.Offering.Category != nil {
		fmt.Printf("      订单: %s\n", orderWithNested.OrderNo)
		fmt.Printf("      服务: %s\n", orderWithNested.Offering.Name)
		fmt.Printf("      分类: %s\n", orderWithNested.Offering.Category.Name)
	}
	fmt.Println()

	// 2.4 预加载时添加条件（fuyelead 项目中的优化技巧）
	fmt.Println("   2.4 预加载时添加条件（只加载活跃的服务）")
	fmt.Println("   在 Preload 中添加 WHERE 条件，只加载符合条件的关联数据")
	fmt.Println("   语法：Preload(\"关联字段\", \"条件\", 参数)")
	fmt.Println("   优点：减少不必要的数据加载，提升性能")
	fmt.Println()
	var orderWithCondition OrderWithRelations
	// Preload("Offering", "status = ?", "active") 表示：
	// 只预加载状态为 "active" 的服务
	// 如果服务状态不是 "active"，则不会被加载（关联字段为 nil）
	db.Preload("Offering", "status = ?", "active").First(&orderWithCondition, order.ID)
	if orderWithCondition.Offering != nil {
		fmt.Printf("      订单的服务状态: %s\n", orderWithCondition.Offering.Status)
	}
	fmt.Println()

	// 2.5 使用函数预加载（更灵活的条件）
	fmt.Println("   2.5 使用函数预加载（fuyelead 项目中的高级用法）")
	fmt.Println("   使用函数可以添加更复杂的查询条件，如排序、多条件等")
	fmt.Println("   语法：Preload(\"关联字段\", func(db *gorm.DB) *gorm.DB { ... })")
	fmt.Println("   优点：可以链式调用 Where、Order、Limit 等方法")
	fmt.Println()
	var orderWithFunc OrderWithRelations
	// 使用函数预加载，可以添加复杂的查询条件
	// 函数接收 *gorm.DB，返回配置好的 *gorm.DB
	// 可以在函数中使用 Where、Order、Limit 等方法
	db.Preload("Offering", func(db *gorm.DB) *gorm.DB {
		// 只加载状态为 active 的服务，并按 ID 升序排序
		return db.Where("status = ?", "active").Order("id ASC")
	}).First(&orderWithFunc, order.ID)
	fmt.Printf("      使用函数预加载，可以添加排序和复杂条件\n")
	fmt.Println()

	// 3. 一对多关联查询
	fmt.Println("3. 一对多关联查询（分类 -> 服务列表）")
	fmt.Println()
	fmt.Println("   一对多关联说明：")
	fmt.Println("   - 一个分类可以有多个服务（一对多）")
	fmt.Println("   - 使用 []Offering 切片类型表示多个关联")
	fmt.Println("   - Preload 会一次性加载该分类下的所有服务")
	fmt.Println("   - 可以在 Preload 中添加条件，只加载符合条件的服务")
	fmt.Println()

	var categoryWithOfferings OfferingCategory
	// Preload("Offerings") 会加载该分类下的所有服务
	// 使用函数添加条件：只加载状态为 active 的服务，并按 ID 排序
	db.Preload("Offerings", func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", "active").Order("id ASC")
	}).First(&categoryWithOfferings, category.ID)

	fmt.Printf("   分类: %s\n", categoryWithOfferings.Name)
	fmt.Printf("   服务数量: %d\n", len(categoryWithOfferings.Offerings))
	for _, o := range categoryWithOfferings.Offerings {
		fmt.Printf("     - %s (价格: %.2f)\n", o.Name, o.UnitPrice)
	}
	fmt.Println()

	// 4. 查询列表时预加载（fuyelead 项目中的分页查询模式）
	fmt.Println("4. 查询列表时预加载（分页查询）")
	fmt.Println()
	fmt.Println("   分页查询最佳实践：")
	fmt.Println("   - 查询列表时，必须使用 Preload 预加载关联数据")
	fmt.Println("   - 避免在循环中查询关联数据（会产生 N+1 问题）")
	fmt.Println("   - 可以预加载多层嵌套关联")
	fmt.Println("   - 配合 Limit 和 Offset 实现分页")
	fmt.Println()

	var orders []OrderWithRelations
	// 查询订单列表时，同时预加载用户和服务分类信息
	// 执行流程：
	// 1. 查询订单列表（带分页）
	// 2. 收集所有 UserID 和 OfferingID
	// 3. 批量查询所有用户
	// 4. 批量查询所有服务及其分类
	// 5. 自动关联到对应的订单
	db.Preload("User").Preload("Offering.Category").
		Order("created_at DESC").
		Limit(10).
		Find(&orders)

	fmt.Printf("   查询到 %d 个订单\n", len(orders))
	for _, o := range orders {
		if o.User != nil && o.Offering != nil {
			fmt.Printf("     订单 %s: 用户=%s, 服务=%s\n", o.OrderNo, o.User.Username, o.Offering.Name)
		}
	}
	fmt.Println()

	// 5. 关联查询的性能优化说明
	fmt.Println("5. 关联查询性能优化（fuyelead 项目中的最佳实践）")
	fmt.Println("   ✓ 使用 Preload 避免 N+1 查询问题")
	fmt.Println("   ✓ 在 Preload 中添加条件，减少不必要的数据加载")
	fmt.Println("   ✓ 使用嵌套 Preload 一次性加载多层关联")
	fmt.Println("   ✓ 分页查询时，先查询主表，再批量预加载关联")
	fmt.Println()

	fmt.Println("=== GORM 关联查询示例完成 ===")
}
