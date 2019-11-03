package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/go-redis/redis"

	log "github.com/sirupsen/logrus"
)

func main() {
	client, err := NewRedisClient("127.0.0.1:6379", "aBc#12pD", 0)
	if err != nil {
		log.Errorln(err)
		return
	}

	values, err := GetByKeyName(client, "rule")
	for _, value := range values {
		fmt.Println(value)
	}
}

// 创建 redis 客户端
func NewRedisClient(address string, password string, index int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       index,
	})

	return client, client.Ping().Err()
}

// 获取 key-value
func Get(client *redis.Client, key string) (string, error) {
	return client.Get(key).Result()
}

// 模糊查询 key，此操作有危险性，在数据量大时非常消耗性能
func GetByKeyName(client *redis.Client, key string) ([]string, error) {
	var values []string
	err := client.Keys(fmt.Sprintf("*%s*", key)).ScanSlice(&values)
	if err != nil {
		return values, err
	}

	return values, nil
}

// 模糊查询 value，自己实现，此操作会有性能问题
func GetByValue(client *redis.Client, value string) (map[string]string, error) {
	var keys []string

	err := client.Keys("*").ScanSlice(&keys)
	if err != nil {
		return nil, err
	}

	result := make(map[string]string)
	for _, k := range keys {
		v, err := client.Get(k).Result()
		if err != nil {
			continue
		}
		if strings.Contains(v, value) {
			result[k] = v

		}
	}

	return result, nil
}

// 创建永不过期的 key-value
func Set(client *redis.Client, key string, value string) error {
	return client.Set(key, value, 0).Err()
}

// 创建过期的 key-value
func SetWithDuration(client *redis.Client, key string, value string, duration time.Duration) error {
	return client.Set(key, value, duration).Err()
}

// 对于 redis，创建和更新的操作效果一样

// 删除一个 key-value
func Delete(client *redis.Client, key string) error {
	return client.Del(key).Err()
}

// 删除一组 key-value
func DeleteKeys(client *redis.Client, keys ...string) error {
	return client.Del(keys...).Err()
}

// 删除 redis 数据库所有数据，使用异步删除，主线程会把删除任务交给异步线程来完成删除任务
func DeleteDB(client *redis.Client) error {
	return client.FlushAllAsync().Err()
}

// 删除当前 redis 数据库，如果想要删除其他的 redis 数据库，需要构造 map 存储各个 redis 数据库的客户端，想删除指定的数据库时，取出其客户端即可
func DeleteCurrentDB(client *redis.Client) error {
	return client.FlushDBAsync().Err()
}
