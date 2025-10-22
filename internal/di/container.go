package di

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"github.com/RodrigoGuerraCortes/ai-backend/internal/ai"
	"github.com/RodrigoGuerraCortes/ai-backend/internal/config"
	appRouter "github.com/RodrigoGuerraCortes/ai-backend/internal/http/router"
)

// Container holds application singletons and lifecycle.
type Container struct {
	Cfg        *config.Config
	Logger     *zap.Logger
	AI         ai.AIClient
	Router     *gin.Engine
	HTTPServer *http.Server
}

// New builds a production container from environment/config.
func New(logger *zap.Logger) *Container {
	// 1) Config
	cfg := config.LoadConfig()

	// 2) AI client (implements ai.AIClient)
	aiClient := ai.NewGeminiClient(cfg.GeminiAPIKey)

	// 3) HTTP router (inject AI interface)
	r := appRouter.NewRouter(aiClient)

	// 4) HTTP server (separate from gin.Run to support graceful shutdown)
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: r,
	}

	return &Container{
		Cfg:        cfg,
		Logger:     logger,
		AI:         aiClient,
		Router:     r,
		HTTPServer: srv,
	}
}

// Start runs the HTTP server in a goroutine.
func (c *Container) Start() error {
	c.Logger.Info("✅ Server starting", zap.String("port", c.HTTPServer.Addr))
	go func() {
		// http.ErrServerClosed is expected on Shutdown
		if err := c.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			c.Logger.Fatal("HTTP server ListenAndServe failed", zap.Error(err))
		}
	}()
	return nil
}

// Stop performs graceful shutdown with a timeout.
func (c *Container) Stop(timeout time.Duration) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	if err := c.HTTPServer.Shutdown(ctx); err != nil {
		c.Logger.Error("HTTP server graceful shutdown failed", zap.Error(err))
	} else {
		c.Logger.Info("🛑 HTTP server stopped gracefully")
	}
}
