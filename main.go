package main

import (
	"errors"
	"gitlab-notificatons/helper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	// Настройка роутера и маршрутов
	http.HandleFunc("/webhook", helper.HandleWebhook)

	port := ":8080"
	log.Printf("🚀 Сервер запущен на http://localhost%s/webhook", port)

	// Запуск сервера в отдельной горутине
	server := &http.Server{Addr: port}
	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Ошибка запуска сервера: %v", err)
		}
	}()

	// Ожидание сигнала завершения (Ctrl+C, SIGTERM и т.д.)
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop
	log.Println("🛑 Завершение работы сервера...")
}
