package gorm

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// OrderLog 订单日志模型（基于 fuyelead 项目）
type OrderLog struct {
	ID        int       `json:"id" gorm:"column:id;primaryKey;autoIncrement"`
	OrderID   int       `json:"orderID" gorm:"column:order_id;not null;index"`
	OrderNo   string    `json:"orderNo" gorm:"column:order_no;not null;size:50"`
	OldStatus *string   `json:"oldStatus,omitempty" gorm:"column:old_status;size:20"`
	NewStatus string    `json:"newStatus" gorm:"column:new_status;not null;size:20"`
	Action    *string   `json:"action,omitempty" gorm:"column:action;size:50"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`

	Order *OrderWithRelations `json:"order,omitempty" gorm:"foreignKey:order_id;references:id"`
}

func (OrderLog) TableName() string {
	return "t_order_log"
}

// GormQueryOptimizationDemo 展示 GORM 查询优化（基于 fuyelead 项目）
func GormQueryOptimizationDemo() {
	fmt.Println("=== GORM 查询优化示例（基于 fuyelead 项目）===")
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
	err = db.AutoMigrate(&User{}, &OrderWithRelations{}, &OrderLog{})
	if err != nil {
		fmt.Printf("迁移失败: %v\n", err)
		return
	}

	// 创建测试数据
	email := "user@example.com"
	user := &User{
		Username:     "testuser",
		Email:        &email,
		PasswordHash: "hash123",
		Status:       "active",
	}
	db.Create(user)

	// 创建多个订单
	for i := 1; i <= 5; i++ {
		order := &OrderWithRelations{
			OrderNo:    fmt.Sprintf("ORD%03d", i),
			UserID:     user.ID,
			TotalPrice: float64(i * 100),
			Status:     "pending",
		}
		db.Create(order)

		// 为每个订单创建多条日志
		for j := 1; j <= 3; j++ {
			oldStatus := "pending"
			newStatus := "processing"
			if j == 2 {
				newStatus = "completed"
			}
			log := &OrderLog{
				OrderID:   order.ID,
				OrderNo:   order.OrderNo,
				OldStatus: &oldStatus,
				NewStatus: newStatus,
				Action:    stringPtr("status_change"),
			}
			db.Create(log)
		}
	}
	fmt.Println("✓ 测试数据创建完成（5个订单，每个订单3条日志）")
	fmt.Println()

	// 1. 分页查询（fuyelead 项目中的标准模式）
	fmt.Println("1. 分页查询（fuyelead 项目标准模式）")

	limit := 3
	offset := 0

	// 1.1 先查询总数
	var total int64
	db.Model(&OrderWithRelations{}).Count(&total)
	fmt.Printf("   总订单数: %d\n", total)

	// 1.2 分页查询数据
	var orders []OrderWithRelations
	db.Preload("User").
		Order("created_at DESC").
		Limit(limit).
		Offset(offset).
		Find(&orders)

	fmt.Printf("   第 1 页（每页 %d 条）: %d 条记录\n", limit, len(orders))
	for _, o := range orders {
		fmt.Printf("     - 订单: %s, 金额: %.2f\n", o.OrderNo, o.TotalPrice)
	}
	fmt.Println()

	// 2. 条件查询和链式调用
	fmt.Println("2. 条件查询和链式调用")

	var pendingOrders []OrderWithRelations
	query := db.Where("status = ?", "pending")
	query = query.Where("total_price > ?", 200)
	query.Order("created_at DESC").Find(&pendingOrders)

	fmt.Printf("   状态为 pending 且金额 > 200 的订单: %d 条\n", len(pendingOrders))
	fmt.Println()

	// 3. 子查询优化（fuyelead 项目中的高级技巧）
	fmt.Println("3. 子查询优化（获取每个订单的最新日志）")
	fmt.Println("   fuyelead 项目使用子查询避免 N+1 问题")

	var allOrders []OrderWithRelations
	db.Find(&allOrders)

	if len(allOrders) > 0 {
		orderIDs := make([]int, len(allOrders))
		for i, o := range allOrders {
			orderIDs[i] = o.ID
		}

		// 使用子查询获取每个订单的最新日志
		subquery := db.Table("t_order_log").
			Select("order_id, MAX(created_at) as max_created_at").
			Where("order_id IN ?", orderIDs).
			Group("order_id")

		var latestLogs []OrderLog
		db.Table("t_order_log as ol").
			Select("ol.*").
			Joins("INNER JOIN (?) as latest ON ol.order_id = latest.order_id AND ol.created_at = latest.max_created_at", subquery).
			Where("ol.order_id IN ?", orderIDs).
			Scan(&latestLogs)

		fmt.Printf("   使用子查询一次性获取 %d 个订单的最新日志\n", len(latestLogs))
		for _, log := range latestLogs {
			fmt.Printf("     订单 %s 最新状态: %s\n", log.OrderNo, log.NewStatus)
		}
	}
	fmt.Println()

	// 4. 批量查询优化
	fmt.Println("4. 批量查询优化")

	// 4.1 使用 IN 查询
	var users []User
	db.Where("id IN ?", []int{user.ID}).Find(&users)
	fmt.Printf("   批量查询用户: %d 条\n", len(users))

	// 4.2 使用 FindInBatches 处理大量数据
	var batchOrders []OrderWithRelations
	db.FindInBatches(&batchOrders, 2, func(tx *gorm.DB, batch int) error {
		fmt.Printf("   批次 %d: %d 条记录\n", batch, len(batchOrders))
		return nil
	})
	fmt.Println()

	// 5. 查询字段选择（Select）
	fmt.Println("5. 查询字段选择（减少数据传输）")

	var ordersWithSelect []OrderWithRelations
	db.Select("id", "order_no", "total_price", "status").
		Find(&ordersWithSelect)

	fmt.Printf("   只查询必要字段: %d 条记录\n", len(ordersWithSelect))
	for _, o := range ordersWithSelect {
		fmt.Printf("     - %s: %.2f (%s)\n", o.OrderNo, o.TotalPrice, o.Status)
	}
	fmt.Println()

	// 6. 统计查询（Count, Sum, Avg）
	fmt.Println("6. 统计查询")

	var count int64
	db.Model(&OrderWithRelations{}).Count(&count)
	fmt.Printf("   订单总数: %d\n", count)

	var totalAmount float64
	db.Model(&OrderWithRelations{}).Select("SUM(total_price)").Scan(&totalAmount)
	fmt.Printf("   订单总金额: %.2f\n", totalAmount)

	var avgAmount float64
	db.Model(&OrderWithRelations{}).Select("AVG(total_price)").Scan(&avgAmount)
	fmt.Printf("   平均订单金额: %.2f\n", avgAmount)
	fmt.Println()

	// 7. 错误处理（fuyelead 项目中的标准模式）
	fmt.Println("7. 错误处理（fuyelead 项目标准模式）")

	var notFoundOrder OrderWithRelations
	err = db.Where("order_no = ?", "NONEXISTENT").First(&notFoundOrder).Error
	if err == gorm.ErrRecordNotFound {
		fmt.Println("   ✓ 正确处理记录不存在的情况")
	} else if err != nil {
		fmt.Printf("   其他错误: %v\n", err)
	}
	fmt.Println()

	// 8. 事务处理
	fmt.Println("8. 事务处理")

	tx := db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			fmt.Println("   事务回滚")
		}
	}()

	// 在事务中执行多个操作
	newOrder := &OrderWithRelations{
		OrderNo:    "ORD_TX001",
		UserID:     user.ID,
		TotalPrice: 500.00,
		Status:     "pending",
	}
	tx.Create(newOrder)

	log := &OrderLog{
		OrderID:   newOrder.ID,
		OrderNo:   newOrder.OrderNo,
		NewStatus: "pending",
		Action:    stringPtr("created"),
	}
	tx.Create(log)

	tx.Commit()
	fmt.Println("   ✓ 事务提交成功")
	fmt.Println()

	// 9. 查询优化最佳实践总结
	fmt.Println("9. fuyelead 项目查询优化最佳实践")
	fmt.Println("   ✓ 使用 Preload 避免 N+1 查询")
	fmt.Println("   ✓ 使用子查询优化批量关联查询")
	fmt.Println("   ✓ 分页查询时先 Count 再 Find")
	fmt.Println("   ✓ 使用 Select 只查询必要字段")
	fmt.Println("   ✓ 使用索引优化常用查询条件")
	fmt.Println("   ✓ 使用事务保证数据一致性")
	fmt.Println("   ✓ 正确处理 gorm.ErrRecordNotFound")
	fmt.Println()

	fmt.Println("=== GORM 查询优化示例完成 ===")
}

// stringPtr 辅助函数，创建字符串指针
func stringPtr(s string) *string {
	return &s
}
