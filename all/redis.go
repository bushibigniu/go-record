package main

import (
	"fmt"
	"github.com/go-redis/redis"
)

func ExampleNewClient()  {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err = client.Set("key_go", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := client.Get("key_go").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)
}

func main()  {
	ExampleNewClient()
}
