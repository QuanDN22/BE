package internal

import (
	"context"
	"fmt"
	"log"

	"github.com/QuanDN22/BE/gRPC/example-go-grpc-gateway/protogen/golang/orders"
)

type OrderService struct {
	db *DB
	orders.UnimplementedOrdersServer
}

func NewOrdersService(db *DB) OrderService {
	return OrderService{db: db}
}

func (s *OrderService) AddOrder(_ context.Context, in *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Recieved an add-order request")

	err := s.db.AddOrder(in.GetOrder())

	// do something
	return &orders.Empty{}, err
}

// GetOrder implements the GetOrder method of the grpc OrdersServer interface to fetch an order for a given orderID
func (o *OrderService) GetOrder(_ context.Context, req *orders.PayloadWithOrderID) (*orders.PayloadWithSingleOrder, error) {
	log.Printf("Received get order request")
	order := o.db.GetOrderByID(req.GetOrderId())
	if order == nil {
		return nil, fmt.Errorf("order not found for orderID: %d", req.GetOrderId())
	}
	return &orders.PayloadWithSingleOrder{Order: order}, nil
}

// UpdateOrder implements the UpdateOrder method of the grpc OrdersServer interface to update an order
func (o *OrderService) UpdateOrder(_ context.Context, req *orders.PayloadWithSingleOrder) (*orders.Empty, error) {
	log.Printf("Received an update order request")
	o.db.UpdateOrder(req.GetOrder())
	return &orders.Empty{}, nil
}

// RemoveOrder implements the RemoveOrder method of the grpc OrdersServer interface to remove an order
func (o *OrderService) RemoveOrder(_ context.Context, req *orders.PayloadWithOrderID) (*orders.Empty, error) {
	log.Printf("Received a remove order request")
	o.db.RemoveOrder(req.GetOrderId())
	return &orders.Empty{}, nil
}