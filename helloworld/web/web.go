package web

import (
	"fmt"
	"net/http"
)

type Handler struct {
	content string
}

func (handler Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, handler.content)
}

func wrap(handler func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") //允许访问所有域
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		w.Header().Set("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers, Cache-Control, Content-Language, Content-Type")
		//w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("content-type", "application/json") //返回数据格式是json
		handler(w, r)
	}
}

func Start(handler map[string]func(w http.ResponseWriter, r *http.Request)) {
	h := &Handler{
		content: "Hello",
	}
	http.Handle("/", h)
	for k, v := range handler {
		fmt.Printf("start listening %s \n", k)
		http.HandleFunc(k, wrap(v))
	}
	http.ListenAndServe("0.0.0.0:8080", nil)
}
