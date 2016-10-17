package handlers

import (
	"encoding/json"
	"fmt"
	"go-websocket-server/types"
	"log"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// GetAnalyticsStream is the handler function to receive the analytics result
// data from Apache Spark
func GetAnalyticsStream(w http.ResponseWriter, r *http.Request) {

	connection, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	defer connection.Close()

	var messageData types.Message

	for {
		messageType, message, err := connection.ReadMessage()
		if err != nil {
			return
		}
		log.Println(string(message))
		log.Println(string(messageType))

		err = json.Unmarshal(message, &messageData)
		if err != nil {
			log.Println(err)
		}
		fmt.Println(messageData)

		if messageData.MessageType == "initial" {
			log.Println("receive initial message")
			if token, ok := messageData.Data["token"].(string); ok {
				token, erro := validateToken(token)
				if erro != nil {
					log.Println("token is not valid")
				} else {
					claims, ok := token.Claims.(jwt.MapClaims)
					if ok && token.Valid {
						fmt.Println(claims["sub"])
						id, err2 := strconv.Atoi(claims["sub"].(string))
						if err2 != nil {
							log.Println(err2)
						}
						if setConnection(id, connection) {
							log.Println("connection successfully set")
						} else {
							log.Println("error on set connection")
						}
					}
				}
			}
		} else {
			log.Println("not an initial message")
		}

		err = connection.WriteMessage(messageType, message)
		if err != nil {
			return
		}
	}
}

func validateToken(tokenString string) (validatedToken *jwt.Token, err error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	return token, err
}
