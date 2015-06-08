package main

import (
	"net/http"
	"time"
)

func main(){
	
	p("goinfo", version(), "started at", config.Address)
	
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(config.Static))
	mux.Handle("/static/",http.StripPrefix("/static/", files))
	
	//defininf in route_main.go
	mux.HandleFunc("/",index)
	
	//defining in route_service.go
	mux.HandleFunc("/service/continente", continente)
	
	//defining in route_mongo.go
	mux.HandleFunc("/mongo/analisis", analisis)	
	
	//starting up the server
	
	server := &http.Server{
		Addr: 			config.Address,
		Handler:			mux,
		ReadTimeout:		time.Duration(config.ReadTimeout * int64(time.Second)),
		WriteTimeout:	time.Duration(config.WriteTimeout * int64(time.Second)),
		MaxHeaderBytes:	1 << 20,
	}
	
	server.ListenAndServe()

}