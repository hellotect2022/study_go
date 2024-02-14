package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool) // websocket 에대한 포인터지정
var broadcast = make(chan Message)           // 클라이언트에서 보낸 메세지 큐잉
// var broadcast = make(chan string)
var upgrader = websocket.Upgrader{}

// var tempRoomList = make(map[int]RoomInfo)
// type RoomInfo struct {
// 	RoomId int      `json: "roomId"`
// 	Member []string `json:"member"`
// }

type ClientsInfo struct {
	userId        string
	joinedRoom    map[string]bool
	onConnected   bool
	websocketConn *websocket.Conn
}

var clientsTest = make(map[string]ClientsInfo)

type Message struct {
	Type     string `json:"type"`
	UserID   string `json:"userId"`
	Email    string `json:"email"` // 직렬화 및 역 직렬화 시에 매핑해주는 필드명
	Username string `json:"username"`
	Message  string `json:"message"`
}

func main() {
	// websocket으로 루팅을 연결
	http.HandleFunc("/ws", handleSocketConnection)
	go broadCastMessages()
	fmt.Println("websocket Server is running on port 7777...")
	log.Fatal(http.ListenAndServe(":7777", nil))
}

// func generateUserID(clients *ClientsInfo) {
// 	// 0 이상 100 미만의 임의의 정수를 생성합니다.
// 	_, exists := clientsTest[randomNumber]
// 	if !exists {
// 		clients.userId = randomNumber
// 		clientsTest[randomNumber] = *clients
// 	}
// }

// 소켓연결
func handleSocketConnection(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("upgrader error :", err)
		return
	}

	defer conn.Close()

	//fmt.Println("socket 연결된 client :", conn)
	for {
		var messageStruct Message
		err := conn.ReadJSON(&messageStruct)
		if err != nil {
			fmt.Println("err (conn.ReadJSON(&messageStruct)) : ", err)
			delete(clients, conn)
			return
		}

		switch messageStruct.Type {
		case "connect":
			connectUser(&messageStruct, conn)
			break
		case "message":
			fmt.Print("메세지를 받았음")
			//broadcast <- messageStruct
			break
		}

	}

}

func connectUser(messageStruct *Message, conn *websocket.Conn) {
	userID := messageStruct.UserID
	clientsTest[userID] = ClientsInfo{userId: userID, onConnected: true, websocketConn: conn}

	fmt.Println("clinet숫자 : ", len(clientsTest))
	fmt.Println("clinet 들 : ", clientsTest)

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
