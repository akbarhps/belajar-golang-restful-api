package main

import (
	"net/http"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/middleware"

	_ "github.com/go-sql-driver/mysql"
)

func NewServer(authMiddleware *middleware.AuthMiddleware) *http.Server {
	return &http.Server{
		Addr:    "localhost:3000",
		Handler: authMiddleware,
	}

}

func main() {
	server := InitializedServer()
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
