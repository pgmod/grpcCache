package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/pgmod/grpcCache/db"
	"github.com/pgmod/logger"
	"google.golang.org/grpc"
)

func StartServer(logLevel int, port string) {
	log := logger.NewLogger(3, "grpcServer.log", false, "")
	defer func() {
		db.CloseDB()
		log.Info("База данных закрыта")
	}()

	// Инициализация базы данных
	if err := db.InitDB(); err != nil {
		log.Error("Не удалось инициализировать базу данных:", err)
		os.Exit(1)
	}
	defer db.CloseDB()

	grpcListener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Error("Не удалось запустить gRPC сервер:", err)
		os.Exit(1)
	}
	grpcServer := grpc.NewServer()
	RegisterGRPCServer(grpcServer, log)

	go func() {
		log.Info("Запуск gRPC сервера на порту", port)
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Error("Ошибка gRPC сервера:", err)
			os.Exit(1)
		}
	}()

	// Обработка завершения программы
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)
	<-sigChan

	log.Info("Выключение сервера...")
	grpcServer.GracefulStop()
}
