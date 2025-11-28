package gin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// UnifiedResponseDemo 演示统一响应格式规范
func UnifiedResponseDemo() {
	fmt.Println("=== Gin 统一响应格式规范示例 ===")
	fmt.Println()

	router := gin.Default()

	// 定义统一响应结构
	type Response struct {
		Code    int         `json:"code"`
		Message string      `json:"message"`
		Data    interface{} `json:"data,omitempty"`
		Error   string      `json:"error,omitempty"`
		Time    int64       `json:"time"`
	}

	// 响应辅助函数
	success := func(c *gin.Context, data interface{}) {
		c.JSON(http.StatusOK, Response{
			Code:    200,
			Message: "操作成功",
			Data:    data,
			Time:    time.Now().Unix(),
		})
	}

	errorResponse := func(c *gin.Context, code int, message string, err error) {
		response := Response{
			Code:    code,
			Message: message,
			Time:    time.Now().Unix(),
		}
		if err != nil {
			response.Error = err.Error()
		}
		c.JSON(http.StatusOK, response)
	}

	// 使用统一响应格式
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		success(c, gin.H{
			"id":   id,
			"name": "John Doe",
		})
	})

	router.POST("/users", func(c *gin.Context) {
		var user struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email" binding:"required,email"`
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			errorResponse(c, 1001, "参数校验失败", err)
			return
		}

		success(c, user)
	})

	fmt.Println("统一响应格式规范:")
	fmt.Println("  成功响应: {\"code\": 200, \"message\": \"操作成功\", \"data\": {...}, \"time\": 1234567890}")
	fmt.Println("  失败响应: {\"code\": 1001, \"message\": \"错误信息\", \"error\": \"详细错误\", \"time\": 1234567890}")
	fmt.Println()
	fmt.Println("响应码规范:")
	fmt.Println("  200  - 操作成功")
	fmt.Println("  1001 - 参数校验失败")
	fmt.Println("  1002 - 业务逻辑错误")
	fmt.Println("  1003 - 认证失败")
	fmt.Println("  1004 - 权限不足")
	fmt.Println("  1005 - 资源不存在")
}

// SensitiveDataFilterDemo 演示敏感参数过滤处理
func SensitiveDataFilterDemo() {
	fmt.Println("=== Gin 敏感参数过滤处理示例 ===")
	fmt.Println()

	router := gin.Default()

	// 敏感字段过滤中间件
	filterSensitiveData := func(c *gin.Context) {
		// 记录原始响应
		c.Next()

		// 获取响应数据并过滤敏感信息
		// 注意：实际应用中需要根据响应类型进行处理
	}

	router.Use(filterSensitiveData)

	// 用户信息结构（包含敏感字段）
	type UserInfo struct {
		ID       int    `json:"id"`
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"-"` // 使用 json:"-" 标签排除序列化
		Token    string `json:"-"` // 敏感信息不返回
		Phone    string `json:"phone,omitempty"` // 可选字段
	}

	router.GET("/profile", func(c *gin.Context) {
		user := UserInfo{
			ID:       1,
			Username: "john",
			Email:    "john@example.com",
			Password: "secret123", // 不会序列化到JSON
			Token:    "token123",  // 不会序列化到JSON
			Phone:    "13800138000",
		}

		c.JSON(http.StatusOK, user)
	})

	fmt.Println("敏感参数过滤方法:")
	fmt.Println("  1. 使用 json:\"-\" 标签排除字段序列化")
	fmt.Println("  2. 使用 json:\"omitempty\" 标签隐藏空值")
	fmt.Println("  3. 创建单独的响应DTO结构体")
	fmt.Println("  4. 使用中间件统一过滤敏感数据")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  Password string `json:\"-\"`        - 完全排除")
	fmt.Println("  Phone    string `json:\"omitempty\"` - 空值时排除")
}

// RateLimitDemo 演示请求频率限制中间件
func RateLimitDemo() {
	fmt.Println("=== Gin 请求频率限制中间件示例 ===")
	fmt.Println()

	router := gin.Default()

	// 简单的内存限流器（生产环境建议使用Redis）
	type RateLimiter struct {
		requests map[string][]time.Time
		limit    int
		window   time.Duration
	}

	limiter := &RateLimiter{
		requests: make(map[string][]time.Time),
		limit:    10,              // 限制10次请求
		window:   time.Minute * 1, // 时间窗口1分钟
	}

	rateLimitMiddleware := func(c *gin.Context) {
		clientIP := c.ClientIP()
		now := time.Now()

		// 清理过期记录
		if requests, exists := limiter.requests[clientIP]; exists {
			validRequests := []time.Time{}
			for _, reqTime := range requests {
				if now.Sub(reqTime) < limiter.window {
					validRequests = append(validRequests, reqTime)
				}
			}
			limiter.requests[clientIP] = validRequests
		}

		// 检查是否超过限制
		if len(limiter.requests[clientIP]) >= limiter.limit {
			c.JSON(http.StatusTooManyRequests, gin.H{
				"code":    1006,
				"message": "请求过于频繁，请稍后再试",
			})
			c.Abort()
			return
		}

		// 记录本次请求
		limiter.requests[clientIP] = append(limiter.requests[clientIP], now)
		c.Next()
	}

	// 应用限流中间件
	api := router.Group("/api")
	api.Use(rateLimitMiddleware)
	{
		api.GET("/data", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "数据获取成功",
			})
		})
	}

	fmt.Println("请求频率限制配置:")
	fmt.Println("  限制: 10次请求/分钟")
	fmt.Println("  作用范围: /api/* 路由组")
	fmt.Println()
	fmt.Println("限流策略:")
	fmt.Println("  1. 固定窗口限流 - 简单但可能有突发流量")
	fmt.Println("  2. 滑动窗口限流 - 更平滑，推荐使用")
	fmt.Println("  3. 令牌桶算法 - 适合突发流量场景")
	fmt.Println("  4. 漏桶算法 - 严格控制速率")
	fmt.Println()
	fmt.Println("生产环境建议:")
	fmt.Println("  使用 Redis 实现分布式限流")
	fmt.Println("  使用 go-redis/redis 或 goredis 库")
	fmt.Println("  支持不同用户/IP的差异化限流策略")
}

// VersionControlDemo 演示路由版本控制方案
func VersionControlDemo() {
	fmt.Println("=== Gin 路由版本控制方案示例 ===")
	fmt.Println()

	router := gin.Default()

	// 方案1: URL路径版本控制
	v1 := router.Group("/api/v1")
	{
		v1.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"version": "v1",
				"message": "用户列表 (v1)",
			})
		})
	}

	v2 := router.Group("/api/v2")
	{
		v2.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"version": "v2",
				"message": "用户列表 (v2)",
			})
		})
	}

	// 方案2: Header版本控制
	router.GET("/api/users", func(c *gin.Context) {
		version := c.GetHeader("API-Version")
		switch version {
		case "v1":
			c.JSON(http.StatusOK, gin.H{
				"version": "v1",
				"message": "用户列表 (v1 via Header)",
			})
		case "v2":
			c.JSON(http.StatusOK, gin.H{
				"version": "v2",
				"message": "用户列表 (v2 via Header)",
			})
		default:
			c.JSON(http.StatusOK, gin.H{
				"version": "latest",
				"message": "用户列表 (默认版本)",
			})
		}
	})

	fmt.Println("路由版本控制方案:")
	fmt.Println()
	fmt.Println("方案1: URL路径版本控制")
	fmt.Println("  /api/v1/users - 版本1")
	fmt.Println("  /api/v2/users - 版本2")
	fmt.Println("  优点: 清晰直观，易于理解")
	fmt.Println("  缺点: URL较长")
	fmt.Println()
	fmt.Println("方案2: Header版本控制")
	fmt.Println("  GET /api/users")
	fmt.Println("  Header: API-Version: v1")
	fmt.Println("  优点: URL简洁")
	fmt.Println("  缺点: 需要客户端支持")
	fmt.Println()
	fmt.Println("推荐方案:")
	fmt.Println("  1. 新版本使用URL路径: /api/v2/...")
	fmt.Println("  2. 保持旧版本兼容性")
	fmt.Println("  3. 设置版本过期时间")
	fmt.Println("  4. 提供版本迁移文档")
}

