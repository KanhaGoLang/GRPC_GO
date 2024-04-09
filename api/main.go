package main

import (
	"log"
	"net/http"

	"github.com/KanhaGoLang/go_common/common"
	"github.com/KanhaGoLang/grpc_go/api/routes"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	routes.InitUserRoutes(router)

	common.MyLogger.Println(color.GreenString("Server started on port %s", common.RestAPI))
	log.Fatal(http.ListenAndServe(common.RestAPI, router))
}
