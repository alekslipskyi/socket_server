package Connect

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
)

type TypeConnect struct {
	Token string
}

func Init(data string) {
	conn := TypeConnect{}
	payload := json.Unmarshal([]byte(data), &conn)

	client := &http.Client{}
	request, _ := http.NewRequest("GET", "/api/v0/user", nil)
	request.Header.Set("Authorization", conn.Token)

	res, _ := client.Do(request)
	var parsedPayload map[string]interface{}
	_ = json.NewDecoder(res.Body).Decode(&parsedPayload)

	fmt.Printf("test payload are %s", parsedPayload)

}

func Destroy(data interface{}, conn net.Conn) {

}