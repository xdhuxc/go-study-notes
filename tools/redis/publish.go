package main

import (
	"encoding/json"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

func main() {
	client, err := newRedisClient("127.0.0.1:6379", "aBc#12pD", 0)
	if err != nil {
		log.Errorln(err)
		return
	}

	channel := "alert_rules"

	for i := 0; i < 100; i++ {
		user := struct {
			ID    int
			Name  string
			Email string
		}{
			ID:    i,
			Name:  "xdhuxc" + strconv.Itoa(i),
			Email: "xdhuxc@163.com",
		}

		data, err := json.Marshal(&user)
		if err != nil {
			log.Println(err)
			return
		}

		err = client.Publish(channel, string(data)).Err()
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("publishing message %d successfully", i)
		}
		time.Sleep(10 * time.Second)
	}
}

// 创建 redis 客户端
func newRedisClient(address string, password string, index int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       index,
	})

	return client, client.Ping().Err()
}
