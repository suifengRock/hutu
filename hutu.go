package main

import (
	"encoding/json"
	"fmt"
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

func CreateHost(rsp http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		fmt.Println("...")
	}
	val := req.FormValue("host")
	fmt.Println(val)
	http.Redirect(rsp, req, "/apilist/", http.StatusFound)
}

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/apilist/", http.StatusFound)
	}

	t, err := template.ParseFiles("src/hutu/views/api_list.html")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(r.URL.Path)
	t.Execute(w, nil)

}

func main() {
	http.HandleFunc("/", NotFoundHandler)
	http.HandleFunc("/create/host/", CreateHost)
	http.HandleFunc("/apilist/", ApiListView)
	http.HandleFunc("/say", SayHello)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./src/hutu/views"))))
	http.ListenAndServe(":8001", nil)
}
