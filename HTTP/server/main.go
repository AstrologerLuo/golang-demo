package main
import (
	"log"
	"net/http"
)

func main(){
	server := http.Server{
		Addr:"127.0.0.1:8080"
	}
	http.HandleFunc("/",func(writer http.ResponseWriter, request *http.Request))
}