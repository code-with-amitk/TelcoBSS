package models

// Customer represents a customer onboarding record.
type Customer struct {
	CustomerID string `json:"customer_id" protobuf:"bytes,1,opt,name=customer_id,json=customerId"`
	Name       string `json:"name" protobuf:"bytes,2,opt,name=name"`
	Email      string `json:"email" protobuf:"bytes,3,opt,name=email"`
	Phone      string `json:"phone" protobuf:"bytes,4,opt,name=phone"`
}

// OrderRequest represents a subscription order.
type OrderRequest struct {
	OrderID     string `json:"order_id" protobuf:"bytes,1,opt,name=order_id,json=orderId"`
	CustomerID  string `json:"customer_id" protobuf:"bytes,2,opt,name=customer_id,json=customerId"`
	ProductCode string `json:"product_code" protobuf:"bytes,3,opt,name=product_code,json=productCode"`
	Quantity    int32  `json:"quantity" protobuf:"varint,4,opt,name=quantity"`
}

// Invoice represents a billing invoice.
type Invoice struct {
	InvoiceID  string `json:"invoice_id" protobuf:"bytes,1,opt,name=invoice_id,json=invoiceId"`
	CustomerID string `json:"customer_id" protobuf:"bytes,2,opt,name=customer_id,json=customerId"`
	TotalCents int64  `json:"total_cents" protobuf:"varint,3,opt,name=total_cents,json=totalCents"`
	Status     string `json:"status" protobuf:"bytes,4,opt,name=status"`
}

// UsageEvent represents a raw usage row.
type UsageEvent struct {
	UsageID  string `json:"usage_id" protobuf:"bytes,1,opt,name=usage_id,json=usageId"`
	CustomerID string `json:"customer_id" protobuf:"bytes,2,opt,name=customer_id,json=customerId"`
	Bytes     int64  `json:"bytes" protobuf:"varint,3,opt,name=bytes"`
}

// RatedEvent is the rated output from the rating engine.
type RatedEvent struct {
	UsageID    string `json:"usage_id" protobuf:"bytes,1,opt,name=usage_id,json=usageId"`
	CustomerID string `json:"customer_id" protobuf:"bytes,2,opt,name=customer_id,json=customerId"`
	AmountCents int64 `json:"amount_cents" protobuf:"varint,3,opt,name=amount_cents,json=amountCents"`
}
