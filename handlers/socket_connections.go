package handlers

import "github.com/gorilla/websocket"

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
	return true
}
