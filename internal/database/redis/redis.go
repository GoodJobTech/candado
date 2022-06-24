package redis

import (
	"context"
	"log"
	"os"

	"github.com/go-redis/redis/v8"
	"github.com/goodjobtech/candado/internal/errors"
)

type Redis struct {
	client   *redis.Client
	Host     string
	Port     string
	Password string
}

func New() *Redis {
	host, ok := os.LookupEnv("REDIS_HOST")
	if !ok {
		host = "localhost"
		log.Println("REDIS_HOST is not set, using default value: 'localhost'")
	}

	port, ok := os.LookupEnv("REDIS_PORT")
	if !ok {
		port = "6379"
		log.Println("REDIS_PORT is not set, using default value: '6379'")
	}

	password, ok := os.LookupEnv("REDIS_PASSWORD")
	if !ok {
		password = ""
		log.Println("REDIS_PASSWORD is not set, using default value: ''")
	}

	rd := &Redis{
		Host:     host,
		Port:     port,
		Password: password,
	}

	rd.connect()

	return rd
}

func (r *Redis) connect() error {
	client := redis.NewClient(&redis.Options{
		Addr:     r.Host + ":" + r.Port,
		Password: r.Password,
	})

	r.client = client

	return nil
}

func (r *Redis) lock(ctx context.Context, id string) error {
	switch r.client.SetNX(ctx, id, "1", 0).Err() {
	case nil:
		return nil
	case redis.Nil:
		return errors.ErrAlreadyLocked
	default:
		return nil
	}
}

func (r *Redis) Lock(id string) error {
	return r.lock(context.Background(), id)
}

func (r *Redis) unlock(ctx context.Context, id string) error {
	return r.client.Del(ctx, id).Err()
}

func (r *Redis) Unlock(id string) error {
	return r.unlock(context.Background(), id)
}

func (r *Redis) Heartbeat(id string) (uint16, error) {
	val, err := r.client.Get(context.Background(), id).Result()
	if err != nil {
		return 0, err
	}

	if val == "1" {
		return 1, nil
	}

	return 0, nil
}
