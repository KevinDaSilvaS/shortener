package customtypes

import "github.com/go-redis/redis"

type NewLink struct {
	Url   string
	Alias string
}

type Conn struct {
	DB *redis.Client
}
