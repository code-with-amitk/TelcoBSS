package service

import (
	"context"
	"database/sql"
	"encoding/json"

	"telcobss/internal/common/models"
)

type BillingRepository struct {
	sqlDB *sql.DB
}

type BillingService struct {
	repo     *BillingRepository
	producer kafkaProducer
}

type kafkaProducer interface {
	Publish(topic, key string, value []byte) error
}

func NewBillingRepository(sqlDB *sql.DB) *BillingRepository {
	return &BillingRepository{sqlDB: sqlDB}
}

func NewBillingService(repo *BillingRepository, producer kafkaProducer) *BillingService {
	return &BillingService{repo: repo, producer: producer}
}

func (s *BillingService) FetchInvoice(ctx context.Context, invoiceID string) (*models.Invoice, error) {
	query := `SELECT invoice_id, customer_id, total_cents, status FROM invoices WHERE invoice_id = ?`
	row := s.repo.sqlDB.QueryRowContext(ctx, query, invoiceID)
	invoice := models.Invoice{}
	if err := row.Scan(&invoice.InvoiceID, &invoice.CustomerID, &invoice.TotalCents, &invoice.Status); err != nil {
		return nil, err
	}
	return &invoice, nil
}

func (s *BillingService) ProcessRatedEvent(ctx context.Context, event *models.RatedEvent) error {
	payload, err := json.Marshal(event)
	if err != nil {
		return err
	}
	return s.producer.Publish("billing_events", event.UsageID, payload)
}

func (r *BillingRepository) CreateInvoice(ctx context.Context, invoice *models.Invoice) error {
	query := `INSERT INTO invoices (invoice_id, customer_id, total_cents, status, created_at) VALUES (?, ?, ?, ?, NOW())`
	_, err := r.sqlDB.ExecContext(ctx, query, invoice.InvoiceID, invoice.CustomerID, invoice.TotalCents, invoice.Status)
	return err
}
