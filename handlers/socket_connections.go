package handlers

import (
	"fmt"

	"github.com/gorilla/websocket"
	"gopkg.in/redis.v5"
)

var connectionPool map[int]*websocket.Conn

func init() {
	connectionPool = map[int]*websocket.Conn{}
}

func deleteConnection(id int) bool {
	if connectionPool[id] != nil {
		connectionPool[id] = nil
		return true
	}
	return false
}

func getConnection(id int) (connection *websocket.Conn) {
	return connectionPool[id]
}

func setConnection(id int, connection *websocket.Conn) bool {
	if connectionPool[id] != nil {
		deleteConnection(id)
	}
	if connectionPool[id] != nil {
		return false
	}
	connectionPool[id] = connection

	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)

	err2 := client.Set(string(id), "value", 0).Err()
	if err2 != nil {
		panic(err2)
	}

	val, err := client.Get(string(id)).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("1234567890", val)

	return true
}
