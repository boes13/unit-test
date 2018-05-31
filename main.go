package main

import (
	"log"
	"net/http"

	"github.com/boes13/unit-test/common"
	"github.com/boes13/unit-test/service/http_handler"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	log.Println("start")

	common.InitDataConnection()

	http.HandleFunc("/get_user_email", http_handler.GetUserEmail)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
