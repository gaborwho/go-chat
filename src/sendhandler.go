package main

import "net/http"
import "fmt"

var sendHandlerImplements http.Handler = (*sendHandler)(nil)

type sendHandler struct {
	messages *messages
}

func (vh *sendHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vh.messages.lastMessage = r.URL.Path
	fmt.Fprint(w, "message was sent")
}
