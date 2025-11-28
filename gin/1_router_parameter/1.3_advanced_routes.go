package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// RouteGroupDemo 演示路由分组配置
func RouteGroupDemo() {
	fmt.Println("=== Gin 路由分组示例 ===")
	fmt.Println()

	router := gin.Default()

	// API版本分组
	v1 := router.Group("/api/v1")
	{
		// 用户相关路由组
		users := v1.Group("/users")
		{
			users.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "获取用户列表",
					"version": "v1",
				})
			})

			users.POST("", func(c *gin.Context) {
				c.JSON(http.StatusCreated, gin.H{
					"message": "创建用户",
					"version": "v1",
				})
			})

			users.GET("/:id", func(c *gin.Context) {
				id := c.Param("id")
				c.JSON(http.StatusOK, gin.H{
					"message": "获取用户信息",
					"id":      id,
					"version": "v1",
				})
			})

			users.PUT("/:id", func(c *gin.Context) {
				id := c.Param("id")
				c.JSON(http.StatusOK, gin.H{
					"message": "更新用户",
					"id":      id,
					"version": "v1",
				})
			})

			users.DELETE("/:id", func(c *gin.Context) {
				id := c.Param("id")
				c.JSON(http.StatusOK, gin.H{
					"message": "删除用户",
					"id":      id,
					"version": "v1",
				})
			})
		}

		// 文章相关路由组
		posts := v1.Group("/posts")
		{
			posts.GET("", func(c *gin.Context) {
				c.JSON(http.StatusOK, gin.H{
					"message": "获取文章列表",
				})
			})

			posts.POST("", func(c *gin.Context) {
				c.JSON(http.StatusCreated, gin.H{
					"message": "创建文章",
				})
			})
		}
	}

	// API v2版本分组
	v2 := router.Group("/api/v2")
	{
		v2.GET("/users", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "获取用户列表 (v2)",
				"version": "v2",
			})
		})
	}

	fmt.Println("路由分组配置完成:")
	fmt.Println("  /api/v1/users      - GET    - 获取用户列表")
	fmt.Println("  /api/v1/users      - POST   - 创建用户")
	fmt.Println("  /api/v1/users/:id  - GET    - 获取用户")
	fmt.Println("  /api/v1/users/:id  - PUT    - 更新用户")
	fmt.Println("  /api/v1/users/:id  - DELETE - 删除用户")
	fmt.Println("  /api/v1/posts      - GET    - 获取文章列表")
	fmt.Println("  /api/v1/posts      - POST   - 创建文章")
	fmt.Println("  /api/v2/users      - GET    - 获取用户列表 (v2)")
	fmt.Println()
	fmt.Println("路由分组优势:")
	fmt.Println("  1. 代码组织更清晰")
	fmt.Println("  2. 便于添加中间件到特定路由组")
	fmt.Println("  3. 支持API版本控制")
	fmt.Println("  4. 减少重复的路径前缀")
}

// RegexRouteDemo 演示正则表达式路由
func RegexRouteDemo() {
	fmt.Println("=== Gin 正则表达式路由示例 ===")
	fmt.Println()

	router := gin.Default()

	// 1. 数字ID路由 - 只匹配数字
	router.GET("/users/:id([0-9]+)", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "数字ID路由",
			"id":      id,
		})
	})

	// 2. UUID路由 - 匹配UUID格式
	router.GET("/users/:uuid([0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12})", func(c *gin.Context) {
		uuid := c.Param("uuid")
		c.JSON(http.StatusOK, gin.H{
			"message": "UUID路由",
			"uuid":    uuid,
		})
	})

	// 3. 字母数字组合路由
	router.GET("/posts/:slug([a-z0-9-]+)", func(c *gin.Context) {
		slug := c.Param("slug")
		c.JSON(http.StatusOK, gin.H{
			"message": "Slug路由",
			"slug":    slug,
		})
	})

	fmt.Println("正则表达式路由示例:")
	fmt.Println("  GET /users/123     → 匹配数字ID")
	fmt.Println("  GET /users/abc    → 不匹配（非数字）")
	fmt.Println("  GET /users/550e8400-e29b-41d4-a716-446655440000 → 匹配UUID")
	fmt.Println("  GET /posts/my-post-123 → 匹配字母数字组合")
	fmt.Println()
	fmt.Println("正则表达式说明:")
	fmt.Println("  :id([0-9]+)                    - 只匹配数字")
	fmt.Println("  :uuid([0-9a-f]{8}-...)         - 匹配UUID格式")
	fmt.Println("  :slug([a-z0-9-]+)              - 匹配小写字母、数字和连字符")
	fmt.Println()
	fmt.Println("注意: 正则表达式必须用括号包裹，放在参数名后面")
}

// MiddlewareRouteDemo 演示路由中间件
func MiddlewareRouteDemo() {
	fmt.Println("=== Gin 路由中间件示例 ===")
	fmt.Println()

	router := gin.Default()

	// 自定义中间件 - 记录请求日志
	requestLogger := func(c *gin.Context) {
		fmt.Printf("[中间件] 请求: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next() // 继续处理下一个中间件或处理函数
	}

	// 自定义中间件 - 认证检查
	authMiddleware := func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "未授权，需要Token",
			})
			c.Abort() // 终止请求处理
			return
		}
		c.Set("user", "authenticated_user") // 设置上下文值
		c.Next()
	}

	// 全局中间件
	router.Use(requestLogger)

	// 公开路由 - 不需要认证
	public := router.Group("/public")
	{
		public.GET("/info", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "公开信息",
			})
		})
	}

	// 受保护的路由组 - 需要认证
	protected := router.Group("/api")
	protected.Use(authMiddleware) // 应用认证中间件
	{
		protected.GET("/profile", func(c *gin.Context) {
			user, _ := c.Get("user")
			c.JSON(http.StatusOK, gin.H{
				"message": "用户资料",
				"user":    user,
			})
		})

		protected.GET("/dashboard", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "仪表盘数据",
			})
		})
	}

	fmt.Println("中间件配置完成:")
	fmt.Println("  全局中间件: 所有请求都会记录日志")
	fmt.Println("  公开路由: /public/info (无需认证)")
	fmt.Println("  受保护路由: /api/profile, /api/dashboard (需要Authorization头)")
	fmt.Println()
	fmt.Println("中间件说明:")
	fmt.Println("  router.Use(middleware)      - 全局中间件")
	fmt.Println("  group.Use(middleware)       - 路由组中间件")
	fmt.Println("  router.GET(path, m1, m2, handler) - 单个路由中间件")
	fmt.Println()
	fmt.Println("中间件函数签名:")
	fmt.Println("  func(c *gin.Context) {")
	fmt.Println("    // 前置处理")
	fmt.Println("    c.Next()  // 继续处理")
	fmt.Println("    // 后置处理")
	fmt.Println("  }")
}

// StaticFilesDemo 演示静态文件服务
func StaticFilesDemo() {
	fmt.Println("=== Gin 静态文件服务示例 ===")
	fmt.Println()

	router := gin.Default()

	// 1. 静态文件目录
	router.Static("/static", "./static")

	// 2. 静态文件服务（带虚拟路径前缀）
	router.StaticFS("/assets", gin.Dir("./assets", false))

	// 3. 单个静态文件
	router.StaticFile("/favicon.ico", "./favicon.ico")
	router.StaticFile("/robots.txt", "./robots.txt")

	// 4. HTML模板渲染
	router.LoadHTMLGlob("templates/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Gin Web应用",
		})
	})

	fmt.Println("静态文件配置完成:")
	fmt.Println("  /static/*          → ./static/ 目录下的文件")
	fmt.Println("  /assets/*          → ./assets/ 目录下的文件")
	fmt.Println("  /favicon.ico       → ./favicon.ico")
	fmt.Println("  /robots.txt        → ./robots.txt")
	fmt.Println()
	fmt.Println("方法说明:")
	fmt.Println("  router.Static(relativePath, root) - 静态文件目录")
	fmt.Println("  router.StaticFS(relativePath, fs) - 使用文件系统")
	fmt.Println("  router.StaticFile(relativePath, filepath) - 单个文件")
	fmt.Println("  router.LoadHTMLGlob(pattern) - 加载HTML模板")
}

