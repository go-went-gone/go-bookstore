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
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
