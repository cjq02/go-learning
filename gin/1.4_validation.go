package gin

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// CustomValidationDemo 演示自定义验证规则
func CustomValidationDemo() {
	fmt.Println("=== Gin 自定义验证规则示例 ===")
	fmt.Println()

	router := gin.Default()

	// 注册自定义验证器
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册手机号验证
		v.RegisterValidation("phone", func(fl validator.FieldLevel) bool {
			phone := fl.Field().String()
			matched, _ := regexp.MatchString(`^1[3-9]\d{9}$`, phone)
			return matched
		})

		// 注册密码强度验证（至少包含字母和数字）
		v.RegisterValidation("strong_password", func(fl validator.FieldLevel) bool {
			password := fl.Field().String()
			hasLetter, _ := regexp.MatchString(`[a-zA-Z]`, password)
			hasDigit, _ := regexp.MatchString(`[0-9]`, password)
			return hasLetter && hasDigit && len(password) >= 8
		})
	}

	// 使用自定义验证
	type RegisterRequest struct {
		Username string `json:"username" binding:"required,min=3,max=20"`
		Phone    string `json:"phone" binding:"required,phone"`
		Email    string `json:"email" binding:"required,email"`
		Password string `json:"password" binding:"required,strong_password"`
	}

	router.POST("/register", func(c *gin.Context) {
		var req RegisterRequest

		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "参数校验失败",
				"details": formatValidationErrors(err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "注册成功",
			"user":    req,
		})
	})

	fmt.Println("自定义验证规则配置完成:")
	fmt.Println("  手机号验证: phone - 匹配中国大陆手机号格式")
	fmt.Println("  密码强度验证: strong_password - 至少8位，包含字母和数字")
	fmt.Println()
	fmt.Println("使用示例:")
	fmt.Println("  POST /register")
	fmt.Println("  {")
	fmt.Println("    \"username\": \"john\",")
	fmt.Println("    \"phone\": \"13800138000\",")
	fmt.Println("    \"email\": \"john@example.com\",")
	fmt.Println("    \"password\": \"Password123\"")
	fmt.Println("  }")
	fmt.Println()
	fmt.Println("注册自定义验证器:")
	fmt.Println("  v.RegisterValidation(\"tag\", func(fl validator.FieldLevel) bool {")
	fmt.Println("    // 验证逻辑")
	fmt.Println("    return true/false")
	fmt.Println("  })")
}

// ValidationErrorHandlingDemo 演示参数验证错误处理标准流程
func ValidationErrorHandlingDemo() {
	fmt.Println("=== Gin 参数验证错误处理示例 ===")
	fmt.Println()

	router := gin.Default()

	type UserRequest struct {
		Username string `json:"username" binding:"required,min=3,max=20"`
		Email    string `json:"email" binding:"required,email"`
		Age      int    `json:"age" binding:"required,gte=18,lte=100"`
		Phone    string `json:"phone" binding:"required,len=11"`
	}

	router.POST("/users", func(c *gin.Context) {
		var req UserRequest

		// 绑定并验证参数
		if err := c.ShouldBindJSON(&req); err != nil {
			// 标准错误处理流程
			errors := err.(validator.ValidationErrors)
			errorMessages := make([]string, 0, len(errors))

			for _, e := range errors {
				errorMessage := formatFieldError(e)
				errorMessages = append(errorMessages, errorMessage)
			}

			// 统一响应格式
			c.JSON(http.StatusBadRequest, gin.H{
				"code":   1001,
				"msg":    "参数校验失败",
				"errors": errorMessages,
			})
			return
		}

		// 处理成功
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"msg":     "操作成功",
			"data":    req,
		})
	})

	fmt.Println("错误处理标准流程:")
	fmt.Println("  1. 使用 ShouldBind 或 ShouldBindJSON 绑定参数")
	fmt.Println("  2. 检查错误类型是否为 ValidationErrors")
	fmt.Println("  3. 遍历所有验证错误，格式化错误信息")
	fmt.Println("  4. 返回统一的错误响应格式")
	fmt.Println()
	fmt.Println("统一响应格式:")
	fmt.Println("  成功: {\"code\": 200, \"msg\": \"操作成功\", \"data\": {...}}")
	fmt.Println("  失败: {\"code\": 1001, \"msg\": \"参数校验失败\", \"errors\": [...]}")
}

// formatValidationErrors 格式化验证错误
func formatValidationErrors(err error) []string {
	errors := err.(validator.ValidationErrors)
	errorMessages := make([]string, 0, len(errors))

	for _, e := range errors {
		errorMessages = append(errorMessages, formatFieldError(e))
	}

	return errorMessages
}

// formatFieldError 格式化单个字段错误
func formatFieldError(e validator.FieldError) string {
	field := e.Field()
	tag := e.Tag()

	// 根据验证标签返回友好的错误信息
	switch tag {
	case "required":
		return fmt.Sprintf("参数 %s 是必填项", field)
	case "min":
		return fmt.Sprintf("参数 %s 长度不能小于 %s", field, e.Param())
	case "max":
		return fmt.Sprintf("参数 %s 长度不能大于 %s", field, e.Param())
	case "email":
		return fmt.Sprintf("参数 %s 必须是有效的邮箱地址", field)
	case "gte":
		return fmt.Sprintf("参数 %s 必须大于等于 %s", field, e.Param())
	case "lte":
		return fmt.Sprintf("参数 %s 必须小于等于 %s", field, e.Param())
	case "len":
		return fmt.Sprintf("参数 %s 长度必须为 %s", field, e.Param())
	case "phone":
		return fmt.Sprintf("参数 %s 必须是有效的手机号", field)
	default:
		return fmt.Sprintf("参数 %s 校验失败：%s", field, tag)
	}
}

// BuiltinValidationTagsDemo 演示内置验证标签
func BuiltinValidationTagsDemo() {
	fmt.Println("=== Gin 内置验证标签示例 ===")
	fmt.Println()

	type ValidationExample struct {
		// 字符串验证
		RequiredString string `json:"required_string" binding:"required"`
		MinString      string `json:"min_string" binding:"min=3"`
		MaxString      string `json:"max_string" binding:"max=10"`
		LenString      string `json:"len_string" binding:"len=5"`
		Email          string `json:"email" binding:"email"`
		URL            string `json:"url" binding:"url"`
		Alpha          string `json:"alpha" binding:"alpha"`
		Alphanum       string `json:"alphanum" binding:"alphanum"`

		// 数字验证
		MinInt    int `json:"min_int" binding:"min=1"`
		MaxInt    int `json:"max_int" binding:"max=100"`
		RangeInt  int `json:"range_int" binding:"gte=18,lte=65"`
		OneOfInt  int `json:"one_of_int" binding:"oneof=1 2 3"`

		// 数组/切片验证
		MinSlice []string `json:"min_slice" binding:"min=1"`
		MaxSlice []string `json:"max_slice" binding:"max=10"`

		// 其他验证
		UUID      string `json:"uuid" binding:"uuid"`
		IP        string `json:"ip" binding:"ip"`
		IPv4      string `json:"ipv4" binding:"ipv4"`
		IPv6      string `json:"ipv6" binding:"ipv6"`
		Datetime  string `json:"datetime" binding:"datetime=2006-01-02 15:04:05"`
		Date      string `json:"date" binding:"date"`
		Time      string `json:"time" binding:"time"`
		Boolean   bool   `json:"boolean" binding:"-"`
	}

	router := gin.Default()

	router.POST("/validate", func(c *gin.Context) {
		var example ValidationExample

		if err := c.ShouldBindJSON(&example); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": formatValidationErrors(err),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "验证通过",
			"data":    example,
		})
	})

	fmt.Println("常用内置验证标签:")
	fmt.Println()
	fmt.Println("字符串验证:")
	fmt.Println("  required  - 必填")
	fmt.Println("  min=3     - 最小长度")
	fmt.Println("  max=10    - 最大长度")
	fmt.Println("  len=5     - 固定长度")
	fmt.Println("  email     - 邮箱格式")
	fmt.Println("  url       - URL格式")
	fmt.Println("  alpha     - 仅字母")
	fmt.Println("  alphanum  - 字母和数字")
	fmt.Println()
	fmt.Println("数字验证:")
	fmt.Println("  min=1     - 最小值")
	fmt.Println("  max=100   - 最大值")
	fmt.Println("  gte=18    - 大于等于")
	fmt.Println("  lte=65    - 小于等于")
	fmt.Println("  oneof=1 2 3 - 只能是其中之一")
	fmt.Println()
	fmt.Println("其他验证:")
	fmt.Println("  uuid      - UUID格式")
	fmt.Println("  ip        - IP地址")
	fmt.Println("  ipv4      - IPv4地址")
	fmt.Println("  ipv6      - IPv6地址")
	fmt.Println("  datetime=2006-01-02 15:04:05 - 日期时间格式")
	fmt.Println("  date      - 日期格式")
	fmt.Println("  time      - 时间格式")
}

