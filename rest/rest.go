package rest

import (
	"coin/blockchain"
	"coin/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

var port string

type urlDescription struct {
	URL         url    `json:"url"` //json 일경우 소문자로 반환 함
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty 값이 없으면 key도 지움
}

type url string

type addBlockBody struct {
	Message string
}

// 인터페이스 구현 - 메서드의 시그니처를 구현 하면 된다
func (u url) MarshalText() ([]byte, error) {
	url := fmt.Sprintf("http://localhost%s%s", port, u)
	return []byte(url), nil
}
func documentation(rw http.ResponseWriter, r *http.Request) {
	data := []urlDescription{
		{
			URL:         url("/"),
			Method:      "GET",
			Description: "see documentation",
		},
		{
			URL:         url("/blocks"),
			Method:      "POST",
			Description: "Add a Block ",
			Payload:     "data:string",
		},
		{
			URL:         url("/blocks/{id}"),
			Method:      "GET",
			Description: "see a block",
		},
	}
	rw.Header().Add("Content-Type", "application-json")

	// b, err := json.Marshal(data)
	// utils.HandleErr(err)
	// fmt.Fprintf(rw, "%s", b)
	//위와 동일함
	json.NewEncoder(rw).Encode(data)

}
func blocks(rw http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		rw.Header().Add("Content-Type", "application-json")
		json.NewEncoder(rw).Encode(blockchain.GetBlockchain().AllBlock())
	case "POST":
		var addBlockBody addBlockBody
		utils.HandleErr(json.NewDecoder(r.Body).Decode(&addBlockBody))
		blockchain.GetBlockchain().AddBlock(addBlockBody.Message)
		rw.WriteHeader(http.StatusCreated)
	}
}
func Start(aport int) {
	handler := http.NewServeMux()
	port = fmt.Sprintf(":%d", aport)
	handler.HandleFunc("/", documentation)
	handler.HandleFunc("/blocks", blocks)
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal((http.ListenAndServe(port, handler)))
}
