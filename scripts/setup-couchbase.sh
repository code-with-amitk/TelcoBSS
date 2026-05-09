#!/usr/bin/env bash
set -euo pipefail

CB_HOST=${COUCHBASE_HOST:-localhost}
USER=${COUCHBASE_USER:-Administrator}
PASS=${COUCHBASE_PASSWORD:-password}
BUCKET=${COUCHBASE_BUCKET:-telco_bss}

cat <<EOF
Creating Couchbase bucket and indexes for TelcoBSS...
EOF

curl -s -u "$USER:$PASS" -X POST "http://$CB_HOST:8091/pools/default/buckets" \
  -d name="$BUCKET" \
  -d ramQuotaMB=256 \
  -d bucketType=couchbase \
  -d evictionPolicy=valueOnly || true

sleep 5

curl -s -u "$USER:$PASS" -X POST "http://$CB_HOST:8091/cli/query/service" \
  -d statement="CREATE PRIMARY INDEX ON \\`$BUCKET\\`;" || true
curl -s -u "$USER:$PASS" -X POST "http://$CB_HOST:8091/cli/query/service" \
  -d statement="CREATE INDEX idx_customers_customer_id ON \\`$BUCKET\\`._default.customers(customer_id);" || true
curl -s -u "$USER:$PASS" -X POST "http://$CB_HOST:8091/cli/query/service" \
  -d statement="CREATE INDEX idx_orders_order_date ON \\`$BUCKET\\`._default.orders(order_date);" || true

echo "Couchbase setup complete."
