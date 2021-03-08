package main

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf := &api.Config{
		Address: "10.20.13.158:8500",
		Scheme:  "http",
	}

	client, err := api.NewClient(conf)
	if err != nil {
		log.Fatalln(err)
	}

	kv := client.KV()

	testKey := "xdhuxc/consul/123"

	p := &api.KVPair{
		Key:   testKey,
		Value: []byte("nihao"),
	}
	_, err = kv.Put(p, nil)
	if err != nil {
		log.Fatalln(err)
	}

	pair, _, err := kv.Get(testKey, nil)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("KV: %v %s\n", pair.Key, pair.Value)
}
