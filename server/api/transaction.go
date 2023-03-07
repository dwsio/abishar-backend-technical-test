package api

import (
	"context"
	"google.golang.org/protobuf/types/known/timestamppb"

	
	"abishar-backend-technical-test/server/pb"
)

func (s *Server) CreateTransactionMultiple(ctx context.Context, req *pb.CreateTransactionMultipleRequest) (*pb.CreateTransactionMultipleResponse, error) {

	result := &pb.CreateTransactionMultipleResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	for _, data := range req.GetData() {
		s.provider.UpdateOrCreateTransaction(ctx, &pb.Transactions{
			RequestId: req.GetRequestId(),
			ProductId: data.GetProductId(),
			CustomerName: data.GetCustomerName(),
			Quatity: data.GetQuatity(),
			TotalPrice: data.GetTotalPrice(),
			CreatedAt: timestamppb.Now(),
		})
	}

	return result, nil

}
