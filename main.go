package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const port string = ":4000"

type URLDescription struct {
	URL         URL    `json:"url"` //json 일경우 소문자로 반환 함
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty 값이 없으면 key도 지움
}

type URL string

// 인터페이스 구현 - 메서드의 시그니처를 구현 하면 된다
func (u URL) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}
func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []URLDescription{
		{
			URL:         URL("/"),
			Method:      "GET",
			Description: "see documentation",
		},
		{
			URL:         URL("/blocks"),
			Method:      "POST",
			Description: "Add a Block ",
			Payload:     "data:string",
		},
	}
	rw.Header().Add("Content-Type", "application-json")

	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
	//위와 동일함
	json.NewEncoder(rw).Encode(data)

}

func main() {
	// explorer.Start()
	http.HandleFunc("/", documentation)
	fmt.Printf("http://localhost%s\n", port)
	log.Fatal((http.ListenAndServe(port, nil)))

}
