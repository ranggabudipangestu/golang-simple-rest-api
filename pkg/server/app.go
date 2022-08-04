package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/ranggabudipangestu/go-book-store/pkg/routes"
	_ "gorm.io/driver/mysql"
)

func InitServer() {
	r := mux.NewRouter()
	fmt.Println("connecting to routes")
	routes.RegiterBookStoreRoutes(r)
	fmt.Println("connected")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9020", r))
}
