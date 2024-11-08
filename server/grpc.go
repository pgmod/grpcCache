package server

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/pgmod/grpcCache/db"
	pb "github.com/pgmod/grpcCache/server/proto"
	"github.com/pgmod/logger"
	"google.golang.org/grpc"
)

// serverImpl представляет реализацию сервера gRPC
type serverImpl struct {
	pb.UnimplementedTokenServiceServer
	log *logger.Logger
}

// RegisterGRPCServer регистрирует gRPC сервер и связывает его с реализацией
func RegisterGRPCServer(grpcServer *grpc.Server, log *logger.Logger) {
	pb.RegisterTokenServiceServer(grpcServer, &serverImpl{log: log})
}

// Save сохраняет токен в базу данных
func (s *serverImpl) Save(ctx context.Context, req *pb.SaveRequest) (*pb.SaveResponse, error) {
	fmt.Println("save ", req.Id, req.AuthToken)
	_, err := db.DB.Exec("INSERT OR REPLACE INTO tokens (id, authToken) VALUES (?, ?)", req.Id, req.AuthToken)
	if err != nil {
		return nil, err
	}

	return &pb.SaveResponse{Message: "Saved successfully"}, nil
}

// Clear очищает базу данных
func (s *serverImpl) Clear(ctx context.Context, req *pb.ClearRequest) (*pb.ClearResponse, error) {
	_, err := db.DB.Exec("DELETE FROM tokens")
	if err != nil {
		return nil, err
	}

	return &pb.ClearResponse{Message: "Cleared successfully"}, nil
}

// Check проверяет наличие id в базе данных
func (s *serverImpl) Check(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	fmt.Println("check ", req.Id)
	var exists bool
	err := db.DB.QueryRow("SELECT EXISTS(SELECT 1 FROM tokens WHERE id = ?)", req.Id).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return nil, err
	}

	return &pb.CheckResponse{Exists: exists}, nil
}

// CheckMultiple проверяет наличие нескольких id в базе данных
func (s *serverImpl) CheckMultiple(ctx context.Context, req *pb.CheckMultipleRequest) (*pb.CheckMultipleResponse, error) {
	fmt.Println("check multiple ", req.Ids)
	if len(req.Ids) == 0 {
		return &pb.CheckMultipleResponse{Count: 0}, nil
	}
	// Создаем плейсхолдеры для SQL запроса
	placeholders := strings.Repeat("?,", len(req.Ids))
	placeholders = placeholders[:len(placeholders)-1] // Убираем последнюю запятую

	// Создаем запрос с использованием плейсхолдеров
	query := fmt.Sprintf("SELECT COUNT(*) FROM tokens WHERE id IN (%s)", placeholders)

	// Преобразуем список ID в интерфейсный срез
	args := make([]interface{}, len(req.Ids))
	for i, id := range req.Ids {
		args[i] = id
	}

	// Выполняем запрос и получаем количество найденных записей
	var count int32
	err := db.DB.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &pb.CheckMultipleResponse{Count: count}, nil
}
