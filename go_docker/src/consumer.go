package main

import (
	"fmt"
    "time"
    "encoding/json"
	"github.com/go-redis/redis"
)

type Payload struct {
    UserId int64        `json:"userid"`
    Message string      `json:"message"`
    MessageType int64   `json:"message_type"`
    EventTime int64     `json:"ts"`
}

func main() {

    time.Sleep(5 * time.Second)

    redisdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "abcd1234",
		DB:       0,  // use default DB
	})

	pong, err := redisdb.Ping().Result()
	fmt.Println(pong, err)

    // for without any condition is an infinite loop.
    for {

        //BLPop: blocking left pop. Push happens from right.
        // use `redisdb.BLPop(0, "queue")` for infinite waiting time
        result, err := redisdb.BLPop(0, "gopy_message_queue").Result()
        if err != nil {
            panic(err)
        }

        fmt.Println("\n", result[0], result[1])

        go processMessage(result[1])

    }
}

func processMessage(message string) {
    var p Payload

    err := json.Unmarshal([]byte(message), &p)

    if err != nil {
        fmt.Println(err)
        return
    }

    fmt.Println(p)
}
