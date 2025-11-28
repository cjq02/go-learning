package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BasicRoutesDemo 演示基础路由定义方式
// 本函数展示了 Gin 框架中最基础的路由注册方法，包括：
// 1. GET 路由 - 处理查询参数
// 2. POST 路由 - 处理表单数据
// 3. Any 路由 - 支持所有 HTTP 方法
func BasicRoutesDemo() {
	fmt.Println("=== Gin 基础路由示例 ===")
	fmt.Println()

	// ========== 创建路由引擎 ==========
	// gin.Default() 创建一个带有默认中间件的路由引擎
	// 默认包含两个中间件：
	//   - Logger: 自动记录 HTTP 请求日志（请求方法、路径、状态码、响应时间等）
	//   - Recovery: 捕获 panic 并返回 500 错误，防止程序崩溃
	//
	// 如果不需要默认中间件，可以使用 gin.New() 创建空的路由引擎
	// 例如: router := gin.New()
	router := gin.Default()

	// ========== 1. 基础 GET 路由示例 ==========
	// router.GET() 注册一个 GET 方法的路由
	// 第一个参数: 路由路径 "/welcome"
	// 第二个参数: 处理函数（Handler），接收 *gin.Context 参数
	//
	// 使用场景: 获取数据、查询信息等只读操作
	// 示例请求: GET http://localhost:8080/welcome?firstname=John&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		// gin.Context 封装了 HTTP 请求和响应的所有信息
		// 提供了获取参数、设置响应、操作 Cookie 等方法

		// c.DefaultQuery() 获取查询参数，如果参数不存在则返回默认值
		// 语法: c.DefaultQuery("参数名", "默认值")
		// 示例: /welcome?firstname=John → firstName = "John"
		//        /welcome → firstName = "Guest" (使用默认值)
		firstName := c.DefaultQuery("firstname", "Guest")

		// c.Query() 获取查询参数，如果参数不存在则返回空字符串 ""
		// 语法: c.Query("参数名")
		// 示例: /welcome?lastname=Doe → lastName = "Doe"
		//        /welcome → lastName = "" (空字符串)
		//
		// 区别:
		//   - DefaultQuery: 有默认值，适合可选参数
		//   - Query: 无默认值，需要自己判断是否为空
		lastName := c.Query("lastname")

		// c.String() 返回纯文本响应
		// 第一个参数: HTTP 状态码 (http.StatusOK = 200)
		// 第二个参数: 格式化字符串（类似 fmt.Sprintf）
		// 后续参数: 格式化参数
		// 自动设置 Content-Type: text/plain
		c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	})

	// ========== 2. POST 路由示例 ==========
	// router.POST() 注册一个 POST 方法的路由
	// 使用场景: 提交表单、创建资源等写操作
	// 示例请求: POST http://localhost:8080/submit
	//           Content-Type: application/x-www-form-urlencoded
	//           Body: name=张三&email=zhangsan@example.com
	router.POST("/submit", func(c *gin.Context) {
		// c.PostForm() 获取 POST 请求中的表单字段
		// 适用于两种 Content-Type:
		//   1. application/x-www-form-urlencoded (标准表单)
		//   2. multipart/form-data (文件上传)
		//
		// 语法: c.PostForm("字段名")
		// 如果字段不存在，返回空字符串 ""
		name := c.PostForm("name")
		email := c.PostForm("email")

		// c.JSON() 返回 JSON 格式的响应
		// 第一个参数: HTTP 状态码
		// 第二个参数: 要序列化的对象（可以是结构体、map 等）
		//
		// gin.H 是 map[string]interface{} 的快捷类型别名
		// 使用 gin.H 可以方便地创建 JSON 对象
		// 自动设置 Content-Type: application/json
		// 自动将 Go 对象序列化为 JSON 字符串
		c.JSON(http.StatusOK, gin.H{
			"message": "提交成功",
			"name":    name,
			"email":   email,
		})
	})

	// ========== 3. Any 路由 - 支持所有 HTTP 方法 ==========
	// router.Any() 注册一个支持所有 HTTP 方法的路由
	// 支持的方法包括: GET, POST, PUT, DELETE, PATCH, HEAD, OPTIONS 等
	//
	// 使用场景:
	//   - 需要处理多种 HTTP 方法的同一路径
	//   - 调试和测试
	//   - 简单的 API 网关
	router.Any("/any", func(c *gin.Context) {
		// c.Request.Method 获取当前请求的 HTTP 方法
		// 返回字符串，如: "GET", "POST", "PUT", "DELETE" 等
		// c.Request 是标准库的 *http.Request 对象
		c.String(http.StatusOK, "支持所有HTTP方法: %s", c.Request.Method)
	})

	fmt.Println("基础路由配置完成:")
	fmt.Println("  GET  /welcome?firstname=John&lastname=Doe")
	fmt.Println("  POST /submit (form-data: name, email)")
	fmt.Println("  ANY  /any (支持所有HTTP方法)")
	fmt.Println()
	fmt.Println("注意: 此示例仅展示路由配置，实际运行需要启动服务器")
	fmt.Println("      可以使用 router.Run(\":8080\") 启动服务器")
	fmt.Println()
	fmt.Println("测试示例:")
	fmt.Println("  curl \"http://localhost:8080/welcome?firstname=John&lastname=Doe\"")
	fmt.Println("  curl -X POST http://localhost:8080/submit -d \"name=张三&email=test@example.com\"")
	fmt.Println("  curl -X GET http://localhost:8080/any")
	fmt.Println("  curl -X POST http://localhost:8080/any")
}

// RESTfulRoutesDemo 演示 RESTful 风格的路由定义
// RESTful (Representational State Transfer) 是一种 API 设计风格
// 核心原则:
//  1. 使用 HTTP 方法表示操作 (GET=查询, POST=创建, PUT=更新, DELETE=删除)
//  2. 使用 URL 表示资源 (名词，如 /users, /products)
//  3. 使用 HTTP 状态码表示结果 (200=成功, 201=创建, 400=错误, 404=不存在)
//  4. 使用 JSON 格式传输数据
//
// RESTful API 设计规范:
//
//	GET    /users      → 获取用户列表 (查询操作，不改变资源状态)
//	POST   /users      → 创建新用户 (创建资源)
//	GET    /users/:id  → 获取指定用户 (查询单个资源)
//	PUT    /users/:id  → 更新指定用户 (完整更新资源)
//	DELETE /users/:id  → 删除指定用户 (删除资源)
func RESTfulRoutesDemo() {
	fmt.Println("=== Gin RESTful 路由示例 ===")
	fmt.Println()

	router := gin.Default()

	// ========== RESTful API 路由设计 ==========
	// 以下展示了标准的 RESTful CRUD 操作
	// CRUD = Create(创建), Read(读取), Update(更新), Delete(删除)

	// ========== 1. 获取用户列表 - GET /users ==========
	// 功能: 查询所有用户
	// HTTP 方法: GET (只读操作，不改变资源状态)
	// 路径: /users (资源集合)
	// 状态码: 200 OK (成功)
	//
	// 实际应用中应该:
	//   - 从数据库查询用户列表
	//   - 支持分页 (page, limit)
	//   - 支持排序 (sort)
	//   - 支持过滤 (filter)
	router.GET("/users", func(c *gin.Context) {
		// 返回用户列表（示例中只返回提示信息）
		// 实际应该返回: []User 数组
		c.JSON(http.StatusOK, gin.H{
			"message": "获取用户列表",
			"method":  "GET",
		})
	})

	// ========== 2. 创建新用户 - POST /users ==========
	// 功能: 创建新的用户资源
	// HTTP 方法: POST (创建操作)
	// 路径: /users (资源集合)
	// 状态码: 201 Created (资源创建成功)
	//
	// 请求体格式: JSON
	// {
	//   "name": "张三",
	//   "email": "zhangsan@example.com"
	// }
	router.POST("/users", func(c *gin.Context) {
		// 定义匿名结构体用于接收 JSON 数据
		// 结构体标签说明:
		//   json:"name"  - JSON 字段名映射到 Go 结构体字段
		//   binding:"required" - 验证规则：字段必填
		//   binding:"email" - 验证规则：必须是有效的邮箱格式
		var user struct {
			Name  string `json:"name" binding:"required"`        // name 字段，必填
			Email string `json:"email" binding:"required,email"` // email 字段，必填且必须是邮箱格式
		}

		// c.ShouldBindJSON() 将请求体中的 JSON 数据绑定到结构体
		// 功能:
		//   1. 解析 JSON 请求体
		//   2. 将数据填充到结构体字段
		//   3. 执行验证规则（binding 标签）
		//   4. 如果验证失败，返回错误
		//
		// 返回值:
		//   - error == nil: 绑定成功
		//   - error != nil: 绑定失败（格式错误、验证失败等）
		if err := c.ShouldBindJSON(&user); err != nil {
			// 绑定失败，返回 400 Bad Request 错误
			// http.StatusBadRequest = 400，表示客户端请求参数错误
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return // 提前返回，不执行后续代码
		}

		// 绑定成功，返回 201 Created 状态码
		// http.StatusCreated = 201，表示资源创建成功
		// RESTful 规范中，创建资源应该返回 201 而不是 200
		c.JSON(http.StatusCreated, gin.H{
			"message": "用户创建成功",
			"user":    user, // 返回创建的用户信息
		})
	})

	// ========== 3. 获取指定用户 - GET /users/:id ==========
	// 功能: 根据 ID 查询单个用户
	// HTTP 方法: GET (只读操作)
	// 路径: /users/:id (单个资源，:id 是路径参数)
	// 状态码: 200 OK (成功) 或 404 Not Found (用户不存在)
	//
	// 路径参数说明:
	//   /users/:id 中的 :id 是路径参数（Path Parameter）
	//   示例: GET /users/123 → id = "123"
	//   示例: GET /users/456 → id = "456"
	//
	// 路径参数 vs 查询参数:
	//   路径参数: /users/:id → /users/123 (资源标识，必需)
	//   查询参数: /users?id=123 (过滤条件，可选)
	router.GET("/users/:id", func(c *gin.Context) {
		// c.Param() 获取路径参数
		// 语法: c.Param("参数名")
		// 参数名对应路由中的 :参数名
		// 示例: 路由 /users/:id，请求 /users/123 → c.Param("id") = "123"
		//
		// 注意: 返回的是字符串类型，如果需要数字类型需要转换
		// 例如: idInt, _ := strconv.Atoi(c.Param("id"))
		id := c.Param("id")

		// 实际应用中应该:
		//   1. 将 id 转换为数字类型
		//   2. 从数据库查询用户
		//   3. 如果用户不存在，返回 404 Not Found
		c.JSON(http.StatusOK, gin.H{
			"message": "获取用户信息",
			"id":      id,
		})
	})

	// ========== 4. 更新指定用户 - PUT /users/:id ==========
	// 功能: 完整更新用户信息（替换整个资源）
	// HTTP 方法: PUT (更新操作，完整更新)
	// 路径: /users/:id (单个资源)
	// 状态码: 200 OK (成功) 或 404 Not Found (用户不存在)
	//
	// PUT vs PATCH:
	//   PUT: 完整更新，需要提供所有字段（即使不修改也要提供）
	//   PATCH: 部分更新，只需要提供要修改的字段
	//
	// 请求体格式: JSON
	// {
	//   "name": "李四",
	//   "email": "lisi@example.com"
	// }
	router.PUT("/users/:id", func(c *gin.Context) {
		// 获取路径参数（用户ID）
		id := c.Param("id")

		// 定义结构体接收更新数据
		// 注意: 这里没有使用 binding:"required"，因为 PUT 是完整更新
		// 实际应用中可能需要根据业务需求决定哪些字段必填
		var user struct {
			Name  string `json:"name"`  // 可选的 name 字段
			Email string `json:"email"` // 可选的 email 字段
		}

		// 绑定 JSON 数据
		if err := c.ShouldBindJSON(&user); err != nil {
			// 绑定失败，返回 400 错误
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// 实际应用中应该:
		//   1. 验证用户是否存在
		//   2. 更新数据库中的用户信息
		//   3. 返回更新后的用户信息
		c.JSON(http.StatusOK, gin.H{
			"message": "用户更新成功",
			"id":      id,
			"user":    user,
		})
	})

	// ========== 5. 删除指定用户 - DELETE /users/:id ==========
	// 功能: 删除用户资源
	// HTTP 方法: DELETE (删除操作)
	// 路径: /users/:id (单个资源)
	// 状态码: 200 OK (成功) 或 204 No Content (成功但无返回内容) 或 404 Not Found (用户不存在)
	//
	// 注意: DELETE 请求通常不需要请求体，只需要路径参数
	router.DELETE("/users/:id", func(c *gin.Context) {
		// 获取要删除的用户ID
		id := c.Param("id")

		// 实际应用中应该:
		//   1. 验证用户是否存在
		//   2. 执行删除操作（软删除或硬删除）
		//   3. 返回删除结果
		//
		// 也可以返回 204 No Content（成功但无返回内容）
		// 例如: c.Status(http.StatusNoContent)
		c.JSON(http.StatusOK, gin.H{
			"message": "用户删除成功",
			"id":      id,
		})
	})

	fmt.Println("RESTful 路由配置完成:")
	fmt.Println("  GET    /users      - 获取用户列表")
	fmt.Println("  POST   /users      - 创建新用户")
	fmt.Println("  GET    /users/:id  - 获取指定用户")
	fmt.Println("  PUT    /users/:id  - 更新指定用户")
	fmt.Println("  DELETE /users/:id  - 删除指定用户")
	fmt.Println()
	fmt.Println("路径参数使用 :param 格式，通过 c.Param(\"param\") 获取")
	fmt.Println()
	fmt.Println("测试示例:")
	fmt.Println("  # 获取用户列表")
	fmt.Println("  curl http://localhost:8080/users")
	fmt.Println()
	fmt.Println("  # 创建用户")
	fmt.Println("  curl -X POST http://localhost:8080/users \\")
	fmt.Println("    -H \"Content-Type: application/json\" \\")
	fmt.Println("    -d '{\"name\":\"张三\",\"email\":\"zhangsan@example.com\"}'")
	fmt.Println()
	fmt.Println("  # 获取指定用户")
	fmt.Println("  curl http://localhost:8080/users/123")
	fmt.Println()
	fmt.Println("  # 更新用户")
	fmt.Println("  curl -X PUT http://localhost:8080/users/123 \\")
	fmt.Println("    -H \"Content-Type: application/json\" \\")
	fmt.Println("    -d '{\"name\":\"李四\",\"email\":\"lisi@example.com\"}'")
	fmt.Println()
	fmt.Println("  # 删除用户")
	fmt.Println("  curl -X DELETE http://localhost:8080/users/123")
	fmt.Println()
	fmt.Println("关键概念:")
	fmt.Println("  1. 路径参数: /users/:id 中的 :id 是路径参数")
	fmt.Println("  2. JSON 绑定: 使用 c.ShouldBindJSON() 自动解析和验证")
	fmt.Println("  3. 状态码: 遵循 RESTful 规范使用正确的 HTTP 状态码")
	fmt.Println("  4. HTTP 方法: GET(查询), POST(创建), PUT(更新), DELETE(删除)")
}
