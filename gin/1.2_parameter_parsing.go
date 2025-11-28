package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PathParameterDemo 演示路径参数获取
// 路径参数（Path Parameter）是 URL 路径中的一部分，用于标识资源
// 例如: /users/123 中的 123 就是路径参数
//
// 路径参数的特点:
//  1. 是 URL 路径的一部分，不是查询字符串
//  2. 用于标识资源（如用户ID、文章ID等）
//  3. 必需参数，缺少会导致路由不匹配
//  4. 使用 :param 格式定义，使用 *param 定义通配符
//
// 路径参数 vs 查询参数:
//
//	路径参数: /users/:id → /users/123 (资源标识，必需)
//	查询参数: /users?id=123 (过滤条件，可选)
func PathParameterDemo() {
	fmt.Println("=== Gin 路径参数解析示例 ===")
	fmt.Println()

	router := gin.Default()

	// ========== 重要提示：路由冲突问题 ==========
	// Gin 的路由树结构中，如果两个路由共享相同的前缀且在同一位置都有路径参数，
	// 会导致路由冲突错误
	//
	// 冲突示例:
	//   /users/:id 和 /users/:userId/posts/:postId 会冲突
	//   因为它们在 /users/ 后的第一个位置都有路径参数
	//
	// 解决方案:
	//   1. 使用不同的路径前缀（推荐）
	//   2. 调整路由结构，避免在同一位置有路径参数
	//
	// 本示例使用不同的路径前缀来避免冲突

	// ========== 1. 单个路径参数 ==========
	// 路由格式: /users/:id
	// 说明: :id 是路径参数占位符，可以匹配任意字符串
	// 使用场景: 获取单个资源（如用户详情、文章详情等）
	//
	// 匹配示例:
	//   ✅ /users/123        → id = "123"
	//   ✅ /users/456        → id = "456"
	//   ✅ /users/abc        → id = "abc" (注意：返回的是字符串)
	//   ❌ /users            → 不匹配（缺少参数）
	//   ❌ /users/123/posts  → 不匹配（路径不匹配）
	router.GET("/users/:id", func(c *gin.Context) {
		// c.Param() 获取路径参数
		// 语法: c.Param("参数名")
		// 参数名对应路由定义中的 :参数名
		// 返回值: string 类型（即使路径中是数字，也会作为字符串返回）
		//
		// 如果需要数字类型，需要手动转换:
		//   idInt, err := strconv.Atoi(c.Param("id"))
		id := c.Param("id")

		c.JSON(http.StatusOK, gin.H{
			"message": "获取用户ID",
			"id":      id,
		})
	})

	// ========== 2. 多个路径参数 ==========
	// 路由格式: /user/:userId/posts/:postId
	// 说明: 可以在一个路由中定义多个路径参数
	// 使用场景: 嵌套资源（如用户的文章、订单的商品等）
	//
	// 注意: 使用 /user/ 而不是 /users/ 来避免与上面的路由冲突
	// 实际应用中可以根据业务需求选择合适的路径结构
	//
	// 匹配示例:
	//   ✅ /user/123/posts/456     → userId = "123", postId = "456"
	//   ✅ /user/789/posts/101      → userId = "789", postId = "101"
	//   ❌ /user/123/posts          → 不匹配（缺少 postId）
	//   ❌ /user/123                → 不匹配（缺少 posts/:postId 部分）
	router.GET("/user/:userId/posts/:postId", func(c *gin.Context) {
		// 获取第一个路径参数
		userId := c.Param("userId")

		// 获取第二个路径参数
		postId := c.Param("postId")

		// 注意: 参数名必须与路由定义中的名称完全一致（区分大小写）
		c.JSON(http.StatusOK, gin.H{
			"userId": userId,
			"postId": postId,
		})
	})

	// ========== 3. 通配符路径参数 ==========
	// 路由格式: /files/*filepath
	// 说明: *filepath 是通配符参数，可以匹配路径的剩余部分
	// 使用场景: 文件服务、静态资源、需要匹配多级路径的场景
	//
	// 通配符特点:
	//   1. 使用 * 而不是 : 定义
	//   2. 可以匹配多级路径
	//   3. 返回的值包含前导斜杠 /
	//
	// 匹配示例:
	//   ✅ /files/images/photo.jpg           → filepath = "/images/photo.jpg"
	//   ✅ /files/docs/readme.md             → filepath = "/docs/readme.md"
	//   ✅ /files/static/css/style.css       → filepath = "/static/css/style.css"
	//   ❌ /files                             → 不匹配（通配符至少需要一个斜杠）
	router.GET("/files/*filepath", func(c *gin.Context) {
		// 获取通配符参数
		// 注意: 返回的值包含前导斜杠 /
		// 例如: /files/images/photo.jpg → filepath = "/images/photo.jpg"
		filepath := c.Param("filepath")

		c.JSON(http.StatusOK, gin.H{
			"filepath": filepath,
		})
	})

	fmt.Println("路径参数示例:")
	fmt.Println("  GET /users/123              → id = \"123\"")
	fmt.Println("  GET /user/123/posts/456     → userId = \"123\", postId = \"456\"")
	fmt.Println("  GET /files/images/photo.jpg → filepath = \"/images/photo.jpg\"")
	fmt.Println()
	fmt.Println("========== 重要概念详解 ==========")
	fmt.Println()
	fmt.Println("1. 通配符参数会包含前导斜杠:")
	fmt.Println("   路由: /files/*filepath")
	fmt.Println("   请求: /files/images/photo.jpg")
	fmt.Println("   结果: filepath = \"/images/photo.jpg\" (注意：包含前导斜杠 /)")
	fmt.Println("   原因: Gin 会保留通配符匹配到的完整路径部分，包括第一个斜杠")
	fmt.Println("   用途: 这样可以直接用于文件路径操作，无需手动添加斜杠")
	fmt.Println()
	fmt.Println("2. 避免路由冲突：同一前缀下不能在同一位置有不同名称的路径参数")
	fmt.Println("   冲突示例:")
	fmt.Println("     ❌ /users/:id 和 /users/:userId/posts")
	fmt.Println("        原因: 两个路由在 /users/ 后的第一个位置都有路径参数")
	fmt.Println("        结果: Gin 无法区分，会报错: 'conflicting route'")
	fmt.Println()
	fmt.Println("   正确做法:")
	fmt.Println("     ✅ 方案1: 使用不同前缀")
	fmt.Println("        /users/:id")
	fmt.Println("        /user/:userId/posts  (注意：user 单数，避免冲突)")
	fmt.Println()
	fmt.Println("     ✅ 方案2: 调整路由结构")
	fmt.Println("        /users/:id")
	fmt.Println("        /users/posts/:postId  (posts 是固定路径，不是参数)")
	fmt.Println()
	fmt.Println("   为什么会有这个限制？")
	fmt.Println("     Gin 使用路由树（Radix Tree）来匹配路由，当两个路由在")
	fmt.Println("     同一位置都有参数时，无法确定应该匹配哪一个")
	fmt.Println()
	fmt.Println("测试示例:")
	fmt.Println("  curl http://localhost:8080/users/123")
	fmt.Println("  curl http://localhost:8080/user/123/posts/456")
	fmt.Println("  curl http://localhost:8080/files/images/photo.jpg")
}

// QueryParameterDemo 演示查询参数获取
// 查询参数（Query Parameter）是 URL 中 ? 后面的键值对
// 例如: /welcome?firstname=John&lastname=Doe
//
// 查询参数的特点:
//  1. 位于 URL 的查询字符串部分（? 后面）
//  2. 用于过滤、排序、分页等操作
//  3. 可选参数，可以省略
//  4. 可以有多个同名参数（数组参数）
//
// 查询参数 vs 路径参数:
//
//	查询参数: /users?page=1&limit=10 (过滤条件，可选)
//	路径参数: /users/:id (资源标识，必需)
func QueryParameterDemo() {
	fmt.Println("=== Gin 查询参数解析示例 ===")
	fmt.Println()

	router := gin.Default()

	// ========== 1. 基础查询参数 ==========
	// 查询参数格式: ?key=value&key2=value2
	// 使用场景: 过滤、排序、分页、搜索等
	//
	// 示例 URL:
	//   /welcome?firstname=John&lastname=Doe&age=25
	router.GET("/welcome", func(c *gin.Context) {
		// c.Query() - 获取查询参数，如果不存在返回空字符串 ""
		// 语法: c.Query("参数名")
		// 返回值: string
		//
		// 使用场景: 可选参数，需要自己判断是否为空
		// 示例: /welcome?firstname=John → firstName = "John"
		//        /welcome → firstName = "" (空字符串)
		firstName := c.Query("firstname")

		// c.DefaultQuery() - 获取查询参数，如果不存在返回默认值
		// 语法: c.DefaultQuery("参数名", "默认值")
		// 返回值: string
		//
		// 使用场景: 可选参数，但有合理的默认值
		// 示例: /welcome?lastname=Doe → lastName = "Doe"
		//        /welcome → lastName = "Guest" (使用默认值)
		//
		// 区别:
		//   Query(): 不存在返回 ""，需要手动判断
		//   DefaultQuery(): 不存在返回默认值，更安全
		lastName := c.DefaultQuery("lastname", "Guest")

		// c.GetQuery() - 获取查询参数，同时返回值和是否存在标志
		// 语法: value, exists := c.GetQuery("参数名")
		// 返回值: (string, bool)
		//   - value: 参数值（不存在时为空字符串）
		//   - exists: 参数是否存在（true/false）
		//
		// 使用场景: 需要明确知道参数是否存在
		// 优势: 可以区分"参数值为空字符串"和"参数不存在"两种情况
		age, exists := c.GetQuery("age")
		if !exists {
			// 参数不存在时的处理
			age = "未知"
		}
		// 注意: 如果参数存在但值为空，exists = true, age = ""

		c.String(http.StatusOK, "Hello %s %s, 年龄: %s", firstName, lastName, age)
	})

	// ========== 2. 数组查询参数 ==========
	// 数组查询参数: 同一个参数名可以有多个值
	// 格式: ?tag=go&tag=gin&tag=web
	// 使用场景: 多选标签、多条件过滤等
	//
	// 示例 URL:
	//   /tags?tag=go&tag=gin&tag=web
	//   结果: tags = ["go", "gin", "web"]
	router.GET("/tags", func(c *gin.Context) {
		// c.QueryArray() - 获取多个同名查询参数
		// 语法: c.QueryArray("参数名")
		// 返回值: []string (字符串切片)
		//
		// 如果参数不存在，返回空切片 []string{}
		// 如果参数只有一个值，返回包含一个元素的切片
		//
		// 示例:
		//   /tags?tag=go&tag=gin → tags = ["go", "gin"]
		//   /tags?tag=go         → tags = ["go"]
		//   /tags                → tags = [] (空切片)
		tags := c.QueryArray("tag")

		c.JSON(http.StatusOK, gin.H{
			"tags": tags,
		})
	})

	// ========== 3. Map 查询参数 ==========
	// Map 查询参数: 使用方括号表示嵌套结构
	// 格式: ?filter[status]=active&filter[type]=user
	// 使用场景: 复杂的过滤条件、嵌套的查询参数
	//
	// 示例 URL:
	//   /filters?filter[status]=active&filter[type]=user&filter[age]=18
	//   结果: filter = {"status": "active", "type": "user", "age": "18"}
	router.GET("/filters", func(c *gin.Context) {
		// c.QueryMap() - 获取查询参数映射
		// 语法: c.QueryMap("前缀")
		// 返回值: map[string]string
		//
		// 工作原理:
		//   查询参数格式: prefix[key]=value
		//   例如: filter[status]=active → map["status"] = "active"
		//
		// 示例:
		//   /filters?filter[status]=active&filter[type]=user
		//   → queryMap = {"status": "active", "type": "user"}
		//
		// 如果参数不存在，返回空 map map[string]string{}
		queryMap := c.QueryMap("filter")

		c.JSON(http.StatusOK, gin.H{
			"filters": queryMap,
		})
	})

	fmt.Println("查询参数示例:")
	fmt.Println("  GET /welcome?firstname=John&lastname=Doe&age=25")
	fmt.Println("  GET /tags?tag=go&tag=gin&tag=web")
	fmt.Println("  GET /filters?filter[status]=active&filter[type]=user")
	fmt.Println()
	fmt.Println("方法说明:")
	fmt.Println("  c.Query(\"key\")           - 获取单个参数，不存在返回空字符串")
	fmt.Println("  c.DefaultQuery(\"key\", \"default\") - 获取参数，不存在返回默认值")
	fmt.Println("  c.GetQuery(\"key\")        - 返回值和是否存在标志")
	fmt.Println("  c.QueryArray(\"key\")      - 获取数组参数")
	fmt.Println("  c.QueryMap(\"key\")        - 获取Map参数")
	fmt.Println()
	fmt.Println("测试示例:")
	fmt.Println("  curl \"http://localhost:8080/welcome?firstname=John&lastname=Doe&age=25\"")
	fmt.Println("  curl \"http://localhost:8080/tags?tag=go&tag=gin&tag=web\"")
	fmt.Println("  curl \"http://localhost:8080/filters?filter[status]=active&filter[type]=user\"")
}

// JSONBindingDemo 演示 JSON 参数绑定
// JSON 绑定是将 HTTP 请求体中的 JSON 数据自动解析并填充到 Go 结构体中
//
// 优势:
//  1. 自动类型转换（JSON 字符串 → Go 类型）
//  2. 自动数据验证（使用 binding 标签）
//  3. 代码更简洁（不需要手动解析 JSON）
//  4. 类型安全（编译时检查）
//
// 使用场景:
//   - RESTful API 的 POST/PUT 请求
//   - 前端发送 JSON 数据
//   - 微服务之间的数据交换
func JSONBindingDemo() {
	fmt.Println("=== Gin JSON 参数绑定示例 ===")
	fmt.Println()

	router := gin.Default()

	// ========== 定义数据结构 ==========
	// 定义登录结构体，用于接收 JSON 数据
	// 结构体标签说明:
	//   json:"user" - JSON 字段名映射到 Go 结构体字段
	//   binding:"required" - 验证规则：字段必填
	//   binding:"min=6" - 验证规则：最小长度为 6
	//
	// 注意: 结构体字段名首字母必须大写（导出字段），才能被 JSON 包访问
	type Login struct {
		User     string `json:"user" binding:"required"`           // 用户名，必填
		Password string `json:"password" binding:"required,min=6"` // 密码，必填且至少 6 位
	}

	// ========== JSON 绑定示例 - 登录接口 ==========
	// 请求格式:
	//   POST /login
	//   Content-Type: application/json
	//   Body: {"user": "admin", "password": "123456"}
	router.POST("/login", func(c *gin.Context) {
		// 声明变量，用于接收绑定的数据
		var login Login

		// c.ShouldBindJSON() - 将请求体中的 JSON 数据绑定到结构体
		// 功能:
		//   1. 读取请求体（Request Body）
		//   2. 解析 JSON 字符串
		//   3. 将 JSON 字段映射到结构体字段（根据 json 标签）
		//   4. 执行数据验证（根据 binding 标签）
		//   5. 类型转换（JSON 类型 → Go 类型）
		//
		// 返回值:
		//   - error == nil: 绑定成功
		//   - error != nil: 绑定失败（JSON 格式错误、验证失败等）
		//
		// 注意: 必须传递指针 &login，因为需要修改结构体的值
		if err := c.ShouldBindJSON(&login); err != nil {
			// 绑定失败，返回 400 Bad Request
			// 错误信息包含详细的验证失败原因
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "参数校验失败",
				"message": err.Error(), // 详细的错误信息
			})
			return // 提前返回，不执行后续代码
		}

		// 绑定成功，login 结构体已填充数据
		// 此时可以安全地使用 login.User 和 login.Password
		// 实际应用中应该在这里执行认证逻辑（验证用户名密码、生成 Token 等）

		// 认证逻辑（示例）
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
			"user":    login.User,
		})
	})

	// ========== 多种绑定方式示例 - 注册接口 ==========
	// 使用匿名结构体，适用于简单的、不需要复用的场景
	router.POST("/register", func(c *gin.Context) {
		// 定义匿名结构体，直接在函数内部使用
		// 结构体标签说明:
		//   json:"username" - JSON 字段名
		//   binding:"required" - 必填验证
		//   binding:"email" - 邮箱格式验证
		//   binding:"gte=18" - 大于等于 18（Greater Than or Equal）
		//   binding:"lte=100" - 小于等于 100（Less Than or Equal）
		var user struct {
			Username string `json:"username" binding:"required"`    // 用户名，必填
			Email    string `json:"email" binding:"required,email"` // 邮箱，必填且必须是邮箱格式
			Age      int    `json:"age" binding:"gte=18,lte=100"`   // 年龄，必须在 18-100 之间
		}

		// c.ShouldBind() - 根据 Content-Type 自动选择绑定方式
		// 支持的 Content-Type:
		//   - application/json → JSON 绑定
		//   - application/xml → XML 绑定
		//   - application/x-www-form-urlencoded → 表单绑定
		//   - multipart/form-data → 表单绑定（文件上传）
		//
		// 与 ShouldBindJSON() 的区别:
		//   ShouldBindJSON(): 只绑定 JSON 格式，Content-Type 必须是 application/json
		//   ShouldBind(): 自动识别 Content-Type，支持多种格式
		//
		// 使用建议:
		//   - 如果确定是 JSON，使用 ShouldBindJSON() 更明确
		//   - 如果需要支持多种格式，使用 ShouldBind()
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 绑定成功，返回 201 Created 状态码
		c.JSON(http.StatusCreated, gin.H{
			"message": "注册成功",
			"user":    user,
		})
	})

	fmt.Println("JSON绑定示例:")
	fmt.Println("  POST /login")
	fmt.Println("  Content-Type: application/json")
	fmt.Println("  Body: {\"user\": \"admin\", \"password\": \"123456\"}")
	fmt.Println()
	fmt.Println("绑定方法说明:")
	fmt.Println("  c.ShouldBindJSON(&struct) - 仅绑定JSON格式，Content-Type必须是application/json")
	fmt.Println("  c.ShouldBind(&struct)     - 根据Content-Type自动选择绑定方式（JSON/XML/表单）")
	fmt.Println("  c.MustBindJSON(&struct)   - 绑定失败会自动返回400错误，不需要手动处理")
	fmt.Println()
	fmt.Println("验证标签说明:")
	fmt.Println("  required  - 必填字段，不能为空")
	fmt.Println("  min=6     - 最小长度6（字符串）或最小值6（数字）")
	fmt.Println("  max=100   - 最大长度100（字符串）或最大值100（数字）")
	fmt.Println("  email     - 必须是有效的邮箱格式")
	fmt.Println("  gte=18    - 大于等于18（Greater Than or Equal）")
	fmt.Println("  lte=100   - 小于等于100（Less Than or Equal）")
	fmt.Println("  gt=0      - 大于0（Greater Than）")
	fmt.Println("  lt=100    - 小于100（Less Than）")
	fmt.Println()
	fmt.Println("测试示例:")
	fmt.Println("  curl -X POST http://localhost:8080/login \\")
	fmt.Println("    -H \"Content-Type: application/json\" \\")
	fmt.Println("    -d '{\"user\":\"admin\",\"password\":\"123456\"}'")
	fmt.Println()
	fmt.Println("  curl -X POST http://localhost:8080/register \\")
	fmt.Println("    -H \"Content-Type: application/json\" \\")
	fmt.Println("    -d '{\"username\":\"test\",\"email\":\"test@example.com\",\"age\":25}'")
}

// FormBindingDemo 演示表单参数绑定
// 表单绑定是将 HTTP 表单数据（application/x-www-form-urlencoded 或 multipart/form-data）
// 自动解析并填充到 Go 结构体中
//
// 表单数据格式:
//  1. application/x-www-form-urlencoded: 标准表单，键值对格式
//  2. multipart/form-data: 支持文件上传的表单
//
// 使用场景:
//   - HTML 表单提交
//   - 文件上传
//   - 简单的数据提交（不需要 JSON 的复杂场景）
func FormBindingDemo() {
	fmt.Println("=== Gin 表单参数绑定示例 ===")
	fmt.Println()

	router := gin.Default()

	// ========== 定义表单结构体 ==========
	// 结构体标签说明:
	//   form:"name" - 表单字段名映射到 Go 结构体字段
	//   binding:"required" - 验证规则：字段必填
	//   binding:"email" - 验证规则：必须是邮箱格式
	//   binding:"min=6" - 验证规则：最小长度为 6
	//
	// form 标签 vs json 标签:
	//   form 标签: 用于表单数据绑定（PostForm、ShouldBind）
	//   json 标签: 用于 JSON 数据绑定（ShouldBindJSON）
	type UserForm struct {
		Name     string `form:"name" binding:"required"`           // 姓名，必填
		Email    string `form:"email" binding:"required,email"`    // 邮箱，必填且必须是邮箱格式
		Password string `form:"password" binding:"required,min=6"` // 密码，必填且至少 6 位
	}

	// ========== 表单绑定示例 - 使用结构体 ==========
	// 请求格式:
	//   POST /register-form
	//   Content-Type: application/x-www-form-urlencoded
	//   Body: name=John&email=john@example.com&password=123456
	router.POST("/register-form", func(c *gin.Context) {
		var form UserForm

		// c.ShouldBind() - 根据 Content-Type 自动选择绑定方式
		// 对于表单数据:
		//   - application/x-www-form-urlencoded → 表单绑定
		//   - multipart/form-data → 表单绑定（支持文件上传）
		//
		// 工作原理:
		//   1. 读取表单数据
		//   2. 根据 form 标签映射字段
		//   3. 执行数据验证（binding 标签）
		//   4. 类型转换（字符串 → Go 类型）
		//
		// 优势:
		//   - 自动类型转换
		//   - 自动数据验证
		//   - 代码简洁
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		// 绑定成功，form 结构体已填充数据
		c.JSON(http.StatusOK, gin.H{
			"message": "表单提交成功",
			"form":    form,
		})
	})

	// ========== 直接获取表单字段 - 不使用结构体绑定 ==========
	// 适用于简单的场景，不需要复杂的数据验证
	// 请求格式:
	//   POST /submit
	//   Content-Type: application/x-www-form-urlencoded
	//   Body: name=John&email=john@example.com&age=25
	router.POST("/submit", func(c *gin.Context) {
		// c.PostForm() - 获取 POST 表单字段
		// 语法: c.PostForm("字段名")
		// 返回值: string（如果字段不存在，返回空字符串 ""）
		//
		// 适用于: application/x-www-form-urlencoded 和 multipart/form-data
		name := c.PostForm("name")
		email := c.PostForm("email")

		// c.DefaultPostForm() - 获取 POST 表单字段，如果不存在返回默认值
		// 语法: c.DefaultPostForm("字段名", "默认值")
		// 返回值: string
		//
		// 使用场景: 可选字段，但有合理的默认值
		age := c.DefaultPostForm("age", "0")

		// 注意: 返回的都是字符串类型，如果需要数字类型需要手动转换
		// 例如: ageInt, _ := strconv.Atoi(age)

		c.JSON(http.StatusOK, gin.H{
			"name":  name,
			"email": email,
			"age":   age,
		})
	})

	fmt.Println("表单绑定示例:")
	fmt.Println("  POST /register-form")
	fmt.Println("  Content-Type: application/x-www-form-urlencoded")
	fmt.Println("  Body: name=John&email=john@example.com&password=123456")
	fmt.Println()
	fmt.Println("方法说明:")
	fmt.Println("  c.PostForm(\"key\")           - 获取POST表单字段，不存在返回空字符串")
	fmt.Println("  c.DefaultPostForm(\"key\", \"default\") - 获取字段，不存在返回默认值")
	fmt.Println("  c.PostFormArray(\"key\")      - 获取表单数组（多个同名字段）")
	fmt.Println("  c.PostFormMap(\"key\")        - 获取表单Map（嵌套字段）")
	fmt.Println("  c.ShouldBind(&struct)        - 自动绑定表单到结构体（推荐）")
	fmt.Println()
	fmt.Println("表单绑定 vs 直接获取:")
	fmt.Println("  ShouldBind(): 自动类型转换、数据验证、代码简洁（推荐）")
	fmt.Println("  PostForm(): 手动处理、灵活但代码较多（简单场景）")
	fmt.Println()
	fmt.Println("测试示例:")
	fmt.Println("  curl -X POST http://localhost:8080/register-form \\")
	fmt.Println("    -d \"name=John&email=john@example.com&password=123456\"")
	fmt.Println()
	fmt.Println("  curl -X POST http://localhost:8080/submit \\")
	fmt.Println("    -d \"name=John&email=john@example.com&age=25\"")
	fmt.Println()
	fmt.Println("关键概念总结:")
	fmt.Println("  1. 路径参数: /users/:id → c.Param(\"id\") - 资源标识，必需")
	fmt.Println("  2. 查询参数: /users?page=1 → c.Query(\"page\") - 过滤条件，可选")
	fmt.Println("  3. JSON绑定: ShouldBindJSON() - 自动解析和验证JSON数据")
	fmt.Println("  4. 表单绑定: ShouldBind() - 自动解析和验证表单数据")
	fmt.Println("  5. 数据验证: 使用 binding 标签定义验证规则")
}

// RouteConflictDemo 演示路由冲突问题和解决方案
// 这个函数展示了 Gin 路由冲突的常见场景和解决方法
func RouteConflictDemo() {
	fmt.Println("=== Gin 路由冲突示例 ===")
	fmt.Println()
	fmt.Println("========== 问题说明 ==========")
	fmt.Println()
	fmt.Println("Gin 路由冲突规则:")
	fmt.Println("  在同一路径前缀下，不能在相同位置定义不同名称的路径参数")
	fmt.Println()
	fmt.Println("冲突示例:")
	fmt.Println("  ❌ 路由1: /users/:id")
	fmt.Println("  ❌ 路由2: /users/:userId/posts")
	fmt.Println("  问题: 两个路由在 /users/ 后的第一个位置都有路径参数")
	fmt.Println("  结果: Gin 会报错: 'conflicting route parameter'")
	fmt.Println()
	fmt.Println("========== 解决方案 ==========")
	fmt.Println()
	fmt.Println("方案1: 使用不同的路径前缀（推荐）")
	fmt.Println("  ✅ /users/:id")
	fmt.Println("  ✅ /user/:userId/posts  (使用单数 user，避免冲突)")
	fmt.Println()
	fmt.Println("方案2: 调整路由结构，使用固定路径")
	fmt.Println("  ✅ /users/:id")
	fmt.Println("  ✅ /users/posts/:postId  (posts 是固定路径，不是参数)")
	fmt.Println()
	fmt.Println("方案3: 使用不同的 HTTP 方法（如果业务逻辑允许）")
	fmt.Println("  ✅ GET  /users/:id")
	fmt.Println("  ✅ POST /users/:userId/posts  (不同方法不会冲突)")
	fmt.Println()
	fmt.Println("========== 通配符参数的前导斜杠 ==========")
	fmt.Println()
	fmt.Println("通配符参数 (*param) 的特点:")
	fmt.Println("  1. 可以匹配多级路径")
	fmt.Println("  2. 返回的值包含前导斜杠 /")
	fmt.Println()
	fmt.Println("示例:")
	fmt.Println("  路由: /files/*filepath")
	fmt.Println("  请求: /files/images/photo.jpg")
	fmt.Println("  结果: filepath = \"/images/photo.jpg\"")
	fmt.Println("  注意: 返回的值以斜杠开头，可以直接用于文件路径操作")
	fmt.Println()
	fmt.Println("为什么包含前导斜杠？")
	fmt.Println("  - 保持路径的完整性，方便直接拼接使用")
	fmt.Println("  - 例如: fullPath := \"/static\" + filepath")
	fmt.Println("  - 结果: \"/static/images/photo.jpg\" (无需手动添加斜杠)")
	fmt.Println()
	fmt.Println("如果需要去除前导斜杠:")
	fmt.Println("  filepath := c.Param(\"filepath\")")
	fmt.Println("  if strings.HasPrefix(filepath, \"/\") {")
	fmt.Println("    filepath = filepath[1:]  // 去除第一个字符")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("========== 实际代码示例 ==========")
	fmt.Println()
	fmt.Println("// 正确的路由定义（避免冲突）")
	fmt.Println("router := gin.Default()")
	fmt.Println()
	fmt.Println("// 方案1: 不同前缀")
	fmt.Println("router.GET(\"/users/:id\", handler1)           // /users/123")
	fmt.Println("router.GET(\"/user/:userId/posts\", handler2) // /user/123/posts")
	fmt.Println()
	fmt.Println("// 方案2: 固定路径")
	fmt.Println("router.GET(\"/users/:id\", handler1)          // /users/123")
	fmt.Println("router.GET(\"/users/posts/:postId\", handler2) // /users/posts/456")
	fmt.Println()
	fmt.Println("// 通配符示例")
	fmt.Println("router.GET(\"/files/*filepath\", func(c *gin.Context) {")
	fmt.Println("  filepath := c.Param(\"filepath\")")
	fmt.Println("  // filepath 包含前导斜杠，例如: \"/images/photo.jpg\"")
	fmt.Println("})")
	fmt.Println()
	fmt.Println("========== 测试建议 ==========")
	fmt.Println("  1. 注册路由时，Gin 会立即检查冲突")
	fmt.Println("  2. 如果发现冲突，程序启动时会报错")
	fmt.Println("  3. 建议在开发阶段就避免冲突，而不是运行时才发现")
	fmt.Println("  4. 使用不同的路径前缀是最简单可靠的解决方案")
}
