package gorm

import (
	"fmt"
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DatabaseConfig 数据库配置（基于 fuyelead 项目）
type DatabaseConfig struct {
	Host     string
	Port     int
	User     string
	Password string
	Name     string
	Charset  string
}

// Database 数据库连接封装（基于 fuyelead 项目）
type Database struct {
	*gorm.DB
}

// GormDatabaseConfigDemo 展示 GORM 数据库配置（基于 fuyelead 项目）
func GormDatabaseConfigDemo() {
	fmt.Println("=== GORM 数据库配置示例（基于 fuyelead 项目）===")
	fmt.Println()

	// 1. MySQL 连接配置（fuyelead 项目使用的方式）
	fmt.Println("1. MySQL 连接配置（fuyelead 项目）")
	fmt.Println("   fuyelead 项目使用 MySQL 数据库，配置如下：")
	fmt.Println()

	config := &DatabaseConfig{
		Host:     "localhost",
		Port:     3306,
		User:     "fuyelead_user",
		Password: "fuyelead_pass",
		Name:     "fuyelead",
		Charset:  "utf8mb4",
	}

	// 构建 DSN（Data Source Name）
	// fuyelead 项目使用 Asia/Shanghai 时区
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Asia%%2FShanghai",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset,
	)
	fmt.Printf("   DSN 示例: %s\n", dsn)
	fmt.Println("   关键参数说明：")
	fmt.Println("     - charset=utf8mb4: 支持完整的 UTF-8 字符集（包括 emoji）")
	fmt.Println("     - parseTime=True: 自动解析时间字段")
	fmt.Println("     - loc=Asia/Shanghai: 使用中国时区（UTC+8）")
	fmt.Println()

	// 2. GORM 配置（基于 fuyelead 项目）
	fmt.Println("2. GORM 配置")
	fmt.Println("   fuyelead 项目根据环境变量配置日志级别：")
	fmt.Println()

	// 开发环境：显示 SQL 日志
	devConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 显示所有 SQL
	}
	fmt.Println("   开发环境配置:")
	fmt.Println("     Logger: Info (显示所有 SQL 语句)")
	fmt.Println()

	// 生产环境：静默模式
	_ = &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // 不显示 SQL
	}
	fmt.Println("   生产环境配置:")
	fmt.Println("     Logger: Silent (不显示 SQL 语句)")
	fmt.Println()

	// 3. 连接池配置（fuyelead 项目的最佳实践）
	fmt.Println("3. 连接池配置（fuyelead 项目）")
	fmt.Println("   fuyelead 项目配置了连接池参数：")
	fmt.Println()

	// 使用 SQLite 内存数据库演示（不需要真实 MySQL）
	db, err := gorm.Open(sqlite.Open(":memory:"), devConfig)
	if err != nil {
		fmt.Printf("连接失败: %v\n", err)
		return
	}

	// 获取底层 *sql.DB 进行连接池配置
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Printf("获取数据库实例失败: %v\n", err)
		return
	}

	// fuyelead 项目的连接池配置
	sqlDB.SetMaxIdleConns(10)        // 最大空闲连接数
	sqlDB.SetMaxOpenConns(100)       // 最大打开连接数
	sqlDB.SetConnMaxLifetime(time.Hour) // 连接最大生存时间

	fmt.Println("   连接池参数:")
	fmt.Printf("     MaxIdleConns: %d (最大空闲连接数)\n", 10)
	fmt.Printf("     MaxOpenConns: %d (最大打开连接数)\n", 100)
	fmt.Printf("     ConnMaxLifetime: %s (连接最大生存时间)\n", time.Hour)
	fmt.Println()

	// 4. 实际连接示例（使用 SQLite 演示）
	fmt.Println("4. 实际连接示例（使用 SQLite 内存数据库）")
	
	// 测试连接
	err = sqlDB.Ping()
	if err != nil {
		fmt.Printf("   ✗ 连接测试失败: %v\n", err)
		return
	}
	fmt.Println("   ✓ 数据库连接成功")
	fmt.Println()

	// 5. MySQL 连接示例代码（fuyelead 项目实际代码）
	fmt.Println("5. MySQL 连接示例代码（fuyelead 项目）")
	fmt.Println("   实际项目中的连接代码：")
	fmt.Println()
	fmt.Println("   ```go")
	fmt.Println("   // 构建 DSN")
	fmt.Println("   dsn := fmt.Sprintf(\"%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Asia%%2FShanghai\",")
	fmt.Println("       config.User, config.Password, config.Host, config.Port, config.Name, config.Charset)")
	fmt.Println()
	fmt.Println("   // 配置 GORM")
	fmt.Println("   logLevel := logger.Silent")
	fmt.Println("   if getEnv(\"APP_ENV\", \"production\") == \"development\" {")
	fmt.Println("       logLevel = logger.Info")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   gormConfig := &gorm.Config{")
	fmt.Println("       Logger: logger.Default.LogMode(logLevel),")
	fmt.Println("   }")
	fmt.Println()
	fmt.Println("   // 打开连接")
	fmt.Println("   db, err := gorm.Open(mysql.Open(dsn), gormConfig)")
	fmt.Println()
	fmt.Println("   // 配置连接池")
	fmt.Println("   sqlDB, _ := db.DB()")
	fmt.Println("   sqlDB.SetMaxIdleConns(10)")
	fmt.Println("   sqlDB.SetMaxOpenConns(100)")
	fmt.Println("   sqlDB.SetConnMaxLifetime(time.Hour)")
	fmt.Println("   ```")
	fmt.Println()

	// 6. 环境变量配置（fuyelead 项目）
	fmt.Println("6. 环境变量配置（fuyelead 项目）")
	fmt.Println("   fuyelead 项目从环境变量读取配置：")
	fmt.Println()
	fmt.Println("   DB_HOST: 数据库主机（默认: localhost）")
	fmt.Println("   DB_PORT: 数据库端口（默认: 3306）")
	fmt.Println("   DB_USER: 数据库用户名（默认: fuyelead_user）")
	fmt.Println("   DB_PASSWORD: 数据库密码（默认: fuyelead_pass）")
	fmt.Println("   DB_NAME: 数据库名称（默认: fuyelead）")
	fmt.Println("   APP_ENV: 应用环境（development/production）")
	fmt.Println()

	// 7. 数据库关闭
	fmt.Println("7. 数据库关闭")
	fmt.Println("   使用 defer 确保数据库连接正确关闭：")
	fmt.Println()
	fmt.Println("   ```go")
	fmt.Println("   defer db.Close()")
	fmt.Println("   ```")
	fmt.Println()

	// 8. 最佳实践总结
	fmt.Println("8. fuyelead 项目数据库配置最佳实践")
	fmt.Println("   ✓ 使用环境变量管理配置")
	fmt.Println("   ✓ 根据环境设置不同的日志级别")
	fmt.Println("   ✓ 配置合理的连接池参数")
	fmt.Println("   ✓ 使用 utf8mb4 字符集支持完整 UTF-8")
	fmt.Println("   ✓ 设置正确的时区（Asia/Shanghai）")
	fmt.Println("   ✓ 使用 defer 确保连接关闭")
	fmt.Println()

	// 关闭连接
	sqlDB.Close()
	fmt.Println("=== GORM 数据库配置示例完成 ===")
}

// NewDatabase 创建数据库连接（fuyelead 项目风格）
func NewDatabase(config *DatabaseConfig) (*Database, error) {
	// 注意：这里只是示例，实际使用时需要 MySQL 驱动
	// 示例中使用 SQLite 演示
	
	// 实际项目中的 MySQL DSN 构建（这里仅作演示，实际使用 SQLite）
	_ = fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Asia%%2FShanghai",
		config.User,
		config.Password,
		config.Host,
		config.Port,
		config.Name,
		config.Charset,
	)

	// 配置日志级别
	logLevel := logger.Silent
	// 实际项目中从环境变量读取: if getEnv("APP_ENV", "production") == "development"

	gormConfig := &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	}

	// 实际项目中使用: db, err := gorm.Open(mysql.Open(dsn), gormConfig)
	// 这里使用 SQLite 演示
	db, err := gorm.Open(sqlite.Open(":memory:"), gormConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	// 配置连接池
	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to get database instance: %w", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	return &Database{DB: db}, nil
}

