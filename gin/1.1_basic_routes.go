package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// BasicRoutesDemo 演示基础路由定义方式
func BasicRoutesDemo() {
	fmt.Println("=== Gin 基础路由示例 ===")
	fmt.Println()

	// 创建默认路由引擎
	router := gin.Default()

	// 1. 基础GET路由示例
	router.GET("/welcome", func(c *gin.Context) {
		// 获取查询参数，如果不存在则使用默认值
		firstName := c.DefaultQuery("firstname", "Guest")
		lastName := c.Query("lastname") // 如果不存在返回空字符串

		c.String(http.StatusOK, "Hello %s %s", firstName, lastName)
	})

	// 2. POST路由示例
	router.POST("/submit", func(c *gin.Context) {
		// 获取POST表单数据
		name := c.PostForm("name")
		email := c.PostForm("email")

		c.JSON(http.StatusOK, gin.H{
			"message": "提交成功",
			"name":    name,
			"email":   email,
		})
	})

	// 3. 多方法路由
	router.Any("/any", func(c *gin.Context) {
		c.String(http.StatusOK, "支持所有HTTP方法: %s", c.Request.Method)
	})

	fmt.Println("基础路由配置完成:")
	fmt.Println("  GET  /welcome?firstname=John&lastname=Doe")
	fmt.Println("  POST /submit (form-data: name, email)")
	fmt.Println("  ANY  /any (支持所有HTTP方法)")
	fmt.Println()
	fmt.Println("注意: 此示例仅展示路由配置，实际运行需要启动服务器")
	fmt.Println("      可以使用 router.Run(\":8080\") 启动服务器")
}

// RESTfulRoutesDemo 演示RESTful风格的路由定义
func RESTfulRoutesDemo() {
	fmt.Println("=== Gin RESTful 路由示例 ===")
	fmt.Println()

	router := gin.Default()

	// RESTful API 路由示例
	// GET    /users      - 获取用户列表
	// POST   /users      - 创建新用户
	// GET    /users/:id  - 获取指定用户
	// PUT    /users/:id  - 更新指定用户
	// DELETE /users/:id  - 删除指定用户

	// 获取用户列表
	router.GET("/users", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "获取用户列表",
			"method":  "GET",
		})
	})

	// 创建新用户
	router.POST("/users", func(c *gin.Context) {
		var user struct {
			Name  string `json:"name" binding:"required"`
			Email string `json:"email" binding:"required,email"`
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{
			"message": "用户创建成功",
			"user":    user,
		})
	})

	// 获取指定用户
	router.GET("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"message": "获取用户信息",
			"id":      id,
		})
	})

	// 更新指定用户
	router.PUT("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		var user struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "用户更新成功",
			"id":      id,
			"user":    user,
		})
	})

	// 删除指定用户
	router.DELETE("/users/:id", func(c *gin.Context) {
		id := c.Param("id")
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
}
