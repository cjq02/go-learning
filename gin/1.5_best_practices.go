package gin

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// UnifiedResponseDemo 演示 RESTful API 标准化响应格式
// 这是 RESTful API 最佳实践的核心部分
func UnifiedResponseDemo() {
	fmt.Println("=== RESTful API 标准化响应格式示例 ===")
	fmt.Println()

	router := gin.Default()

	// ========== 标准化响应结构 ==========
	// 统一的响应格式，便于前端统一处理
	// 设计原则:
	//   - code: 业务状态码，0 表示成功，非0表示失败
	//   - data: 成功时返回的数据
	//   - message: 错误时的提示信息
	type Response struct {
		Code    int         `json:"code"`    // 业务状态码：0=成功，非0=失败
		Data    interface{} `json:"data"`    // 成功时返回的数据
		Message string      `json:"message"` // 错误时的提示信息
	}

	// ========== 响应辅助函数 ==========
	// Success: 成功响应
	// 参数:
	//   - c: Gin 上下文
	//   - data: 要返回的数据（可以是任意类型）
	// 返回格式: {"code": 0, "data": {...}}
	Success := func(c *gin.Context, data interface{}) {
		c.JSON(http.StatusOK, Response{
			Code: 0,
			Data: data,
		})
	}

	// Error: 错误响应
	// 参数:
	//   - c: Gin 上下文
	//   - code: 业务错误码（非0）
	//   - msg: 错误提示信息
	// 返回格式: {"code": 1001, "message": "错误信息"}
	Error := func(c *gin.Context, code int, msg string) {
		c.JSON(http.StatusOK, Response{
			Code:    code,
			Message: msg,
		})
	}

	// ========== 使用示例 ==========
	// 成功响应示例
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		// 模拟查询用户信息
		user := gin.H{
			"id":    id,
			"name":  "John Doe",
			"email": "john@example.com",
		}
		Success(c, user)
	})

	// 错误响应示例
	router.POST("/users", func(c *gin.Context) {
		var user struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email" binding:"required,email"`
		}

		// 参数校验失败
		if err := c.ShouldBindJSON(&user); err != nil {
			Error(c, 1001, "参数校验失败: "+err.Error())
			return
		}

		// 业务逻辑验证（示例）
		if user.Name == "admin" {
			Error(c, 1002, "用户名不能为 admin")
			return
		}

		// 成功创建
		Success(c, gin.H{
			"id":    1,
			"name":  user.Name,
			"email": user.Email,
		})
	})

	fmt.Println("========== 标准化响应格式 ==========")
	fmt.Println()
	fmt.Println("响应结构:")
	fmt.Println("  type Response struct {")
	fmt.Println("      Code    int         `json:\"code\"`")
	fmt.Println("      Data    interface{} `json:\"data\"`")
	fmt.Println("      Message string      `json:\"message\"`")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("成功响应示例:")
	fmt.Println("  GET /users/123")
	fmt.Println("  响应: {\"code\": 0, \"data\": {\"id\": \"123\", \"name\": \"John Doe\"}}")
	fmt.Println()
	fmt.Println("错误响应示例:")
	fmt.Println("  POST /users (参数错误)")
	fmt.Println("  响应: {\"code\": 1001, \"message\": \"参数校验失败\"}")
	fmt.Println()
	fmt.Println("========== 错误代码规范 ==========")
	fmt.Println()
	fmt.Println("业务错误码规范:")
	fmt.Println("  0     - 操作成功")
	fmt.Println()
	fmt.Println("  1xxx  - 参数/请求错误")
	fmt.Println("    1001 - 参数校验失败")
	fmt.Println("    1002 - 认证失败")
	fmt.Println("    1003 - 权限不足")
	fmt.Println("    1004 - 资源不存在")
	fmt.Println()
	fmt.Println("  2xxx  - 服务端错误")
	fmt.Println("    2001 - 数据库错误")
	fmt.Println("    2002 - 缓存错误")
	fmt.Println("    2003 - 第三方服务错误")
	fmt.Println()
	fmt.Println("  3xxx  - 业务逻辑错误")
	fmt.Println("    3001 - 业务规则违反")
	fmt.Println("    3002 - 状态不允许")
	fmt.Println()
	fmt.Println("设计原则:")
	fmt.Println("  1. HTTP 状态码用于表示请求状态（200, 400, 500等）")
	fmt.Println("  2. 业务状态码用于表示业务逻辑结果（0=成功，非0=失败）")
	fmt.Println("  3. 所有响应都返回 HTTP 200，通过 code 字段区分成功/失败")
	fmt.Println("  4. 便于前端统一处理，无需判断 HTTP 状态码")
	fmt.Println()
	fmt.Println("使用建议:")
	fmt.Println("  1. 将 Success 和 Error 函数放在公共包中")
	fmt.Println("  2. 统一错误码定义在常量文件中")
	fmt.Println("  3. 使用枚举或常量避免硬编码错误码")
	fmt.Println("  4. 提供错误码文档给前端团队")
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
		Password string `json:"-"`               // 使用 json:"-" 标签排除序列化
		Token    string `json:"-"`               // 敏感信息不返回
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

// SwaggerDocumentationDemo 演示接口文档生成（Swagger）
// Swagger 是 RESTful API 文档生成工具，可以自动生成交互式 API 文档
func SwaggerDocumentationDemo() {
	fmt.Println("=== RESTful API 接口文档生成示例 ===")
	fmt.Println()

	fmt.Println("========== Swagger 简介 ==========")
	fmt.Println()
	fmt.Println("Swagger 是一个强大的 API 文档工具，可以:")
	fmt.Println("  1. 自动生成交互式 API 文档")
	fmt.Println("  2. 支持在线测试接口")
	fmt.Println("  3. 生成多种语言的客户端 SDK")
	fmt.Println("  4. 提供 API 版本管理")
	fmt.Println()

	fmt.Println("========== 安装 Swagger ==========")
	fmt.Println()
	fmt.Println("1. 安装 swag 工具:")
	fmt.Println("   go install github.com/swaggo/swag/cmd/swag@latest")
	fmt.Println()
	fmt.Println("2. 验证安装:")
	fmt.Println("   swag --version")
	fmt.Println()

	fmt.Println("========== 在项目中使用 ==========")
	fmt.Println()
	fmt.Println("1. 安装 Gin Swagger 依赖:")
	fmt.Println("   go get -u github.com/swaggo/gin-swagger")
	fmt.Println("   go get -u github.com/swaggo/files")
	fmt.Println()
	fmt.Println("2. 在 main.go 中添加注释:")
	fmt.Println("   // @title           API 文档标题")
	fmt.Println("   // @version         1.0")
	fmt.Println("   // @description     这是 API 文档描述")
	fmt.Println("   // @termsOfService  http://swagger.io/terms/")
	fmt.Println("   // @contact.name    API Support")
	fmt.Println("   // @contact.url     http://www.example.com/support")
	fmt.Println("   // @contact.email   support@example.com")
	fmt.Println("   // @license.name    Apache 2.0")
	fmt.Println("   // @license.url     http://www.apache.org/licenses/LICENSE-2.0.html")
	fmt.Println("   // @host            localhost:8080")
	fmt.Println("   // @BasePath        /api/v1")
	fmt.Println()
	fmt.Println("3. 在路由处理函数上添加注释:")
	fmt.Println("   // @Summary         获取用户信息")
	fmt.Println("   // @Description     根据用户ID获取详细信息")
	fmt.Println("   // @Tags            users")
	fmt.Println("   // @Accept          json")
	fmt.Println("   // @Produce         json")
	fmt.Println("   // @Param           id path int true \"用户ID\"")
	fmt.Println("   // @Success         200 {object} Response")
	fmt.Println("   // @Failure         400 {object} Response")
	fmt.Println("   // @Router          /users/{id} [get]")
	fmt.Println()
	fmt.Println("4. 生成文档:")
	fmt.Println("   swag init -g main.go")
	fmt.Println()
	fmt.Println("5. 在代码中引入 Swagger:")
	fmt.Println("   import (")
	fmt.Println("       swaggerFiles \"github.com/swaggo/files\"")
	fmt.Println("       ginSwagger \"github.com/swaggo/gin-swagger\"")
	fmt.Println("   )")
	fmt.Println()
	fmt.Println("   // 注册 Swagger 路由")
	fmt.Println("   router.GET(\"/swagger/*any\", ginSwagger.WrapHandler(swaggerFiles.Handler))")
	fmt.Println()
	fmt.Println("6. 访问文档:")
	fmt.Println("   浏览器打开: http://localhost:8080/swagger/index.html")
	fmt.Println()

	fmt.Println("========== 注释示例 ==========")
	fmt.Println()
	fmt.Println("// @Summary      创建用户")
	fmt.Println("// @Description  创建新用户")
	fmt.Println("// @Tags         users")
	fmt.Println("// @Accept       json")
	fmt.Println("// @Produce      json")
	fmt.Println("// @Param        user body UserRequest true \"用户信息\"")
	fmt.Println("// @Success      200 {object} Response{data=User}")
	fmt.Println("// @Failure      400 {object} Response")
	fmt.Println("// @Router       /users [post]")
	fmt.Println("func CreateUser(c *gin.Context) {")
	fmt.Println("    // ...")
	fmt.Println("}")
	fmt.Println()

	fmt.Println("========== 常用注释标签 ==========")
	fmt.Println()
	fmt.Println("@title          - API 标题")
	fmt.Println("@version        - API 版本")
	fmt.Println("@description    - API 描述")
	fmt.Println("@host           - 服务器地址")
	fmt.Println("@BasePath       - 基础路径")
	fmt.Println("@Summary        - 接口摘要")
	fmt.Println("@Description    - 接口详细描述")
	fmt.Println("@Tags           - 接口分组标签")
	fmt.Println("@Accept         - 接受的请求类型 (json, xml, form)")
	fmt.Println("@Produce        - 返回的数据类型 (json, xml)")
	fmt.Println("@Param          - 参数说明 (path/query/body/header)")
	fmt.Println("@Success        - 成功响应")
	fmt.Println("@Failure        - 失败响应")
	fmt.Println("@Router         - 路由定义")
	fmt.Println("@Security       - 安全认证")
	fmt.Println()

	fmt.Println("========== 错误代码规范 ==========")
	fmt.Println()
	fmt.Println("标准错误码分类:")
	fmt.Println()
	fmt.Println("成功:")
	fmt.Println("  0     - 操作成功")
	fmt.Println()
	fmt.Println("参数/请求错误 (1xxx):")
	fmt.Println("  1001  - 参数校验失败")
	fmt.Println("  1002  - 认证失败")
	fmt.Println("  1003  - 权限不足")
	fmt.Println("  1004  - 资源不存在")
	fmt.Println("  1005  - 请求方法不允许")
	fmt.Println("  1006  - 请求过于频繁")
	fmt.Println()
	fmt.Println("服务端错误 (2xxx):")
	fmt.Println("  2001  - 数据库错误")
	fmt.Println("  2002  - 缓存错误")
	fmt.Println("  2003  - 第三方服务错误")
	fmt.Println("  2004  - 内部服务器错误")
	fmt.Println()
	fmt.Println("业务逻辑错误 (3xxx):")
	fmt.Println("  3001  - 业务规则违反")
	fmt.Println("  3002  - 状态不允许")
	fmt.Println("  3003  - 余额不足")
	fmt.Println("  3004  - 操作冲突")
	fmt.Println()
	fmt.Println("错误码定义建议:")
	fmt.Println("  // 定义错误码常量")
	fmt.Println("  const (")
	fmt.Println("      CodeSuccess          = 0")
	fmt.Println("      CodeParamError       = 1001")
	fmt.Println("      CodeAuthFailed       = 1002")
	fmt.Println("      CodePermissionDenied = 1003")
	fmt.Println("      CodeNotFound         = 1004")
	fmt.Println("      CodeDatabaseError    = 2001")
	fmt.Println("      CodeBusinessError    = 3001")
	fmt.Println("  )")
	fmt.Println()
	fmt.Println("使用示例:")
	fmt.Println("  Error(c, CodeParamError, \"用户名不能为空\")")
	fmt.Println("  Error(c, CodeAuthFailed, \"登录已过期，请重新登录\")")
	fmt.Println()
	fmt.Println("最佳实践:")
	fmt.Println("  1. 错误码统一管理，避免硬编码")
	fmt.Println("  2. 错误码与错误信息分离，支持国际化")
	fmt.Println("  3. 提供错误码对照表给前端")
	fmt.Println("  4. 记录错误日志，便于排查问题")
	fmt.Println("  5. 错误信息要清晰明确，便于用户理解")
}
