package config

import (
	"os"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

const (
	JWTSecretKey             = "SECRET_KEY"
	JWTSecretKeyDefaultValue = "SECRET_VALUE"
)

var DB = clientv3.Config{ // should be plain connection string and not concrete type
	Endpoints:   []string{"localhost:2379"},
	DialTimeout: 30 * time.Second,
}

func init() {
	if os.Getenv(JWTSecretKey) == "" {
		if err := os.Setenv(JWTSecretKey, JWTSecretKeyDefaultValue); err != nil {
			panic(err)
		}
	}
}
