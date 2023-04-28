package redis_cache

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v8"
	"go.opentelemetry.io/otel"
	"log"
	"practice_optelem/internal/models"
)

type Cache struct {
	client *redis.Client
}

func NewCache(client *redis.Client) *Cache {
	return &Cache{client: client}
}

func (c *Cache) SetHash(ctx context.Context) error {
	//provide cache trace
	ctxt, span := otel.Tracer("practice-service").Start(ctx, "Repository.SetHash")
	defer span.End()
	data := models.RedisData{
		Name:  "hello",
		Value: "bro",
		Filed: models.SomeThing{
			Fields: []string{"how", "are", "you"},
		},
	}
	value, err := ConvertToMap(data)
	if err != nil {
		return err
	}
	fmt.Println(value)
	if err = c.client.HMSet(ctxt, "qwerty", value).Err(); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func ConvertToMap(data models.RedisData) (map[string]interface{}, error) {
	jsData, ErrStr := MarshalStruct(data.Filed)
	if ErrStr != nil {
		return nil, ErrStr
	}

	return map[string]interface{}{
		"name":   data.Name,
		"value":  data.Value,
		"fields": jsData,
	}, nil
}

func MarshalStruct(data interface{}) ([]byte, error) {
	response, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	return response, nil
}
