package app

import (
	"context"
	"crypto/tls"
	"fmt"
	"github.com/nekohor/mangosteen/internal/app/config"
	"github.com/nekohor/mangosteen/pkg/logger/zap/logger"
	"net/http"
	"time"
)

// InitHTTPServer 初始化http服务
func InitHTTPServer(ctx context.Context, handler http.Handler) func() {
	cfg := config.C.HTTP
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	srv := &http.Server{
		Addr:         addr,
		Handler:      handler,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	go func() {
		logger.Infof("HTTP server is running at %s.", addr)

		var err error
		if cfg.CertFile != "" && cfg.KeyFile != "" {
			srv.TLSConfig = &tls.Config{MinVersion: tls.VersionTLS12}
			err = srv.ListenAndServeTLS(cfg.CertFile, cfg.KeyFile)
		} else {
			err = srv.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			panic(err)
		}

	}()

	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(cfg.ShutdownTimeout))
		defer cancel()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctx); err != nil {
			logger.Errorf(err.Error())
		}
	}
}