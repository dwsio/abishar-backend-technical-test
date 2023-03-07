package api

import (
	"context"
	"os"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"

	"abishar-backend-technical-test/server/db"
	svc "abishar-backend-technical-test/server/lib/stubs"
	pb "abishar-backend-technical-test/server/pb"
)

const apiServicePath string = "/transaction.service.v1.ApiService/"

type Server struct {
	provider *db.GormProvider
	svcConn  *svc.ServiceConnection

	pb.ApiServiceServer
}

type DataPublish struct {
	DataType string
	Data     string
}

func New(
	db01 *gorm.DB,
	svcConn *svc.ServiceConnection,
) *Server {

	return &Server{
		provider:         db.NewProvider(db01),
		svcConn:          svcConn,
		ApiServiceServer: nil,
	}
}

func (s *Server) NotImplementedError() error {
	st := status.New(codes.Unimplemented, "Not implemented yet")
	return st.Err()
}

func (s *Server) UnauthorizedError() error {
	st := status.New(codes.Unauthenticated, "Unauthorized")
	return st.Err()
}

func (s *Server) HealthCheck(ctx context.Context, _ *pb.HealthCheckRequest) (*pb.HealthCheckResponse, error) {
	return &pb.HealthCheckResponse{Message: "API Running !"}, nil
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
