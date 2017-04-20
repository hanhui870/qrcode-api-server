package main

import (
	"net/http"
	"log"
	"qrcode/handler"
)

func main(){
	api := handler.NewApi();

	http.HandleFunc("/qrcode", api.Create)

	log.Println("server startd http://127.0.0.1:8787")
	log.Fatal(http.ListenAndServe(":8787", nil))
}
