package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gofiber/fiber/v3"

	"github.com/HemanthGowda/library-management-go/internal/app"
	"github.com/HemanthGowda/library-management-go/internal/config"
)

func main() {
	if err := run(); err != nil {
		slog.Error("application stopped", "error", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	application, err := app.New(ctx, cfg)
	if err != nil {
		return err
	}
	defer application.Close()

	if err := application.Seed(ctx); err != nil {
		return err
	}

	addr := ":" + cfg.Port
	slog.Info("listening", "addr", addr)
	return application.App.Listen(addr, fiber.ListenConfig{
		GracefulContext:       ctx,
		DisableStartupMessage: true,
		ListenerNetwork:       fiber.NetworkTCP,
		ShutdownTimeout:       10 * time.Second,
	})
}
