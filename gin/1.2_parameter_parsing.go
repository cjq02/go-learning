package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PathParameterDemo 演示路径参数获取
func PathParameterDemo() {
	fmt.Println("=== Gin 路径参数解析示例 ===")
	fmt.Println()

	router := gin.Default()

	// 1. 单个路径参数
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "获取用户ID",
			"id":      id,
		})
	})

	// 2. 多个路径参数
	router.GET("/users/:userId/posts/:postId", func(c *gin.Context) {
		userId := c.Param("userId")
		postId := c.Param("postId")
		c.JSON(http.StatusOK, gin.H{
			"userId": userId,
			"postId": postId,
		})
	})

	// 3. 通配符路径参数
	router.GET("/files/*filepath", func(c *gin.Context) {
		filepath := c.Param("filepath")
		c.JSON(http.StatusOK, gin.H{
			"filepath": filepath,
		})
	})

	fmt.Println("路径参数示例:")
	fmt.Println("  GET /users/123              → id = \"123\"")
	fmt.Println("  GET /users/123/posts/456    → userId = \"123\", postId = \"456\"")
	fmt.Println("  GET /files/images/photo.jpg → filepath = \"/images/photo.jpg\"")
	fmt.Println()
	fmt.Println("注意: 通配符参数会包含前导斜杠")
}

// QueryParameterDemo 演示查询参数获取
func QueryParameterDemo() {
	fmt.Println("=== Gin 查询参数解析示例 ===")
	fmt.Println()

	router := gin.Default()

	// 1. 基础查询参数
	router.GET("/welcome", func(c *gin.Context) {
		// c.Query() - 如果不存在返回空字符串
		firstName := c.Query("firstname")

		// c.DefaultQuery() - 如果不存在返回默认值
		lastName := c.DefaultQuery("lastname", "Guest")

		// c.GetQuery() - 返回值和是否存在标志
		age, exists := c.GetQuery("age")
		if !exists {
			age = "未知"
		}

		c.String(http.StatusOK, "Hello %s %s, 年龄: %s", firstName, lastName, age)
	})

	// 2. 数组查询参数
	router.GET("/tags", func(c *gin.Context) {
		// 获取多个同名查询参数
		tags := c.QueryArray("tag")
		c.JSON(http.StatusOK, gin.H{
			"tags": tags,
		})
	})

	// 3. Map查询参数
	router.GET("/filters", func(c *gin.Context) {
		// 获取查询参数映射
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
}

// JSONBindingDemo 演示JSON参数绑定
func JSONBindingDemo() {
	fmt.Println("=== Gin JSON 参数绑定示例 ===")
	fmt.Println()

	router := gin.Default()

	// 定义登录结构体
	type Login struct {
		User     string `json:"user" binding:"required"`
		Password string `json:"password" binding:"required,min=6"`
	}

	// JSON绑定示例
	router.POST("/login", func(c *gin.Context) {
		var login Login

		// ShouldBindJSON 自动解析JSON并验证
		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error":   "参数校验失败",
				"message": err.Error(),
			})
			return
		}

		// 认证逻辑（示例）
		c.JSON(http.StatusOK, gin.H{
			"message": "登录成功",
			"user":    login.User,
		})
	})

	// 多种绑定方式示例
	router.POST("/register", func(c *gin.Context) {
		var user struct {
			Username string `json:"username" binding:"required"`
			Email    string `json:"email" binding:"required,email"`
			Age      int    `json:"age" binding:"gte=18,lte=100"`
		}

		// 根据Content-Type自动选择绑定方式
		if err := c.ShouldBind(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

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
	fmt.Println("  c.ShouldBindJSON(&struct) - 仅绑定JSON格式")
	fmt.Println("  c.ShouldBind(&struct)     - 根据Content-Type自动选择")
	fmt.Println("  c.MustBindJSON(&struct)   - 绑定失败会返回400错误")
	fmt.Println()
	fmt.Println("验证标签说明:")
	fmt.Println("  required  - 必填字段")
	fmt.Println("  min=6     - 最小长度6")
	fmt.Println("  email     - 邮箱格式")
	fmt.Println("  gte=18    - 大于等于18")
	fmt.Println("  lte=100   - 小于等于100")
}

// FormBindingDemo 演示表单参数绑定
func FormBindingDemo() {
	fmt.Println("=== Gin 表单参数绑定示例 ===")
	fmt.Println()

	router := gin.Default()

	// 表单绑定示例
	type UserForm struct {
		Name     string `form:"name" binding:"required"`
		Email    string `form:"email" binding:"required,email"`
		Password string `form:"password" binding:"required,min=6"`
	}

	router.POST("/register-form", func(c *gin.Context) {
		var form UserForm

		// ShouldBind 会根据Content-Type自动选择绑定方式
		// application/x-www-form-urlencoded 或 multipart/form-data
		if err := c.ShouldBind(&form); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "表单提交成功",
			"form":    form,
		})
	})

	// 直接获取表单字段
	router.POST("/submit", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")
		age := c.DefaultPostForm("age", "0")

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
	fmt.Println("  c.PostForm(\"key\")           - 获取POST表单字段")
	fmt.Println("  c.DefaultPostForm(\"key\", \"default\") - 获取字段，不存在返回默认值")
	fmt.Println("  c.PostFormArray(\"key\")      - 获取表单数组")
	fmt.Println("  c.PostFormMap(\"key\")        - 获取表单Map")
	fmt.Println("  c.ShouldBind(&struct)        - 自动绑定表单到结构体")
}
