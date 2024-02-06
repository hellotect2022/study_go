package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // websocket 에대한 포인터지정
// var broadcast = make(chan Message)           // 클라이언트에서 보낸 메세지 큐잉
var broadcast = make(chan string)

var upgrader = websocket.Upgrader{}

type Message struct {
	Email    string `json:"email"` // 직렬화 및 역 직렬화 시에 매핑해주는 필드명
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	fs := http.FileServer(http.Dir("./static")) // 정적 파일 서버 생성
	http.Handle("/", fs)
	// websocket으로 루팅을 연결
	http.HandleFunc("/ws", handleSocketConnection)
	go broadCastMessages()

	err := http.ListenAndServe(":7777", nil)
	if err != nil {
		log.Fatal("ListenAndserve: ", err)
	}

}

// 소켓연결
func handleSocketConnection(w http.ResponseWriter, r *http.Request) {
	//fmt.Fprint(w, "Hello, World!")
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrader error :", err)
		return
	}

	defer conn.Close()

	clients[conn] = true

	for {
		//var msg Message
		//_ = msg
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Received message:", msg)
		fmt.Println("Received message:", string(msg))

		err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, client!"))
		if err != nil {
			log.Println(err)
			return
		}

		broadcast <- string(msg)
	}

}

// 참여사용자에게 broadcasting 하기
// redis 가 있는 경우 pubsub 으로 대체가능할듯
func broadCastMessages() {
	for {
		// 채널을 통해서 다음 메세지를 받는다.
		msg := <-broadcast
		// 현재 접속중인 클라이언트에게 메세지를 보낸다.
		for client := range clients {
			fmt.Println("client는 과연 어떻게 생겼을까? : ", client)
			err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			if err != nil {
				log.Fatal("broadCastMessages error :", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}