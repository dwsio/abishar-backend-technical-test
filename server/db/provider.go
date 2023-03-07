package db

import (
	"context"

	pb "abishar-backend-technical-test/server/pb"
	"gorm.io/gorm"
)

type Provider interface {
	CreateTransaction(ctx context.Context, data *pb.Transactions) (*pb.Transactions, error)
}

type GormProvider struct {
	db_main *gorm.DB
}

func NewProvider(db *gorm.DB) *GormProvider {
	return &GormProvider{db_main: db}
}
