package main

import "net/http"

func main() {
	messages := &messages{}

	viewHandler := &viewHandler{messages: messages}
	sendHandler := &sendHandler{messages: messages}

    http.Handle("/", viewHandler)
    http.Handle("/send/", sendHandler)
    http.ListenAndServe(":8081", nil)
}

