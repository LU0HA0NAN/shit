package lock

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"sync"
	"time"
)

var mutex sync.Mutex

func Lock(key string, rdb *redis.Client) bool {
	return TryLock(key, "1", 60, rdb)
}

func TryLock(key, val string, lockSeconds int, rdb *redis.Client) bool {
	mutex.Lock()
	defer mutex.Unlock()
	flag, err := rdb.SetNX(context.TODO(), key, val, time.Second*time.Duration(lockSeconds)).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return flag
}

func UnLock(key string, rdb *redis.Client) int64 {
	nums, err := rdb.Del(context.TODO(), key).Result()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return nums
}
