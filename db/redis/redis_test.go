package redis

import (
	"syt-crawler/config"
	"testing"
)

func connect() (*RedisDB, error) {
	return Connect(&config.Redis{
		Host:    "127.0.0.1",
		Port:    6379,
		Db:      0,
		Enabled: 1,
	})
}

func TestRedisDB_TTL(t *testing.T) {
	db, err := connect()
	if err != nil {
		t.Fatal(err)
	}
	db.TTL("syt-crawler:093c8a0e-518c-49fd-9a8e-0553d1739bb12")
}