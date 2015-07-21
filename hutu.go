package main

import (
	"encoding/json"
	"net/http"
	"text/template"
)

type Server struct {
	ServerName string
	ServerIP   string
}

func SayHello(rsp http.ResponseWriter, req *http.Request) {
	s := Server{ServerName: "localhost", ServerIP: "127.0.0.1"}
	str, _ := json.Marshal(s)
	rsp.Write(str)
}

func ApiListView(rsp http.ResponseWriter, req *http.Request) {
	t, _ := template.ParseFiles("src/hutu/views/api_list.html")
	t.Execute(rsp, nil)
}

func main() {
	http.HandleFunc("/apilist/", ApiListView)
	http.HandleFunc("/say", SayHello)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./src/hutu/views"))))
	http.ListenAndServe(":8001", nil)
}
