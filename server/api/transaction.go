package api

import (
	"context"

	
	"abishar-backend-technical-test/server/pb"
)

func (s *Server) CreateTransactionMultiple(ctx context.Context, req *pb.CreateTransactionMultipleRequest) (*pb.CreateTransactionMultipleResponse, error) {

	result := &pb.CreateTransactionMultipleResponse{
		Error:   false,
		Code:    200,
		Message: "Success",
	}

	// for i, v := range req.GetData() {
	// 	i
	// }

	return result, nil

}
