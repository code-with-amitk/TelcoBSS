package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	"telcobss/internal/common/db"
	"telcobss/internal/common/models"
)

type OrderRepository struct {
	couchbase *db.CouchbaseClient
	sqlDB     *sql.DB
}

type OrderService struct {
	repo     *OrderRepository
	producer kafkaProducer
}

type kafkaProducer interface {
	Publish(topic, key string, value []byte) error
}

func NewOrderRepository(cb *db.CouchbaseClient, sqlDB *sql.DB) *OrderRepository {
	return &OrderRepository{couchbase: cb, sqlDB: sqlDB}
}

func NewOrderService(repo *OrderRepository, producer kafkaProducer) *OrderService {
	return &OrderService{repo: repo, producer: producer}
}

func (s *OrderService) PlaceOrder(ctx context.Context, order *models.OrderRequest) error {
	if err := s.validateInventory(ctx, order.ProductCode); err != nil {
		return err
	}

	if err := s.repo.SaveOrder(ctx, order); err != nil {
		return err
	}

	eventPayload, err := json.Marshal(order)
	if err != nil {
		return err
	}
	return s.producer.Publish("order_events", order.OrderID, eventPayload)
}

func (s *OrderService) validateInventory(ctx context.Context, productCode string) error {
	// Circuit breaker demo: wrap an external inventory check.
	// Example using go-hystrix:
	// hystrix.Do("inventory_check", func() error {
	//     // callInventorySystem(productCode)
	//     return nil
	// }, nil)
	return nil
}

func (r *OrderRepository) SaveOrder(ctx context.Context, order *models.OrderRequest) error {
	// Persist order in MySQL and Couchbase.
	if err := db.InsertOrderItem(r.sqlDB, order.OrderID, order.ProductCode, 1000); err != nil {
		return err
	}

	docID := fmt.Sprintf("order_%s", order.OrderID)
	return r.couchbase.UpsertDocument(ctx, "orders", docID, order)
}
