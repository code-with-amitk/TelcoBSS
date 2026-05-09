package service

import (
	"context"
	"encoding/json"
	"fmt"

	"telcobss/internal/common/db"
	"telcobss/internal/common/models"
)

type OnboardingRepository struct {
	couchbase *db.CouchbaseClient
}

type OnboardingService struct {
	repo     *OnboardingRepository
	producer kafkaProducer
}

type kafkaProducer interface {
	Publish(topic, key string, value []byte) error
}

func NewOnboardingRepository(cb *db.CouchbaseClient) *OnboardingRepository {
	return &OnboardingRepository{couchbase: cb}
}

func NewOnboardingService(repo *OnboardingRepository, producer kafkaProducer) *OnboardingService {
	return &OnboardingService{repo: repo, producer: producer}
}

func (s *OnboardingService) CreateCustomer(ctx context.Context, customer *models.Customer) error {
	docID := fmt.Sprintf("customer_%s", customer.CustomerID)

	// Store customer in Couchbase collection: customers
	if err := s.repo.StoreCustomer(ctx, docID, customer); err != nil {
		return err
	}

	eventPayload, err := json.Marshal(customer)
	if err != nil {
		return err
	}

	return s.producer.Publish("customer_events", customer.CustomerID, eventPayload)
}

func (r *OnboardingRepository) StoreCustomer(ctx context.Context, docID string, customer *models.Customer) error {
	// StoreCustomer inserts a new customer into Couchbase.
	// Bucket: telco_bss, Collection: customers, Document ID: customer_<UUID>
	return r.couchbase.UpsertCustomer(ctx, docID, customer)
}

func (r *OnboardingRepository) GetCustomer(ctx context.Context, docID string) (*models.Customer, error) {
	var customer models.Customer
	if err := r.couchbase.GetCustomer(ctx, docID, &customer); err != nil {
		return nil, err
	}
	return &customer, nil
}
