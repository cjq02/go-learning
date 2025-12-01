package middleware

import (
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// TestJWTMiddleware 演示中间件单元测试
// 注意: 这是一个示例文件，实际测试需要安装测试依赖
func TestJWTMiddleware(t *testing.T) {
	// 设置测试环境变量
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	// 创建测试路由
	router := gin.New()
	router.Use(JWTAuth())
	router.GET("/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// 有效令牌测试
	t.Run("valid token", func(t *testing.T) {
		// 生成有效 Token
		token, err := GenerateToken("user123", []string{"admin"})
		assert.NoError(t, err)

		// 创建测试请求
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		// 执行请求
		router.ServeHTTP(w, req)

		// 验证响应
		assert.Equal(t, http.StatusOK, w.Code)
	})

	// 无效令牌测试
	t.Run("invalid token", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)
		req.Header.Set("Authorization", "Bearer invalid_token")

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})

	// 缺少令牌测试
	t.Run("missing token", func(t *testing.T) {
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/test", nil)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusUnauthorized, w.Code)
	})
}

// TestRequireRole 测试 RBAC 权限中间件
// RBAC: Role-Based Access Control (基于角色的访问控制)
func TestRequireRole(t *testing.T) {
	os.Setenv("JWT_SECRET", "test-secret-key")
	defer os.Unsetenv("JWT_SECRET")

	router := gin.New()
	router.Use(JWTAuth())
	router.Use(RequireRole("admin"))
	router.GET("/admin/test", func(c *gin.Context) {
		c.Status(http.StatusOK)
	})

	// 有 admin 角色的测试
	t.Run("has admin role", func(t *testing.T) {
		token, _ := GenerateToken("user123", []string{"admin", "user"})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/admin/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
	})

	// 没有 admin 角色的测试
	t.Run("no admin role", func(t *testing.T) {
		token, _ := GenerateToken("user123", []string{"user"})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/admin/test", nil)
		req.Header.Set("Authorization", "Bearer "+token)

		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusForbidden, w.Code)
	})
}


