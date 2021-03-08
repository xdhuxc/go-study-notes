package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strconv"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/elastic/go-elasticsearch/v7/esapi"
)

type User struct {
	ID      string
	Name    string
	Address string
}

func (u User) String() string {
	if dataInBytes, err := json.Marshal(&u); err == nil {
		return string(dataInBytes)
	}

	return ""
}

func main() {
	cfg := elasticsearch.Config{
		Addresses:            []string{"http://10.20.13.158:9200"},
		Username:             "elastic",
		Password:             "Shareit@2020",
		DisableRetry:         false,
		EnableRetryOnTimeout: false,
		MaxRetries:           3,
		EnableMetrics:        true,
		EnableDebugLogger:    true,
	}
	esClient, err := elasticsearch.NewClient(cfg)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = esClient.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}

	users := generateUsers()
	var buf bytes.Buffer
	for _, user := range users {
		metadata := []byte(fmt.Sprintf(`{ "index" : { "_id" : "%s" } }%s`, user.ID, "\n"))

		data, err := json.Marshal(user)
		if err != nil {
			log.Println(err)
			continue
		}

		data = append(data, "\n"...)
		buf.Grow(len(metadata) + len(data))
		buf.Write(metadata)
		buf.Write(data)
	}

	resp, err := esClient.Bulk(bytes.NewReader(buf.Bytes()), esClient.Bulk.WithIndex("xdhuxc-bulk"))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp.String())
}

func Insert(es *elasticsearch.Client, users []User) error {
	for i, user := range users {
		req := esapi.IndexRequest{
			Index:      "xdhuxc-test",
			DocumentID: strconv.Itoa(i + 1),
			Body:       strings.NewReader(user.String()),
			Refresh:    "true",
		}

		resp, err := req.Do(context.Background(), es)
		if err != nil {
			log.Println(err)
		}
		defer resp.Body.Close()

		if resp.IsError() {
			log.Printf("[%s] Error indexing document ID=%d", resp.Status(), i+1)
		} else {
			// Deserialize the response into a map.
			var r map[string]interface{}
			if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
				log.Printf("Error parsing the response body: %s", err)
			} else {
				// print the response body
				dataInBytes, err := json.Marshal(&r)
				if err != nil {
					log.Println(err)
				}
				fmt.Println(string(dataInBytes))

				// Print the response status and indexed document version.
				log.Printf("[%s] %s; version=%d", resp.Status(), r["result"], int(r["_version"].(float64)))
			}
		}
	}

	return nil
}

func generateUsers() []User {
	var users []User
	for i := 0; i < 100; i++ {
		u := User{
			ID:      uuid.New().String(),
			Name:    uuid.New().String(),
			Address: "xdhuxc" + uuid.New().String(),
		}
		users = append(users, u)
	}

	return users
}
