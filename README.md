# go-bookstore

This repo can be used as a bookstore API to create, delete, update and get information about books stored in a Maria DB (mysql) database. 

There are some dependicies for runing this project, they are the following:

1) in the main folder run the mod initialization command: go mod init github.com/bh1995/go-bookstore
2) make sure the following three packages are installed with the folling commands: go get "github.com/jinzhu/gorm", go get "github.com/jinzhu/gorm/dialects/mysql", go get "github.com/gorilla/mux".

The project can be tested for example with Postman, see my blog article @ "" about this project for examples on how to use the bookstore API but also on how to set the database up locally.  

