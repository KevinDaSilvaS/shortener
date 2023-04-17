package repository

import (
	"fmt"
	"shortener/customtypes"

	"github.com/go-redis/redis"
)

func Conn() customtypes.Conn {
	db := redis.NewClient(&redis.Options{
		Addr:     "172.17.0.3:6379",
		Password: "",
		DB:       0,
	})

	return customtypes.Conn{DB: db}
}

func GetKey(conn customtypes.Conn, key string) (string, error) {
	result := conn.DB.Get(key)
	value, err := result.Result()

	if err == nil || fmt.Sprint(err) == "redis: nil" {
		return value, nil
	}

	return "", err
}

func SetExKey(conn customtypes.Conn, key string, value string) (bool, error) {
	result := conn.DB.SetNX(key, value, 0)

	inserted, err := result.Result()

	if err != nil {
		return false, err
	}

	fmt.Println("From SetExKey", value, key, inserted)

	return inserted, nil
}
