package server

import (
	"ecommerce/common/constant"
	"ecommerce/common/log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Server struct {
	Engine *gin.Engine
}

func NewServer() *Server {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.MaxMultipartMemory = 100 << 20
	engine.Use(CORSMiddleware())
	engine.Use(allowOptionsMethod())
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"service": constant.SERVICE_NAME,
			"version": constant.SERVICE_VERSION,
			"time":    time.Now().Unix(),
		})
	})
	server := &Server{Engine: engine}
	return server
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT, PATCH, DELETE, OPTIONS")
		c.Next()
	}
}

func allowOptionsMethod() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
		}
	}
}

func (server *Server) Start(port string) {
	v := make(chan struct{})
	go func() {
		if err := server.Engine.Run(":" + port); err != nil {
			log.Errorf("[API] service %v error: %v", constant.SERVICE_NAME, port)
			close(v)
		}
	}()
	log.Infof("[API] service %v listening on port: %v", constant.SERVICE_NAME, port)
	<-v
}
