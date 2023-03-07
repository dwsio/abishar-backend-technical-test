package api

import (
	"context"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"

	
	"abishar-backend-technical-test/server/pb"
)

func (s *Server) CreateTransactionMultiple(ctx context.Context, req *pb.CreateTransactionMultipleRequest) (*pb.CreateTransactionMultipleResponse, error) {

	result := &pb.CreateTransactionMultipleResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	if len(req.GetData()) == 0 {
		logrus.Println("[api][CreateTransactionMultiple] No Data Inserted")
		return nil, status.Errorf(codes.InvalidArgument, "No Data Inserted")
	}

	for _, data := range req.GetData() {
		_, err := s.provider.UpdateOrCreateTransaction(ctx, &pb.Transactions{
			RequestId: req.GetRequestId(),
			ProductId: data.GetProductId(),
			CustomerName: data.GetCustomerName(),
			Quatity: data.GetQuatity(),
			TotalPrice: data.GetTotalPrice(),
			// CreatedAt: timestamppb.Now(),
		})
		if err != nil {
			return nil, err
		}
	}

	return result, nil

}
