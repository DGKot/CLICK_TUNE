package app

import (
	"click_tune/internal/http/rest"
	"click_tune/internal/selector"
	"click_tune/internal/service"
	"click_tune/internal/storage"
	"click_tune/internal/storage/inmemory"
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type App struct{}

func NewApp() App {
	return App{}
}

func (a *App) Start() {
	fmt.Println("Project started")
	selector := selector.NewSelector()
	storage := buildStorage()

	serviceDeps := service.Deps{
		Storage:  storage,
		Selector: selector,
	}
	service := service.NewService(serviceDeps)
	// TODO Сделать подтяжку параметров сервера через конфиг
	server := rest.NewServer(rest.ServerDeps{
		Service:      service,
		Host:         "localhost",
		Port:         "8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  30 * time.Second,
	})

	go func() {
		err := server.Start()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("server open error", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	_ = server.Stop(ctx)
}

// TODO Сделать создание разных реализаций в зависимости от конфигов.
func buildStorage() storage.Storage {
	return inmemory.NewStorage()
}
