package middleware

import (
	"context"
	"fmt"
	"net/http"
	"voucher_system/database"
	"voucher_system/helper"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Cacher database.Cacher
}

func NewMiddleware(cacher database.Cacher) Middleware {
	_, err := cacher.GetClient().Ping(context.Background()).Result()
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	return Middleware{
		Cacher: cacher,
	}
}

func (m *Middleware) Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("token")
		id := c.GetHeader("id_key")
		val, err := m.Cacher.Get(id)
		if err != nil {
			helper.ResponseError(c, "server error", err.Error(), http.StatusInternalServerError)
			c.Abort()
			return
		}

		if val == "" || val != token {
			helper.ResponseError(c, "Unauthorized", "Unauthorized", http.StatusUnauthorized)
			c.Abort()
			return
		}

		c.Next()

	}
}
