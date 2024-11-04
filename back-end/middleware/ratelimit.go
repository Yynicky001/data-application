package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/ratelimit"
	"net"
	"net/http"
	"sync"
)

// RateLimitByIP 是一个 Gin 中间件，用于限制每个 IP 的请求速率
func RateLimitByIP(limiter *IPRateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		ip, _, err := net.SplitHostPort(c.Request.RemoteAddr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Unable to parse IP address"})
			return
		}

		limiter := limiter.GetLimiter(ip)
		limiter.Take() // 等待直到允许通过一个请求

		c.Next()
	}
}

// IPRateLimiter 存储每个 IP 的速率限制器
type IPRateLimiter struct {
	visitors map[string]*ratelimit.Limiter
	mu       sync.Mutex
}

// GetLimiter 获取或创建给定 IP 的速率限制器
func (limiter *IPRateLimiter) GetLimiter(ip string) ratelimit.Limiter {
	limiter.mu.Lock()
	defer limiter.mu.Unlock()

	if l, exists := limiter.visitors[ip]; exists {
		return *l
	}

	// 创建一个新的速率限制器，每秒最多 10 个请求
	newLimiter := ratelimit.New(10) // 注意：ratelimit.New 参数是每秒允许的请求数
	limiter.visitors[ip] = &newLimiter
	return newLimiter
}

// NewIPRateLimiter 创建一个新的 IPRateLimiter
func NewIPRateLimiter() *IPRateLimiter {
	return &IPRateLimiter{
		visitors: make(map[string]*ratelimit.Limiter),
	}
}
