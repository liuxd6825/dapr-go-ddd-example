package service

import (
	"context"
	"fmt"
	pd "github.com/liuxd6825/dapr-go-ddd-example/command-service/userinterface/grpc/service/proto"
)

type ProdService struct {
}

func (ps *ProdService) GetProductStock(ctx context.Context, req *pd.ProductRequest) (*pd.ProductResponse, error) {
	fmt.Println(req.ProdId)
	return &pd.ProductResponse{ProdStock: 1}, nil
}
