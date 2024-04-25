/*
to run before running main.go:
go mod init github.com/bh1995/go-bookstore
go get "github.com/jinzhu/gorm"
go get "github.com/jinzhu/gorm/dialects/mysql"
go get "github.com/gorilla/mux"
*/

package main

import (
	"log"
	"net/http"

	"github.com/bh1995/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	router := mux.NewRouter()

	routes.RegisterBookStoreRoutes(router)

	http.Handle("/", router)

	log.Fatal(http.ListenAndServe("localhost:9011", router))
}
