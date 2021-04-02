package main

import (
	"log"
	"net/http"
	"push"
	"strconv"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func hello(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if e := recover(); e != nil {
			log.Println(e)
		}
	}()
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	text := r.Form.Get("text")
	result, err := push.PushText(text, "config/config.json", "config/token.json")
	if result == nil {
		return
	}
	_, err = w.Write([]byte(strconv.Itoa(result.Errcode)))
	if err != nil {
		panic(err)
	}
}
