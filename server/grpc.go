package server

import (
	"context"
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

// Save сохраняет токен в базу данных с идентификатором клуба
func (s *serverImpl) Save(ctx context.Context, req *pb.SaveRequest) (*pb.SaveResponse, error) {
	s.log.Info("save ", req.Id, req.ClubId, "with club ID", req.ClubId)
	_, err := db.DB.Exec("INSERT OR REPLACE INTO tokens (id, club_id) VALUES (?, ?)", req.Id, req.ClubId)
	if err != nil {
		return nil, err
	}

	return &pb.SaveResponse{Message: "Saved successfully"}, nil
}
func (s *serverImpl) Clear(ctx context.Context, req *pb.ClearRequest) (*pb.ClearResponse, error) {
	s.log.Info("clear ", req.ClubId, "with club ID", req.ClubId)
	_, err := db.DB.Exec("DELETE FROM tokens WHERE club_id = ?", req.ClubId)
	if err != nil {
		return nil, err
	}

	return &pb.ClearResponse{Message: "Cleared successfully"}, nil
}

// CheckMultiple проверяет наличие нескольких id в базе данных, включая проверку по идентификатору клуба
func (s *serverImpl) CheckMultiple(ctx context.Context, req *pb.CheckMultipleRequest) (*pb.CheckMultipleResponse, error) {
	s.log.Info("check multiple ", req.Ids, "with club ID", req.ClubId)
	if len(req.Ids) == 0 {
		return &pb.CheckMultipleResponse{Count: 0}, nil
	}
	// Создаем плейсхолдеры для SQL запроса
	placeholders := strings.Repeat("?,", len(req.Ids))
	placeholders = placeholders[:len(placeholders)-1] // Убираем последнюю запятую

	// Создаем запрос с использованием плейсхолдеров и условия для clubId
	query := fmt.Sprintf("SELECT COUNT(*) FROM tokens WHERE id IN (%s) AND club_id = ?", placeholders)

	// Преобразуем список ID в интерфейсный срез и добавляем clubId
	args := make([]interface{}, len(req.Ids)+1)
	for i, id := range req.Ids {
		args[i] = id
	}
	args[len(req.Ids)] = req.ClubId // Добавляем идентификатор клуба в параметры запроса

	// Выполняем запрос и получаем количество найденных записей с соответствующим club_id
	var count int32
	err := db.DB.QueryRow(query, args...).Scan(&count)
	if err != nil {
		return nil, err
	}

	return &pb.CheckMultipleResponse{Count: count}, nil
}
