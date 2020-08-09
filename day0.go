package main

import (
	"encoding/json"
	"net/http"
)
type AddJsonRequest struct{
	A int `json:"p"`
	B int `json:"q"`
}

type AddJsonResponse struct {
	Ans int `json:"Whatever"`
	Answer int `json:"This"`
}


func main() {
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/add",add)
http.ListenAndServe(":8080", nil)
}

func hello(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Hello World"))

}

func add(writer http.ResponseWriter, request *http.Request) {
	 var req AddJsonRequest
	 json.NewDecoder(request.Body).Decode(&req)
	 var sum int
	 sum = req.A + req.B
	 var resp *AddJsonResponse
	 resp=new(AddJsonResponse)
	 resp.Ans=sum
	 resp.Answer=sum+1
	 json.NewEncoder(writer).Encode(resp)

}
