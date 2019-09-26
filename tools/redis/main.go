package main

import (
	"github.com/go-redis/redis"
	"log"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr:     "10.21.77.227:6379",
		Password: "qZ2JXiGi#LSf0bHBL",
		DB:       0,
	})

	err := client.Ping().Err()
	if err != nil {
		log.Fatalln(err)
	}

	// key := "/sgt_hawkeye/alertmanager_yaml/receivers/1000"
	/*
		key := "/sgt_hawkeye/alertmanager_yaml/receivers/1001"
		value := "{\"name\":\"send_to_dingding_web\",\"resolved\":false,\"type\":\"dingtalk\",\"url\":\"https://oapi.dingtalk.com/robot/send?access_token=3f033821d1ca8839d8075301c3843f4f27df16a75e11e239522feb199cee1b59\"}"
		client.Set(key, value, 0)

	*/

	client.Get("*").Err()
}

func NewRedisClient(address string, password string, index int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     "10.21.77.227:6379",
		Password: "qZ2JXiGi#LSf0bHBL",
		DB:       0,
	})

	return client, client.Ping().Err()
}
