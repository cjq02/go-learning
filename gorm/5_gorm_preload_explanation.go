package gorm

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// GormPreloadExplanationDemo 详细解释预加载的概念
func GormPreloadExplanationDemo() {
	fmt.Println("=== GORM 预加载（Preload）详解 ===")
	fmt.Println()

	// 连接数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Printf("连接失败: %v\n", err)
		return
	}

	// 创建表
	db.AutoMigrate(&User{}, &OrderWithRelations{})

	// 创建测试数据：1个用户，3个订单
	email := "user@example.com"
	user := &User{
		Username:     "testuser",
		Email:        &email,
		PasswordHash: "hash123",
		Status:       "active",
	}
	db.Create(user)

	// 创建3个订单
	for i := 1; i <= 3; i++ {
		order := &OrderWithRelations{
			OrderNo:    fmt.Sprintf("ORD%03d", i),
			UserID:     user.ID,
			TotalPrice: float64(i * 100),
			Status:     "pending",
		}
		db.Create(order)
	}

	fmt.Println("📊 测试场景：查询 3 个订单及其用户信息")
	fmt.Println()

	// ========== 方式1：不使用 Preload（N+1 问题）==========
	fmt.Println("❌ 方式1：不使用 Preload（会产生 N+1 查询问题）")
	fmt.Println("   代码：")
	fmt.Println("   var orders []OrderWithRelations")
	fmt.Println("   db.Find(&orders)  // 第1次查询：获取订单")
	fmt.Println("   for _, order := range orders {")
	fmt.Println("       db.First(&order.User, order.UserID)  // 每个订单都查一次用户")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   实际执行的 SQL：")
	fmt.Println("   1. SELECT * FROM t_order;                    (查询订单)")
	fmt.Println("   2. SELECT * FROM t_sys_user WHERE id = 1;  (订单1的用户)")
	fmt.Println("   3. SELECT * FROM t_sys_user WHERE id = 1;    (订单2的用户)")
	fmt.Println("   4. SELECT * FROM t_sys_user WHERE id = 1;    (订单3的用户)")
	fmt.Println()
	fmt.Println("   ⚠️  问题：总共执行了 4 次查询（1 + 3 = N+1）")
	fmt.Println("   ⚠️  如果订单数量是 100，就会执行 101 次查询！")
	fmt.Println()

	// ========== 方式2：使用 Preload（优化后）==========
	fmt.Println("✅ 方式2：使用 Preload（优化后）")
	fmt.Println("   代码：")
	fmt.Println("   var orders []OrderWithRelations")
	fmt.Println("   db.Preload(\"User\").Find(&orders)  // 一次性加载所有关联")
	fmt.Println()
	fmt.Println("   实际执行的 SQL：")
	fmt.Println("   1. SELECT * FROM t_order;                    (查询订单)")
	fmt.Println("   2. SELECT * FROM t_sys_user WHERE id IN (1); (批量查询所有用户)")
	fmt.Println()
	fmt.Println("   ✅ 优势：只执行了 2 次查询，无论订单数量多少！")
	fmt.Println()

	// ========== 实际演示 ==========
	fmt.Println("📝 实际演示：")
	fmt.Println()

	// 演示 Preload
	var orders []OrderWithRelations
	fmt.Println("执行: db.Preload(\"User\").Find(&orders)")
	db.Preload("User").Find(&orders)

	fmt.Printf("\n   查询结果：找到 %d 个订单\n", len(orders))
	for i, order := range orders {
		if order.User != nil {
			fmt.Printf("   订单 %d: %s - 用户: %s\n", i+1, order.OrderNo, order.User.Username)
		}
	}
	fmt.Println()

	// ========== Preload 的其他用法 ==========
	fmt.Println("🔧 Preload 的其他用法：")
	fmt.Println()

	// 1. 预加载多个关联
	fmt.Println("1. 预加载多个关联：")
	fmt.Println("   db.Preload(\"User\").Preload(\"Offering\").Find(&orders)")
	fmt.Println("   → 一次性加载用户和服务信息")
	fmt.Println()

	// 2. 预加载嵌套关联
	fmt.Println("2. 预加载嵌套关联：")
	fmt.Println("   db.Preload(\"Offering.Category\").Find(&orders)")
	fmt.Println("   → 订单 -> 服务 -> 分类，三层关联一次性加载")
	fmt.Println()

	// 3. 预加载时添加条件
	fmt.Println("3. 预加载时添加条件：")
	fmt.Println("   db.Preload(\"User\", \"status = ?\", \"active\").Find(&orders)")
	fmt.Println("   → 只加载状态为 active 的用户")
	fmt.Println()

	// 4. 使用函数预加载（更灵活）
	fmt.Println("4. 使用函数预加载：")
	fmt.Println("   db.Preload(\"User\", func(db *gorm.DB) *gorm.DB {")
	fmt.Println("       return db.Where(\"status = ?\", \"active\").Order(\"id ASC\")")
	fmt.Println("   }).Find(&orders)")
	fmt.Println("   → 可以添加复杂的查询条件")
	fmt.Println()

	// ========== 性能对比 ==========
	fmt.Println("⚡ 性能对比：")
	fmt.Println()
	fmt.Println("   场景：查询 100 个订单及其用户信息")
	fmt.Println()
	fmt.Println("   不使用 Preload：")
	fmt.Println("     - 查询次数：101 次（1 + 100）")
	fmt.Println("     - 网络往返：101 次")
	fmt.Println("     - 执行时间：~101ms（假设每次查询 1ms）")
	fmt.Println()
	fmt.Println("   使用 Preload：")
	fmt.Println("     - 查询次数：2 次（1 + 1）")
	fmt.Println("     - 网络往返：2 次")
	fmt.Println("     - 执行时间：~2ms")
	fmt.Println()
	fmt.Println("   🚀 性能提升：约 50 倍！")
	fmt.Println()

	// ========== 总结 ==========
	fmt.Println("📚 总结：")
	fmt.Println()
	fmt.Println("   预加载（Preload）是什么？")
	fmt.Println("   → 在查询主表数据时，同时查询并加载关联表的数据")
	fmt.Println()
	fmt.Println("   为什么需要预加载？")
	fmt.Println("   → 避免 N+1 查询问题，大幅提升性能")
	fmt.Println()
	fmt.Println("   什么时候使用预加载？")
	fmt.Println("   → 需要访问关联数据时，都应该使用 Preload")
	fmt.Println("   → 特别是在循环中访问关联数据时，必须使用 Preload")
	fmt.Println()
	fmt.Println("   fuyelead 项目中的使用：")
	fmt.Println("   → 查询订单时预加载用户和服务信息")
	fmt.Println("   → 查询分类时预加载服务列表")
	fmt.Println("   → 分页查询时预加载所有关联数据")
	fmt.Println()

	fmt.Println("=== 预加载详解完成 ===")
}

