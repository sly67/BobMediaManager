package main

import (
	"context"
	"embed"
	"io/fs"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"bobmediamanager/api"
	"bobmediamanager/config"
)

//go:embed static/*
var staticFiles embed.FS

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	staticFS, err := fs.Sub(staticFiles, "static")
	if err != nil {
		log.Fatalf("failed to sub static FS: %v", err)
	}

	router := api.NewRouter(cfg, staticFS)

	addr := ":48040"
	if port := os.Getenv("PORT"); port != "" {
		addr = ":" + port
	}

	srv := &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 60 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	go func() {
		log.Printf("BobMediaManager listening on %s", addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("server error: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown error: %v", err)
	}
	log.Println("server stopped")
}
