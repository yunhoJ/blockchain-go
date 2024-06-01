package rest

import (
	"coin/blockchain"
	"coin/utils"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var port string

type urlDescription struct {
	URL         url    `json:"url"` //json 일경우 소문자로 반환 함
	Method      string `json:"method"`
	Description string `json:"description"`
	Payload     string `json:"payload,omitempty"` // omitempty 값이 없으면 key도 지움
}
type errorResponse struct {
	ErrorMessage string `json:"errorMessage"`
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
			URL:         url("/blocks/{height}"),
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
func block(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	height := vars["height"]
	id, err := strconv.Atoi(height)
	utils.HandleErr(err)
	block, err := blockchain.GetBlockchain().GetBlock(id)
	encoder := json.NewEncoder(rw)

	if err == blockchain.ErrNotFound {
		encoder.Encode(errorResponse{fmt.Sprint(err)})
	} else {
		encoder.Encode(block)
	}

}

func jsonContentTypeMiddleware(next http.Handler) http.Handler {
	// 직접 struct, type을 만들어서 serveHttp를 구현 하는대신
	// 알맞은 매개변수로 구현 해줌 - adapter
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}
func Start(aport int) {
	router := mux.NewRouter()
	router.Use(jsonContentTypeMiddleware)
	port = fmt.Sprintf(":%d", aport)
	router.HandleFunc("/", documentation).Methods("GET")
	router.HandleFunc("/blocks", blocks).Methods("GET", "POST")
	router.HandleFunc("/blocks/{height:[0-9]+}", block).Methods("GET")
	fmt.Printf("Listening on http://localhost%s\n", port)
	log.Fatal((http.ListenAndServe(port, router)))
}
