package main

import "net/http"
import "fmt"

var viewHandlerImplements http.Handler = (*viewHandler)(nil)

type viewHandler struct {
	messages* messages
}

func (vh *viewHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "the message was: %s", vh.messages.lastMessage);
}

