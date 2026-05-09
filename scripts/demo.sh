#!/usr/bin/env bash
set -euo pipefail

echo "Running TelcoBSS demo script..."

docker compose -f deployments/docker-compose/docker-compose.yml up -d
sleep 15

echo "Creating sample customer..."
curl -s -X POST http://localhost:8081/v1/customers \
  -H 'Content-Type: application/json' \
  -d '{"customer_id":"cust-001","name":"Jane Doe","email":"jane@example.com","phone":"+1234567890"}'

echo "\nPlacing sample order..."
curl -s -X POST http://localhost:8082/v1/orders \
  -H 'Content-Type: application/json' \
  -d '{"order_id":"order-001","customer_id":"cust-001","product_code":"mobile-plan-1","quantity":1}'

echo "\nInjecting usage event into Kafka..."
cat <<EOF | docker exec -i kafka kafka-console-producer.sh --broker-list localhost:9092 --topic usage_events --property parse.key=true --property key.separator=:
usage-001:{"usage_id":"usage-001","customer_id":"cust-001","bytes":10240}
EOF

sleep 5

echo "Fetching invoice (may not exist until billing pipeline is implemented)..."
curl -s http://localhost:8083/v1/invoices/invoice-001 || true

echo "Demo script complete."
