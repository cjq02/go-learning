package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

// MiddlewareFlowDemo 演示中间件执行流程
// 展示中间件链式调用的完整流程
func MiddlewareFlowDemo() {
	fmt.Println("=== Gin 中间件执行流程示例 ===")
	fmt.Println()

	// ========== gin.Default() 说明 ==========
	// gin.Default() 创建一个带有默认中间件的路由器
	// 它等价于: gin.New() + Logger中间件 + Recovery中间件
	//
	// 内置中间件:
	//   1. Logger 中间件: 自动记录HTTP请求日志
	//   2. Recovery 中间件: 自动恢复panic，避免程序崩溃
	//
	// 与 gin.New() 的区别:
	//   - gin.Default(): 包含Logger和Recovery中间件（适合开发环境）
	//   - gin.New(): 不包含任何中间件，需要手动添加（适合生产环境，更灵活）
	//
	// 使用建议:
	//   - 开发环境: 使用 gin.Default()，方便调试
	//   - 生产环境: 使用 gin.New()，然后按需添加中间件，性能更好
	router := gin.Default()

	// ========== 中间件执行流程说明 ==========
	// 中间件按照注册顺序依次执行，形成执行链:
	//
	//  [客户端请求]
	//      ↓
	//  [Logger中间件] → 记录请求开始时间
	//      ↓
	//  [CORS中间件] → 处理跨域请求
	//      ↓
	//  [JWT鉴权] → 验证访问令牌
	//      ↓
	//  [RBAC鉴权] → 校验用户权限 (RBAC: Role-Based Access Control 基于角色的访问控制)
	//      ↓
	//  [业务处理] → 核心业务逻辑
	//      ↓
	//  [Logger中间件] ← 记录响应耗时（c.Next()后执行）

	// Logger 中间件 - 记录请求和响应时间
	loggerMiddleware := func(c *gin.Context) {
		start := time.Now()
		requestID := uuid.New().String()
		c.Set("requestID", requestID)

		log.Printf("[%s] 请求开始: %s %s", requestID, c.Request.Method, c.Request.URL.Path)

		c.Next() // 执行下一个中间件或处理函数

		// c.Next() 后的代码在所有后续中间件执行完后才执行
		latency := time.Since(start)
		log.Printf("[%s] 请求完成: 状态码=%d, 耗时=%v", requestID, c.Writer.Status(), latency)
	}

	// 模拟 CORS 中间件
	corsMiddleware := func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		log.Println("CORS 中间件: 设置跨域头")
		c.Next()
	}

	// 模拟 JWT 鉴权中间件（简化版）
	jwtMiddleware := func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			log.Println("JWT 中间件: 未提供 Token")
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "未提供认证令牌"})
			return
		}
		log.Println("JWT 中间件: Token 验证通过")
		c.Set("userID", "user123")
		c.Set("roles", []string{"admin", "user"})
		c.Next()
	}

	// 模拟 RBAC 权限中间件
	// RBAC: Role-Based Access Control (基于角色的访问控制)
	rbacMiddleware := func(c *gin.Context) {
		roles, exists := c.Get("roles")
		if !exists {
			log.Println("RBAC 中间件: 未找到用户角色")
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "访问被拒绝"})
			return
		}
		log.Printf("RBAC 中间件: 用户角色=%v", roles)
		c.Next()
	}

	// 注册中间件（按顺序）
	router.Use(loggerMiddleware)
	router.Use(corsMiddleware)
	router.Use(jwtMiddleware)
	router.Use(rbacMiddleware)

	// 业务处理函数
	router.GET("/api/data", func(c *gin.Context) {
		requestID, _ := c.Get("requestID")
		userID, _ := c.Get("userID")
		roles, _ := c.Get("roles")

		c.JSON(http.StatusOK, gin.H{
			"message":   "业务处理完成",
			"requestID": requestID,
			"userID":    userID,
			"roles":     roles,
		})
	})

	fmt.Println("中间件执行流程说明:")
	fmt.Println("  1. Logger 中间件: 记录请求开始时间")
	fmt.Println("  2. CORS 中间件: 处理跨域请求")
	fmt.Println("  3. JWT 中间件: 验证访问令牌")
	fmt.Println("  4. RBAC 中间件: 校验用户权限 (RBAC: Role-Based Access Control)")
	fmt.Println("  5. 业务处理: 执行核心业务逻辑")
	fmt.Println("  6. Logger 中间件: 记录响应耗时（c.Next()后执行）")
	fmt.Println()
	fmt.Println("关键概念:")
	fmt.Println()
	fmt.Println("1. c.Next() - 执行下一个中间件或处理函数")
	fmt.Println("   作用: 将控制权传递给下一个中间件或路由处理函数")
	fmt.Println("   执行顺序:")
	fmt.Println("     - c.Next() 前的代码: 在后续中间件执行前运行")
	fmt.Println("     - c.Next() 后的代码: 在所有后续中间件执行完后才运行")
	fmt.Println("   示例:")
	fmt.Println("     func middleware(c *gin.Context) {")
	fmt.Println("         fmt.Println(\"1. 中间件开始\")")
	fmt.Println("         c.Next()  // 执行下一个中间件")
	fmt.Println("         fmt.Println(\"3. 中间件结束\")  // 最后执行")
	fmt.Println("     }")
	fmt.Println()
	fmt.Println("2. c.Abort() - 终止后续中间件执行")
	fmt.Println("   作用: 立即终止请求处理，不再执行后续中间件和路由处理函数")
	fmt.Println("   使用场景:")
	fmt.Println("     - 认证失败时终止请求")
	fmt.Println("     - 权限不足时终止请求")
	fmt.Println("     - 参数校验失败时终止请求")
	fmt.Println("   示例:")
	fmt.Println("     if !isAuthenticated {")
	fmt.Println("         c.AbortWithStatusJSON(401, gin.H{\"error\": \"未授权\"})")
	fmt.Println("         return  // 必须 return，否则会继续执行")
	fmt.Println("     }")
	fmt.Println()
	fmt.Println("3. c.Set()/c.Get() - 在中间件间传递数据")
	fmt.Println("   c.Set(key, value): 存储数据到上下文")
	fmt.Println("   c.Get(key): 从上下文获取数据")
	fmt.Println("   使用场景:")
	fmt.Println("     - JWT 中间件存储用户ID和角色")
	fmt.Println("     - Logger 中间件存储请求ID")
	fmt.Println("     - 在业务处理函数中使用这些数据")
	fmt.Println("   示例:")
	fmt.Println("     // 在中间件中存储")
	fmt.Println("     c.Set(\"userID\", \"123\")")
	fmt.Println("     c.Set(\"roles\", []string{\"admin\"})")
	fmt.Println()
	fmt.Println("     // 在业务处理中获取")
	fmt.Println("     userID, _ := c.Get(\"userID\")")
	fmt.Println("     roles, _ := c.Get(\"roles\")")
	fmt.Println()
	fmt.Println("测试示例:")
	fmt.Println("  curl -H \"Authorization: Bearer token123\" http://localhost:8080/api/data")
}

// GenerateToken 生成JWT令牌
// 参数:
//   - userID: 用户ID
//   - roles: 用户角色列表
//
// 返回: JWT令牌字符串和错误
func GenerateToken(userID string, roles []string) (string, error) {
	// 设置 JWT 密钥（实际应用中应从环境变量获取）
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		secret = "your-secret-key" // 默认密钥（仅用于演示）
	}

	// 创建 JWT Claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID": userID,
		"roles":  roles,
		"exp":    time.Now().Add(8 * time.Hour).Unix(), // 8小时后过期
		"iat":    time.Now().Unix(),                    // 签发时间
	})

	// 签名并返回令牌
	return token.SignedString([]byte(secret))
}

// JWTAuth JWT鉴权中间件
//
// 功能说明:
//
//	这是一个 Gin 中间件函数，用于验证 HTTP 请求中的 JWT (JSON Web Token) 令牌。
//	如果 Token 有效，将用户信息存储到 Gin 上下文中，供后续中间件和路由处理函数使用。
//
// 工作流程:
//  1. 从请求头中提取 Authorization 字段
//  2. 解析并验证 Token 的签名和有效性
//  3. 从 Token 中提取用户信息（userID、roles等）
//  4. 将用户信息存储到 Gin 上下文，供后续使用
//  5. 如果验证失败，终止请求并返回 401 错误
//
// 使用方式:
//
//	router.Use(JWTAuth())                    // 全局中间件
//	router.GET("/api/users", JWTAuth(), handler)  // 路由级中间件
//
// 请求头格式:
//
//	Authorization: Bearer <token_string>
//
// 返回值:
//
//	gin.HandlerFunc: Gin 中间件函数
func JWTAuth() gin.HandlerFunc {
	// 返回一个 Gin 中间件函数
	// 这个函数会在每个请求到达时被调用
	return func(c *gin.Context) {
		// ========== 步骤1: 从请求头获取 Token ==========
		// Authorization 头是 HTTP 标准认证头，格式通常是: "Bearer <token>"
		// 使用 c.GetHeader() 方法获取请求头值
		authHeader := c.GetHeader("Authorization")

		// 检查是否提供了 Authorization 头
		if authHeader == "" {
			// 如果没有提供认证头，终止请求并返回 401 未授权错误
			// c.AbortWithStatusJSON() 会:
			//   1. 设置响应状态码为 401
			//   2. 返回 JSON 格式的错误信息
			//   3. 终止后续中间件和处理函数的执行
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    1002, // 业务错误码：认证失败
				"message": "未提供认证令牌",
			})
			return // 必须 return，否则会继续执行后续代码
		}

		// ========== 步骤2: 解析 Token 字符串 ==========
		// Authorization 头的格式是 "Bearer <token>"，需要移除 "Bearer " 前缀
		// strings.TrimPrefix() 会移除字符串开头的指定前缀
		// 例如: "Bearer eyJhbGciOiJIUzI1NiIs..." -> "eyJhbGciOiJIUzI1NiIs..."
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 检查 Token 字符串是否为空
		// 如果移除前缀后为空，说明格式不正确（可能没有 "Bearer " 前缀）
		if tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    1002,
				"message": "Token 格式错误，请使用格式: Bearer <token>",
			})
			return
		}

		// ========== 步骤3: 获取 JWT 签名密钥 ==========
		// JWT 使用密钥来签名和验证 Token
		// 实际应用中，密钥应该从环境变量或配置文件中读取，不要硬编码
		secret := os.Getenv("JWT_SECRET")
		if secret == "" {
			// 如果环境变量未设置，使用默认密钥（仅用于演示，生产环境必须设置）
			secret = "your-secret-key"
		}

		// ========== 步骤4: 解析和验证 Token ==========
		// jwt.Parse() 用于解析和验证 JWT Token
		// 参数说明:
		//   - tokenString: 要解析的 Token 字符串
		//   - keyFunc: 一个函数，用于获取签名密钥
		//     这个函数会在解析时被调用，用于验证 Token 的签名
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 这个函数会在解析 Token 时被调用
			// 用于验证 Token 的签名方法和获取签名密钥

			// 验证签名方法是否为 HS256 (HMAC-SHA256)
			// 这是最常用的 JWT 签名方法
			// token.Method 是 Token 中声明的签名方法
			// jwt.SigningMethodHMAC 是 HMAC 签名方法的接口
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				// 如果签名方法不是 HMAC，返回错误
				// 这可以防止算法替换攻击（Algorithm Confusion Attack）
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			// 返回签名密钥
			// 密钥必须与生成 Token 时使用的密钥相同
			// 返回 []byte 类型，因为 HMAC 签名需要字节数组
			return []byte(secret), nil
		})

		// ========== 步骤5: 检查解析结果 ==========
		// jwt.Parse() 可能返回两种错误:
		//   1. Token 格式错误（无法解析）
		//   2. Token 签名验证失败
		//   3. Token 已过期（如果 Claims 中包含 exp 字段）
		if err != nil {
			// 如果解析或验证失败，返回 401 错误
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    1002,
				"message": "Token 验证失败: " + err.Error(), // 包含详细错误信息，便于调试
			})
			return
		}

		// ========== 步骤6: 提取 Claims（Token 中的用户信息）==========
		// Claims 是 JWT Token 中存储的数据
		// jwt.MapClaims 是一个 map[string]interface{} 类型，用于存储自定义数据
		// token.Valid 检查 Token 是否有效（未过期、签名正确等）
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// Token 验证成功，提取用户信息

			// 提取 userID（用户ID）
			// claims["userID"] 是从 Token 中获取 userID 字段
			// .(string) 是类型断言，将 interface{} 转换为 string
			if userID, ok := claims["userID"].(string); ok {
				// 将 userID 存储到 Gin 上下文
				// 后续的中间件和路由处理函数可以通过 c.Get("userID") 获取
				c.Set("userID", userID)
			}

			// 提取 roles（用户角色列表）
			// roles 在 Token 中存储为数组，类型是 []interface{}
			if roles, ok := claims["roles"].([]interface{}); ok {
				// 将 []interface{} 转换为 []string
				// 因为后续的 RBAC 中间件需要 []string 类型
				roleStrings := make([]string, 0, len(roles))
				for _, r := range roles {
					// 遍历每个角色，转换为 string
					if role, ok := r.(string); ok {
						roleStrings = append(roleStrings, role)
					}
				}
				// 将角色列表存储到上下文
				c.Set("roles", roleStrings)
			}

			// ========== 步骤7: 继续执行后续中间件和路由处理函数 ==========
			// c.Next() 将控制权传递给下一个中间件或路由处理函数
			// 如果没有调用 c.Next()，请求会被终止
			c.Next()
		} else {
			// Token 无效（可能是格式错误、签名错误、已过期等）
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    1002,
				"message": "无效的 Token",
			})
		}
	}
}

// RequireRole RBAC权限中间件
// RBAC: Role-Based Access Control (基于角色的访问控制)
// 检查用户是否具有指定角色
// 参数:
//   - role: 需要的角色名称
func RequireRole(role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从上下文获取用户角色
		roles, exists := c.Get("roles")
		if !exists {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    1003,
				"message": "访问被拒绝：未找到用户角色",
			})
			return
		}

		// 检查是否有指定角色
		roleList, ok := roles.([]string)
		if !ok {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"code":    1003,
				"message": "访问被拒绝：角色数据格式错误",
			})
			return
		}

		// 遍历角色列表查找匹配的角色
		for _, r := range roleList {
			if r == role {
				c.Next() // 权限通过，继续执行
				return
			}
		}

		// 权限不足
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
			"code":    1003,
			"message": "权限不足：需要 " + role + " 角色",
		})
	}
}

// JWTAuthDemo 演示JWT鉴权完整实现
func JWTAuthDemo() {
	fmt.Println("=== JWT 鉴权完整实现示例 ===")
	fmt.Println()

	router := gin.Default()

	// 登录接口 - 生成 Token
	router.POST("/api/login", func(c *gin.Context) {
		var login struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"code":    1001,
				"message": "参数校验失败",
			})
			return
		}

		// 模拟用户验证（实际应用中应查询数据库）
		if login.Username == "admin" && login.Password == "admin123" {
			// 生成 Token
			token, err := GenerateToken("user123", []string{"admin", "user"})
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"code":    2001,
					"message": "Token 生成失败",
				})
				return
			}

			c.JSON(http.StatusOK, gin.H{
				"code":    0,
				"data":    gin.H{"token": token},
				"message": "登录成功",
			})
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{
				"code":    1002,
				"message": "用户名或密码错误",
			})
		}
	})

	// 需要认证的接口
	api := router.Group("/api")
	api.Use(JWTAuth()) // 应用 JWT 中间件
	{
		// 获取用户信息
		api.GET("/profile", func(c *gin.Context) {
			userID, _ := c.Get("userID")
			roles, _ := c.Get("roles")

			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"data": gin.H{
					"userID": userID,
					"roles":  roles,
				},
			})
		})

		// 需要 admin 角色的接口
		admin := api.Group("/admin")
		admin.Use(RequireRole("admin")) // 应用 RBAC 中间件
		{
			admin.GET("/users", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"code":    0,
					"data":    []string{"user1", "user2", "user3"},
					"message": "获取用户列表成功",
				})
			})
		}
	}

	fmt.Println("JWT 鉴权实现说明:")
	fmt.Println()
	fmt.Println("1. 生成JWT令牌:")
	fmt.Println("   func GenerateToken(userID string, roles []string) (string, error)")
	fmt.Println("   - 使用 HS256 签名方法")
	fmt.Println("   - 包含 userID、roles、过期时间等信息")
	fmt.Println()
	fmt.Println("2. JWT鉴权中间件:")
	fmt.Println("   func JWTAuth() gin.HandlerFunc")
	fmt.Println("   - 从 Authorization 头提取 Token")
	fmt.Println("   - 验证 Token 签名和有效性")
	fmt.Println("   - 将用户信息存储到上下文")
	fmt.Println()
	fmt.Println("3. RBAC权限中间件:")
	fmt.Println("   RBAC: Role-Based Access Control (基于角色的访问控制)")
	fmt.Println("   func RequireRole(role string) gin.HandlerFunc")
	fmt.Println("   - 检查用户是否具有指定角色")
	fmt.Println("   - 权限不足时终止请求")
	fmt.Println()
	fmt.Println("测试示例:")
	fmt.Println("  1. 登录获取 Token:")
	fmt.Println("     curl -X POST http://localhost:8080/api/login \\")
	fmt.Println("       -H \"Content-Type: application/json\" \\")
	fmt.Println("       -d '{\"username\":\"admin\",\"password\":\"admin123\"}'")
	fmt.Println()
	fmt.Println("  2. 使用 Token 访问受保护接口:")
	fmt.Println("     curl -H \"Authorization: Bearer <token>\" http://localhost:8080/api/profile")
	fmt.Println()
	fmt.Println("  3. 访问需要 admin 角色的接口:")
	fmt.Println("     curl -H \"Authorization: Bearer <token>\" http://localhost:8080/api/admin/users")
}

// CORSMiddlewareDemo 演示跨域中间件配置
// 注意: 实际使用需要安装 github.com/gin-contrib/cors
func CORSMiddlewareDemo() {
	fmt.Println("=== CORS 跨域中间件配置示例 ===")
	fmt.Println()

	fmt.Println("CORS (Cross-Origin Resource Sharing) 跨域资源共享配置")
	fmt.Println()
	fmt.Println("安装依赖:")
	fmt.Println("  go get github.com/gin-contrib/cors")
	fmt.Println()
	fmt.Println("代码示例:")
	fmt.Println("  import \"github.com/gin-contrib/cors\"")
	fmt.Println()
	fmt.Println("  func CORSMiddleware() gin.HandlerFunc {")
	fmt.Println("      return cors.New(cors.Config{")
	fmt.Println("          AllowOrigins: []string{")
	fmt.Println("              \"https://prod.com\",")
	fmt.Println("              \"http://localhost:3000\",")
	fmt.Println("          },")
	fmt.Println("          AllowMethods: []string{\"GET\", \"POST\", \"PUT\", \"PATCH\", \"DELETE\"},")
	fmt.Println("          AllowHeaders: []string{")
	fmt.Println("              \"Origin\",")
	fmt.Println("              \"Content-Type\",")
	fmt.Println("              \"Authorization\",")
	fmt.Println("          },")
	fmt.Println("          ExposeHeaders:    []string{\"Content-Length\"},")
	fmt.Println("          AllowCredentials: true,")
	fmt.Println("          MaxAge:           12 * time.Hour,")
	fmt.Println("      })")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("配置说明:")
	fmt.Println("  AllowOrigins     - 允许的源地址列表（* 表示允许所有）")
	fmt.Println("  AllowMethods     - 允许的HTTP方法")
	fmt.Println("  AllowHeaders     - 允许的请求头")
	fmt.Println("  ExposeHeaders    - 暴露给客户端的响应头")
	fmt.Println("  AllowCredentials - 是否允许携带凭证（Cookie等）")
	fmt.Println("  MaxAge           - 预检请求（OPTIONS）缓存时间")
	fmt.Println()
	fmt.Println("使用示例:")
	fmt.Println("  router.Use(CORSMiddleware())  // 全局中间件")
	fmt.Println("  router.GET(\"/api/data\", CORSMiddleware(), handler)  // 路由级中间件")
}

// MiddlewareDebugDemo 演示中间件调试技巧
func MiddlewareDebugDemo() {
	fmt.Println("=== 中间件调试技巧示例 ===")
	fmt.Println()

	router := gin.Default()

	// 1. 上下文数据追踪
	requestIDMiddleware := func(c *gin.Context) {
		// 生成唯一请求ID
		requestID := uuid.New().String()
		c.Set("requestID", requestID)

		// 记录请求信息
		log.Printf("[%s] %s %s", requestID, c.Request.Method, c.Request.URL.Path)
		log.Printf("[%s] 请求头: %v", requestID, c.Request.Header)

		c.Next()

		// 记录响应信息（c.Next()后执行）
		log.Printf("[%s] 响应状态码: %d", requestID, c.Writer.Status())
	}

	// 2. 中间件执行顺序验证
	router.Use(requestIDMiddleware)

	router.GET("/api/debug", func(c *gin.Context) {
		requestID, _ := c.Get("requestID")
		c.JSON(http.StatusOK, gin.H{
			"requestID": requestID,
			"message":   "调试信息",
		})
	})

	fmt.Println("调试技巧:")
	fmt.Println()
	fmt.Println("1. 上下文数据追踪:")
	fmt.Println("   - 使用 c.Set() 存储调试信息")
	fmt.Println("   - 使用 c.Get() 获取调试信息")
	fmt.Println("   - 使用 UUID 生成唯一请求ID")
	fmt.Println()
	fmt.Println("2. 中间件执行顺序验证:")
	fmt.Println("   - 在每个中间件中添加日志输出")
	fmt.Println("   - 观察日志顺序确认执行流程")
	fmt.Println()
	fmt.Println("3. 性能监控:")
	fmt.Println("   - 在 c.Next() 前后记录时间")
	fmt.Println("   - 计算每个中间件的执行耗时")
	fmt.Println()
	fmt.Println("测试示例:")
	fmt.Println("  curl http://localhost:8080/api/debug")
	fmt.Println("  查看控制台日志输出，观察中间件执行顺序")
}

// MiddlewareBestPracticesDemo 演示中间件最佳实践
func MiddlewareBestPracticesDemo() {
	fmt.Println("=== 中间件最佳实践 ===")
	fmt.Println()

	fmt.Println("========== 1. 中间件链式调用示例 ==========")
	fmt.Println()
	fmt.Println("  router.Use(")
	fmt.Println("      middleware.Recovery(),           // 异常恢复")
	fmt.Println("      middleware.CORSMiddleware(),    // 跨域处理")
	fmt.Println("      middleware.JWTAuth(),           // JWT鉴权")
	fmt.Println("      middleware.RequireRole(\"admin\"), // 权限检查")
	fmt.Println("      middleware.RequestLogger(),      // 请求日志")
	fmt.Println("  )")
	fmt.Println()

	fmt.Println("========== 2. 敏感信息过滤 ==========")
	fmt.Println()
	fmt.Println("  // 在日志中间件中过滤敏感字段")
	fmt.Println("  func SensitiveDataLogger() gin.HandlerFunc {")
	fmt.Println("      return func(c *gin.Context) {")
	fmt.Println("          // 过滤认证相关路径的请求体")
	fmt.Println("          if strings.Contains(c.Request.URL.Path, \"/auth\") {")
	fmt.Println("              // 不记录敏感信息")
	fmt.Println("              c.Next()")
	fmt.Println("              return")
	fmt.Println("          }")
	fmt.Println()
	fmt.Println("          // 记录其他请求")
	fmt.Println("          body, _ := ioutil.ReadAll(c.Request.Body)")
	fmt.Println("          log.Printf(\"Request Body: %%s\", body)")
	fmt.Println("          c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(body))")
	fmt.Println("          c.Next()")
	fmt.Println("      }")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("========== 3. 中间件性能优化 ==========")
	fmt.Println()
	fmt.Println("  - 避免在中间件中进行耗时操作（如数据库查询）")
	fmt.Println("  - 使用缓存减少重复计算")
	fmt.Println("  - 合理使用 c.Abort() 提前终止不必要的处理")
	fmt.Println("  - 使用连接池管理数据库连接")
	fmt.Println()

	fmt.Println("========== 4. 错误处理 ==========")
	fmt.Println()
	fmt.Println("  func ErrorHandler() gin.HandlerFunc {")
	fmt.Println("      return func(c *gin.Context) {")
	fmt.Println("          c.Next()")
	fmt.Println()
	fmt.Println("          // 检查是否有错误")
	fmt.Println("          if len(c.Errors) > 0 {")
	fmt.Println("              err := c.Errors.Last()")
	fmt.Println("              // 统一错误响应格式")
	fmt.Println("              c.JSON(http.StatusInternalServerError, gin.H{")
	fmt.Println("                  \"code\":    500,")
	fmt.Println("                  \"message\": err.Error(),")
	fmt.Println("              })")
	fmt.Println("          }")
	fmt.Println("      }")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("========== 5. 中间件注册建议 ==========")
	fmt.Println()
	fmt.Println("  1. 全局中间件: router.Use() - 所有路由生效")
	fmt.Println("  2. 路由组中间件: group.Use() - 特定路由组生效")
	fmt.Println("  3. 单路由中间件: router.GET(\"/path\", middleware, handler) - 单个路由生效")
	fmt.Println()
	fmt.Println("  推荐顺序:")
	fmt.Println("    1. Recovery (异常恢复)")
	fmt.Println("    2. Logger (日志记录)")
	fmt.Println("    3. CORS (跨域处理)")
	fmt.Println("    4. Auth (认证)")
	fmt.Println("    5. Permission (权限)")
	fmt.Println("    6. Business Logic (业务逻辑)")
	fmt.Println()

	fmt.Println("========== 6. 完整示例代码 ==========")
	fmt.Println()
	fmt.Println("  func main() {")
	fmt.Println("      router := gin.Default()")
	fmt.Println()
	fmt.Println("      // 全局中间件")
	fmt.Println("      router.Use(")
	fmt.Println("          Recovery(),")
	fmt.Println("          CORSMiddleware(),")
	fmt.Println("          RequestLogger(),")
	fmt.Println("      )")
	fmt.Println()
	fmt.Println("      // 公开路由（无需认证）")
	fmt.Println("      public := router.Group(\"/api/public\")")
	fmt.Println("      {")
	fmt.Println("          public.POST(\"/login\", LoginHandler)")
	fmt.Println("          public.POST(\"/register\", RegisterHandler)")
	fmt.Println("      }")
	fmt.Println()
	fmt.Println("      // 需要认证的路由")
	fmt.Println("      api := router.Group(\"/api\")")
	fmt.Println("      api.Use(JWTAuth()) // 路由组中间件")
	fmt.Println("      {")
	fmt.Println("          api.GET(\"/profile\", GetProfileHandler)")
	fmt.Println()
	fmt.Println("          // 需要 admin 权限")
	fmt.Println("          admin := api.Group(\"/admin\")")
	fmt.Println("          admin.Use(RequireRole(\"admin\"))")
	fmt.Println("          {")
	fmt.Println("              admin.GET(\"/users\", GetUsersHandler)")
	fmt.Println("              admin.DELETE(\"/users/:id\", DeleteUserHandler)")
	fmt.Println("          }")
	fmt.Println("      }")
	fmt.Println()
	fmt.Println("      router.Run(\":8080\")")
	fmt.Println("  }")
}

// MiddlewareTestDemo 演示单元测试方案
func MiddlewareTestDemo() {
	fmt.Println("=== 中间件单元测试方案 ===")
	fmt.Println()

	fmt.Println("========== 1. 安装测试依赖 ==========")
	fmt.Println()
	fmt.Println("  go get github.com/stretchr/testify/assert")
	fmt.Println()

	fmt.Println("========== 2. 中间件测试示例 ==========")
	fmt.Println()
	fmt.Println("  func TestJWTMiddleware(t *testing.T) {")
	fmt.Println("      // 设置测试环境变量")
	fmt.Println("      os.Setenv(\"JWT_SECRET\", \"test-secret-key\")")
	fmt.Println("      defer os.Unsetenv(\"JWT_SECRET\")")
	fmt.Println()
	fmt.Println("      // 创建测试路由")
	fmt.Println("      router := gin.New()")
	fmt.Println("      router.Use(JWTAuth())")
	fmt.Println("      router.GET(\"/test\", func(c *gin.Context) {")
	fmt.Println("          c.Status(http.StatusOK)")
	fmt.Println("      })")
	fmt.Println()
	fmt.Println("      // 有效令牌测试")
	fmt.Println("      t.Run(\"valid token\", func(t *testing.T) {")
	fmt.Println("          token, _ := GenerateToken(\"user123\", []string{\"admin\"})")
	fmt.Println("          w := httptest.NewRecorder()")
	fmt.Println("          req, _ := http.NewRequest(\"GET\", \"/test\", nil)")
	fmt.Println("          req.Header.Set(\"Authorization\", \"Bearer \"+token)")
	fmt.Println("          router.ServeHTTP(w, req)")
	fmt.Println("          assert.Equal(t, http.StatusOK, w.Code)")
	fmt.Println("      })")
	fmt.Println()
	fmt.Println("      // 无效令牌测试")
	fmt.Println("      t.Run(\"invalid token\", func(t *testing.T) {")
	fmt.Println("          w := httptest.NewRecorder()")
	fmt.Println("          req, _ := http.NewRequest(\"GET\", \"/test\", nil)")
	fmt.Println("          req.Header.Set(\"Authorization\", \"Bearer invalid_token\")")
	fmt.Println("          router.ServeHTTP(w, req)")
	fmt.Println("          assert.Equal(t, http.StatusUnauthorized, w.Code)")
	fmt.Println("      })")
	fmt.Println("  }")
	fmt.Println()

	fmt.Println("========== 3. 测试覆盖率统计 ==========")
	fmt.Println()
	fmt.Println("  # 生成测试覆盖率报告")
	fmt.Println("  go test -coverprofile=coverage.out")
	fmt.Println()
	fmt.Println("  # 查看HTML格式的覆盖率报告")
	fmt.Println("  go tool cover -html=coverage.out")
	fmt.Println()
	fmt.Println("  # 查看覆盖率百分比")
	fmt.Println("  go test -cover")
	fmt.Println()

	fmt.Println("========== 4. 测试最佳实践 ==========")
	fmt.Println()
	fmt.Println("  1. 使用表格驱动测试（Table-Driven Tests）")
	fmt.Println("  2. 测试边界条件和异常情况")
	fmt.Println("  3. 使用 mock 对象隔离依赖")
	fmt.Println("  4. 保持测试代码简洁可读")
	fmt.Println("  5. 测试覆盖率目标: 80%%+")
	fmt.Println()

	fmt.Println("========== 5. 运行测试 ==========")
	fmt.Println()
	fmt.Println("  # 运行所有测试")
	fmt.Println("  go test ./...")
	fmt.Println()
	fmt.Println("  # 运行特定包的测试")
	fmt.Println("  go test ./gin/2_middleware")
	fmt.Println()
	fmt.Println("  # 运行测试并显示详细输出")
	fmt.Println("  go test -v ./gin/2_middleware")
	fmt.Println()
	fmt.Println("  # 运行测试并显示覆盖率")
	fmt.Println("  go test -cover ./gin/2_middleware")
}

// GinRouterDemo 演示 gin.Default() 和 gin.New() 的区别
func GinRouterDemo() {
	fmt.Println("=== gin.Default() 和 gin.New() 详解 ===")
	fmt.Println()

	fmt.Println("========== gin.Default() ==========")
	fmt.Println()
	fmt.Println("gin.Default() 创建一个带有默认中间件的路由器")
	fmt.Println()
	fmt.Println("等价代码:")
	fmt.Println("  router := gin.New()")
	fmt.Println("  router.Use(gin.Logger())    // 添加Logger中间件")
	fmt.Println("  router.Use(gin.Recovery())  // 添加Recovery中间件")
	fmt.Println()
	fmt.Println("内置中间件:")
	fmt.Println("  1. Logger 中间件")
	fmt.Println("     - 自动记录HTTP请求日志")
	fmt.Println("     - 包括: 请求方法、路径、状态码、响应时间等")
	fmt.Println("     - 输出格式: [GIN] 2024/01/01 - 10:00:00 | 200 | 1.234ms | 127.0.0.1 | GET \"/api/users\"")
	fmt.Println()
	fmt.Println("  2. Recovery 中间件")
	fmt.Println("     - 自动捕获panic并恢复")
	fmt.Println("     - 避免程序因panic而崩溃")
	fmt.Println("     - 返回500错误响应，而不是让程序退出")
	fmt.Println()
	fmt.Println("使用场景:")
	fmt.Println("  ✅ 开发环境: 方便调试，自动记录日志")
	fmt.Println("  ✅ 快速原型: 快速搭建API，无需手动配置")
	fmt.Println("  ⚠️  生产环境: 性能略低，日志格式固定")
	fmt.Println()

	fmt.Println("========== gin.New() ==========")
	fmt.Println()
	fmt.Println("gin.New() 创建一个不包含任何中间件的路由器")
	fmt.Println()
	fmt.Println("特点:")
	fmt.Println("  - 轻量级，性能更好")
	fmt.Println("  - 完全自定义，按需添加中间件")
	fmt.Println("  - 适合生产环境")
	fmt.Println()
	fmt.Println("使用示例:")
	fmt.Println("  router := gin.New()")
	fmt.Println()
	fmt.Println("  // 按需添加中间件")
	fmt.Println("  router.Use(gin.Logger())")
	fmt.Println("  router.Use(gin.Recovery())")
	fmt.Println("  router.Use(CustomMiddleware())")
	fmt.Println()

	fmt.Println("========== 对比总结 ==========")
	fmt.Println()
	fmt.Println("| 特性 | gin.Default() | gin.New() |")
	fmt.Println("|------|---------------|-----------|")
	fmt.Println("| 内置中间件 | Logger + Recovery | 无 |")
	fmt.Println("| 性能 | 略低（有日志开销） | 更高 |")
	fmt.Println("| 灵活性 | 较低 | 高 |")
	fmt.Println("| 适用场景 | 开发/原型 | 生产环境 |")
	fmt.Println("| 代码量 | 少 | 需要手动添加中间件 |")
	fmt.Println()

	fmt.Println("========== 实际使用建议 ==========")
	fmt.Println()
	fmt.Println("开发环境:")
	fmt.Println("  router := gin.Default()  // 简单快速")
	fmt.Println()
	fmt.Println("生产环境:")
	fmt.Println("  router := gin.New()")
	fmt.Println("  router.Use(gin.Recovery())  // 必须添加，防止panic")
	fmt.Println("  router.Use(CustomLogger())  // 使用自定义日志中间件")
	fmt.Println("  router.Use(CORSMiddleware())")
	fmt.Println("  router.Use(JWTAuth())")
	fmt.Println()
	fmt.Println("测试环境:")
	fmt.Println("  router := gin.New()  // 不记录日志，测试更干净")
	fmt.Println()

	fmt.Println("========== 代码示例 ==========")
	fmt.Println()
	fmt.Println("// 使用 gin.Default()")
	fmt.Println("func main() {")
	fmt.Println("    router := gin.Default()")
	fmt.Println("    router.GET(\"/api/users\", GetUsersHandler)")
	fmt.Println("    router.Run(\":8080\")")
	fmt.Println("}")
	fmt.Println()
	fmt.Println("// 使用 gin.New()")
	fmt.Println("func main() {")
	fmt.Println("    router := gin.New()")
	fmt.Println("    router.Use(gin.Recovery())")
	fmt.Println("    router.Use(CustomLogger())")
	fmt.Println("    router.GET(\"/api/users\", GetUsersHandler)")
	fmt.Println("    router.Run(\":8080\")")
	fmt.Println("}")
}
