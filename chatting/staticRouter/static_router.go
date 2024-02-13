package staticRouter

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"newProj/chatting/model"
)

func RoutingStaticPage(filePath string) {
	fs := http.FileServer(http.Dir(filePath)) // 정적 파일 서버 생성
	http.Handle("/", fs)
	http.HandleFunc("/postLogin", postLogin)
	fmt.Println("Server is running on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func postLogin(w http.ResponseWriter, r *http.Request) {
	var user User
	responseBody := model.ResponseData{}
	err := requestPostReturnJson(w, r, user)
	if err != nil {
		responseBody.Status = "fail"
		responseBody.Message = err.Error()
	}
	// 클라이언트에게 응답을 보냅니다.

	w.WriteHeader(http.StatusOK)
	responseBody.Status = "success"
	responseBody.Message = "nice to meet you"

	// JSON으로 직렬화한 후에 응답을 보냅니다.
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(responseBody)
	if err != nil {
		http.Error(w, "Failed to encode JSON response", http.StatusInternalServerError)
		return
	}

}

func requestPostReturnJson(w http.ResponseWriter, r *http.Request, requestBody interface{}) error {
	// POST 요청의 바디를 읽기 위해 ioutil.ReadAll()을 사용합니다.
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST requests are allowed", http.StatusMethodNotAllowed)
		return errors.New("cannot divide by zero")

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
