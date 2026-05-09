package db

import (
	"context"
	"fmt"

	gocb "github.com/couchbase/gocb/v2"
)

const (
	CouchbaseBucket = "telco_bss"
)

// CouchbaseClient wraps the Couchbase cluster handle.
type CouchbaseClient struct {
	cluster *gocb.Cluster
	bucket  *gocb.Bucket
}

func NewCouchbaseClient(url, username, password string) (*CouchbaseClient, error) {
	cluster, err := gocb.Connect(url, gocb.ClusterOptions{Username: username, Password: password})
	if err != nil {
		return nil, fmt.Errorf("connect Couchbase: %w", err)
	}

	bucket := cluster.Bucket(CouchbaseBucket)
	return &CouchbaseClient{cluster: cluster, bucket: bucket}, nil
}

func (c *CouchbaseClient) QueryCustomersByField(ctx context.Context, field, value string, dest interface{}) error {
	query := fmt.Sprintf("SELECT `%s`.* FROM `%s`.%s.%s WHERE `%s` = $1", CouchbaseBucket, CouchbaseBucket, "_default", "customers", field)
	rows, err := c.cluster.Query(query, &gocb.QueryOptions{PositionalParameters: []interface{}{value}})
	if err != nil {
		return err
	}
	defer rows.Close()
	return rows.One(dest)
}

func (c *CouchbaseClient) Collection(collectionName string) *gocb.Collection {
	scope := c.bucket.DefaultScope()
	return scope.Collection(collectionName)
}

func (c *CouchbaseClient) UpsertCustomer(ctx context.Context, docID string, document interface{}) error {
	collection := c.Collection("customers")
	_, err := collection.Upsert(docID, document, nil)
	return err
}

func (c *CouchbaseClient) GetCustomer(ctx context.Context, docID string, dest interface{}) error {
	collection := c.Collection("customers")
	result, err := collection.Get(docID, nil)
	if err != nil {
		return err
	}
	return result.Content(dest)
}

func (c *CouchbaseClient) UpsertDocument(ctx context.Context, collectionName, docID string, document interface{}) error {
	collection := c.Collection(collectionName)
	_, err := collection.Upsert(docID, document, nil)
	return err
}
