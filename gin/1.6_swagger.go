package gin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SwaggerIntegrationDemo 演示Swagger集成规范
func SwaggerIntegrationDemo() {
	fmt.Println("=== Gin Swagger 集成规范示例 ===")
	fmt.Println()

	router := gin.Default()

	// Swagger注释示例（需要在main.go中配置）
	// 这些注释会被swag工具解析生成swagger.json

	// @Summary 用户登录
	// @Description 用户登录接口，需要提供用户名和密码
	// @Tags auth
	// @Accept json
	// @Produce json
	// @Param login body LoginRequest true "登录凭证"
	// @Success 200 {object} LoginResponse "登录成功"
	// @Failure 400 {object} ErrorResponse "参数错误"
	// @Failure 401 {object} ErrorResponse "认证失败"
	// @Router /api/v1/login [post]
	router.POST("/api/v1/login", func(c *gin.Context) {
		var login struct {
			Username string `json:"username" binding:"required"`
			Password string `json:"password" binding:"required"`
		}

		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "参数校验失败",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token": "jwt_token_here",
			"user":  login.Username,
		})
	})

	// @Summary 获取用户信息
	// @Description 根据用户ID获取用户详细信息
	// @Tags users
	// @Accept json
	// @Produce json
	// @Param id path int true "用户ID"
	// @Param Authorization header string true "Bearer Token"
	// @Success 200 {object} UserResponse "获取成功"
	// @Failure 404 {object} ErrorResponse "用户不存在"
	// @Router /api/v1/users/{id} [get]
	router.GET("/api/v1/users/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.JSON(http.StatusOK, gin.H{
			"id":   id,
			"name": "John Doe",
		})
	})

	fmt.Println("Swagger集成步骤:")
	fmt.Println()
	fmt.Println("1. 安装swag工具:")
	fmt.Println("   go install github.com/swaggo/swag/cmd/swag@latest")
	fmt.Println()
	fmt.Println("2. 在main.go顶部添加Swagger配置注释:")
	fmt.Println("   // @title Gin Web API")
	fmt.Println("   // @version 1.0")
	fmt.Println("   // @description RESTful API 文档")
	fmt.Println("   // @host localhost:8080")
	fmt.Println("   // @BasePath /api/v1")
	fmt.Println()
	fmt.Println("3. 在路由处理函数上方添加API注释:")
	fmt.Println("   // @Summary 接口摘要")
	fmt.Println("   // @Tags 标签名")
	fmt.Println("   // @Accept json")
	fmt.Println("   // @Produce json")
	fmt.Println("   // @Param 参数定义")
	fmt.Println("   // @Success 200 {object} Response")
	fmt.Println("   // @Router /path [method]")
	fmt.Println()
	fmt.Println("4. 在main函数中添加Swagger路由:")
	fmt.Println("   import \"github.com/swaggo/gin-swagger\"")
	fmt.Println("   import \"github.com/swaggo/files\"")
	fmt.Println("   router.GET(\"/docs/*any\", ginSwagger.WrapHandler(swaggerFiles.Handler))")
	fmt.Println()
	fmt.Println("5. 生成Swagger文档:")
	fmt.Println("   swag init -g main.go -o ./docs --parseDependency")
	fmt.Println()
	fmt.Println("6. 访问Swagger UI:")
	fmt.Println("   http://localhost:8080/docs/index.html")
}

// SwaggerAnnotationsDemo 演示Swagger注释规范
func SwaggerAnnotationsDemo() {
	fmt.Println("=== Swagger 注释规范详解 ===")
	fmt.Println()

	fmt.Println("1. 主配置注释（main.go顶部）:")
	fmt.Println("   // @title           API标题")
	fmt.Println("   // @version         版本号")
	fmt.Println("   // @description     API描述")
	fmt.Println("   // @termsOfService  服务条款URL")
	fmt.Println("   // @contact.name    联系人")
	fmt.Println("   // @contact.email   联系邮箱")
	fmt.Println("   // @license.name    许可证")
	fmt.Println("   // @license.url     许可证URL")
	fmt.Println("   // @host            主机地址")
	fmt.Println("   // @BasePath        基础路径")
	fmt.Println("   // @schemes         http https")
	fmt.Println()
	fmt.Println("2. API接口注释:")
	fmt.Println("   // @Summary         接口摘要（必填）")
	fmt.Println("   // @Description     详细描述")
	fmt.Println("   // @Tags            标签（用于分组）")
	fmt.Println("   // @Accept          接受的Content-Type")
	fmt.Println("   // @Produce         返回的Content-Type")
	fmt.Println("   // @Param           参数定义")
	fmt.Println("   // @Success         成功响应")
	fmt.Println("   // @Failure         失败响应")
	fmt.Println("   // @Router          路由定义")
	fmt.Println()
	fmt.Println("3. 参数定义格式:")
	fmt.Println("   // @Param name type location required \"description\"")
	fmt.Println("   示例:")
	fmt.Println("   // @Param id path int true \"用户ID\"")
	fmt.Println("   // @Param name query string false \"用户名\"")
	fmt.Println("   // @Param user body UserRequest true \"用户信息\"")
	fmt.Println("   // @Param Authorization header string true \"Bearer Token\"")
	fmt.Println()
	fmt.Println("4. 响应定义格式:")
	fmt.Println("   // @Success 200 {object} Response \"描述\"")
	fmt.Println("   // @Success 200 {array} User \"用户列表\"")
	fmt.Println("   // @Success 200 {string} string \"纯文本\"")
	fmt.Println("   // @Failure 400 {object} ErrorResponse \"错误信息\"")
	fmt.Println()
	fmt.Println("5. 路由定义格式:")
	fmt.Println("   // @Router /path [method]")
	fmt.Println("   示例:")
	fmt.Println("   // @Router /users [get]")
	fmt.Println("   // @Router /users/{id} [put]")
	fmt.Println("   // @Router /login [post]")
}

// SwaggerSecurityDemo 演示Swagger安全配置
func SwaggerSecurityDemo() {
	fmt.Println("=== Swagger 安全配置示例 ===")
	fmt.Println()

	router := gin.Default()

	// 方案1: 基础认证保护Swagger文档
	authMiddleware := gin.BasicAuth(gin.Accounts{
		"admin": "swagger123",
	})

	// 保护Swagger文档访问
	router.GET("/docs", authMiddleware, func(c *gin.Context) {
		c.String(http.StatusOK, "Swagger文档（需要认证）")
	})

	// 方案2: JWT Token认证（在Swagger中配置）
	// 在main.go的Swagger配置中添加:
	// // @securityDefinitions.apikey ApiKeyAuth
	// // @in header
	// // @name Authorization
	// // @description Bearer {token}

	router.GET("/api/protected", func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "需要认证",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "受保护的数据",
		})
	})

	fmt.Println("Swagger安全配置方案:")
	fmt.Println()
	fmt.Println("方案1: 基础认证保护文档访问")
	fmt.Println("  使用gin.BasicAuth中间件")
	fmt.Println("  配置用户名密码")
	fmt.Println("  访问 /docs 时需要输入认证信息")
	fmt.Println()
	fmt.Println("方案2: JWT Token认证（API接口）")
	fmt.Println("  在Swagger配置中添加securityDefinitions")
	fmt.Println("  客户端在Swagger UI中输入Token")
	fmt.Println("  自动在请求头中添加Authorization")
	fmt.Println()
	fmt.Println("配置示例:")
	fmt.Println("  // @securityDefinitions.apikey ApiKeyAuth")
	fmt.Println("  // @in header")
	fmt.Println("  // @name Authorization")
	fmt.Println("  // @description Bearer {token}")
	fmt.Println()
	fmt.Println("在接口中使用:")
	fmt.Println("  // @Security ApiKeyAuth")
	fmt.Println("  // @Router /api/protected [get]")
}

