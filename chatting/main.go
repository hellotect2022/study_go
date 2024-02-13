package main

import (
	"fmt"
	"log"
	"net/http"
	"newProj/chatting/staticRouter"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // websocket 에대한 포인터지정
var broadcast = make(chan Message)           // 클라이언트에서 보낸 메세지 큐잉
//var broadcast = make(chan string)

var upgrader = websocket.Upgrader{}

var tempRoomList = make(map[int]RoomInfo)

type RoomInfo struct {
	RoomId int      `json: "roomId"`
	Member []string `json:"member"`
}

type Message struct {
	Email    string `json:"email"` // 직렬화 및 역 직렬화 시에 매핑해주는 필드명
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	// 정적 파일 출력 Routing
	staticRouter.RoutingStaticPage("./staticRouter/html")

	// websocket으로 루팅을 연결
	http.HandleFunc("/ws", handleSocketConnection)
	go broadCastMessages()

	log.Fatal(http.ListenAndServe(":7777", nil))
}

// 소켓연결
func handleSocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrader error :", err)
		return
	}

	defer conn.Close()

	clients[conn] = true
	fmt.Println("socket 연결된 client 수 :", len(clients))

	// 임이의 room number 를 가져와서 대입하고 roomId 를 전달한다.
	err = conn.WriteMessage(websocket.TextMessage, []byte("msg"))
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Println("socket 연결된 client :", conn)
	for {
		var messageStruct Message
		//_ = msg
		//_, msg, err := conn.ReadMessage()
		err := conn.ReadJSON(&messageStruct)
		if err != nil {
			fmt.Println("err (conn.ReadJSON(&messageStruct)) : ", err)
			delete(clients, conn)
			return
		}
		//fmt.Println("Received message:", msg)
		//fmt.Println("Received message:", string(*messageStruct))

		fmt.Println("Message : ", messageStruct)

		broadcast <- messageStruct
	}

}

// 참여사용자에게 broadcasting 하기
// redis 가 있는 경우 pubsub 으로 대체가능할듯
func broadCastMessages() {
	for {
		// 채널을 통해서 다음 메세지를 받는다.
		messageStruct := <-broadcast
		// 현재 접속중인 클라이언트에게 메세지를 보낸다.
		for client := range clients {
			//fmt.Println("client는 과연 어떻게 생겼을까? : ", client)
			//err := client.WriteMessage(websocket.TextMessage, []byte(msg))
			err := client.WriteJSON(messageStruct)
			if err != nil {
				log.Fatal("broadCastMessages error :", err)
				client.Close()
				delete(clients, client)
			}
		}
	}
}
