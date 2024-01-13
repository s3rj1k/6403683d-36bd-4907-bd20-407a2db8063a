package repositories

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"

	"code.local/test/config"
	"code.local/test/models"
)

const (
	keyPrefix               = "signatures"
	defaultOperationTimeout = 30 * time.Second
)

var ErrNotFound = errors.New("not found")

type SignatureRepository struct {
	etcdClient *clientv3.Client
}

func NewSignatureRepository(cfg clientv3.Config) *SignatureRepository {
	c, err := clientv3.New(config.DB)
	if err != nil {
		panic(err)
	}

	return &SignatureRepository{
		etcdClient: c,
	}
}

func getKey(userID string) string {
	return fmt.Sprintf("%s/%s", keyPrefix, userID)
}

func (r *SignatureRepository) Close() error {
	return r.Close()
}

func (r *SignatureRepository) SaveSignature(userID string, signature models.Data) error {
	signatureBytes, err := json.Marshal(signature)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), defaultOperationTimeout)
	defer cancel()

	_, err = r.etcdClient.Put(ctx, getKey(userID), string(signatureBytes))
	return err
}

func (r *SignatureRepository) GetSignature(userID string) (models.Data, error) {
	ctx, cancel := context.WithTimeout(context.Background(), defaultOperationTimeout)
	defer cancel()

	resp, err := r.etcdClient.Get(ctx, getKey(userID))
	if err != nil {
		return models.Data{}, err
	}

	if len(resp.Kvs) == 0 {
		return models.Data{}, fmt.Errorf("signature for user ID %q: %w", userID, ErrNotFound)
	}

	var signature models.Data
	err = json.Unmarshal(resp.Kvs[0].Value, &signature)
	if err != nil {
		return models.Data{}, err
	}

	return signature, nil
}
