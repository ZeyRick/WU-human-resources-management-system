package controller

import (
	"fmt"
	"net/http"
)

type HelloWorld struct {
}

func NewHelloWorl() *HelloWorld {
	return &HelloWorld{}
}

func GetHelloWorld(w http.ResponseWriter, r *http.Request){
	fmt.Println(r.Header)
	w.Write([]byte("This Hello World "))
}