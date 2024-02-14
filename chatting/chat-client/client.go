package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	chatClient "newProj/chatting/chat-client/model"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// 임의의 키를 사용하여 쿠키 스토어를 만듭니다.
var store = sessions.NewCookieStore([]byte("your-secret-key"))

var tempRedis_userId = make(map[string]string)
var tempRedis_roomId = make(map[string]interface{})

func main() {
	// 1. 정적 파일 서버 생성
	staticFileServer := http.FileServer(http.Dir("./static/html"))
	// 2. 서버 Router 지정
	r := mux.NewRouter()
	// 	2.1.http.StripPrefix() 함수를 사용하여 경로에서 /static/ 접두사를 제거하고 파일 서버 핸들러에 전달합니다.
	r.PathPrefix("/view/").Handler(http.StripPrefix("/view/", staticFileServer))
	// 	2.2 api router 설정
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("GET")
	r.HandleFunc("/getUserListAll", getUserListAllHandler).Methods("GET")
	r.HandleFunc("/createRoom", createRoomHandler).Methods("POST")

	http.Handle("/", r)
	fmt.Println("client Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := chatClient.User{}
	responseBody := chatClient.ResponseData{}

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// 받은 JSON 데이터 사용 예시
	fmt.Printf("Received username: %s\n", requestBody.Username)
	fmt.Printf("Received password: %s\n", requestBody.Password)

	// 여기에 로직을 추가해야함
	responseBody.Status = "success"
	// 여기에 session 관련로직을 추가
	//setSessionHandler(w, r)
	// 대신 임시 redis 에 값을 저장해서 상요한다. DB X

	if _, exists := tempRedis_userId[requestBody.Username]; requestBody.Username != "" && exists {
		fmt.Println("여기에 이미 있다.")
	} else {
		tempRedis_userId[requestBody.Username] = requestBody.Password
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)

}

func setSessionHandler(w http.ResponseWriter, r *http.Request) {
	// 세션을 설정합니다.
	session, err := store.Get(r, "session-test")
	if err != nil {
		fmt.Println("세션 생성시 에러발생 :", err)
	}
	// 세션에 값을 저장합니다.
	session.Values["key"] = "value"
	err = session.Save(r, w)
	if err != nil {
		// 세션을 저장합니다.
		log.Fatal("세션 저장시 에러발생 :", err)
	}

}

// GET
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session-test")
	val := session.Values["key"]

	fmt.Println("session key is : ", val)
	fmt.Println("session is : ", session, *session)

	// 세션에서 특정 키의 값을 삭제
	delete(session.Values, "key")

	// 변경된 세션을 저장
	err := session.Save(r, w)
	if err != nil {
		// 세션을 저장하는 중에 오류가 발생한 경우 처리
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func getUserListAllHandler(w http.ResponseWriter, r *http.Request) {
	responseBody := chatClient.ResponseData{}

	responseBody.Status = "success"
	responseBody.Result = tempRedis_userId

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)

}

func createRoomHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := chatClient.RoomInfo{}
	responseBody := chatClient.ResponseData{}

	json.NewDecoder(r.Body).Decode(&requestBody)
	requestBody.RoomId = strconv.Itoa(len(tempRedis_roomId))
	fmt.Println("새로 생성된 Room -> ", responseBody)
	tempRedis_roomId[requestBody.RoomId] = requestBody
	fmt.Println("서버에 생성된 RoomList -> ", tempRedis_roomId)
	// 아마 여기서 redis pubsub 으로 구현할듯함

	// 성공일때의 얘기
	responseBody.Status = "success"
	responseBody.Message = "채팅방을 생성하였습니다."
	responseBody.Result = requestBody

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseBody)

}

/*
func loginHandler(w http.ResponseWriter, r *http.Request) {
	requestBody := chatClient.User{}
	responseBody := chatClient.ResponseData{}
	err := requestPostReturnJson(w, r, user, &responseBody)

	if err != nil {
		responseBody.Status = "fail"
		responseBody.Message = err.Error()
	}
	// 클라이언트에게 응답을 보냅니다.

	w.WriteHeader(http.StatusOK)
	responseBody.Status = "success"
	responseBody.Message = "nice to meet you"

	//JSON으로 직렬화한 후에 응답을 보냅니다.
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

}

func requestPostReturnJson(w http.ResponseWriter, r *http.Request, requestBody interface{}, responseBody interface{}) error {
	// POST 요청의 바디를 읽기 위해 ioutil.ReadAll()을 사용합니다.
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return errors.New("only POST requests are allowed")
	}

	// responseBody를 ResponseData 타입으로 타입 어설션합니다.
	respData, ok := responseBody.(*chatClient.ResponseData)
	if !ok {
		return errors.New("responseBody is not of type *ResponseData")
	}



		// 요청 바디를 읽어옵니다.
		body, err := io.ReadAll(r.Body)
	//fmt.Println("body : ", string(body))
	if err != nil {
		http.Error(w, "Failed to read request body", http.StatusInternalServerError)
		return err
	}
	defer r.Body.Close()

	// JSON 데이터를 파싱하여 LoginForm 구조체에 저장합니다.
	err = json.Unmarshal(body, &requestBody)
	//fmt.Println("json.Unmarshal : ", requestBody)
	if err != nil {
		http.Error(w, "Failed to parse JSON body", http.StatusBadRequest)
		return err
	}
	return nil
}
*/
