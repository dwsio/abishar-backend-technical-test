package db

import (
	"context"
	"errors"

	pb "abishar-backend-technical-test/server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

func (p *GormProvider) UpdateOrCreateMapping(ctx context.Context, data *pb.Transactions) (*pb.Transactions, error) {
	if data.Id > 0 {
		model := &pb.Transactions{
			Id: data.Id,
		}
		if err := p.db_main.Model(&model).Updates(&data).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, status.Error(codes.NotFound, "ID Not Found")
			} else {
				return nil, status.Error(codes.Internal, "Internal Error : "+err.Error())
			}
		}

		return model, nil
	} else {
		if err := p.db_main.Create(&data).Error; err != nil {
			return nil, status.Error(codes.Internal, "Internal Error : "+err.Error())
		}

		return data, nil
	}
}
