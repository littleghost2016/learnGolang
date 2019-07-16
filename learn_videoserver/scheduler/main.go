package main

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
	"learnGolang/learn_videoserver/scheduler/taskrunner"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()

	router.GET("/video-delete-record/:vid-id", vidDelRecHandler)

	return router
}

func main() {
	// c := make(chan int)
	go taskrunner.Start()
	r := RegisterHandlers()
	// 加入没有下面这一句，就要把两句注释去掉
	http.ListenAndServe(":9001", r)
	// <- c
}