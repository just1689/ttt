package io

import "net/http"

func StartServer(listenAddr string) {
	http.HandleFunc("/list", HandleList)
	http.HandleFunc("/new", HandleNew)
	http.HandleFunc("/join", HandleJoin)
	http.HandleFunc("/move", HandleMove)
	http.HandleFunc("/poll", HandlePoll)
	go panic(http.ListenAndServe(listenAddr, nil))
}
