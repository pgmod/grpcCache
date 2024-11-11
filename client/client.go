package client

import (
	"context"
	"fmt"
	"time"

	pb "github.com/pgmod/grpcCache/server/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// GRPCClient представляет собой структуру клиента gRPC.
type GRPCClient struct {
	conn   *grpc.ClientConn
	client pb.TokenServiceClient
}

// Connect устанавливает соединение с gRPC сервером.
func (g *GRPCClient) Connect(address string) error {
	// Создаем контекст с тайм-аутом для подключения
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Устанавливаем соединение с сервером
	conn, err := grpc.DialContext(ctx, address, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return fmt.Errorf("не удалось подключиться: %v", err)
	}

	// Сохраняем соединение и создаем клиента
	g.conn = conn
	g.client = pb.NewTokenServiceClient(conn)
	return nil
}

// Close закрывает соединение с сервером.
func (g *GRPCClient) Close() error {
	if g.conn != nil {
		return g.conn.Close()
	}
	return nil
}

// Save вызывает метод сохранения на сервере.
func (g *GRPCClient) Save(ctx context.Context, id string, clubId string) (string, error) {
	resp, err := g.client.Save(ctx, &pb.SaveRequest{Id: id, ClubId: clubId})
	if err != nil {
		return "", fmt.Errorf("ошибка сохранения: %v", err)
	}
	return resp.Message, nil
}

// CheckMultiple вызывает метод проверки нескольких ID на сервере.
func (g *GRPCClient) CheckMultiple(ctx context.Context, ids []string, myClubId string) (int32, error) {
	resp, err := g.client.CheckMultiple(ctx, &pb.CheckMultipleRequest{Ids: ids, ClubId: myClubId})
	if err != nil {
		return 0, fmt.Errorf("ошибка проверки нескольких ID: %v", err)
	}
	return resp.Count, nil
}

func CreateClient(address string) GRPCClient {
	grpcClient := GRPCClient{}
	grpcClient.Connect(address)
	return grpcClient
}
